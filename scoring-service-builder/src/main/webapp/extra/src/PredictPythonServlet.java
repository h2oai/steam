import java.io.*;
import java.net.MalformedURLException;
import javax.servlet.http.*;
import javax.servlet.*;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import com.google.gson.GsonBuilder;
import hex.genmodel.easy.exception.PredictException;
import hex.genmodel.easy.prediction.AbstractPrediction;
import hex.genmodel.easy.prediction.BinomialModelPrediction;
import hex.genmodel.easy.prediction.MultinomialModelPrediction;
import hex.genmodel.easy.*;
import hex.genmodel.*;

import com.google.gson.Gson;

public class PredictPythonServlet extends HttpServlet {
  private final Logger logger = Logging.getLogger(this.getClass());

  private static GenModel rawModel = ServletUtil.rawModel;
  private static EasyPredictModelWrapper model = ServletUtil.model;

  private static Process p = null;
  private static ProcessBuilder pb = null;
  private static OutputStream stdin;
  private static BufferedReader reader, err_reader;

  private static final Gson gson = new GsonBuilder().serializeSpecialFloatingPointValues().create();

  private static File servletPath = null;
  String[] colNames;

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

  private void startPython() throws Exception {
    logger.debug("startPython");
    String program = servletPath.getAbsolutePath() + "/WEB-INF/python.py";
    logger.debug("program {}", program);
    // start the python process
    try {
      // score.py
      pb = new ProcessBuilder("python", program);
      File pythonProcessDir = new File(servletPath, "/WEB-INF");
      pb.directory(pythonProcessDir);
      p = pb.start();
      stdin = p.getOutputStream();
      InputStream stdout = p.getInputStream();
      InputStream stderr = p.getErrorStream();
      reader = new BufferedReader(new InputStreamReader(stdout));
      err_reader = new BufferedReader(new InputStreamReader(stderr));
      logger.info("Python started");
    } catch (Exception ex) {
      logger.error("Python failed", ex);
      throw new Exception("Python failed");
    }
  }

  public void destroy() {
    if (p != null) {
      p.destroy();
      logger.info("Python destroyed");
    }
    super.destroy();
  }

  static private String jsonModel() {
    Gson gson = new Gson();
    String modelJson = gson.toJson(model);
    return modelJson;
  }

  private static final byte[] NewlineByteArray = "\n".getBytes();

  private synchronized String sendPython(String queryString) {
    String result = null;

    try {
      // restart if python failed
      if (p == null) {
        startPython();

      }
      try {
        logger.debug("send {}", queryString);
        stdin.write(queryString.getBytes());
        stdin.write(NewlineByteArray);
        stdin.flush();
        result = reader.readLine();
        logger.debug("result {}", result);
      }
      catch (IOException e) {
        String msg = "ERROR IOException in sendPython restarting python";
        result = msg;
        logger.error(msg, e);
        showStderr();
        // it failed so we restart it and retry
//        if (p != null) p.destroy();
//        startPython();
//        stdin.write(queryString.getBytes());
//        stdin.write(NewlineByteArray);
//        stdin.flush();
//        result = reader.readLine();
      }
//        showStderr();
    }
    catch (Exception ex) {
      logger.error("sendPython failed", ex);
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
      logger.debug("queryString {}", queryString);

      String result = sendPython(queryString.replaceAll("%20", " "));
      logger.debug("result {}", result);

      // should now be in sparse format from python
      RowData row = strMapToRowData(result);
      logger.debug("row: {}", row);

      AbstractPrediction pr = ServletUtil.predict(row);
      logger.debug("pr: {}", pr);

      // assemble json result
      String prJson = gson.toJson(pr);

      response.getWriter().write(prJson);
      response.setStatus(HttpServletResponse.SC_OK);

    }
    catch (Exception e) {
      // Prediction failed.
      logger.error("doGet failed", e);
      response.sendError(HttpServletResponse.SC_NOT_ACCEPTABLE, e.getMessage());
    }
    long done = System.nanoTime();
    ServletUtil.getPythonTimes.add(start, done);
    logger.debug("Python Get time {}", ServletUtil.getPythonTimes);
  }

  private void showStderr() {
    String line;
    try {
      while ((line=err_reader.readLine())!=null) {
        logger.info(line);
      }
    } catch (Exception ex2) {
      logger.error("showStderr failed", ex2);
    }
  }

  private RowData sparseToRowData(String[] colNames, String result) throws Exception {
    RowData row = new RowData();
    if (result == null || result.length() == 0)
      return row;
    String[] pairs = result.split(" ");
    try {
      for (String p : pairs) {
        String[] a = p.split(":");
        int index = Integer.parseInt(a[0]);
        double value = Float.parseFloat(a[1]);
        row.put(colNames[index], value);
      }
    }
    catch (NumberFormatException e) {
      logger.error("Failed to parse {}", result);
    }
    return row;
  }

  private RowData strMapToRowData(String result) throws Exception {
    RowData row = new RowData();
    if (result == null || result.length() == 0)
      return row;
    String[] pairs = result.split(" ");
    try {
      for (String p : pairs) {
        String[] a = p.split(":");
        String term = a[0];
        double value = Float.parseFloat(a[1]);
        row.put(term, value);
      }
    }
    catch (NumberFormatException e) {
      logger.error("Failed to parse {}", result);
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
      String result;
      String endingMultipartBoundary = null;
      String batchFileName = null;
      final String filenameString = "filename=";
      final int filenameStringLength = filenameString.length();
      boolean outputResult = false;
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
        logger.debug("line after batch optional code {}", line);

        result = sendPython(line);
        logger.debug("from python: {}", result);
        if (result == null)
          result = "ERROR null result from python";
        if (result.startsWith("ERROR"))
          throw new Exception(result);

        // should now be in sparse format from python
        row = strMapToRowData(result);
        logger.debug("row: {}", row);

        // do the prediction
        pr = ServletUtil.predict(row);

        // assemble json result
        prJson = gson.toJson(pr);
        logger.debug("prJson: {}", prJson);

        // Emit the prediction to the servlet response.
        writer.write(prJson);
        writer.write('\n');
        count += 1;
        outputResult = true;
      }
      if (outputResult)
        response.setStatus(HttpServletResponse.SC_OK);
      else
        response.sendError(HttpServletResponse.SC_NOT_ACCEPTABLE, "Empty input to POST");
    }
    catch (Exception e) {
      // Prediction failed.
      logger.error("doPost failed", e);
      response.sendError(HttpServletResponse.SC_NOT_ACCEPTABLE, e.getMessage());
    }
    if (count > 0) {
      long done = System.nanoTime();
      ServletUtil.postPythonTimes.add(start, done, count);
      logger.debug("Python Post time {}", ServletUtil.postPythonTimes);
    }
  }

}
