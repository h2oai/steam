from textblob import TextBlob
import pickle

def split_into_lemmas(message):
    message = unicode(message, 'utf8').lower()
    words = TextBlob(message).words
    return [word.lemma for word in words if len(word) > 0 and word.isalpha() ]

def saveModel(model, destination):
    with open(destination, 'wb') as f:
        pickle.dump(model, f)

def loadModel(source):
    with open(source, 'rb') as f:
        return pickle.load(f)


