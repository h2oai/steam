import java.io.*;
import java.util.Map;
import java.util.Arrays;
import javax.servlet.http.*;
import javax.servlet.*;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import hex.genmodel.easy.prediction.AbstractPrediction;
import hex.genmodel.easy.*;
import hex.genmodel.*;

import com.google.gson.Gson;

public class InfoServlet extends HttpServlet {

  private static final Logger logger = LoggerFactory.getLogger("InfoServlet");

  static EasyPredictModelWrapper model = ServletUtil.model;

  public void doGet(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
    try {
      if (model == null)
        throw new Exception("No predictor model");

      // Emit the prediction to the servlet response.
      Gson gson = new Gson();
      String modelJson = gson.toJson(model);
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
