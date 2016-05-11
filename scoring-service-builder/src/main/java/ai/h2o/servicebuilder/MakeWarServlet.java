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
import java.io.*;
import java.util.Arrays;
import java.util.Collection;
import java.util.List;

import static ai.h2o.servicebuilder.Util.*;

/**
 * Compile server for POJO to war file
 * <p>
 * curl -X POST --form pojo=@pojo/gbm_3f258f27_f0ad_4520_b6a5_3d2bb4a9b0ff.java --form jar=@pojo/h2o-genmodel.jar --form extra=@makewar-files.tar "localhost:8080/makewar" > model.war
 * java -jar jetty-runner.jar model.war
 * curl "localhost:8080/pred?DayOfMonth=1&Distance=2"
 * <p>
 * <p>
 * Input is form with pojo java file and h2o-genmodel.jar and extra tar file
 * Output is the war file of the compiled code
 * Errors are sent back if any
 */
public class MakeWarServlet extends HttpServlet {

  public void doPost(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
    File tmpDir = null;
    try {
      //create temp directory
      tmpDir = createTempDirectory("makeWar");
      System.out.println("tmp dir " + tmpDir);

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
      String extrafile = null;
      String predictorClassName = null;
      for (FileItem i : items) {
        String field = i.getFieldName();
        String filename = i.getName();
        if (filename != null && filename.length() > 0) {
          if (field.equals("pojo")) {
            pojofile = filename;
            predictorClassName = filename.replace(".java", "");
            System.out.println("predictorClassName " + predictorClassName);
            FileUtils.copyInputStreamToFile(i.getInputStream(), new File(tmpDir, filename));
          }
          if (field.equals("jar")) {
            jarfile = "WEB-INF" + File.separator + "lib" + File.separator + filename;
            FileUtils.copyInputStreamToFile(i.getInputStream(), new File(libDir, filename));
          }
          if (field.equals("extra")) {
            extrafile = filename;
            FileUtils.copyInputStreamToFile(i.getInputStream(), new File(tmpDir, filename));
          }
        }
      }
      System.out.printf("jar %s  pojo %s  extra %s\n", jarfile, pojofile, extrafile);
      if (pojofile == null || jarfile == null || extrafile == null)
        throw new Exception("need pojo, jar and extra");

      // Compile the pojo
      runCmd(tmpDir, Arrays.asList("javac", "-target", JAVA_TARGET_VERSION, "-source", JAVA_TARGET_VERSION, "-J-Xmx" + MEMORY_FOR_JAVA_PROCESSES,
          "-cp", jarfile, "-d", outDir.getPath(), pojofile), "Compilation of pojo failed");

      // possible way to get files included with this servlet
      // instead of inclusing extras
      // would add /makewar-files.jar to this war file
      // seems to be addressed relative to root of this war file
      // request.getRequestDispatcher("/included.html").include(request, response)

      // unpack the extras file
      runCmd(tmpDir, Arrays.asList("jar", "xf", extrafile), "Unpack of extra failed");

      // change the class name in the predictor template file to the predictor we have
      InstantiateJavaTemplateFile(tmpDir, predictorClassName, "PredictServlet-TEMPLATE.java", "PredictServlet.java");
      InstantiateJavaTemplateFile(tmpDir, predictorClassName, "InfoServlet-TEMPLATE.java", "InfoServlet.java");
      InstantiateJavaTemplateFile(tmpDir, predictorClassName, "StatsServlet.java", "StatsServlet.java");
      // now have a correct PredictorServlet.java and InfoServlet.java files

      // compile extra
      runCmd(tmpDir, Arrays.asList("javac", "-target", JAVA_TARGET_VERSION, "-source", JAVA_TARGET_VERSION, "-J-Xmx" + MEMORY_FOR_JAVA_PROCESSES,
          "-cp", "WEB-INF/lib/*:WEB-INF/classes", "-d", outDir.getPath(),
          "src/PredictServlet.java", "src/InfoServlet.java", "src/StatsServlet.java"), "Compilation of extra failed");

      // create the war jar file
      Collection<File> filesc = FileUtils.listFilesAndDirs(webInfDir, TrueFileFilter.INSTANCE, TrueFileFilter.INSTANCE);
      filesc.add(new File(tmpDir, "index.html"));
      filesc.add(new File(tmpDir, "jquery.js"));
      filesc.add(new File(tmpDir, "predict.js"));
      File[] files = filesc.toArray(new File[]{});
      if (files.length == 0)
        throw new Exception("Can't list compiler output files (out)");

      byte[] resjar = createJarArchiveByteArray(files, tmpDir.getPath() + File.separator);
      if (resjar == null)
        throw new Exception("Can't create jar of compiler output");

      System.out.println("jar created from " + files.length + " files, size " + resjar.length);

      String outputFilename = predictorClassName.length() > 0 ? predictorClassName : "h2o-predictor";

      // send jar back
      ServletOutputStream sout = response.getOutputStream();
      response.setContentType("application/octet-stream");
      response.setHeader("Content-disposition", "attachment; filename=" + outputFilename + ".war");
      response.setContentLength(resjar.length);
      sout.write(resjar);
      sout.close();
      response.setStatus(HttpServletResponse.SC_OK);

      System.out.println("Done war creation");
    }
    catch (Exception e) {
      e.printStackTrace();
      // send the error message back
      String message = e.getMessage();
      if (message == null) message = "no message";
      System.out.println(message);
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
          System.err.println("Can't delete tmp directory");
        }
      }
    }

  }

}

