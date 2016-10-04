import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import hex.genmodel.*;
import hex.genmodel.easy.*;
import hex.genmodel.easy.prediction.AbstractPrediction;
import org.apache.commons.fileupload.FileItem;
import org.apache.commons.fileupload.disk.DiskFileItemFactory;
import org.apache.commons.fileupload.servlet.ServletFileUpload;
import org.apache.commons.io.FileUtils;
import org.apache.commons.io.IOUtils;
import org.slf4j.Logger;

import javax.servlet.ServletConfig;
import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.File;
import java.io.InputStream;
import java.io.IOException;
import java.net.MalformedURLException;
import java.util.List;

public class PredictBinaryServlet extends HttpServlet {
  private final Logger logger = Logging.getLogger(this.getClass());

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

  public void doPost(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
    long start = System.nanoTime();
    int count = 0;
    File tmpDir = null;
    try {
      if (model == null)
        throw new Exception("No predictor model");

      // fill row with parameters, some of which are binary, like images
      RowData row = new RowData();
      List<FileItem> items = new ServletFileUpload(new DiskFileItemFactory()).parseRequest(request);
      for (FileItem i : items) {
        String field = i.getFieldName();
        String value = i.getString();
        logger.debug("field {}  value {}", field, value);
        if (field.startsWith("binary_")) {
          String binFieldName = field.substring("binary_".length());
          if (binFieldName == null || binFieldName.length() == 0)
            throw new Exception("empty binary field name for " + field);
          InputStream inputStream = i.getInputStream();
          if (inputStream == null)
            throw new Exception("null input stream for " + field);
          byte[] bindata = IOUtils.toByteArray(inputStream);
          if (bindata.length == 0)
            throw new Exception("empty binary field value for " + field);
          logger.debug("binary field {} size {}", binFieldName, bindata.length);
          row.put(binFieldName, bindata);
        }
        else if (field.equals("data")) {
          RowData r = gson.fromJson(value, ServletUtil.ROW_DATA_TYPE);
          logger.debug("data {}", r);
          row.putAll(r);
        }
        else {
          logger.debug("text field {} value {}", field, value);
          row.put(field, value);
        }
      }
      // now have parameters in row
      logger.debug("row size {}  keys {}", row.size(), row.keySet());

      AbstractPrediction pr;
      String prJson;

      if (transform == null) { // no jar transformation
        logger.debug("no transformation of input data");
      }
      else {
        logger.debug("transformation of input data");
      }

      if (row != null) {
        // do the prediction
        pr = ServletUtil.predict(row);

        // assemble json result
        prJson = gson.toJson(pr);
        logger.debug(prJson);

        // Emit the prediction to the servlet response.
        response.getWriter().write(prJson);
      }

      response.setStatus(HttpServletResponse.SC_OK);
    }
    catch (Exception e) {
      // Prediction failed.
      logger.error("post failed", e);
      response.sendError(HttpServletResponse.SC_NOT_ACCEPTABLE, e.getMessage());
    }
    finally {
      // if the temp directory is still there we delete it
      if (tmpDir != null && tmpDir.exists()) {
        try {
          FileUtils.deleteDirectory(tmpDir);
        }
        catch (IOException e) {
          logger.error("Can't delete tmp directory");
        }
      }
    }
    long done = System.nanoTime();
    ServletUtil.postTimes.add(start, done, count);
    logger.debug("Post time {}", ServletUtil.postTimes);
  }

}

