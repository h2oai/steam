import com.google.gson.Gson;
import hex.genmodel.*;
import hex.genmodel.easy.*;
import org.slf4j.Logger;
import javax.servlet.http.*;
import javax.servlet.*;
import java.io.*;
import java.net.MalformedURLException;
import java.util.Map;
import hex.genmodel.easy.exception.PredictException;

public class InfoServlet extends HttpServlet {
  private final Logger logger = Logging.getLogger(this.getClass());

  static EasyPredictModelWrapper model = null;
  public static Map<String, EasyPredictModelWrapper> models = null;
  private File servletPath = null;
  public void init(ServletConfig servletConfig) throws ServletException {
    super.init(servletConfig);
    try {
      servletPath = new File(servletConfig.getServletContext().getResource("/").getPath());
      logger.debug("servletPath {}", servletPath);
      ServletUtil.loadModels(servletPath);
      model = ServletUtil.model;
      models = ServletUtil.models;
      logger.debug("model {}  models size {}", model, models.size());
    }
    catch (MalformedURLException e) {
      logger.error("init failed", e);
    }
  }

  public void doGet(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
    response.setHeader("Access-Control-Allow-Origin", "*");
    try {
      String pathInfo = request.getPathInfo();
      logger.debug("pathInfo {}", pathInfo);
      EasyPredictModelWrapper mod = null;

      if (pathInfo != null) {
        if (models.size() == 0)
          throw new PredictException("no models");

        String modelName = pathInfo.replace("/", "");
        mod = models.get(modelName);
        if (mod == null)
          throw new PredictException("can't find model " + modelName);
      }
      else {
        if (model == null)
          throw new Exception("No predictor model");
        mod = model;
      }

      // Emit the prediction to the servlet response.
      Gson gson = new Gson();
      String modelJson = gson.toJson(mod);
      response.getWriter().write(modelJson);
      response.setStatus(HttpServletResponse.SC_OK);
    }
    catch (Exception e) {
      // Prediction failed.
      logger.error("doGet failed", e);
      response.sendError(HttpServletResponse.SC_NOT_ACCEPTABLE, e.getMessage());
    }
  }

}
