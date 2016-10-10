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
import hex.genmodel.easy.prediction.BinomialModelPrediction;
import hex.genmodel.easy.prediction.MultinomialModelPrediction;
import hex.genmodel.easy.exception.PredictException;
import hex.genmodel.easy.*;
import hex.genmodel.*;

import com.google.gson.Gson;

public class PredictServlet extends HttpServlet {
  private final Logger logger = Logging.getLogger(this.getClass());

  private static final Class ROW_DATA_TYPE = new RowData().getClass();

  private static final Gson gson = new GsonBuilder().serializeSpecialFloatingPointValues().create();

  private static GenModel rawModel = null;
  private static EasyPredictModelWrapper model = null;
  private static Transform transform = ServletUtil.transform;

  private File servletPath = null;

  public void init(ServletConfig servletConfig) throws ServletException {
    super.init(servletConfig);
    try {
      servletPath = new File(servletConfig.getServletContext().getResource("/").getPath());
      logger.debug("servletPath {}", servletPath);
      ServletUtil.loadModels(servletPath);
      model = ServletUtil.model;
      logger.debug("model {}", model);
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
    String pathInfo = request.getPathInfo();
    logger.debug("pathInfo {}", pathInfo);
    String modelName = null;
    if (pathInfo != null) {
      modelName = pathInfo.replace("/", "");
    }
    if (transform == null) { // no jar transformation
      logger.debug("no transform");
      fillRowDataFromHttpRequest(request, row);
    }
    else {
      logger.debug("transform {}", transform);
      // transform with the jar
      String req = request.getQueryString().replaceAll("%20", " ");
      //System.out.println(req);
      byte[] bytes = req.getBytes();
      logger.debug("req bytes {}", bytes);
      Map<String, Object> tr = transform.fit(new String(bytes));
      for (String k : tr.keySet()) {
        logger.debug("{} = {}", k, tr.get(k));
        row.put(k, tr.get(k));
      }
    }
    try {
      if (model == null)
        throw new Exception("No predictor model");

      // we have a model loaded, do the prediction
      AbstractPrediction pr = null;
      if (modelName == null)
        pr = ServletUtil.predict(row);
      else
        pr = ServletUtil.predictModel(modelName, row);

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
      String endingMultipartBoundary = null;
      String batchFileName = null;
      final String filenameString = "filename=";
      final int filenameStringLength = filenameString.length();
      while (r.ready()) {
        line = r.readLine();
        logger.debug("line {}", line);

        // Check if we're receving a file form the UI which has to be encoded multipart/form
        // if not, we're just receiving lines of text, each of which is json
        if (endingMultipartBoundary != null && endingMultipartBoundary.equals(line)) {
          logger.debug("ending multipart form, line {}", line);
          break;
        }
        else if (line.startsWith("--")) {
          endingMultipartBoundary = line + "--";
          logger.debug("starting multipart form, line {}", line);
          // skip Content-disposition line
          if (r.ready()) {
            line = r.readLine();
            logger.debug("skipped line {}", line);
          }
          // extract the file name if we can
          if (line.contains(filenameString)) {
            int p = line.indexOf(filenameString);
            logger.debug("p = {}", p);
            if (p != -1) {
              int p1 = line.indexOf("\"", p + filenameStringLength + 1);
              logger.debug("p1 = {}", p1);
              if (p1 != -1) {
                batchFileName = line.substring(p + filenameStringLength + 1,  p1);
                logger.debug("batch file name {}", batchFileName);
              }
            }
          }
          else
            batchFileName = "noname";
          // skip Content-Type
          if (r.ready()) {
            line = r.readLine();
            logger.debug("skipped line {}", line);
          }
          // Set Content-disposition to download a file and use the file name + text
          response.setHeader("Content-disposition", "attachment; filename=" + batchFileName + "_prediction-results.txt");
          continue;
        }

        if (transform == null) { // no jar transformation
          logger.debug("no transformation of input data");
          row = gson.fromJson(line, ServletUtil.ROW_DATA_TYPE);
        }
        else {
          logger.debug("transformation of input data");
          row = new RowData();
          Map<String, Object> tr = transform.fit(line);
          for (String k : tr.keySet()) {
            logger.debug("{} = {}", k, tr.get(k));
            row.put(k, tr.get(k));
          }
        }
        logger.debug("row {}", row);
        if (row != null) {
          // do the prediction
          pr = ServletUtil.predict(row);

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

