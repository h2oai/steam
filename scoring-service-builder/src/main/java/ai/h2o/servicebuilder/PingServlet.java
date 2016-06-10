package ai.h2o.servicebuilder;

import org.apache.commons.fileupload.FileItem;
import org.apache.commons.fileupload.disk.DiskFileItemFactory;
import org.apache.commons.fileupload.servlet.ServletFileUpload;
import org.apache.commons.io.FileUtils;
import org.apache.commons.io.filefilter.TrueFileFilter;

import javax.servlet.ServletException;
import javax.servlet.ServletOutputStream;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.File;
import java.io.IOException;
import java.util.Arrays;
import java.util.Collection;
import java.util.List;

import static ai.h2o.servicebuilder.Util.*;

/**
 * Compile server for POJO
 * <p>
 * Input is form with pojo java file and h2o-genmodel.jar
 * Output is the jar file of the compiled code
 * Errors are sent back if any
 */
public class PingServlet extends HttpServlet  {

  public void doGet(HttpServletRequest request, HttpServletResponse response) throws IOException {
    try {
      response.getWriter().write("pong");
      response.setStatus(HttpServletResponse.SC_OK);
    }
    catch (IOException e) {
      System.out.println(e.getMessage());
      response.sendError(HttpServletResponse.SC_NOT_ACCEPTABLE, e.getMessage());
    }
  }

}

