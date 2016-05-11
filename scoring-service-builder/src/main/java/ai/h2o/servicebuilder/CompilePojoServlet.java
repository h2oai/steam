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
import java.util.*;

import static ai.h2o.servicebuilder.Util.*;

/**
 * Compile server for POJO
 * <p>
 * Input is form with pojo java file and h2o-genmodel.jar
 * Output is the jar file of the compiled code
 * Errors are sent back if any
 */
public class CompilePojoServlet extends HttpServlet {

  public void doPost(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
    File tmp = null;
    try {
      //create temp directory
      tmp = createTempDirectory("compilePojo");
      System.out.println("tmp dir " + tmp);

      // get input files
      List<FileItem> items = new ServletFileUpload(new DiskFileItemFactory()).parseRequest(request);
      String pojofile = null;
      String jarfile = null;
      for (FileItem i : items) {
        String field = i.getFieldName();
        String filename = i.getName();
        if (filename != null && filename.length() > 0) {
          if (field.equals("pojo")) pojofile = filename;
          if (field.equals("jar")) jarfile = filename;
          FileUtils.copyInputStreamToFile(i.getInputStream(), new File(tmp, filename));
        }
      }
      System.out.printf("jar %s  pojo %s\n", jarfile, pojofile);
      if (pojofile == null || jarfile == null)
        throw new Exception("need pojofile and jarfile");

      //  create output directory
      File out = new File(tmp.getPath(), "out");
      boolean mkDirResult = out.mkdir();
      if (!mkDirResult)
        throw new Exception("Can't create output directory (out)");

      // Compile the pojo
      runCmd(tmp, Arrays.asList("javac", "-target", JAVA_TARGET_VERSION, "-source", JAVA_TARGET_VERSION, "-J-Xmx" + MEMORY_FOR_JAVA_PROCESSES,
          "-cp", jarfile, "-d", "out", pojofile), "Compilation of pojo failed");

      // unpack jar file
      List<String> cmd2 = Arrays.asList("jar", "xf", tmp + File.separator + jarfile);
      runCmd(tmp, cmd2, "jar extraction failed");

      // create jar result file
      runCmd(out, Arrays.asList("jar", "xf", tmp + File.separator + jarfile), "jar extraction failed");

      Collection<File> filesc = FileUtils.listFilesAndDirs(out, TrueFileFilter.INSTANCE, TrueFileFilter.INSTANCE);
      File[] files = filesc.toArray(new File[]{});

      byte[] resjar = createJarArchiveByteArray(files, out.getPath() + File.separator);
      if (resjar == null)
        throw new Exception("Can't create jar of compiler output");

      System.out.println("jar created from " + files.length + " files, size " + resjar.length);

      // send jar back
      ServletOutputStream sout = response.getOutputStream();
      response.setContentType("application/octet-stream");
      response.setContentLength(resjar.length);
      sout.write(resjar);
      sout.close();
      response.setStatus(HttpServletResponse.SC_OK);

      System.out.println("Done compile and jar");
    } catch (Exception e) {
      e.printStackTrace();
      // send the error message back
      String message = e.getMessage();
      if (message == null) message = "no message";
      System.out.println(message);
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
        System.err.println("Can't delete tmp directory");
      }
    }
  }

}

