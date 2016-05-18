import java.io.*;
import java.util.Map;
import java.util.HashMap;
import java.util.Arrays;
import java.text.SimpleDateFormat;
import java.util.Date;
import java.util.TimeZone;
import javax.servlet.http.*;
import javax.servlet.*;

import hex.genmodel.easy.prediction.AbstractPrediction;
import hex.genmodel.easy.*;
import hex.genmodel.*;

import com.google.gson.Gson;

public class StatsServlet extends HttpServlet {

  static EasyPredictModelWrapper model = PredictServlet.model;

  public void doGet(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
    try {
      if (model == null)
        throw new Exception("No predictor model");

      // Emit the prediction to the servlet response.
      final long now = System.currentTimeMillis();
      final long n = PredictServlet.numberOfPredictions;
      final double avgTimeMs = n > 0 ? PredictServlet.totalTimeMs / n : 0.0;
      final int warmupN = PredictServlet.warmupNumber;
      final double totalPredictionTimeAfterWarmupMs = n > warmupN ? (PredictServlet.totalTimeMs - PredictServlet.warmupTimeMs) : 0.0;
      final double avgTimeAfterWarmupMs = n > warmupN ? totalPredictionTimeAfterWarmupMs / (n - warmupN) : 0.0;
      final long upTimeMs = now - PredictServlet.startTime;
      final long lastTimeAgoMs = now - PredictServlet.lastTime;
      SimpleDateFormat sdf = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss z");
      sdf.setTimeZone(TimeZone.getTimeZone("UTC"));
      final String startUTC = sdf.format(new Date(PredictServlet.startTime));
      final String lastPredictionUTC = sdf.format(new Date(PredictServlet.lastTime));
      Map<String, String> js = new HashMap<String, String>() {
        {
          put("numberOfPredictions", Long.toString(n));
          put("startTime", Long.toString(PredictServlet.startTime));
          put("lastTime", Long.toString(PredictServlet.lastTime));
          put("lastTimeUTC", lastPredictionUTC);
          put("startTimeUTC", startUTC);
          put("upTimeMs", Long.toString(upTimeMs));
          put("lastTimeAgoMs", Long.toString(lastTimeAgoMs));
          put("lastPredictionMs", Double.toString(PredictServlet.lastPredictionMs));
          put("avgPredictionTimeMs", Double.toString(avgTimeMs));
          put("avgPredictionTimeAfterWarmupMs", Double.toString(avgTimeAfterWarmupMs));
          put("totalPredictionTimeMs", Double.toString(PredictServlet.totalTimeMs));
          put("totalPredictionTimeSquareMs", Double.toString(PredictServlet.totalTimeSquareMs));
          put("totalPredictionTimeAfterWarmupMs", Double.toString(totalPredictionTimeAfterWarmupMs));
          put("warmupTimeMs", Double.toString(PredictServlet.warmupTimeMs));
          put("warmupTimeSquareMs", Double.toString(PredictServlet.warmupTimeSquareMs));
          put("warmupFirstN", Integer.toString(PredictServlet.warmupNumber));
        }
      };
      StringBuilder sb = new StringBuilder("{ ");
      for (String k : js.keySet()) {
        String v = js.get(k);
        sb.append("\"" + k + "\": \"" + v + "\", ");
      }
      sb.append("}");
      String json = sb.toString().replace(", }", " }");
//      Gson gson = new Gson();
//      String json = gson.toJson(js);
      response.getWriter().write(json);
      response.setStatus(HttpServletResponse.SC_OK);
    }
    catch (Exception e) {
      // Prediction failed.
      System.out.println(e.getMessage());
      response.sendError(HttpServletResponse.SC_NOT_ACCEPTABLE, e.getMessage());
    }
  }

}
