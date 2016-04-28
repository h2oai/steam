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
import java.nio.file.Files;
import java.util.Arrays;
import java.util.Collection;
import java.util.List;
import java.util.jar.JarEntry;
import java.util.jar.JarOutputStream;
import java.util.jar.Manifest;

/**
 * Compile server for POJO to war file
 *
 * curl -X POST --form pojo=@pojo/gbm_3f258f27_f0ad_4520_b6a5_3d2bb4a9b0ff.java --form jar=@pojo/h2o-genmodel.jar --form extra=@makewar-files.tar "localhost:8080/makewar" > model.war
 * java -jar jetty-runner.jar model.war
 * curl "localhost:8080/pred?DayOfMonth=1&Distance=2"
 *
 * <p>
 * Input is form with pojo java file and h2o-genmodel.jar and extra tar file
 * Output is the war file of the compiled code
 * Errors are sent back if any
 */
public class MakeWarServlet extends HttpServlet {

  public static final String JAVA_TEMPLATE_REPLACE_WITH_CLASS_NAME = "REPLACE_THIS_WITH_PREDICTOR_CLASS_NAME";
  public static String MEMORY_FOR_JAVA_PROCESSES = "4g";

  public void doPost(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
    File tmpDir = null;
    try {
      //create temp directory
      tmpDir = Files.createTempDirectory("makeWar").toFile();
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
            predictorClassName = filename.replace(".java" ,"");
            System.out.println("predictorClassName " + predictorClassName);
            Files.copy(i.getInputStream(), new File(tmpDir, filename).toPath());
          }
          if (field.equals("jar")) {
            jarfile = "WEB-INF" + File.separator + "lib" + File.separator + filename;
            Files.copy(i.getInputStream(), new File(libDir, filename).toPath());
          }
          if (field.equals("extra")) {
            extrafile = filename;
            Files.copy(i.getInputStream(), new File(tmpDir, filename).toPath());
          }
        }
      }
      System.out.printf("jar %s  pojo %s  extra %s\n", jarfile, pojofile, extrafile);
      if (pojofile == null || jarfile == null || extrafile == null)
        throw new Exception("need pojo, jar and extra");

      // Compile the pojo
      runCmd(tmpDir, Arrays.asList("javac", "-J-Xmx" + MEMORY_FOR_JAVA_PROCESSES, "-cp", jarfile, "-d", outDir.getPath(), pojofile), "Compilation of pojo failed");

      // possible way to get files included with this servlet
      // instead of inclusing extras
      // would add /makewar-files.jar to this war file
      // seems to be addressed relative to root of this war file
      // request.getRequestDispatcher("/included.html").include(request, response)

      // unpack the extras file
      runCmd(tmpDir, Arrays.asList("jar", "xf", extrafile), "Unpack of extra failed");

      // change the class name in the predictor template file to the predictor we have
      InstantiateJavaTemplateFile(tmpDir, predictorClassName, "Model-TEMPLATE.java", "Model.java");
      InstantiateJavaTemplateFile(tmpDir, predictorClassName, "PredictServlet-TEMPLATE.java", "PredictServlet.java");
      InstantiateJavaTemplateFile(tmpDir, predictorClassName, "InfoServlet-TEMPLATE.java", "InfoServlet.java");
      // now have a correct PredictorServlet.java and InfoServlet.java files

      // compile extra
      runCmd(tmpDir, Arrays.asList("javac", "-J-Xmx" + MEMORY_FOR_JAVA_PROCESSES, "-cp", "WEB-INF/lib/*:WEB-INF/classes", "-d", outDir.getPath(),
          "src/PredictServlet.java", "src/InfoServlet.java"), "Compilation of extra failed");

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
    } catch (Exception e) {
      e.printStackTrace();
      // send the error message back
      String message = e.getMessage();
      if (message == null) message = "no message";
      System.out.println(message);
      response.setStatus(HttpServletResponse.SC_INTERNAL_SERVER_ERROR);
      response.getWriter().write(message);
      response.getWriter().write(Arrays.toString(e.getStackTrace()));
      response.sendError(HttpServletResponse.SC_INTERNAL_SERVER_ERROR, e.getMessage());
    } finally {
      // if the temp directory is still there we delete it
      if (tmpDir != null && Files.exists(tmpDir.toPath())) {
//        try {
//          FileUtils.deleteDirectory(tmpDir);
//        }
//        catch (IOException e) {
//          System.err.println("Can't delete tmp directory");
//        }

      }
    }

  }

  private void InstantiateJavaTemplateFile(File tmpDir, String javaClassName, String templateFileName, String resultFileName) throws IOException {
    File srcDir = new File(tmpDir, "src");
    byte[] templateJava = Files.readAllBytes(new File(srcDir, templateFileName).toPath());
    String java = new String(templateJava).replace(JAVA_TEMPLATE_REPLACE_WITH_CLASS_NAME, javaClassName);
    Files.write(new File(srcDir, resultFileName).toPath(), java.getBytes());
  }

  /**
   * Run command cmd in separate process in directory
   *
   * @param directory     run in this directory
   * @param cmd           command to run
   * @param errorMessage  error message if process didn't finish with exit value 0
   * @return stdout combined with stderr
   * @throws Exception
   */
  private String runCmd(File directory, List<String> cmd, String errorMessage) throws Exception {
    ProcessBuilder pb = new ProcessBuilder(cmd);
    pb.directory(directory);
    pb.redirectErrorStream(true); // error sent to output stream
    Process p = pb.start();

    // get output stream to string
    String s;
    StringBuilder sb = new StringBuilder();
    BufferedReader stdout = new BufferedReader(new InputStreamReader(p.getInputStream()));
    while ((s = stdout.readLine()) != null) {
      System.out.println(s);
      sb.append(s);
      sb.append('\n');
    }
    String sbs = sb.toString();
    int exitValue = p.waitFor();
    if (exitValue != 0 || sbs.length() > 0)
      throw new Exception(errorMessage + " exit value " + exitValue + "  " + sbs);
    return sbs;
  }

  /**
   * Create jar archive out of files list. Names in archive have paths starting from relativeToDir
   *
   * @param tobeJared list of files
   * @param relativeToDir starting directory for paths
   * @return jar as byte array
   * @throws IOException
   */
  private byte[] createJarArchiveByteArray(File[] tobeJared, String relativeToDir) throws IOException {
    int BUFFER_SIZE = 10240;
    byte buffer[] = new byte[BUFFER_SIZE];
    ByteArrayOutputStream stream = new ByteArrayOutputStream();
    JarOutputStream out = new JarOutputStream(stream, new Manifest());

    for (File t : tobeJared) {
      if (t == null || !t.exists() || t.isDirectory())
        continue;

      // Create jar entry
      String filename = t.getPath().replace(relativeToDir, "").replace("\\", "/");
      JarEntry jarAdd = new JarEntry(filename);
      jarAdd.setTime(t.lastModified());
      out.putNextEntry(jarAdd);

      // Write file to archive
      FileInputStream in = new FileInputStream(t);
      while (true) {
        int nRead = in.read(buffer, 0, buffer.length);
        if (nRead <= 0)
          break;
        out.write(buffer, 0, nRead);
      }
      in.close();
    }

    out.close();
    stream.close();
    return stream.toByteArray();
  }

}

