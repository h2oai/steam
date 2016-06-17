import java.io.*;
import java.net.MalformedURLException;
import java.util.Map;
import java.util.HashMap;
import java.lang.reflect.Type;

import javax.servlet.http.*;
import javax.servlet.*;

import com.google.gson.reflect.TypeToken;
import hex.genmodel.easy.prediction.AbstractPrediction;
import hex.genmodel.easy.exception.PredictException;
import hex.genmodel.easy.*;
import hex.genmodel.*;

import com.google.gson.Gson;

public class PredictServlet extends HttpServlet {
  private static final boolean VERBOSE = false;
  public static final int warmUpCount = 5;
  private static final Class ROW_DATA_TYPE = new RowData().getClass();

  public static final Gson gson = new Gson();
  public static final Type mapType = new TypeToken<HashMap<String, Object>>(){}.getType();
  public static final Type rowDataType = new TypeToken<RowData>(){}.getType();

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

//    public double sdev() {
//
//    }

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

  public static class Statistics {
    // https://en.wikipedia.org/wiki/Algorithms_for_calculating_variance
    // def online_variance(data):
    //          n = 0
    //          mean = 0.0
    //          M2 = 0.0
    //
    //          for x in data:
    //          n += 1
    //          delta = x - mean
    //          mean += delta/n
    //          M2 += delta*(x - mean)
    //
    //          if n < 2:
    //          return float('nan')
    //          else:
    //          return M2 / (n - 1)
    public long n = 0;
    public double mean = 0.0;
    public double M2 = 0.0;

    public synchronized void add(Statistics s) {
      double delta = s.mean - mean; // mean_b - mean_a
      long n_a = n;
      n += s.n; // n_x = n_a + n_b
//      mean += delta * s.n / n; // mean_x = mean_a + delta * n_b / n_x
//      M2 += s.M2 + delta * delta * n_a * s.n / n; // M2,x = M2,a + M2,b * delta^2 * n_a * n_b / n_x
      double delta_mean = delta * s.n / n;
      mean += delta_mean;
      M2 += s.M2 + delta * delta_mean * n_a;
    }

    public synchronized void add(double x) {
      n += 1;
      double delta = x - mean;
      mean += delta / n;
      M2 += delta * (x - mean);
    }

    public void add(double x, int count) {
      for (int i = 0; i < count; i ++) {
        add(x);
      }
    }

//      public synchronized void add(double x, int count) {
//      // x is for count items
//      n += count;
//      double delta = x - mean;
//      mean += delta / n;
//      M2 += delta * (x - mean);
//    }

    public double avg() { return mean; }

    public double var() {
      if (n < 2)
        return 0; //Double.NaN;
      else
        return M2 / (n - 1);
    }

    public double sdev() { return Math.sqrt(var()); }
  }

  public static class Times2 {
    private double lastMs = 0.0;
    private Statistics s = new Statistics();
    int warmup = 0;

    public void add(long startNs, long endNs, int n) {
      double elapsed = (endNs - startNs) / 1.0e6;
      add(elapsed / n, n);
    }

    public void add(long startNs, long endNs) {
      add(startNs, endNs, 1);
    }

    public synchronized void add(double timeMs, int n) {
      if (warmup < 5) {
        warmup += 1;
        return;
      }
      double t = timeMs / n;
      for (int i = 0; i < n; i++) {
        s.add(t);
      }
      lastMs = t;
    }

    public long n() {
      if (s.mean > 0)
        return s.n + warmUpCount;
      else
        return 0;
    }

    public double avg() { return s.mean; }

    public double avgAfterWarmup() { return s.mean; }

    public double sdev() {
      double sdev = s.sdev();
      if (Double.isNaN(sdev))
        return 0.0;
      else
        return sdev;
    }

    public String toJson() {
      return gson.toJson(toMap());
    }

    public Map<String, Object> toMap() {
      Map<String, Object> map = new HashMap<String, Object>(); //classToMap();
      map.put("n", s.n);
      map.put("mean", s.mean);
      map.put("M2", s.M2);
      map.put("lastMs", lastMs);
      map.put("sdev", s.sdev());
      map.put("count", s.n);
      map.put("averageMs", s.mean);
      return map;
    }

    private Map<String, Object> classToMap() {
      return PredictServlet.gson.fromJson(gson.toJson(this), mapType);
    }

    public String toString() {
      return String.format("n %d  last %.3f  avg %.3f after warmup %.3f [ms]", n(), lastMs, avg(), avgAfterWarmup());
    }
  }

  static {
    GenModel rawModel = new REPLACE_THIS_WITH_PREDICTOR_CLASS_NAME();
    model = new EasyPredictModelWrapper(rawModel);
  }

  private File servletPath = null;

  public void init(ServletConfig servletConfig) throws ServletException {
    super.init(servletConfig);
    try {
      servletPath = new File(servletConfig.getServletContext().getResource("/").getPath());
      if (VERBOSE) System.out.println("servletPath " + servletPath);
    }
    catch (MalformedURLException e) {
      e.printStackTrace();
    }
  }

  public static EasyPredictModelWrapper model;
  public static long startTime = System.currentTimeMillis();
  public static long lastTime = 0;
  public static Times predictionTimes = new Times();
  public static Times getTimes = new Times();
  public static Times postTimes = new Times();
  public static Times getPythonTimes = new PredictServlet.Times();
  public static Times postPythonTimes = new PredictServlet.Times();

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

  public static synchronized AbstractPrediction predict(RowData row) throws PredictException {
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
    int count = 0;
    try {
      if (model == null)
        throw new Exception("No predictor model");

      BufferedReader r = request.getReader();
      PrintWriter writer = response.getWriter();
      String line;
      RowData row;
      AbstractPrediction pr;
      String prJson;
      while (r.ready()) {
        line = r.readLine();
        if (VERBOSE) System.out.println("line " + line);

        row = gson.fromJson(line,  rowDataType);
        if (VERBOSE) System.out.println("row " + row);
        if (row != null) {
          // do the prediction
          pr = predict(row);

          // assemble json result
          prJson = gson.toJson(pr);
          if (VERBOSE) System.out.println(prJson);

          // Emit the prediction to the servlet response.
          writer.write(prJson);
          writer.write('\n');
          count += 1;
        }
      }
      response.setStatus(HttpServletResponse.SC_OK);
    }
    catch (Exception e) {
      // Prediction failed.
      System.out.println(e.getMessage());
      e.printStackTrace();
      response.sendError(HttpServletResponse.SC_NOT_ACCEPTABLE, e.getMessage());
    }
    long done = System.nanoTime();
    postTimes.add(start, done, count);
    if (VERBOSE) System.out.println("Post time " + postTimes);
  }

}

