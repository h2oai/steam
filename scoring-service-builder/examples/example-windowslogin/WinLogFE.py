import re
import numpy as np
import sys
import itertools
import datetime
from dateutil import parser
from memsql.common import database

PROJECT_PATH = "/home/markc/H2O/experiments-mklechan/ciso.ai"
HEADER_DATA = ["ComputerName","Security ID2","Source Network Address","datetime","Logon Type","Logon Process","Authentication Package","Account Domain2","Key Length","Package Name (NTLM only)"]

class FeatureExtraction:

    def __init__(self):
        self.lastTime = None

    def firstDiff(self,dt):
        # must run on 2nd element of array
        if self.lastTime == None:
            raise ValueError("lastTime not initialized")
        res = dt - self.lastTime
        self.lastTime = dt
        return res.seconds

    def getAccountType(self,s):
        sp = s.split("\\")
        if len(sp) > 1:
            sp = sp[1]
        else:
            sp = sp[0]

        if re.search("A-",sp):
            return "Admin"
        elif re.search("^[a-zA-Z]{3}\\d{3}$",sp):
            return "User"
        else:
            return "Service"

def prep_memsql(conn):
    # MemSQL Prep Comment if not needed
    conn.execute('CREATE DATABASE IF NOT EXISTS winsyslog;')
    conn.execute("USE winsyslog;")
    conn.execute('''CREATE TABLE IF NOT EXISTS raw_events (
                    computername varchar(256),
                    security_id2 varchar(256),
                    source varchar(64),
                    datetime datetime,
                    logon_type varchar(4),
                    logon_process varchar(128),
                    authentication_package varchar(64),
                    account_domain2 varchar(128),
                    key_length varchar(16),
                    package_name_ntlm varchar(64),
                    id INT AUTO_INCREMENT,
                    PRIMARY KEY (id)
                 )
                 ''')
    conn.execute('''CREATE TABLE IF NOT EXISTS processed_events (
                    computername varchar(256),
                    security_id2 varchar(256),
                    source varchar(64),
                    account_type varchar(64),
                    logon_type varchar(4),
                    tbtw_skew float,
                    count_skew float,
                    PRIMARY KEY (security_id2)
                )
                 ''')
    conn.execute('''CREATE TABLE IF NOT EXISTS pca_processed_events (
                    security_id2 varchar(256),
                    account_type varchar(256),
                    logon_type varchar(4),
                    tbtw_skew float,
                    count_skew float,
                    PC1 float,
                    PC2 float,
                    PRIMARY KEY (security_id2)
                   )''')

    conn.execute('''CREATE TABLE IF NOT EXISTS last_query (
                    security_id2 varchar(256),
                    last_query datetime,
                    PRIMARY KEY (security_id2)
                )
            ''')


def training_main():

    import sys
    import datetime
    sys.path.insert(1,"/home/markc/H2O/h2o-3/h2o-py/")
    import h2o
    from h2o.transforms.decomposition import H2OPCA
    h2o.init(ip="172.16.2.56", strict_version_check=False) # make sure to start h2o with the java -cp h2o.jar:/usr/share/java/mysql.jar water.H2OApp

    df = h2o.import_sql_table(connection_url="jdbc:mysql://localhost:3306/winsyslog?&useSSL=false", table="processed_events", username="root",password="")
    for c in ["account_type","logon_type"]:
        df[c] = df[c].asfactor()
    print len(df)
    # In this case no feature munging is needed. Its already been done on the data ingestion side!

    model = H2OPCA(k=2, transform="STANDARDIZE", pca_method="GramSVD", use_all_factor_levels=True, max_iterations=10000, model_id="WinLogPCA")

    model.train(x=["account_type","logon_type","tbtw_skew","count_skew"], training_frame=df)
    print('Download generated POJO for model')
    model.download_pojo(path=PROJECT_PATH+'/h2o/models')
    #h2o.shutdown(prompt=False)

def production_main(FE, logString, conn):


    # Parse string to dictionary
    logValues = logString.replace("]","").replace("[","").split(",")
    logData = dict([(k,v) for k,v in zip(HEADER_DATA,logValues)])

    # Parse datetime
    logData["datetime"] = str(parser.parse(logData["datetime"]))

    # load raw data
    load_qry = '''INSERT INTO raw_events
                  (computername,
                  security_id2,
                  source,
                  datetime,
                  logon_type,
                  logon_process,
                  authentication_package,
                  account_domain2,
                  key_length,
                  package_name_ntlm)
                  VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s)'''

    conn.execute(load_qry, logData["ComputerName"],
                 logData["Security ID2"],
                 logData["Source Network Address"],
                 logData["datetime"],
                 logData["Logon Type"],
                 logData["Logon Process"],
                 logData["Authentication Package"],
                 logData["Account Domain2"],
                 logData["Key Length"],
                 logData["Package Name (NTLM only)"])

    # For runtime efficiency check if we should do the full skew analysis
    COUNT_THRESHOLD = 1000  # number of events at which we analyze periodically
    HIGH_COUNT_ANALYSIS_PERIODICITY = 0  # every 5 minutes for high login count security ids
    count = conn.query("SELECT count(*) FROM raw_events;")
    if count[0]["count(*)"] > COUNT_THRESHOLD:
        last_query = conn.query("SELECT * FROM last_query WHERE security_id2 = %s", logData["Security ID2"])
        now = datetime.datetime.now()
        if len(last_query) > 0:
            if (now - last_query[0]["last_query"]).seconds < HIGH_COUNT_ANALYSIS_PERIODICITY:
                return {}
        # Update the last query table
        conn.execute('''INSERT INTO last_query (security_id2, last_query) VALUES (%s, %s)
                     ON DUPLICATE KEY UPDATE
                     last_query = %s''',
                        logData["Security ID2"],
                        str(now),
                        str(now))



    # Get time series data for account,source and destination
    qry = "SELECT source, datetime FROM raw_events WHERE security_id2 = %s ORDER BY datetime ASC;"
    data = conn.query(qry, logData["Security ID2"])

    # Get account type
    _account_type = FE.getAccountType(logData["Security ID2"])

    if len(data) > 0:
        # Calculate skewness of the login count per source for this account host pair
        _counts = []
        for k,g in itertools.groupby(data, lambda x: x["source"]):
            _counts.append(len([i for i in g]))
        def get_skew(l):
            _mean = np.mean(l)
            _median = np.median(l)
            _std = np.std(l)
            if _std > 0:
                _skew = (_mean - _median)/_std
            else:
                # we dont have enough data, start with 0
                _skew = 0.0
            return _skew

        try:
            _count_skew = get_skew(_counts)
        except:
            _count_skew = 0.0
        

        # Caculate time between skew
        _dts = [i["datetime"] for i in data]
        FE.lastTime = _dts[0]
        _time_diffs = map(lambda x: FE.firstDiff(x), _dts[1:] )
        FE.lastTime = None
        try:
            _tbtw_skew = get_skew(_time_diffs)
        except:
            _tbtw_skew = 0.0

    else:
        _count_skew = 0.0
        _tbtw_skew = 0.0


    # load processed data
    load_qry =  '''INSERT INTO processed_events
                    (computername,
                    security_id2,
                    source,
                    account_type,
                    logon_type,
                    tbtw_skew,
                    count_skew)
                    VALUES (%s,%s,%s,%s,%s,%s,%s)
                    ON DUPLICATE KEY UPDATE
                     account_type = %s,
                     logon_type = %s,
                     tbtw_skew = %s,
                     count_skew = %s,
                     computername = %s,
                     source = %s'''
    conn.execute(load_qry, logData["ComputerName"],
                           logData["Security ID2"],
                           logData["Source Network Address"],
                           _account_type,
                           logData["Logon Type"],
                           _tbtw_skew,
                           _count_skew,
                           _account_type,  # repeated for duplicate key update
                           logData["Logon Type"],
                           _tbtw_skew,
                           _count_skew,
                           logData["ComputerName"],
                           logData["Source Network Address"])

    # These field names should match exactly to the training variables above
    return { "account_type":_account_type,
             "logon_type":logData["Logon Type"],
             "tbtw_skew":_tbtw_skew,
             "count_skew":_count_skew}


if __name__ == "__main__":
    
    conn = database.connect(host="127.0.0.1", user="root")
    # comment if tables already created
    prep_memsql(conn)
    FE = FeatureExtraction()
    while True:
        logString = raw_input()
        if len(logString) > 0:
            res = production_main(FE, logString, conn)
            print(str(res).replace("}","").replace("{",""))
            sys.stdout.flush()
# for storm pipeline
#    if sys.argv[1] == "training":
#        training_main()
#    elif sys.argv[1] == "production":
#        conn = database.connect(host="127.0.0.1", user="root")
#        # comment if tables already created
#        prep_memsql(conn)
#        FE = FeatureExtraction()
#        while True:
#            logString = raw_input()
#            if len(logString) > 0:
#                res = production_main(FE, logString, conn)
#                print str(res)
#                sys.stdout.flush()
#
