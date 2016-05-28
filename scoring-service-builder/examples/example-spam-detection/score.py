# this is run within the predictor
import sys
import pickle
import json
import argparse
from textblob import TextBlob

# FIXME: shared library with training
#import lib.modelling as modelling

# Should be input parameter
#MODELS_DESTINATION_DIR = "./models"
MODELS_DESTINATION_DIR = "/tmp/models"

## FIXME: should be provided by shared library
def loadModel(source):
    with open(source, 'rb') as f:
        return pickle.load(f)

def split_into_lemmas(message):
    message = unicode(message, 'utf8').lower()
    words = TextBlob(message).words
    return [word.lemma for word in words if len(word) > 0 and word.isalpha() ]
## END of FIXME

def score1(vectorizer, message):
    v = vectorizer.transform([message])
    nz = v.nonzero()[1]
    dat = v.data
    r = ""
    i = 0
    len = nz.size
    while i < len:
        index = nz[i]
        value = dat[i]
        r += repr(index) + ":" + repr(value) + " "
        i += 1
    return r


class Scorer(object):
    def __init__(self, model_file):
        self._init_model(model_file)

    def _init_model(self, model_file):
        self.model = loadModel(model_file)

    def score(self, message):
        return score1(self.model, message)

#
# Main entry point. Accepts parameters
#  for example:
#    ipython score.py -- --verbose --models-dir /tmp/models
#
if __name__ == "__main__":
    parser = argparse.ArgumentParser(description = "Detect Spam messages")
    parser.add_argument('--models-dir', help = 'Directory with saved models', type=str, default = MODELS_DESTINATION_DIR)
    parser.add_argument('--verbose', help = 'More detailed output', dest='verbose', action='store_true')
    cfg = parser.parse_args()
    scorer = Scorer('{}/vectorizer.pickle'.format(cfg.models_dir))
    print >> sys.stderr, "python ready"

    while True:
        logString = raw_input()
        if len(logString) > 0:
            res = scorer.score(logString)
            print res
            sys.stdout.flush()
