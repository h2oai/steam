# This is run within the predictor
import sys
import h2o
from h2o.estimators.gbm import H2OGradientBoostingEstimator
import numpy as np
import lib.modelling as modelling

from sklearn.feature_extraction.text import TfidfVectorizer

verbose = False

# Load data
def load_data(filename):
     # Note: skiping wrong rows
     return np.genfromtxt(filename, dtype=None, delimiter='\t', names=['label', 'message'], skip_header=0, invalid_raise=False)

def tf_idf(corpus):
    vectorizer = TfidfVectorizer(
            #analyzer = 'word',
            analyzer = modelling.split_into_lemmas,
            stop_words = 'english',
            min_df=0,
            decode_error = 'ignore',
            strip_accents = 'ascii',
            ngram_range=(1,3))
    # Fit and transform input corpus
    model = vectorizer.fit_transform(corpus)
    return (vectorizer, model)

def train(datafile):
    # Load data
    messages = load_data(datafile)
    # Prepare tf-idf to feature vectorization and also transform input data
    (vectorizer, train) = tf_idf(messages['message'])
    # Save Tf-Idf model
    modelling.saveModel(vectorizer, './models/vectorizer.pickle')
    h2o.init()
    train_table = h2o.H2OFrame(np.column_stack((messages['label'], train.toarray()))).set_names(['label'] + vectorizer.get_feature_names())
    gbm_model= H2OGradientBoostingEstimator(ntrees=1, learn_rate=0.01, max_depth=6, min_rows=10, distribution="bernoulli")
    gbm_model.train(x = range(1, train_table.shape[1]), y = 0, training_frame = train_table)
    if verbose: print "GBM Model", gbm_model
    h2o.download_pojo(gbm_model, "./models/")
    h2o.shutdown()

if __name__ == '__main__':
    train('data/smsData.txt')

