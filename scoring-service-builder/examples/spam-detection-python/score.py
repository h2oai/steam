# this is run within the predictor
import sys
import pickle
import argparse
from textblob import TextBlob

MODELS_DESTINATION_DIR = "./lib"

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
    return sparsify(v)

def sparsify(v):
    nz = v.nonzero()[1]
    dat = v.data
    r = ""
    i = 0
    len = nz.size
    while i < len:
        index = nz[i]
        value = dat[i]
        r += repr(index) + ":" + repr(round(value, 6)) + " "
        i += 1
    return r

def mappify(v, labels):
    nz = v.nonzero()[1]
    dat = v.data
    r = ""
    i = 0
    len = nz.size
    while i < len:
        index = nz[i]
        value = dat[i]
        r += str(labels[index]) + ":" + repr(round(value, 6)) + " "
        i += 1
    return r

def score2(vectorizer, message):
    v = vectorizer.transform([message])
    return sparsify(v)


class Scorer(object):
    def __init__(self, model_file):
        self._init_model(model_file)

    def _init_model(self, model_file):
        self.model = loadModel(model_file)
        self.labels = self.model.get_feature_names()

    # def score0(self, message):
    #     return score2(self.model, message)

    def score(self, message):
        v = self.model.transform([message])
        # s = sparsify(v)
        s = mappify(v, self.labels)
        return s

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
        input = raw_input()
        if len(input) > 0:
            res = scorer.score(input)
            print res
            sys.stdout.flush()
