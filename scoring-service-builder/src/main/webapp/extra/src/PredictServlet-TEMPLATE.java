import java.io.*;
import java.lang.reflect.Type;
import java.util.HashMap;
import java.util.Map;

import javax.servlet.http.*;
import javax.servlet.*;

import com.google.gson.reflect.TypeToken;
import hex.genmodel.easy.prediction.AbstractPrediction;
import hex.genmodel.easy.exception.PredictException;
import hex.genmodel.easy.*;
import hex.genmodel.*;

import com.google.gson.Gson;

public class PredictServlet extends HttpServlet {

  static boolean VERBOSE = false;
  public static int warmUpCount = 5;

  public static Gson gson = new Gson();
  public static Type mapType = new TypeToken<HashMap<String, Object>>(){}.getType();

  public static class Times {
    public long count = 0;
    public double totalTimeMs = 0;
    public double totalTimeSquaredMs = 0;
    public double warmupTimeMs = 0;
    public double warmupTimeSquaredMs = 0;
    public double lastMs = 0;

    public void add(long startNs, long endNs) {
      double elapsed = (endNs - startNs) / 1.0e6;
      add(elapsed);
    }

    public synchronized void add(double timeMs) {
      count += 1;
      totalTimeMs += timeMs;
      double tt = timeMs * timeMs;
      totalTimeSquaredMs += tt;
      if (count <= warmUpCount) {
        warmupTimeMs += timeMs;
        warmupTimeSquaredMs += tt;
      }
      lastMs = timeMs;
    }

    public double avg() {
      return count > 0 ? totalTimeMs / count : 0.0;
    }

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
      return PredictServlet.gson.fromJson(gson.toJson(this), mapType);
    }

    public String toString() {
      return String.format("n %d  last %.3f  avg %.3f after warmup %.3f [ms]", count, lastMs, avg(), avgAfterWarmup());
    }
  }

  static {
    GenModel rawModel = new REPLACE_THIS_WITH_PREDICTOR_CLASS_NAME();
    model = new EasyPredictModelWrapper(rawModel);
  }


  public static EasyPredictModelWrapper model;
  public static long startTime = System.currentTimeMillis();
  public static long lastTime = 0;
  public static Times predictionTimes = new Times();
  public static Times getTimes = new Times();
  public static Times postTimes = new Times();

  static private String jsonModel() {

    String modelJson = gson.toJson(model);
    return modelJson;
  }

  @SuppressWarnings("unchecked")
  private void fillRowDataFromHttpRequest(HttpServletRequest request, RowData row) {
    Map<String, String[]> parameterMap;
    parameterMap = request.getParameterMap();
    if (VERBOSE) System.out.println();
    for (Map.Entry<String, String[]> entry : parameterMap.entrySet()) {
      String key = entry.getKey();
      String[] values = entry.getValue();
      for (String value : values) {
        if (VERBOSE) System.out.println("Key: " + key + " Value: " + value);
        if (value.length() > 0) {
          row.put(key, value);
        }
      }
    }
  }

  public void doGet(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
    long start = System.nanoTime();
    RowData row = new RowData();
    fillRowDataFromHttpRequest(request, row);
    try {
      if (model == null)
        throw new Exception("No predictor model");

      // we have a model loaded, do the prediction
      AbstractPrediction pr = predict(row);

      // assemble json result
      Gson gson = new Gson();
      String prJson = gson.toJson(pr);

      response.getWriter().write(prJson);
      response.setStatus(HttpServletResponse.SC_OK);

    }
    catch (Exception e) {
      // Prediction failed.
      System.out.println(e.getMessage());
      response.sendError(HttpServletResponse.SC_NOT_ACCEPTABLE, e.getMessage());
    }
    long done = System.nanoTime();
    getTimes.add(start, done);
    if (VERBOSE) System.out.println("Get time " + getTimes);
  }

  private AbstractPrediction predict(RowData row) throws PredictException {
    long start = System.nanoTime();
    AbstractPrediction p = model.predict(row);
    long done = System.nanoTime();
    lastTime = System.currentTimeMillis();
    predictionTimes.add(start, done);
    if (VERBOSE) System.out.println("Prediction time " + predictionTimes);
    return p;
  }

  public void doPost(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
    long start = System.nanoTime();
    try {
      if (model == null)
        throw new Exception("No predictor model");

      Gson gson = new Gson();
      RowData row = gson.fromJson(request.getReader(), new RowData().getClass());

      // do the prediction
      AbstractPrediction pr = model.predict(row);

      // assemble json result
      String prJson = gson.toJson(pr);

      // Emit the prediction to the servlet response.
      response.getWriter().write(prJson);
      response.setStatus(HttpServletResponse.SC_OK);
    }
    catch (Exception e) {
      // Prediction failed.
      System.out.println(e.getMessage());
      response.sendError(HttpServletResponse.SC_NOT_ACCEPTABLE, e.getMessage());
    }
    long done = System.nanoTime();
    postTimes.add(start, done);
    if (VERBOSE) System.out.println("Post time " + postTimes);
  }

}

