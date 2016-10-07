
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;

/**
 * Ping!
 * Check for aliveness.
 */
public class PingServlet extends HttpServlet  {
//  private final Logger logger = Logging.getLogger(this.getClass());

  public void doGet(HttpServletRequest request, HttpServletResponse response) throws IOException {
    response.setStatus(HttpServletResponse.SC_OK);
  }

}
