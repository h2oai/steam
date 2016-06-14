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
  private static boolean VERBOSE = false;

  public static GenModel rawModel;
  public static EasyPredictModelWrapper model;

  private static Process p = null;
  private static ProcessBuilder pb = null;
  private static OutputStream stdin;
  private static BufferedReader reader, err_reader;

  static Gson gson = new Gson();

  private static File servletPath = null;
  String[] colNames;

  public void init(ServletConfig servletConfig) throws ServletException {
    super.init(servletConfig);
    try {
      servletPath = new File(servletConfig.getServletContext().getResource("/").getPath());
      if (VERBOSE) System.out.println("servletPath " + servletPath);

      rawModel = new REPLACE_THIS_WITH_PREDICTOR_CLASS_NAME();
      model = new EasyPredictModelWrapper(rawModel);

      if (rawModel == null || model == null)
        throw new ServletException("can't load model");

      colNames = rawModel.getNames();

      startPython();
    }
    catch (Exception e) {
      e.printStackTrace();
    }
  }

  private void startPython() throws Exception {
    String program = servletPath.getAbsolutePath() + "/WEB-INF/python.py";
    if (VERBOSE) System.out.println(program);
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
      System.out.println("Python started");
    } catch (Exception ex) {
      System.out.println("Python failed");
      ex.printStackTrace();
      throw new Exception("Python failed");
    }
  }

  public void destroy() {
    if (p != null) {
      p.destroy();
      System.out.println("Python destroyed");
    }
    super.destroy();
  }

  static private String jsonModel() {
    Gson gson = new Gson();
    String modelJson = gson.toJson(model);
    return modelJson;
  }

  static final byte[] NewlineByteArray = "\n".getBytes();

  public synchronized String sendPython(String queryString) {
    String result = null;

    try {
      // restart if python failed
      if (p == null)
        startPython();
//      else if (stdin == null || stdin.|| !reader.ready() || !err_reader.ready()) {
//        p.destroy();
//        startPython();
//      }
      // send to python
      try {
        stdin.write(queryString.getBytes());
        stdin.write(NewlineByteArray);
        stdin.flush();
        result = reader.readLine();
      }
      catch (IOException e) {
        System.out.println("IOException in sendPython restarting python");
        e.printStackTrace();
        showStderr();
        // it failed so we restart it and retry
        if (p != null) p.destroy();
        startPython();
        stdin.write(queryString.getBytes());
        stdin.write(NewlineByteArray);
        stdin.flush();
        result = reader.readLine();
      }
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

      String result = sendPython(queryString.replaceAll("%20", " "));
      if (VERBOSE) System.out.println("result " + result);

      // should now be in sparse format from python
      RowData row = sparseToRowData(colNames, result);
      if (VERBOSE) System.out.println("row: " + row);

      AbstractPrediction pr = PredictServlet.predict(row);
      if (VERBOSE) System.out.println("pr: " + pr);

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
//      throw new Exception("Failed to parse " + result);
      System.out.println("Failed to parse " + result);
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
    try {
      if (model == null)
        throw new Exception("No predictor model");

      System.out.println(request);

      BufferedReader r = request.getReader();
      String line = r.readLine();
      r.close();
      if (VERBOSE) System.out.println("to python: " + line);

      String result = "";
      if (line == null) {
        line = "";
//        throw new Exception("null input to python");
        System.out.println("null input to python");
      }
      else {
        result = sendPython(line);
        if (VERBOSE) System.out.println("from python: " + result);
        if (result == null) {
//        throw new Exception("null result from python");
          System.out.println("null result from python");
        }
      }
      if (result.startsWith("ERROR"))
        throw new Exception(result);

      // should now be in sparse format from python
      RowData row = sparseToRowData(colNames, result);
      if (VERBOSE) System.out.println("row: " + row);

      // do the prediction
      AbstractPrediction pr = PredictServlet.predict(row);

      // assemble json result
      String prJson = gson.toJson(pr);
      if (VERBOSE) System.out.println("prJson: " + prJson);

      // Emit the prediction to the servlet response.
      response.getWriter().write(prJson);
      response.setStatus(HttpServletResponse.SC_OK);
    }
    catch (Exception e) {
      // Prediction failed.
      System.out.println(e.getMessage());
      e.printStackTrace();
      response.sendError(HttpServletResponse.SC_NOT_ACCEPTABLE, e.getMessage());
    }
    long done = System.nanoTime();
    PredictServlet.postPythonTimes.add(start, done);
    if (VERBOSE) System.out.println("Python Get time " + PredictServlet.postPythonTimes);

  }

}
