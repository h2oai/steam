# Run H2O models in R

library("rJava")
.jinit()
.jaddClassPath(".")

h2opredict <- function(jarFile, modelName, json) {
  jar <- paste(getwd(), "/", jarFile, sep="")
  .jaddClassPath(jar)
  res <- .jcall("H2OPredictor", "S", "predict", jar, modelName, json)
  return(res)
}

h2opredict("example.jar", "gbm_3f258f27_f0ad_4520_b6a5_3d2bb4a9b0ff", "{Dest:JFK}")

library("rjson")

h2opredictor <- function(directory, jarFile) {
  jar <- paste(directory, "/", jarFile, sep="")
  return(jar)
}
predict <- function(jar, modelName, json) {
  .jaddClassPath(jar)
  res <- .jcall("H2OPredictor", "S", "predict", jar, modelName, json)
  return(fromJSON(res))
}

models <- h2opredictor(getwd(), "example.jar")
predict(models, "gbm_3f258f27_f0ad_4520_b6a5_3d2bb4a9b0ff", "{Dest:JFK}")
res <- predict(models, "gbm_3f258f27_f0ad_4520_b6a5_3d2bb4a9b0ff", "{Dest:SFO}")
res

