import java.io.*;
import java.util.Map;
import java.util.HashMap;
import java.text.SimpleDateFormat;
import java.util.Date;
import java.util.TimeZone;
import javax.servlet.http.*;
import javax.servlet.*;

public class StatsServlet extends HttpServlet {

  public void doGet(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
    try {
      if (PredictServlet.model == null)
        throw new Exception("No predictor model");

      final long now = System.currentTimeMillis();
      final long upTimeMs = now - PredictServlet.startTime;
      final long lastTimeAgoMs = now - PredictServlet.lastTime;
      SimpleDateFormat sdf = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss z");
      sdf.setTimeZone(TimeZone.getTimeZone("UTC"));
      final String startUTC = sdf.format(new Date(PredictServlet.startTime));
      final String lastPredictionUTC = PredictServlet.lastTime > 0 ? sdf.format(new Date(PredictServlet.lastTime)) : "";
      final long warmUpCount = PredictServlet.warmUpCount;

      Map<String, Object> js = new HashMap<String, Object>() {
        {
          put("startTime", PredictServlet.startTime);
          put("lastTime", PredictServlet.lastTime);
          put("lastTimeUTC", lastPredictionUTC);
          put("startTimeUTC", startUTC);
          put("upTimeMs", upTimeMs);
          put("lastTimeAgoMs", lastTimeAgoMs);
          put("lastTimeAgoMs", lastTimeAgoMs);
          put("warmUpCount", warmUpCount);

          put("prediction", PredictServlet.predictionTimes.toMap());
          put("get", PredictServlet.getTimes.toMap());
          put("post", PredictServlet.postTimes.toMap());
          put("pythonget", PredictServlet.getPythonTimes.toMap());
          put("pythonpost", PredictServlet.postPythonTimes.toMap());
        }
      };
      String json = PredictServlet.gson.toJson(js, PredictServlet.mapType);

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
