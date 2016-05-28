import java.io.*;
import java.net.MalformedURLException;

import javax.servlet.http.*;
import javax.servlet.*;

import hex.genmodel.easy.prediction.AbstractPrediction;
import hex.genmodel.easy.*;
import hex.genmodel.*;

import com.google.gson.Gson;

public class PredictPythonServlet extends HttpServlet {
  // Set to true for demo mode (to print the predictions to stdout).
  // Set to false to get better throughput.
  static boolean VERBOSE = false;

  public static GenModel rawModel;
  public static EasyPredictModelWrapper model;

//  public static long numberOfPredictions = 0;
//  public static long startTime = System.currentTimeMillis();
//  public static long lastTime = 0;
//  public static double lastPredictionMs = 0;
//  public static double totalTimeMs = 0;
//  public static double totalTimeSquareMs = 0;
//  public static double warmupTimeMs = 0;
//  public static double warmupTimeSquareMs = 0;
//  public static int warmupNumber = 5;

  private static Process p;
  private static ProcessBuilder pb;
  private static OutputStream stdin;
  private static BufferedReader reader, err_reader;

//  public static PredictServlet.Times getPythonTimes = new PredictServlet.Times();
//  public static PredictServlet.Times postPythonTimes = new PredictServlet.Times();

  static Gson gson = new Gson();

  private static File servletPath = null;
  String[] colNames;

  public void init(ServletConfig servletConfig) throws ServletException {
    super.init(servletConfig);
    try {
      servletPath = new File(servletConfig.getServletContext().getResource("/").getPath());
      System.out.println("path = " + servletPath);

      rawModel = new REPLACE_THIS_WITH_PREDICTOR_CLASS_NAME();
      model = new EasyPredictModelWrapper(rawModel);

      if (rawModel == null || model == null)
        throw new ServletException("can't load model");

      colNames = rawModel.getNames();

      String program = servletPath.getAbsolutePath() + "/WEB-INF/lib/python.py";
      if (VERBOSE) System.out.println(program);
      // start the python process
      try {
        // score.py -- --verbose --models-dir /tmp/models
//        pb = new ProcessBuilder("python", program, "--models-dir", "/Users/magnus/Git/steamY/scoring-service-builder/examples/example-spam-detection/models");
//        pb = new ProcessBuilder("python", program, "--datafile", "/Users/magnus/Git/steamY/scoring-service-builder/examples/example-spam-detection/data/smsData.txt");
        pb = new ProcessBuilder("python", program);
//        pb = new ProcessBuilder("python", program);
//        Map<String, String> env = pb.environment();
//        env.put("PYTHONPATH", env.get("PYTHONPATH") + ":/Users/magnus/Git/steamY/scoring-service-builder/examples/example-spam-detection/lib");
//        System.out.println(env);
        p = pb.start();
        stdin = p.getOutputStream();
        InputStream stdout = p.getInputStream();
        InputStream stderr = p.getErrorStream();
        reader = new BufferedReader(new InputStreamReader(stdout));
        err_reader = new BufferedReader(new InputStreamReader(stderr));
        System.out.println("Python started");
      } catch (Exception ex) {
        System.out.println("Python failed");
        ex.printStackTrace();
      }

    }
    catch (MalformedURLException e) {
      e.printStackTrace();
    }
  }

  static private String jsonModel() {
    Gson gson = new Gson();
    String modelJson = gson.toJson(model);
    return modelJson;
  }

//  @SuppressWarnings("unchecked")
//  private void fillRowDataFromHttpRequest(HttpServletRequest request, RowData row) {
//    Map<String, String[]> parameterMap;
//    parameterMap = request.getParameterMap();
//    if (VERBOSE) System.out.println();
//    for (Map.Entry<String, String[]> entry : parameterMap.entrySet()) {
//      String key = entry.getKey();
//      String[] values = entry.getValue();
//      for (String value : values) {
//        if (VERBOSE) System.out.println("Key: " + key + " Value: " + value);
//        if (value.length() > 0) {
//          row.put(key, value);
//        }
//      }
//    }
//  }

  static final byte[] NewlineByteArray = "\n".getBytes();

  public synchronized String sendPython(String queryString) {
    String result = null;

    try {
//      String res = queryString + "\n";
      stdin.write(queryString.getBytes());
      stdin.write(NewlineByteArray);
      stdin.flush();
      result = reader.readLine();
//        showStderr();
    }
    catch (Exception ex) {
      ex.printStackTrace();
      showStderr();
    }
    return result;
  }

  public void doGet(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
    long start = System.nanoTime();
    try {
      if (model == null)
        throw new Exception("No predictor model");

      String queryString = request.getQueryString();
      if (VERBOSE) System.out.println("queryString " + queryString);

//      String result = null;
//      try {
//        String res = queryString + "\n";
//        stdin.write(res.getBytes());
//        stdin.flush();
//        result = reader.readLine();
////        showStderr();
//      }
//      catch (Exception ex) {
//        ex.printStackTrace();
//        showStderr();
//      }

      String result = sendPython(queryString.replaceAll("%20", " "));
      if (VERBOSE) System.out.println("result " + result);

//      // should now be in CSV from python
//      RowData row = csvToRowData(colNames, result);

      RowData row = sparseToRowData(colNames, result);


      if (VERBOSE) System.out.println("row " + row);

      AbstractPrediction pr = PredictServlet.predict(row);

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
    PredictServlet.getPythonTimes.add(start, done);
    if (VERBOSE) System.out.println("Python Get time " + PredictServlet.getPythonTimes);
  }

  private void showStderr() {
    String line;
    try {
      while ((line=err_reader.readLine())!=null) {
        System.err.println(line);
      }
    } catch (Exception ex2) {
      ex2.printStackTrace();
    }
  }

  private RowData sparseToRowData(String[] colNames, String result) throws Exception {
//    System.out.println("result in sp " + result);
    RowData row = new RowData();
    String[] pairs = result.split(" ");
    for (String p : pairs) {
//      System.out.println(" pair " + p);
      String[] a = p.split(":");
//      System.out.println(a[0] + " = " + a[1]);
      int index = Integer.parseInt(a[0]);
//      if (index < 0 || index > colNames.length)
//        throw new Exception("index out of range " + index);
      double value = Float.parseFloat(a[1]);
//      System.out.println(index + " = " + value);
      row.put(colNames[index], value);
    }
    return row;
  }

  private RowData csvToRowData(String[] colNames, String result) throws Exception {
    String[] vals = result.split(",");
    if (colNames.length != vals.length)
      throw new Exception("CSV fields not same length " + vals.length + " as model expects " + colNames.length);

    RowData row = new RowData();
    for (int i = 0; i < vals.length; i++) {
      String v = vals[i];
      if (v != null && v.length() > 0) {
        row.put(colNames[i], v);
      }
    }
    return row;
  }

  private void setToNaN(double[] arr) {
    for (int i = 0; i < arr.length; i++) {
      arr[i] = Double.NaN;
    }
  }

//  private AbstractPrediction predict(RowData row) throws PredictException {
//    long start = System.nanoTime();
//    AbstractPrediction p = model.predict(row);
//    long done = System.nanoTime();
//    lastTime = System.currentTimeMillis();
//    double elapsedMs = (done - start) / 1.0e6;
//    lastPredictionMs = elapsedMs;
//    totalTimeMs += elapsedMs;
//    totalTimeSquareMs += elapsedMs * elapsedMs;
//    numberOfPredictions += 1;
//    if (numberOfPredictions <= warmupNumber) {
//      warmupTimeMs += elapsedMs;
//      warmupTimeSquareMs += elapsedMs * elapsedMs;
//    }
//    return p;
//  }

  public void doPost(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
    long start = System.nanoTime();
    try {
      if (model == null)
        throw new Exception("No predictor model");

      BufferedReader r = request.getReader();
      StringBuilder sb = new StringBuilder();
      String line = r.readLine();
      r.close();
//      System.out.println("line " + line);
      String result = sendPython(line);
      if (VERBOSE) System.out.println("result " + result);

//      // should now be in CSV from python
//      RowData row = csvToRowData(colNames, result);

      RowData row = sparseToRowData(colNames, result);
//      System.out.println("row " + row);
      // do the prediction
      AbstractPrediction pr = PredictServlet.predict(row);

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
    PredictServlet.postPythonTimes.add(start, done);
    if (VERBOSE) System.out.println("Python Get time " + PredictServlet.postPythonTimes);

  }


}

