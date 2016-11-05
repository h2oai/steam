import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import com.google.gson.reflect.TypeToken;
import hex.genmodel.*;
import hex.genmodel.easy.*;
import hex.genmodel.easy.exception.PredictException;
import hex.genmodel.easy.prediction.AbstractPrediction;
import org.apache.commons.io.FileUtils;
import org.slf4j.Logger;

import javax.servlet.ServletConfig;
import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import java.io.File;
import java.io.IOException;
import java.lang.reflect.Type;
import java.net.MalformedURLException;
import java.util.HashMap;
import java.util.List;
import java.util.Map;


class ServletUtil {
  private final static Logger logger = Logging.getLogger(ServletUtil.class);

  public static String pythonEnvironmentFile = "REPLACE_WITH_PYTHON_ENVIRONMENT_FILE";

  private static List<String> modelNames = null;

  public static synchronized void loadModels(File servletPath) {
    if (modelNames == null) {
      try {
        modelNames = FileUtils.readLines(new File(servletPath, "modelnames.txt"));
        logger.info("modelNames size {}", modelNames.size());
        models = new HashMap<String, EasyPredictModelWrapper>();
        EasyPredictModelWrapper mod = null;
        for (String m : modelNames) {
          if (m.endsWith(".java"))
            mod = addPojoModel(m.replace(".java", ""));
          else if (m.endsWith(".zip"))
            mod = addMojoModel(m.replace(".zip", ""), servletPath);
          if (mod == null)
            throw new Exception("Failed to load model "  + m);
          if (model == null)
            model = mod;
        }
        logger.info("added {} models", models.size());
      }
      catch (Exception e) {
        logger.error("can't load model using modelnames.txt", e);
      }
    }
    else
      logger.debug("models already loaded");
  }

  static GenModel rawModel = null;
  static String modelName = null;
  public static EasyPredictModelWrapper model = null;
  public static Map<String, EasyPredictModelWrapper> models = null;

  // this is how to do easy predict that doesn't error on wrong categorical levels
  //    rawModel = new
  //        new EasyPredictModelWrapper(
  //            new EasyPredictModelWrapper.Config().setModel(rawModel).setConvertUnknownCategoricalLevelsToNa(true)
  //        );

  static EasyPredictModelWrapper addPojoModel(String modelName) {
    EasyPredictModelWrapper model = null;
    try {
      Class<?> clazz = Class.forName(modelName);
      GenModel rawModel = (GenModel) clazz.newInstance();
      model = new EasyPredictModelWrapper(rawModel);
      models.put(modelName, model);
      logger.debug("added model {}  new size {}", modelName, models.size());

    }
    catch (Exception e) {
      logger.error("error {}", e);
    }

    logger.info("loaded {} models", models.size());
    return model;
  }

  static EasyPredictModelWrapper addMojoModel(String modelName, File servletPath) {
    EasyPredictModelWrapper model = null;
    try {
      String fileName = servletPath + File.separator + modelName + ".zip";
      GenModel rawModel = REPLACE_THIS_WITH_MODEL;
      model = new EasyPredictModelWrapper(rawModel);
      models.put(modelName, model);
      logger.info("added model {}  new size {}", modelName, models.size());
    }
    catch (Exception e) {
      logger.error("error {}", e);
    }
    return model;
  }

  // load preprocessing
  public static Transform transform = REPLACE_THIS_WITH_TRANSFORMER_OBJECT;

  public static final Type MAP_TYPE = new TypeToken<HashMap<String, Object>>(){}.getType();
  public static final Type ROW_DATA_TYPE = new TypeToken<RowData>(){}.getType();
  public static final int warmUpCount = 5;

  private static Gson gson = new GsonBuilder().serializeSpecialFloatingPointValues().create();

  public static class Times {
    private long count = 0;
    private double totalTimeMs = 0;
    private double totalTimeSquaredMs = 0;
    private double warmupTimeMs = 0;
    private double warmupTimeSquaredMs = 0;
    private double lastMs = 0;

    public void add(long startNs, long endNs, int n) {
      double elapsed = (endNs - startNs) / 1.0e6;
      add(elapsed, n);
    }

    public void add(long startNs, long endNs) {
      add(startNs, endNs, 1);
    }

    public synchronized void add(double timeMs, int n) {
      count += n;
      totalTimeMs += timeMs; // n * timeMs/n
      double tt = timeMs * timeMs / n; // n * (timeMs/n)^2
      totalTimeSquaredMs += tt;
      if (count <= warmUpCount) {
        warmupTimeMs += timeMs;
        warmupTimeSquaredMs += tt;
      }
      lastMs = timeMs / n;
    }

    public double avg() {
      return count > 0 ? totalTimeMs / count : 0.0;
    }

//    public double sdev() { }

    public double avgAfterWarmup() {
      return count > warmUpCount ? (totalTimeMs - warmupTimeMs) / (count - warmUpCount) : 0.0;
    }

    public String toJson() {
      return gson.toJson(toMap());
    }

    public Map<String, Object> toMap() {
      Map<String, Object> map = classToMap();
      map.put("averageTime", avg());
      map.put("averageAfterWarmupTime", avgAfterWarmup());
      return map;
    }

    private Map<String, Object> classToMap() {
      return gson.fromJson(gson.toJson(this), MAP_TYPE);
    }

    public String toString() {
      return String.format("n %d  last %.3f  avg %.3f after warmup %.3f [ms]", count, lastMs, avg(), avgAfterWarmup());
    }
  }

  public static synchronized AbstractPrediction predict(RowData row) throws PredictException {
    long start = System.nanoTime();
    AbstractPrediction pr = model.predict(row);
    long done = System.nanoTime();
    ServletUtil.lastTime = System.currentTimeMillis();
    ServletUtil.predictionTimes.add(start, done);

//    String label = null;
//    if (pr instanceof BinomialModelPrediction) {
//      label = ((BinomialModelPrediction) pr).label;
//    } else if (pr instanceof MultinomialModelPrediction) {
//      label = ((MultinomialModelPrediction) pr).label;
//    }
//    if (label != null) {
//      ServletUtil.incrementOutputLabel(label);
//    }

    logger.debug("Prediction time {}", ServletUtil.predictionTimes);
    return pr;
  }

  public static synchronized AbstractPrediction predictModel(String modelName, RowData row) throws PredictException {
    long start = System.nanoTime();
    if (models.size() == 0)
      throw new PredictException("no models");
    EasyPredictModelWrapper mod = models.get(modelName);
    if (mod == null)
      throw new PredictException("unknown model " + modelName);
    AbstractPrediction pr = mod.predict(row);
    long done = System.nanoTime();
    ServletUtil.lastTime = System.currentTimeMillis();
    ServletUtil.predictionTimes.add(start, done);

//    String label = null;
//    if (pr instanceof BinomialModelPrediction) {
//      label = ((BinomialModelPrediction) pr).label;
//    } else if (pr instanceof MultinomialModelPrediction) {
//      label = ((MultinomialModelPrediction) pr).label;
//    }
//    if (label != null) {
//      ServletUtil.incrementOutputLabel(label);
//    }

    logger.debug("Prediction time {}", ServletUtil.predictionTimes);
    return pr;
  }



//  public static Map<String, Integer> outputLabels = new ConcurrentHashMap<String, Integer>();
//
//  public static synchronized void incrementOutputLabel(String key) {
//    Integer value = outputLabels.putIfAbsent(key, 1);
//    if (value != null)
//      outputLabels.replace(key, value + 1);
//  }

  public static long startTime = System.currentTimeMillis();
  public static long lastTime = 0;
  public static Times predictionTimes = new Times();
  public static Times getTimes = new Times();
  public static Times postTimes = new Times();
  public static Times getPythonTimes = new Times();
  public static Times postPythonTimes = new Times();

}