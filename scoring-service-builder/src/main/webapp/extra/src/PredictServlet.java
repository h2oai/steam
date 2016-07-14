import java.io.*;
import java.net.MalformedURLException;
import java.util.Map;
import java.util.HashMap;
import java.lang.reflect.Type;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
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
  private static final Logger logger = LoggerFactory.getLogger("PredictServlet");

  private static final Class ROW_DATA_TYPE = new RowData().getClass();

  private static final Gson gson = new GsonBuilder().serializeSpecialFloatingPointValues().create();

  private static GenModel rawModel = ServletUtil.rawModel;
  private static EasyPredictModelWrapper model = ServletUtil.model;
  private static Transform transform = ServletUtil.transform;

  private File servletPath = null;

  public void init(ServletConfig servletConfig) throws ServletException {
    super.init(servletConfig);
    try {
      servletPath = new File(servletConfig.getServletContext().getResource("/").getPath());
      logger.debug("servletPath {}", servletPath);

     }
    catch (MalformedURLException e) {
      logger.error("init failed", e);
    }
  }

  @SuppressWarnings("unchecked")
  private void fillRowDataFromHttpRequest(HttpServletRequest request, RowData row) {
    Map<String, String[]> parameterMap;
    parameterMap = request.getParameterMap();
    for (Map.Entry<String, String[]> entry : parameterMap.entrySet()) {
      String key = entry.getKey();
      String[] values = entry.getValue();
      for (String value : values) {
        logger.debug("Key: {}  Value: {}", key, value);
        if (value.length() > 0) {
          row.put(key, value);
        }
      }
    }
  }

  public void doGet(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
    long start = System.nanoTime();
    RowData row = new RowData();
    if (transform == null) { // no jar transformation
      fillRowDataFromHttpRequest(request, row);
    }
    else {
      // transform with the jar
      String req = request.getQueryString().replaceAll("%20", " ");
      //System.out.println(req);
      byte[] bytes = req.getBytes();
      Map<String, Object> tr = transform.fit(new String(bytes));
      for (String k : tr.keySet()) {
        logger.debug("{} = {}", k, tr.get(k));
        row.put(k, tr.get(k));
      }
      System.out.println(row);
    }
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
      logger.error("get failed", e);
      response.sendError(HttpServletResponse.SC_NOT_ACCEPTABLE, e.getMessage());
    }
    long done = System.nanoTime();
    ServletUtil.getTimes.add(start, done);
    logger.debug("Get time {}", ServletUtil.getTimes);
  }

  public static synchronized AbstractPrediction predict(RowData row) throws PredictException {
    long start = System.nanoTime();
    AbstractPrediction p = model.predict(row);
    long done = System.nanoTime();
    ServletUtil.lastTime = System.currentTimeMillis();
    ServletUtil.predictionTimes.add(start, done);

    logger.debug("Prediction time {}", ServletUtil.predictionTimes);
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
        logger.debug("line {}", line);
        if (transform == null) { // no jar transformation
          row = gson.fromJson(line, ServletUtil.ROW_DATA_TYPE);
        }
        else {
          row = new RowData();
          Map<String, Object> tr = transform.fit(line);
          for (String k : tr.keySet()) {
            logger.debug("{} = {}", k, tr.get(k));
            row.put(k, tr.get(k));
          }
          System.out.println(row);
        }
        logger.debug("row {}", row);
        if (row != null) {
          // do the prediction
          pr = predict(row);

          // assemble json result
          prJson = gson.toJson(pr);
          logger.debug(prJson);

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
      logger.error("post failed", e);
      response.sendError(HttpServletResponse.SC_NOT_ACCEPTABLE, e.getMessage());
    }
    long done = System.nanoTime();
    ServletUtil.postTimes.add(start, done, count);
    logger.debug("Post time {}", ServletUtil.postTimes);
  }

}

