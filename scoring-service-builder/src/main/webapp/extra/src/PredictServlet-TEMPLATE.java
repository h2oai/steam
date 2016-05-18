import java.io.*;
import java.util.Map;
import java.util.HashMap;
import java.util.Arrays;
import java.lang.reflect.Type;

import javax.servlet.http.*;
import javax.servlet.*;

import hex.genmodel.easy.prediction.AbstractPrediction;
import hex.genmodel.easy.exception.PredictException;
import hex.genmodel.easy.*;
import hex.genmodel.*;

import com.google.gson.Gson;

public class PredictServlet extends HttpServlet {
  // Set to true for demo mode (to print the predictions to stdout).
  // Set to false to get better throughput.
  static boolean VERBOSE = false;

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

  static {
    GenModel rawModel = new REPLACE_THIS_WITH_PREDICTOR_CLASS_NAME();
    model = new EasyPredictModelWrapper(rawModel);
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
  }


}

