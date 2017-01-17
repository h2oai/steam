import com.google.gson.Gson;
import hex.genmodel.*;
import hex.genmodel.easy.*;
import org.slf4j.Logger;
import javax.servlet.http.*;
import javax.servlet.*;
import java.io.*;
import java.net.MalformedURLException;
import java.util.Map;
import java.util.HashMap;
import hex.genmodel.easy.exception.PredictException;

public class InfoServlet extends HttpServlet {
  private final Logger logger = Logging.getLogger(this.getClass());

  private Gson gson = new Gson();

  private static EasyPredictModelWrapper model = null;
  private static GenModel genModel = null;
  private static Map<String, EasyPredictModelWrapper> models = null;
  private static Map<String, GenModel> genModels = null;
  private String modelName = null;
  private File servletPath = null;

  public void init(ServletConfig servletConfig) throws ServletException {
    super.init(servletConfig);
    try {
      servletPath = new File(servletConfig.getServletContext().getResource("/").getPath());
      logger.debug("servletPath {}", servletPath);
      ServletUtil.loadModels(servletPath);
      model = ServletUtil.model;
      models = ServletUtil.models;
      genModels = ServletUtil.genModels;
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
      model = ServletUtil.model;
      genModel = ServletUtil.genModel;
      models = ServletUtil.models;
      genModels = ServletUtil.genModels;

      if (pathInfo != null) {
        if (models.size() == 0)
          throw new PredictException("no models");

        modelName = pathInfo.replace("/", "");
        mod = models.get(modelName);
        if (mod == null)
          throw new PredictException("can't find model " + modelName);
        genModel = genModels.get(modelName);
        if (genModel == null)
          throw new PredictException("can't find genmodel " + modelName);
      }
      else {
        if (model == null)
          throw new Exception("No predictor model");
        if (genModel == null)
          throw new Exception("No predictor genmodel");
        mod = model;
      }

      Map<String, Object> map = new HashMap<String, Object>();
      map.put("modelName", modelName);
      map.put("isSupervised", genModel.isSupervised());
      map.put("nfeatures", genModel.nfeatures());
      map.put("nclasses", genModel.nclasses());
      map.put("modelCategory", genModel.getModelCategory());

      map.put("UUID", genModel.getUUID());
      int numCols = genModel.getNumCols();
      map.put("numCols", numCols);
      map.put("names", genModel.getNames());
      map.put("responseIdx", genModel.getResponseIdx());
      map.put("domainValues", genModel.getDomainValues());
      map.put("numResponseClasses", genModel.getNumResponseClasses());
      map.put("isClassifier", genModel.isClassifier());
      map.put("isAutoEncoder", genModel.isAutoEncoder());
      map.put("predsSize", genModel.getPredsSize());

      map.put("responseDomainValues", mod.getResponseDomainValues());
      map.put("header", mod.getHeader());

      // get extra parameters is it's a Deepwater model
      Class cls = genModel.getClass();
      if (cls.getSimpleName().equals("DeepwaterMojoModel")) {
        map.put("problemType", cls.getField("_problem_type").get(genModel));
        map.put("height", cls.getField("_height").get(genModel));
        map.put("width", cls.getField("_width").get(genModel));
        map.put("channels", cls.getField("_channels").get(genModel));
      }

      String modelJson = gson.toJson(map);
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
