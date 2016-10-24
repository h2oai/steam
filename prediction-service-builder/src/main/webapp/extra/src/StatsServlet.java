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
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class StatsServlet extends HttpServlet {
  private final Logger logger = Logging.getLogger(this.getClass());

  public static final Gson gson = new GsonBuilder().serializeSpecialFloatingPointValues().create();

  public void doGet(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
    response.setHeader("Access-Control-Allow-Origin", "*");
    try {
      final long now = System.currentTimeMillis();
      final long upTimeMs = now - ServletUtil.startTime;
      final long lastTimeAgoMs = now - ServletUtil.lastTime;
      SimpleDateFormat sdf = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss z");
      sdf.setTimeZone(TimeZone.getTimeZone("UTC"));
      final String startUTC = sdf.format(new Date(ServletUtil.startTime));
      final String lastPredictionUTC = ServletUtil.lastTime > 0 ? sdf.format(new Date(ServletUtil.lastTime)) : "";
      final long warmUpCount = ServletUtil.warmUpCount;

      Map<String, Object> js = new HashMap<String, Object>() {
        {
          put("startTime", ServletUtil.startTime);
          put("lastTime", ServletUtil.lastTime);
          put("lastTimeUTC", lastPredictionUTC);
          put("startTimeUTC", startUTC);
          put("upTimeMs", upTimeMs);
          put("lastTimeAgoMs", lastTimeAgoMs);
          put("lastTimeAgoMs", lastTimeAgoMs);
          put("warmUpCount", warmUpCount);

          put("prediction", ServletUtil.predictionTimes.toMap());
          put("get", ServletUtil.getTimes.toMap());
          put("post", ServletUtil.postTimes.toMap());
          put("pythonget", ServletUtil.getPythonTimes.toMap());
          put("pythonpost", ServletUtil.postPythonTimes.toMap());

//          put("outputLabels", ServletUtil.outputLabels);
        }
      };
      String json = gson.toJson(js, ServletUtil.MAP_TYPE);

      response.getWriter().write(json);
      response.setStatus(HttpServletResponse.SC_OK);
    }
    catch (Exception e) {
      // Prediction failed.
      logger.error("get failed", e);
      response.sendError(HttpServletResponse.SC_NOT_ACCEPTABLE, e.getMessage());
    }
  }

}
