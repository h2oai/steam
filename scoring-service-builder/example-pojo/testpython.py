# this is run within the predictor
import sys

def oneLine(text):
#    return '{"text": "' + text + '" }'
    return text

if __name__ == "__main__":
    while True:
        logString = raw_input()
        if len(logString) > 0:
#            res = production_main(FE, logString, conn)
            res = oneLine(logString)
            print str(res)
            sys.stdout.flush()
