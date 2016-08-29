import h2o
from h2o.estimators.gbm import H2OGradientBoostingEstimator as gbm
from h2o.estimators.glm import H2OGeneralizedLinearEstimator as glm
from h2o.estimators.deeplearning import H2ODeepLearningEstimator as dlm

h2o.connect(ip="localhost", port="54535", cluster_name="SteamUItest")

arr = h2o.import_file(path="https://s3.amazonaws.com/h2o-public-test-data/smalldata/flow_examples/arrhythmia.csv.gz")
train = h2o.import_file(path="https://s3.amazonaws.com/h2o-public-test-data/smalldata/gbm_test/50_cattest_test.csv")
test = h2o.import_file(path="https://s3.amazonaws.com/h2o-public-test-data/smalldata/gbm_test/50_cattest_train.csv")
bank = h2o.import_file(path="https://s3.amazonaws.com/h2o-public-test-data/smalldata/gbm_test/bank-full.csv.zip")

model = glm(family="binomial", model_id="first")
multi = glm(family="multinomial", model_id="multinom")
valid = glm(family="multinomial", model_id="valimon")
regr = glm(model_id="regress")
grad = gbm(model_id="gradi")


#bank data columns
#1 for multinomial
#16 for binomial
#5 for regression

model.train(y=16, x=range(16), training_frame=bank)
multi.train(y=1, x=[i for i in range(17) if i != 1], training_frame=bank)
model = glm(family="binomial", model_id="second")
model.train(y=16, x=[i for i in range(16) if not i % 2 == 0], training_frame=bank)
model = glm(family="binomial", model_id="third")
model.train(y=16, x=[i for i in range(16) if i % 2 == 0], training_frame=bank)
model = glm(family="binomial", model_id="fourth")
model.train(y=16, x=[i for i in range(16) if not i % 3 == 0], training_frame=bank)




btr, btst = bank.split_frame(ratios=[0.75], seed=18319023)

valid.train(y=1, x=[i for i in range(17) if not i % 3 == 1], training_frame=bank)

multi = glm(family="multinomial", model_id="multimiss")
multi.train(y=1, x=[i for i in range(17) if i % 2 == 0], training_frame=bank)

valid = glm(family="multinomial", model_id="vamiss")
valid.train(y=1, x=[i for i in range(17) if not i % 5 == 1], training_frame=bank)

regr.train(y=5, x=[i for i in range(17) if i != 5], training_frame=bank)

grad.train(y=5, x=[i for i in range(17) if i != 5], training_frame=bank)

grad = gbm(model_id="missin")
regr = glm(model_id="linmiss")
grad.train(y=5, x=[i for i in range(17) if i % 2 == 0], training_frame=bank)
regr.train(y=5, x=[i for i in range(17) if i % 2 == 0], training_frame=bank)


#model.train(x=range(2), y=3, model_id="first", training_frame=train, validation_frame=test, family="binomial")

