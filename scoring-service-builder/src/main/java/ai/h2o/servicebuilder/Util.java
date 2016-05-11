package ai.h2o.servicebuilder;

import org.apache.commons.io.FileUtils;

import java.io.*;
import java.util.List;
import java.util.jar.JarEntry;
import java.util.jar.JarOutputStream;
import java.util.jar.Manifest;

/**
 * Util methods for make war servlet and compile pojo servlet
 *
 * Created by magnus on 5/10/16.
 */
class Util {

  static final String MEMORY_FOR_JAVA_PROCESSES = "4g";
  static final String JAVA_TARGET_VERSION = "1.6";

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
  static void InstantiateJavaTemplateFile(File tmpDir, String javaClassName, String templateFileName, String resultFileName) throws IOException {
    File srcDir = new File(tmpDir, "src");
    byte[] templateJava = FileUtils.readFileToByteArray(new File(srcDir, templateFileName));
    String java = new String(templateJava).replace(JAVA_TEMPLATE_REPLACE_WITH_CLASS_NAME, javaClassName);
    FileUtils.writeStringToFile(new File(srcDir, resultFileName), java);
  }

  /**
   * Run command cmd in separate process in directory
   *
   * @param directory    run in this directory
   * @param cmd          command to run
   * @param errorMessage error message if process didn't finish with exit value 0
   * @return stdout combined with stderr
   * @throws Exception
   */
  static String runCmd(File directory, List<String> cmd, String errorMessage) throws Exception {
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
    if (exitValue != 0)
      throw new Exception(errorMessage + " exit value " + exitValue + "  " + sbs);
    return sbs;
  }

  /**
   * Create jar archive out of files list. Names in archive have paths starting from relativeToDir
   *
   * @param tobeJared     list of files
   * @param relativeToDir starting directory for paths
   * @return jar as byte array
   * @throws IOException
   */
  static byte[] createJarArchiveByteArray(File[] tobeJared, String relativeToDir) throws IOException {
    int BUFFER_SIZE = 10240;
    byte buffer[] = new byte[BUFFER_SIZE];
    ByteArrayOutputStream stream = new ByteArrayOutputStream();
    JarOutputStream out = new JarOutputStream(stream, new Manifest());

    for (File t : tobeJared) {
      if (t == null || !t.exists() || t.isDirectory())
        continue;

      // Create jar entry
      String filename = t.getPath().replace(relativeToDir, "").replace("\\", "/");
      if (filename.endsWith("MANIFEST.MF")) { // skip to avoid duplicates
        continue;
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

  static File createTempDirectory(String prefix) throws IOException {
    File temp = File.createTempFile(prefix, Long.toString(System.nanoTime()));
    if (!(temp.delete())) {
      throw new IOException("Could not delete temp file: " + temp.getAbsolutePath());
    }
    if (!(temp.mkdir())) {
      throw new IOException("Could not create temp directory: " + temp.getAbsolutePath());
    }
    return temp;
  }

}
