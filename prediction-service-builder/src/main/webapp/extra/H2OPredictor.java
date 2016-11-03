import com.google.gson.Gson;
import com.google.gson.JsonSyntaxException;
import hex.genmodel.GenModel;
import hex.genmodel.easy.exception.PredictException;
import hex.genmodel.easy.prediction.AbstractPrediction;
import hex.genmodel.easy.EasyPredictModelWrapper;
import hex.genmodel.easy.RowData;
import hex.genmodel.MojoModel;

import java.io.File;
import java.io.FileNotFoundException;
import java.net.MalformedURLException;
import java.net.URL;
import java.net.URLClassLoader;

/**
 * Created by magnus on 5/5/16.
 */
class H2OPredictor {

  private EasyPredictModelWrapper model = null;

  private Gson gson = new Gson();

  private GenModel loadClassFromJar(String jarFileName, String modelName) throws Exception {
    if (!new File(jarFileName).isFile()) {
      throw new FileNotFoundException("Can't read " + jarFileName);
    }
    try {
      URL url = new File(jarFileName).toURI().toURL();
      ClassLoader loader = URLClassLoader.newInstance(
          new URL[]{url},
          getClass().getClassLoader()
      );
      String packagePrefix = "";
      String className = packagePrefix + modelName;
      Class<?> clazz = loader.loadClass(className);
      Class<? extends GenModel> modelClass = clazz.asSubclass(GenModel.class);
      return modelClass.newInstance();
    }
    catch (MalformedURLException e) {
      throw new Exception("Can't use Jar file" + jarFileName);
    }
    catch (ClassNotFoundException e) {
      throw new Exception("Can't find model " + modelName + " in jar file " + jarFileName);
    }
    catch (InstantiationException e) {
      throw new Exception("Can't find model " + modelName + " in jar file " + jarFileName);
    }
    catch (IllegalAccessException e) {
      throw new Exception("Can't find model " + modelName + " in jar file " + jarFileName);
    }
  }

  private void loadPojo(String jarFileName, String modelName)
      throws Exception {
    GenModel rawModel = loadClassFromJar(jarFileName, modelName);
    model = new EasyPredictModelWrapper(rawModel);
  }

  private void loadMojo(String zipFileName)
      throws Exception {
    GenModel rawModel = MojoModel.load(zipFileName);
    model = new EasyPredictModelWrapper(rawModel);
  }

  private RowData jsonToRowData(String json) {
    try {
      return gson.fromJson(json, RowData.class);
    }
    catch (JsonSyntaxException e) {
      throw new JsonSyntaxException("Malformed JSON");
    }
  }

  private String predict2(RowData row) throws PredictException {
    if (model == null)
      throw new PredictException("No model loaded");
    if (gson == null)
      throw new PredictException("Gson not available");
    if (row == null)
      throw new PredictException("No row data");
    AbstractPrediction pr = model.predict(row);
    return gson.toJson(pr);
  }

  public static String predict(String ojoFileName, String modelName, String jsonArgs) {
    try {
      H2OPredictor p = new H2OPredictor();
      if (ojoFileName == null)
        throw new Exception("file name can't be null");
      else if (ojoFileName.endsWith(".jar"))
        p.loadPojo(ojoFileName, modelName);
      else if (ojoFileName.endsWith(".zip"))
        p.loadMojo(ojoFileName);
      else
        throw new Exception("unknown model archive type");
      return p.predict2(p.jsonToRowData(jsonArgs));
    }
    catch (Exception e) {
      return "{ \"error\": \"" + e.getMessage() + "\" }";
    }
  }

  public static void main(String[] args) {
    if (args.length != 3) {
      System.out.println("{ \"error\": \"Neeed 3 args have " + args.length +
          ". Usage: jarFile modelName JsonString\" }");
    }
    else
      System.out.println(predict(args[0], args[1], args[2]));
  }

}
