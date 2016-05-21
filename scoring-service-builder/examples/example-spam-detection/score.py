# this is run within the predictor
import sys
import pickle
import json
from textblob import TextBlob
import lib.modelling as modelling

class Scorer(object):
    def __init__(self, model_file):
        self._init_model(model_file)

    def _init_model(self, model_file):
        self.model = modelling.loadModel(model_file)

    def score(self, message):
        return json.dumps(self.model.transform([message]).toarray()[0].tolist())

if __name__ == "__main__":
    scorer = Scorer("./models/vectorizer.pickle")
    while True:
        logString = raw_input()
        if len(logString) > 0:
            res = scorer.score(logString)
            print str(res)
            sys.stdout.flush()
