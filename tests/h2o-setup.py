#!/usr/bin/env python

from sys import argv
import h2o
from h2o.estimators.deeplearning import H2ODeepLearningEstimator
from h2o.estimators.gbm import H2OGradientBoostingEstimator
from h2o.estimators.glm import H2OGeneralizedLinearEstimator

address = argv[1]
ip, port = address.split(":")

h2o.init(ip=ip, port=port)

# Create three different frames for testing model categories
bin_frame = h2o.create_frame(
    has_response=True,
    response_factors=2,
    frame_id="bin_hex")
mul_frame = h2o.create_frame(
    has_response=True,
    response_factors=10,
    frame_id="mul_hex")
reg_frame = h2o.create_frame(
    cols=11,
    categorical_fraction=0,
    binary_fraction=0,
    frame_id="reg_hex")

# Create multiple models for testing
# gbm models
bin_gbm = H2OGradientBoostingEstimator(model_id="bin_gbm")
bin_gbm.train(x=list(range(1, bin_frame.ncol)), y=0, training_frame=bin_frame)
mul_gbm = H2OGradientBoostingEstimator(model_id="mul_gbm")
mul_gbm.train(x=list(range(1, mul_frame.ncol)), y=0, training_frame=mul_frame)
reg_gbm = H2OGradientBoostingEstimator(model_id="reg_gbm")
reg_gbm.train(x=list(range(1, reg_frame.ncol)), y=0, training_frame=reg_frame)

# glm models
bin_glm = H2OGeneralizedLinearEstimator(model_id="bin_glm", family="binomial")
bin_glm.train(x=list(range(1, bin_frame.ncol)), y=0, training_frame=bin_frame)
reg_glm = H2OGeneralizedLinearEstimator(model_id="reg_glm")
reg_glm.train(x=list(range(1, reg_frame.ncol)), y=0, training_frame=reg_frame)

# deeplearning models
bin_dpl = H2ODeepLearningEstimator(model_id="bin_dpl")
bin_dpl.train(x=list(range(1, bin_frame.ncol)), y=0, training_frame=bin_frame)
mul_dpl = H2ODeepLearningEstimator(model_id="mul_dpl")
mul_dpl.train(x=list(range(1, mul_frame.ncol)), y=0, training_frame=mul_frame)
reg_dpl = H2ODeepLearningEstimator(model_id="reg_dpl")
reg_dpl.train(x=list(range(1, reg_frame.ncol)), y=0, training_frame=reg_frame)
