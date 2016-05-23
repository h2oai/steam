# This is run within the predictor
import os
import sys
import h2o
import numpy as np
import argparse
from sklearn.feature_extraction.text import TfidfVectorizer
from h2o.estimators.gbm import H2OGradientBoostingEstimator

# Shared user library between training and scoring
from lib.modelling import split_into_lemmas, saveModel

# Should be input parameter
#MODELS_DESTINATION_DIR = "./models"
MODELS_DESTINATION_DIR = "/tmp/models"

# Load data
def load_data(filename):
     # Note: skiping wrong rows
     return np.genfromtxt(filename, dtype=None, delimiter='\t', names=['label', 'message'], skip_header=0, invalid_raise=False)

def tf_idf(corpus):
    vectorizer = TfidfVectorizer(
            #analyzer = 'word',
            analyzer = split_into_lemmas,
            stop_words = 'english',
            min_df=0,
            decode_error = 'ignore',
            strip_accents = 'ascii',
            ngram_range=(1,3))
    # Fit and transform input corpus
    model = vectorizer.fit_transform(corpus)
    return (vectorizer, model)

def train(cfg):
    # Load data
    messages = load_data(cfg.datafile)
    # Prepare tf-idf to feature vectorization and also transform input data
    (vectorizer, train) = tf_idf(messages['message'])
    # Save Tf-Idf model
    h2o.init()
    train_table = h2o.H2OFrame(np.column_stack((messages['label'], train.toarray()))).set_names(['label'] + vectorizer.get_feature_names())
    gbm_model= H2OGradientBoostingEstimator(ntrees=1, learn_rate=0.01, max_depth=6, min_rows=10, distribution="bernoulli")
    gbm_model.train(x = range(1, train_table.shape[1]), y = 0, training_frame = train_table)
    if cfg.verbose: print "GBM Model", gbm_model
    # Save models
    if not os.path.exists(cfg.models_dir):
        os.makedirs(cfg.models_dir)
    saveModel(vectorizer, '{}/vectorizer.pickle'.format(cfg.models_dir))
    h2o.download_pojo(gbm_model, "{}/".format(cfg.models_dir))
    h2o.shutdown()
    return vectorizer

#
# Main entry point. Accepts parameters
#  for example:
#    ipython train.py -- --verbose --models-dir /tmp/models
#
# if __name__ == '__main__':
#     parser = argparse.ArgumentParser(description = "Train models for Spam detection")
#     parser.add_argument('--datafile', help = 'Input data file', type=str, default = 'data/smsData.txt')
#     parser.add_argument('--models-dir', help = 'Directory to save generated models', type=str, default = MODELS_DESTINATION_DIR)
#     parser.add_argument('--verbose', help = 'More detailed output', dest='verbose', action='store_true')
#     cfg = parser.parse_args()
#     train(cfg)

# this is run within the predictor
# import sys
# import pickle
# import json
# import argparse
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

class Scorer(object):
    def __init__(self, model_file):
        self._init_model(model_file)

    def _init_model(self, model_file):
        self.model = loadModel(model_file)

    def score(self, message):
        return json.dumps(self.model.transform([message]).toarray()[0].tolist())
#
# Main entry point. Accepts parameters
#  for example:
#    ipython score.py -- --verbose --models-dir /tmp/models
#
if __name__ == "__main__":
    parser = argparse.ArgumentParser(description = "Train models for Spam detection")
    parser.add_argument('--datafile', help = 'Input data file', type=str, default = 'data/smsData.txt')
    parser.add_argument('--models-dir', help = 'Directory to save generated models', type=str, default = MODELS_DESTINATION_DIR)
    parser.add_argument('--verbose', help = 'More detailed output', dest='verbose', action='store_true')
    cfg = parser.parse_args()
    vectorizer = train(cfg)

    # parser = argparse.ArgumentParser(description = "Detect Spam messages")
    # parser.add_argument('--models-dir', help = 'Directory with saved models', type=str, default = MODELS_DESTINATION_DIR)
    # parser.add_argument('--verbose', help = 'More detailed output', dest='verbose', action='store_true')
    # cfg = parser.parse_args()
    # scorer = Scorer('{}/vectorizer.pickle'.format(cfg.models_dir))
    while True:
        logString = raw_input()
        if len(logString) > 0:
            # res = scorer.score(logString)
            res = vectorizer(logString)
            print str(res)
            sys.stdout.flush()
