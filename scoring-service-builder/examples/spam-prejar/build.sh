
javac -cp 'lib/*' Transform.java PreProcess.java Tfidf.java

jar cf pre.jar Transform.class PreProcess.class  Tfidf.class


curl -X POST \
--form pojo=@gbm_cf6fdeef_cad1_4e85_b644_6358166076ca.java \
--form jar=@h2o-genmodel.jar \
--form prejar=@pre.jar \
--form preclass=PreProcess \
localhost:55000/makewar > example.war

echo "Created example.war"
echo "Run with run-example.sh"

