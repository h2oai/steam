//package ai.h2o.servicebuilder;

import java.io.*;
import java.util.Map;
import java.util.Arrays;

import javax.servlet.http.*;
import javax.servlet.*;

import hex.genmodel.easy.prediction.BinomialModelPrediction;
import hex.genmodel.easy.prediction.RegressionModelPrediction;
import hex.genmodel.easy.prediction.AbstractPrediction;
import hex.genmodel.easy.*;
import hex.genmodel.*;

import com.google.gson.Gson;

public class PredictServlet extends HttpServlet {
  // Set to true for demo mode (to print the predictions to stdout).
  // Set to false to get better throughput.
  static boolean VERBOSE = false;

//  EasyPredictModelWrapper model = Model.getInstance().model;

//  public final class Model {
//    private static final Model INSTANCE = new Model();
//
//    public static EasyPredictModelWrapper model = null;
//    public static GenModel rawModel = null;
//
//    static {
//      rawModel = new REPLACE_THIS_WITH_PREDICTOR_CLASS_NAME();
//      model = new EasyPredictModelWrapper(rawModel);
//    }
//
//    private Model() {}
//
//    public static Model getInstance() {
//      return INSTANCE;
//    }
//  }

  public static EasyPredictModelWrapper model;

  static {
    GenModel rawModel = new REPLACE_THIS_WITH_PREDICTOR_CLASS_NAME();
    model = new EasyPredictModelWrapper(rawModel);

//    if (VERBOSE)
//	System.out.println(jsonModel());
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

      if (key.equals("VERBOSE")) {
	  if (values != null && values.length > 0) {
	      String tf = values[0];
	      if (tf.equals("true"))
		  VERBOSE = true;
	      else if (tf.equals("false"))
		  VERBOSE = false;
	  }
	  continue;
      }

      for (String value : values) {
        if (VERBOSE) System.out.println("Key: " + key + " Value: " + value);
        if (value.length() > 0) {
          row.put(key, value);
        }
      }
    }
  }

  public void doGet (HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
    RowData row = new RowData();
    fillRowDataFromHttpRequest(request, row);

    try {
	if (model == null)
	    throw new Exception("No predictor model");

        // we have a model loaded, do the prediction
        AbstractPrediction pr = model.predict(row);

	// assemble json result
	Gson gson = new Gson();
	String prJson = gson.toJson(pr);
	if (VERBOSE)
	    System.out.println(prJson);

	// Emit the prediction to the servlet response.
	if (VERBOSE)
	    response.getWriter().write(jsonModel());

	response.getWriter().write(prJson);
  response.setHeader("Access-Control-Allow-Origin", "*");
	response.setStatus(HttpServletResponse.SC_OK);

    } 
    catch (Exception e) {
	// Prediction failed.
	System.out.println(e.getMessage());
	response.sendError(HttpServletResponse.SC_NOT_ACCEPTABLE, e.getMessage());
    }
  }

}

