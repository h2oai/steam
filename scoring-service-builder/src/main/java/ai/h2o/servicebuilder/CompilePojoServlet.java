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
import java.io.ByteArrayOutputStream;
import java.nio.file.Files;
import java.util.*;
import java.util.jar.JarEntry;
import java.util.jar.JarOutputStream;
import java.util.jar.Manifest;

import javax.tools.*;

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
      tmp = Files.createTempDirectory("compilePojo").toFile();
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
          Files.copy(i.getInputStream(), new File(tmp, filename).toPath());
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

      // compile with system compiler

      List<String> cmd = Arrays.asList("javac", "-J-Xmx4g", "-cp", jarfile, "-d", "out", pojofile);
      runCmd(tmp, cmd, "Compilation failed");

      // unpack jar file
      List<String> cmd2 = Arrays.asList("jar", "xf", tmp + File.separator + jarfile);
      runCmd(tmp, cmd2, "jar extraction failed");

      // create jar result file
      runCmd(out, Arrays.asList("jar", "xf", tmp + File.separator + jarfile), "jar extraction failed");

      Collection<File> filesc = FileUtils.listFilesAndDirs(out, TrueFileFilter.INSTANCE, TrueFileFilter.INSTANCE);
      File[] files = filesc.toArray(new File[]{});
      if (files == null)
        throw new Exception("Can't list compiler output files (out)");

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
    } finally {
      // if the temp directory is still there we delete it
      if (tmp != null && Files.exists(tmp.toPath())) {
        FileUtils.deleteDirectory(tmp);
      }
    }

  }

  /**
   * Run command as separate process
   *
   * @param directory dir to run in
   * @param cmd command list to run
   * @param errorMessage on fail throw Exception with this error message
   * @return stdout and stderr combined
   * @throws Exception
   */
  private String runCmd(File directory, List<String> cmd, String errorMessage) throws Exception {
    ProcessBuilder pb = new ProcessBuilder(cmd);
    pb.directory(directory);
    pb.redirectErrorStream(true); // error sent to output stream
    Process p = pb.start();

    // get output stream to string
    String s;
    StringBuffer sb = new StringBuffer();
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
   * Create jar archive
   *
   * @param tobeJared file list
   * @param relativeToDir path names become relative to this dir
   * @return jar file as byte array
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
      if (!filename.endsWith(".class")) {
        System.out.println("create jar skipping file " + filename);
        continue; // skip unless it's a class file
      }

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

