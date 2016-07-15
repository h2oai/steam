curl -X POST \
--form pojo=@gbm_3f258f27_f0ad_4520_b6a5_3d2bb4a9b0ff.java \
--form jar=@h2o-genmodel.jar \
localhost:55000/compile > example.jar

echo "compiled to example.jar"


