import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import com.google.gson.reflect.TypeToken;

import java.io.*;
import java.lang.reflect.Type;
import java.util.Map;
import java.util.HashMap;
import java.text.SimpleDateFormat;
import java.util.Date;
import java.util.TimeZone;
import javax.servlet.http.*;
import javax.servlet.*;

public class StatsServlet extends HttpServlet {

  private static final Gson gson = new GsonBuilder().serializeSpecialFloatingPointValues().create();
  private static final Type mapType = new TypeToken<HashMap<String, Object>>(){}.getType();

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
      final long startTime = PredictServlet.startTime;
      final long lastTime = PredictServlet.lastTime;

      Map<String, Object> js = new HashMap<String, Object>() {
        {
          put("startTime", startTime);
          put("lastTime", lastTime);
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
      String json = gson.toJson(js, mapType);

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
