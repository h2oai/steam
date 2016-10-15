# this is run within the predictor
import sys
import pickle
import argparse
from textblob import TextBlob

MODELS_DESTINATION_DIR = "./lib"

def loadModel(source):
    with open(source, 'rb') as f:
        return pickle.load(f)

def split_into_lemmas(message):
    message = unicode(message, 'utf8').lower()
    words = TextBlob(message).words
    return [word.lemma for word in words if len(word) > 0 and word.isalpha() ]


class Scorer(object):
    def __init__(self, model_file):
        self._init_model(model_file)

    def _init_model(self, model_file):
        self.model = loadModel(model_file)
        self.labels = self.model.get_feature_names()

    def score(self, message):
        v = self.model.transform([message])
        s = mappify(v, self.labels)
        return s


# Necessary methods

# Write your main program to first do any kind of setup and then call the main_loop with the transformer function
# you want to apply to each line of input.
# The input is always a string which is one line of input
# The output is a number of COLUMN_NAME:VALUE separated by space.
# Each label is the column name and the value is the value for this column and row

def mappify(v, labels):
    """Convert vector v to RowData type Map<String, Float>
    coded as space-separated LABEL:NUMBER """
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

def main_loop(transformer):
    """This infinite loop reads one line, then transforms it and prints the line"""
    while True:
        input = raw_input()
        res = transformer(input)
        print res
        sys.stdout.flush()

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

    main_loop(scorer.score)
