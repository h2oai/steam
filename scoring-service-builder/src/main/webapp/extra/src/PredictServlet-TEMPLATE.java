import java.io.*;
import java.net.MalformedURLException;
import java.util.Map;
import java.util.HashMap;
import java.lang.reflect.Type;

import javax.servlet.http.*;
import javax.servlet.*;

import com.google.gson.GsonBuilder;
import com.google.gson.reflect.TypeToken;
import hex.genmodel.easy.prediction.AbstractPrediction;
import hex.genmodel.easy.exception.PredictException;
import hex.genmodel.easy.*;
import hex.genmodel.*;

import com.google.gson.Gson;

public class PredictServlet extends HttpServlet {
  private static final boolean VERBOSE = false;
  private static final Class ROW_DATA_TYPE = new RowData().getClass();

  private static final Gson gson = new GsonBuilder().serializeSpecialFloatingPointValues().create();

  private static GenModel rawModel = new REPLACE_THIS_WITH_PREDICTOR_CLASS_NAME();
  private static EasyPredictModelWrapper model = new EasyPredictModelWrapper(rawModel);
  private static ServletUtil.Transform transform = REPLACE_THIS_WITH_TRANSFORMER_OBJECT;

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
    ServletUtil.getTimes.add(start, done);
    if (VERBOSE) System.out.println("Get time " + ServletUtil.getTimes);
  }

  public static synchronized AbstractPrediction predict(RowData row) throws PredictException {
    long start = System.nanoTime();
    AbstractPrediction p = model.predict(row);
    long done = System.nanoTime();
    ServletUtil.lastTime = System.currentTimeMillis();
    ServletUtil.predictionTimes.add(start, done);

    if (VERBOSE) System.out.println("Prediction time " + ServletUtil.predictionTimes);
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

        row = gson.fromJson(line, ServletUtil.rowDataType);
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
    ServletUtil.postTimes.add(start, done, count);
    if (VERBOSE) System.out.println("Post time " + ServletUtil.postTimes);
  }

}

