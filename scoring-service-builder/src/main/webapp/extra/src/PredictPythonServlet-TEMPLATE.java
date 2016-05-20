import java.io.*;
import java.net.MalformedURLException;
import java.util.Map;
import java.util.HashMap;
import java.util.Arrays;
import java.lang.reflect.Type;

import javax.servlet.http.*;
import javax.servlet.*;

import hex.ModelCategory;
import hex.genmodel.easy.prediction.AbstractPrediction;
import hex.genmodel.easy.exception.PredictException;
import hex.genmodel.easy.*;
import hex.genmodel.*;

import com.google.gson.Gson;

public class PredictPythonServlet extends HttpServlet {
  // Set to true for demo mode (to print the predictions to stdout).
  // Set to false to get better throughput.
  static boolean VERBOSE = false;

  public static GenModel rawModel;
  public static EasyPredictModelWrapper model;
  public static long numberOfPredictions = 0;
  public static long startTime = System.currentTimeMillis();
  public static long lastTime = 0;
  public static double lastPredictionMs = 0;
  public static double totalTimeMs = 0;
  public static double totalTimeSquareMs = 0;
  public static double warmupTimeMs = 0;
  public static double warmupTimeSquareMs = 0;
  public static int warmupNumber = 5;

  static Process p;
  static ProcessBuilder pb;
  static OutputStream stdin;
  static BufferedReader reader, err_reader;

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
        pb = new ProcessBuilder("python", program);
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
    try {
      if (model == null)
        throw new Exception("No predictor model");

      String queryString = request.getQueryString();
      if (VERBOSE) System.out.println("queryString " + queryString);

      String result = null;
      try {
        String res = queryString + "\n";
        stdin.write(res.getBytes());
        stdin.flush();
        result = reader.readLine();
      }
      catch (Exception ex) {
        ex.printStackTrace();
        String line;
        try {
          while ((line=err_reader.readLine())!=null) {
            System.out.println(line);
          }
        } catch (Exception ex2) {
          ex2.printStackTrace();
        }
      }
      if (VERBOSE) System.out.println("result " + result); // should now be in CSV from python

      RowData row = csvToRowData(colNames, result);
      if (VERBOSE) System.out.println("row " + row);

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

  private AbstractPrediction predict(RowData row) throws PredictException {
    long start = System.nanoTime();
    AbstractPrediction p = model.predict(row);
    long done = System.nanoTime();
    lastTime = System.currentTimeMillis();
    double elapsedMs = (done - start) / 1.0e6;
    lastPredictionMs = elapsedMs;
    totalTimeMs += elapsedMs;
    totalTimeSquareMs += elapsedMs * elapsedMs;
    numberOfPredictions += 1;
    if (numberOfPredictions <= warmupNumber) {
      warmupTimeMs += elapsedMs;
      warmupTimeSquareMs += elapsedMs * elapsedMs;
    }
    return p;
  }

  public void doPost(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
    try {
      if (model == null)
        throw new Exception("No predictor model");

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
  }


}

