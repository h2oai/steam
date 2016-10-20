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
import org.apache.commons.io.filefilter.TrueFileFilter;
import org.slf4j.Logger;

import javax.servlet.ServletConfig;
import javax.servlet.ServletException;
import javax.servlet.ServletOutputStream;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.File;
import java.io.IOException;
import java.net.MalformedURLException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collection;
import java.util.List;

import static ai.h2o.servicebuilder.Util.*;

/**
 * Compile server for POJO to war file
 * <p>
 * curl -X POST --form pojo=@pojo/gbm_3f258f27_f0ad_4520_b6a5_3d2bb4a9b0ff.java --form jar=@pojo/h2o-genmodel.jar "localhost:8080/makewar" > model.war
 * java -jar jetty-runner.jar model.war
 * curl "localhost:8080/pred?DayOfMonth=1&Distance=2"
 * <p>
 * <p>
 * Input is form with pojo java file and h2o-genmodel.jar and extra tar file
 * Output is the war file of the compiled code
 * Errors are sent back if any
 */
public class MakeWarServlet extends HttpServlet {
  private final Logger logger = Logging.getLogger(this.getClass());

  private File servletPath = null;

  public void init(ServletConfig servletConfig) throws ServletException {
    super.init(servletConfig);
    try {
      servletPath = new File(servletConfig.getServletContext().getResource("/").getPath());
      logger.info("servletPath = {}", servletPath);
    }
    catch (MalformedURLException e) {
      logger.error("Can't init servlet", e);
    }
  }

  public void doPost(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
    Long startTime = System.currentTimeMillis();
    File tmpDir = null;
    try {
      //create temp directory
      tmpDir = createTempDirectory("makeWar");
      logger.info("tmpDir {}", tmpDir);

      //  create output directories
      File webInfDir = new File(tmpDir.getPath(), "WEB-INF");
      if (!webInfDir.mkdir())
        throw new Exception("Can't create output directory (WEB-INF)");
      File outDir = new File(webInfDir.getPath(), "classes");
      if (!outDir.mkdir())
        throw new Exception("Can't create output directory (WEB-INF/classes)");
      File libDir = new File(webInfDir.getPath(), "lib");
      if (!libDir.mkdir())
        throw new Exception("Can't create output directory (WEB-INF/lib)");

      // get input files
      List<FileItem> items = new ServletFileUpload(new DiskFileItemFactory()).parseRequest(request);
      String pojofile = null;
      String jarfile = null;
      String prejarfile = null;
      String deepwaterjarfile = null;
      String rawfile = null;
      String predictorClassName = null;
      String transformerClassName = null;
      ArrayList<String> pojos = new ArrayList<String>();
      ArrayList<String> rawfiles = new ArrayList<String>();
      for (FileItem i : items) {
        String field = i.getFieldName();
        String filename = i.getName();
        if (filename != null && filename.length() > 0) { // file fields
          if (field.equals("pojo")) {
            pojofile = filename;
            pojos.add(pojofile);
            predictorClassName = filename.replace(".java", "");
            FileUtils.copyInputStreamToFile(i.getInputStream(), new File(tmpDir, filename));
            logger.info("added pojo model {}", filename);
          }
          if (field.equals("jar")) {
            jarfile = "WEB-INF" + File.separator + "lib" + File.separator + filename;
            FileUtils.copyInputStreamToFile(i.getInputStream(), new File(libDir, filename));
          }
          if (field.equals("deepwater")) {
            deepwaterjarfile = "WEB-INF" + File.separator + "lib" + File.separator + filename;
            FileUtils.copyInputStreamToFile(i.getInputStream(), new File(libDir, filename));
          }
          if (field.equals("prejar")) {
            prejarfile = "WEB-INF" + File.separator + "lib" + File.separator + filename;
            FileUtils.copyInputStreamToFile(i.getInputStream(), new File(libDir, filename));
          }
          if (field.equals("mojo")) { // a raw model zip file, a mojo file
            rawfile = filename;
            rawfiles.add(rawfile);
            predictorClassName = filename.replace(".zip", "");
            FileUtils.copyInputStreamToFile(i.getInputStream(), new File(tmpDir, filename));
            logger.info("added mojo model {}", filename);
          }
        }
        else { // form text field
          if (field.equals("preclass")) {
            transformerClassName = i.getString();
          }
        }
      }
      logger.debug("genmodeljar {}  deepwaterjar {}  pojo {}  raw {}", jarfile, deepwaterjarfile, pojofile, rawfile);
      if ((pojofile == null || jarfile == null) && (rawfile == null || jarfile == null))
        throw new Exception("need either pojo and genmodel jar, or raw file and genmodel jar ");

      logger.info("prejar {}  preclass {}", prejarfile, transformerClassName);
      if (prejarfile != null && transformerClassName == null || prejarfile == null && transformerClassName != null)
        throw new Exception("need both prejar and preclass");

      if (pojofile != null) {
        // Compile the pojo
        runCmd(tmpDir, Arrays.asList("javac", "-target", JAVA_TARGET_VERSION, "-source", JAVA_TARGET_VERSION,
            "-J-Xmx" + MEMORY_FOR_JAVA_PROCESSES, "-cp", jarfile + ";" + deepwaterjarfile, "-d", outDir.getPath(), pojofile),
            "Compilation of pojo failed");
        logger.info("compiled pojo {}", pojofile);
      }

      if (servletPath == null)
        throw new Exception("servletPath is null");

      FileUtils.copyDirectoryToDirectory(new File(servletPath, "extra"), tmpDir);
      String extraPath = "extra" + File.separator;
      String webInfPath = extraPath + File.separator + "WEB-INF" + File.separator;
      String srcPath = extraPath + "src" + File.separator;

      if (transformerClassName == null)
        copyExtraFile(servletPath, extraPath, tmpDir, "index.html", "index.html");
      else
        copyExtraFile(servletPath, extraPath, tmpDir, "jarindex.html", "index.html");
      copyExtraFile(servletPath, extraPath, tmpDir, "jquery.js", "jquery.js");
      copyExtraFile(servletPath, extraPath, tmpDir, "predict.js", "predict.js");
      copyExtraFile(servletPath, extraPath, tmpDir, "custom.css", "custom.css");
      copyExtraFile(servletPath, webInfPath, webInfDir, "web-predict.xml", "web.xml");
      FileUtils.copyDirectoryToDirectory(new File(servletPath, webInfPath + "lib"), webInfDir);
      FileUtils.copyDirectoryToDirectory(new File(servletPath, extraPath + "bootstrap"), tmpDir);
      FileUtils.copyDirectoryToDirectory(new File(servletPath, extraPath + "fonts"), tmpDir);

      // change the class name in the predictor template file to the predictor we have
      String replaceTransform;
      if (transformerClassName == null)
        replaceTransform = "null";
      else
        replaceTransform = "new " + transformerClassName + "()";

      String modelCode = null;
      if (!pojos.isEmpty()) {
        FileUtils.writeLines(new File(tmpDir, "modelnames.txt"), pojos);
        modelCode = "null";
      }
      else if (!rawfiles.isEmpty()) {
        FileUtils.writeLines(new File(tmpDir, "modelnames.txt"), rawfiles);
        modelCode = "MojoModel.load(fileName)";
      }
      InstantiateJavaTemplateFile(tmpDir, modelCode, predictorClassName, replaceTransform, srcPath + "ServletUtil-TEMPLATE.java", "ServletUtil.java");

      copyExtraFile(servletPath, srcPath, tmpDir, "PredictServlet.java", "PredictServlet.java");
      copyExtraFile(servletPath, srcPath, tmpDir, "PredictBinaryServlet.java", "PredictBinaryServlet.java");
      copyExtraFile(servletPath, srcPath, tmpDir, "InfoServlet.java", "InfoServlet.java");
      copyExtraFile(servletPath, srcPath, tmpDir, "StatsServlet.java", "StatsServlet.java");
      copyExtraFile(servletPath, srcPath, tmpDir, "PingServlet.java", "PingServlet.java");
      copyExtraFile(servletPath, srcPath, tmpDir, "Transform.java", "Transform.java");
      copyExtraFile(servletPath, srcPath, tmpDir, "Logging.java", "Logging.java");

      // compile extra
      List<String> cmd = Arrays.asList("javac", "-target", JAVA_TARGET_VERSION, "-source", JAVA_TARGET_VERSION, "-J-Xmx" + MEMORY_FOR_JAVA_PROCESSES,
          "-cp", "WEB-INF/lib/*:WEB-INF/classes:extra/WEB-INF/lib/*", "-d", outDir.getPath(),
          "PredictServlet.java", "PredictBinaryServlet.java", "InfoServlet.java", "StatsServlet.java", "ServletUtil.java",
          "PingServlet.java", "Transform.java", "Logging.java");
      runCmd(tmpDir, cmd, "Compilation of extra failed");

      // create the war jar file
      Collection<File> filesc = FileUtils.listFilesAndDirs(webInfDir, TrueFileFilter.INSTANCE, TrueFileFilter.INSTANCE);
      filesc.add(new File(tmpDir, "index.html"));
      filesc.add(new File(tmpDir, "jquery.js"));
      filesc.add(new File(tmpDir, "predict.js"));
      filesc.add(new File(tmpDir, "custom.css"));
      filesc.add(new File(tmpDir, "modelnames.txt"));
      for (String m : pojos) {
        filesc.add(new File(tmpDir, m));
      }
      for (String m : rawfiles) {
        filesc.add(new File(tmpDir, m));
      }
      Collection<File> dirc = FileUtils.listFilesAndDirs(new File(tmpDir, "bootstrap"), TrueFileFilter.INSTANCE, TrueFileFilter.INSTANCE);
      filesc.addAll(dirc);
      dirc = FileUtils.listFilesAndDirs(new File(tmpDir, "fonts"), TrueFileFilter.INSTANCE, TrueFileFilter.INSTANCE);
      filesc.addAll(dirc);

      File[] files = filesc.toArray(new File[]{});
      if (files.length == 0)
        throw new Exception("Can't list compiler output files (out)");

      byte[] resjar = createJarArchiveByteArray(files, tmpDir.getPath() + File.separator);
      if (resjar == null)
        throw new Exception("Can't create war of compiler output");
      logger.info("war created from {} files, size {}", files.length, resjar.length);

      // send jar back
      ServletOutputStream sout = response.getOutputStream();
      response.setContentType("application/octet-stream");
      String outputFilename = predictorClassName.length() > 0 ? predictorClassName : "h2o-predictor";
      response.setHeader("Content-disposition", "attachment; filename=" + outputFilename + ".war");
      response.setContentLength(resjar.length);
      sout.write(resjar);
      sout.close();
      response.setStatus(HttpServletResponse.SC_OK);

      Long elapsedMs = System.currentTimeMillis() - startTime;
      logger.info("Done war creation in {} ms", elapsedMs);
    }
    catch (Exception e) {
      logger.error("doPost failed ", e);
      // send the error message back
      String message = e.getMessage();
      if (message == null) message = "no message";
      logger.error(message);
      response.setStatus(HttpServletResponse.SC_INTERNAL_SERVER_ERROR);
      response.getWriter().write(message);
      response.getWriter().write(Arrays.toString(e.getStackTrace()));
      response.sendError(HttpServletResponse.SC_INTERNAL_SERVER_ERROR, e.getMessage());
    }
    finally {
      // if the temp directory is still there we delete it
      if (tmpDir != null && tmpDir.exists()) {
        try {
          FileUtils.deleteDirectory(tmpDir);
        }
        catch (IOException e) {
          logger.error("Can't delete tmp directory");
        }
      }
    }

  }

}

