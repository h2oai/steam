/*
  Copyright (C) 2016 H2O.ai, Inc. <http://h2o.ai/>

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU Affero General Public License as
  published by the Free Software Foundation, either version 3 of the
  License, or (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU Affero General Public License for more details.

  You should have received a copy of the GNU Affero General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package ai.h2o.servicebuilder;

import org.apache.commons.fileupload.FileItem;
import org.apache.commons.fileupload.disk.DiskFileItemFactory;
import org.apache.commons.fileupload.servlet.ServletFileUpload;
import org.apache.commons.io.FileUtils;
import org.apache.commons.io.IOUtils;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import javax.servlet.ServletConfig;
import javax.servlet.ServletException;
import javax.servlet.ServletOutputStream;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.*;
import java.net.MalformedURLException;
import java.util.*;

import static ai.h2o.servicebuilder.Util.*;

//import static ai.h2o.servicebuilder.Util.*;

/**
 * Compile server for POJO
 * <p>
 * Input is form with pojo java file and h2o-genmodel.jar
 * Output is the jar file of the compiled code
 * Errors are sent back if any
 */
public class CompilePojoServlet extends HttpServlet {
  private final Logger logger = Logging.getLogger(this.getClass());

  private static boolean VERBOSE = true;

  private File servletPath = null;

  public void init(ServletConfig servletConfig) throws ServletException {
    super.init(servletConfig);
    try {
      servletPath = new File(servletConfig.getServletContext().getResource("/").getPath());
      if (VERBOSE) logger.info("servletPath = {}", servletPath);
    }
    catch (MalformedURLException e) {
      logger.error("init failed", e);
    }
  }

  public void doPost(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
    File tmp = null;
    try {
      //create temp directory
      tmp = createTempDirectory("compilePojo");
      logger.info("tmp dir {}", tmp);

      // get input files
      List<FileItem> items = new ServletFileUpload(new DiskFileItemFactory()).parseRequest(request);
      List<String> pojofiles = new ArrayList<String>();
      String jarfile = null;
      for (FileItem i : items) {
        String field = i.getFieldName();
        String filename = i.getName();
        if (filename != null && filename.length() > 0) {
          if (field.equals("pojo")) {
            pojofiles.add(filename);
          }
          if (field.equals("jar")) {
            jarfile = filename;
          }
          FileUtils.copyInputStreamToFile(i.getInputStream(), new File(tmp, filename));
        }
      }
      if (pojofiles.isEmpty() || jarfile == null)
        throw new Exception("need pojofile(s) and jarfile");

      //  create output directory
      File out = new File(tmp.getPath(), "out");
      boolean mkDirResult = out.mkdir();
      if (!mkDirResult)
        throw new Exception("Can't create output directory (out)");

      if (servletPath == null)
        throw new Exception("servletPath is null");

      copyExtraFile(servletPath, "extra" + File.separator, tmp, "H2OPredictor.java", "H2OPredictor.java");
      FileUtils.copyDirectoryToDirectory(new File(servletPath, "extra" + File.separator + "WEB-INF" + File.separator + "lib"), tmp);
      copyExtraFile(servletPath, "extra" + File.separator, new File(out, "META-INF"), "MANIFEST.txt", "MANIFEST.txt");

      // Compile the pojo(s)
      for (String pojofile : pojofiles) {
        runCmd(tmp, Arrays.asList("javac", "-target", JAVA_TARGET_VERSION, "-source", JAVA_TARGET_VERSION, "-J-Xmx" + MEMORY_FOR_JAVA_PROCESSES,
            "-cp", jarfile + ":lib/*", "-d", "out", pojofile, "H2OPredictor.java"), "Compilation of pojo failed: " + pojofile);
      }

      // create jar result file
      runCmd(out, Arrays.asList("jar", "xf", tmp + File.separator + jarfile), "jar extraction of h2o-genmodel failed");

      runCmd(out, Arrays.asList("jar", "xf", tmp + File.separator + "lib" + File.separator + "gson-2.6.2.jar"), "jar extraction of gson failed");

      runCmd(out, Arrays.asList("jar", "cfm", tmp + File.separator + "result.jar", "META-INF" + File.separator + "MANIFEST.txt", "."), "jar creation failed");

      byte[] resjar = IOUtils.toByteArray(new FileInputStream(tmp + File.separator + "result.jar"));
      if (resjar == null)
        throw new Exception("Can't create jar of compiler output");

      logger.info("jar created, size {}", resjar.length);

      // send jar back
      ServletOutputStream sout = response.getOutputStream();
      response.setContentType("application/octet-stream");
      String outputFilename = pojofiles.get(0).replace(".java", "");
      response.setHeader("Content-disposition", "attachment; filename=" + outputFilename + ".jar");
      response.setContentLength(resjar.length);
      sout.write(resjar);
      sout.close();
      response.setStatus(HttpServletResponse.SC_OK);
    } catch (Exception e) {
      logger.error("post failed", e);
      // send the error message back
      String message = e.getMessage();
      if (message == null) message = "no message";
      logger.error(message);
      response.getWriter().write(message);
      response.getWriter().write(Arrays.toString(e.getStackTrace()));
      response.sendError(HttpServletResponse.SC_BAD_REQUEST, e.getMessage());
    }
    finally {
      // if the temp directory is still there we delete it
      try {
        if (tmp != null && tmp.exists())
          FileUtils.deleteDirectory(tmp);
      }
      catch (IOException e) {
        logger.error("Can't delete tmp directory");
      }
    }
  }

}

