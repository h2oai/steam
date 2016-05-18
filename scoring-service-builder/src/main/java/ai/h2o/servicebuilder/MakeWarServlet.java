package ai.h2o.servicebuilder;

import org.apache.commons.fileupload.FileItem;
import org.apache.commons.fileupload.disk.DiskFileItemFactory;
import org.apache.commons.fileupload.servlet.ServletFileUpload;
import org.apache.commons.io.FileUtils;
import org.apache.commons.io.filefilter.TrueFileFilter;

import javax.servlet.ServletConfig;
import javax.servlet.ServletException;
import javax.servlet.ServletOutputStream;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.*;
import java.net.MalformedURLException;
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

  private File servletPath = null;

  public void init(ServletConfig servletConfig) throws ServletException {
    super.init(servletConfig);
    try {
      servletPath = new File(servletConfig.getServletContext().getResource("/").getPath());
      System.out.println("path = " + servletPath);
    }
    catch (MalformedURLException e) {
      e.printStackTrace();
    }
  }

  public void doPost(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
    Long startTime = System.currentTimeMillis();
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
        }
      }
      System.out.printf("jar %s  pojo %s\n", jarfile, pojofile);
      if (pojofile == null || jarfile == null)
        throw new Exception("need pojo and jar");

      // Compile the pojo
      runCmd(tmpDir, Arrays.asList("javac", "-target", JAVA_TARGET_VERSION, "-source", JAVA_TARGET_VERSION, "-J-Xmx" + MEMORY_FOR_JAVA_PROCESSES,
          "-cp", jarfile, "-d", outDir.getPath(), pojofile), "Compilation of pojo failed");

      if (servletPath == null)
        throw new Exception("servletPath is null");

      FileUtils.copyDirectoryToDirectory(new File(servletPath, "extra"), tmpDir);
      String extraPath = "extra" + File.separator;
      String webInfPath = extraPath + File.separator + "WEB-INF" + File.separator;
      String srcPath = extraPath + "src" + File.separator;
      copyExtraFile(extraPath, tmpDir, "index.html");
      copyExtraFile(extraPath, tmpDir, "jquery.js");
      copyExtraFile(extraPath, tmpDir, "predict.js");
      copyExtraFile(webInfPath, webInfDir, "web.xml");
      FileUtils.copyDirectoryToDirectory(new File(servletPath, webInfPath + "lib"), webInfDir);

      // change the class name in the predictor template file to the predictor we have
      InstantiateJavaTemplateFile(tmpDir, predictorClassName, srcPath + "PredictServlet-TEMPLATE.java", "PredictServlet.java");
      copyExtraFile(srcPath, tmpDir, "InfoServlet.java");
      copyExtraFile(srcPath, tmpDir, "StatsServlet.java");

      // compile extra
      runCmd(tmpDir, Arrays.asList("javac", "-target", JAVA_TARGET_VERSION, "-source", JAVA_TARGET_VERSION, "-J-Xmx" + MEMORY_FOR_JAVA_PROCESSES,
          "-cp", "WEB-INF/lib/*:WEB-INF/classes:extra/WEB-INF/lib/*", "-d", outDir.getPath(),
          "PredictServlet.java", "InfoServlet.java", "StatsServlet.java"),
          "Compilation of extra failed");

      // create the war jar file
      Collection<File> filesc = FileUtils.listFilesAndDirs(webInfDir, TrueFileFilter.INSTANCE, TrueFileFilter.INSTANCE);
      filesc.add(new File(tmpDir, "index.html"));
      filesc.add(new File(tmpDir, "jquery.js"));
      filesc.add(new File(tmpDir, "predict.js"));
      File[] files = filesc.toArray(new File[]{});
      if (files.length == 0)
        throw new Exception("Can't list compiler output files (out)");

//      System.out.println(filesc);

      byte[] resjar = createJarArchiveByteArray(files, tmpDir.getPath() + File.separator);
      if (resjar == null)
        throw new Exception("Can't create war of compiler output");
      System.out.println("war created from " + files.length + " files, size " + resjar.length);

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
      System.out.println("Done war creation in " + elapsedMs + " ms");
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



  private static final String JAVA_TEMPLATE_REPLACE_WITH_CLASS_NAME = "REPLACE_THIS_WITH_PREDICTOR_CLASS_NAME";

  /**
   * The Java template file has a placeholder for the model name -- we replace that here
   *
   * @param tmpDir            run in this directory
   * @param javaClassName     model name
   * @param templateFileName  template file
   * @param resultFileName    restult file
   * @throws IOException
   */
  private static void InstantiateJavaTemplateFile(File tmpDir, String javaClassName, String templateFileName, String resultFileName) throws IOException {
//    File srcDir = new File(tmpDir, "src");
    byte[] templateJava = FileUtils.readFileToByteArray(new File(tmpDir, templateFileName));
    String java = new String(templateJava).replace(JAVA_TEMPLATE_REPLACE_WITH_CLASS_NAME, javaClassName);
    FileUtils.writeStringToFile(new File(tmpDir, resultFileName), java);
  }



  private void copyExtraFile(String extraPath, File toDir, String fileName) throws IOException {
    FileUtils.copyFile(new File(servletPath, extraPath + fileName), new File(toDir, fileName));
  }

}

