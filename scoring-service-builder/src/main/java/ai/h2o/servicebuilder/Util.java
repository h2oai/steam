package ai.h2o.servicebuilder;

import org.apache.commons.io.FileUtils;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

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
public class Util {
  private static final Logger logger = Logging.getLogger(Util.class);

  public static final String JAVA_TEMPLATE_REPLACE_WITH_PREDICTOR_CLASS_NAME = "REPLACE_THIS_WITH_PREDICTOR_CLASS_NAME";
  public static final String JAVA_TEMPLATE_REPLACE_WITH_TRANSFORMER_OBJECT = "REPLACE_THIS_WITH_TRANSFORMER_OBJECT";
//  public static final String JAVA_TEMPLATE_REPLACE_WITH_POJO_BOOLEAN = "REPLACE_THIS_WITH_POJO_BOOLEAN";
  public static final String REPLACE_THIS_WITH_MODEL = "REPLACE_THIS_WITH_MODEL";

  protected static final String MEMORY_FOR_JAVA_PROCESSES = "4g";
  protected static final String JAVA_TARGET_VERSION = "1.6";

//  public static class Times {
//    private long count = 0;
//    private double totalTimeMs = 0;
//    private double totalTimeSquaredMs = 0;
//    private double warmupTimeMs = 0;
//    private double warmupTimeSquaredMs = 0;
//    private double lastMs = 0;
//
//    private Gson gson = new Gson();
//
//    public static final int warmUpCount = 5;
//    public static final Type MapType = new TypeToken<HashMap<String, Object>>(){}.getType();
////    public static final Type ROW_DATA_TYPE = new TypeToken<RowData>(){}.getType();
//
//
//    public void add(long startNs, long endNs, int n) {
//      double elapsed = (endNs - startNs) / 1.0e6;
//      add(elapsed, n);
//    }
//
//    public void add(long startNs, long endNs) {
//      add(startNs, endNs, 1);
//    }
//
//    public synchronized void add(double timeMs, int n) {
//      count += n;
//      totalTimeMs += timeMs; // n * timeMs/n
//      double tt = timeMs * timeMs / n; // n * (timeMs/n)^2
//      totalTimeSquaredMs += tt;
//      if (count <= warmUpCount) {
//        warmupTimeMs += timeMs;
//        warmupTimeSquaredMs += tt;
//      }
//      lastMs = timeMs / n;
//    }
//
//    public double avg() {
//      return count > 0 ? totalTimeMs / count : 0.0;
//    }
//
////    public double sdev() {
////
////    }
//
//    public double avgAfterWarmup() {
//      return count > warmUpCount ? (totalTimeMs - warmupTimeMs) / (count - warmUpCount) : 0.0;
//    }
//
//    public String toJson() {
//      return gson.toJson(toMap());
//    }
//
//    public Map<String, Object> toMap() {
//      Map<String, Object> map = classToMap();
//      map.put("averageTime", avg());
//      map.put("averageAfterWarmupTime", avgAfterWarmup());
//      return map;
//    }
//
//    private Map<String, Object> classToMap() {
//      return gson.fromJson(gson.toJson(this), MapType);
//    }
//
//    public String toString() {
//      return String.format("n %d  last %.3f  avg %.3f after warmup %.3f [ms]", count, lastMs, avg(), avgAfterWarmup());
//    }
//  }
//
//  public static long startTime = System.currentTimeMillis();
//  public static long lastTime = 0;
//  public static Times predictionTimes = new Times();
//  public static Times getTimes = new Times();
//  public static Times postTimes = new Times();
//  public static Times getPythonTimes = new Times();
//  public static Times postPythonTimes = new Times();
//

//  public static synchronized AbstractPrediction predict(EasyPredictModelWrapper model, RowData row)
//      throws PredictException {
//    long start = System.nanoTime();
//    AbstractPrediction p = model.predict(row);
//    long done = System.nanoTime();
//    lastTime = System.currentTimeMillis();
//    predictionTimes.add(start, done);
//
////    if (VERBOSE) System.out.println("Prediction time " + predictionTimes);
//    return p;
//  }


  /**
   * Run command cmd in separate process in directory
   *
   * @param directory    run in this directory
   * @param cmd          command to run
   * @param errorMessage error message if process didn't finish with exit value 0
   * @return stdout combined with stderr
   * @throws Exception
   */
 public static String runCmd(File directory, List<String> cmd, String errorMessage) throws Exception {
    ProcessBuilder pb = new ProcessBuilder(cmd);
    pb.directory(directory);
    pb.redirectErrorStream(true); // error sent to output stream
    Process p = pb.start();

    // get output stream to string
    String s;
    StringBuilder sb = new StringBuilder();
    BufferedReader stdout = new BufferedReader(new InputStreamReader(p.getInputStream()));
    while ((s = stdout.readLine()) != null) {
      logger.info(s);
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
  public static byte[] createJarArchiveByteArray(File[] tobeJared, String relativeToDir) throws IOException {
    int BUFFER_SIZE = 10240;
    byte buffer[] = new byte[BUFFER_SIZE];
    ByteArrayOutputStream stream = new ByteArrayOutputStream();
    JarOutputStream out = new JarOutputStream(stream, new Manifest());

    for (File t : tobeJared) {
      if (t == null || !t.exists() || t.isDirectory()) {
        if (t != null && !t.isDirectory())
          logger.error("Can't add to jar {}", t);
        continue;
      }
      // Create jar entry
      String filename = t.getPath().replace(relativeToDir, "").replace("\\", "/");
//      if (filename.endsWith("MANIFEST.MF")) { // skip to avoid duplicates
//        continue;
//      }
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

  static void copyExtraFile(File servletPath, String extraPath, File toDir, String fromFileName, String toFileName) throws IOException {
    FileUtils.copyFile(new File(servletPath, extraPath + fromFileName), new File(toDir, toFileName));
  }

  /**
   * The Java template file has a placeholder for the model name -- we replace that here
   *
   * @param tmpDir            run in this directory
   * @param javaClassName     model name
   * @param templateFileName  template file
   * @param resultFileName    restult file
   * @throws IOException
   */
  public static void InstantiateJavaTemplateFile(File tmpDir, String modelCode, String javaClassName, String replaceTransform, String templateFileName, String resultFileName) throws IOException {
    byte[] templateJava = FileUtils.readFileToByteArray(new File(tmpDir, templateFileName));
    String java = new String(templateJava)
        .replace(JAVA_TEMPLATE_REPLACE_WITH_PREDICTOR_CLASS_NAME, javaClassName);
    if (replaceTransform != null)
      java = java.replace(JAVA_TEMPLATE_REPLACE_WITH_TRANSFORMER_OBJECT, replaceTransform);
    if (modelCode != null)
      java = java.replace(REPLACE_THIS_WITH_MODEL, modelCode);
    FileUtils.writeStringToFile(new File(tmpDir, resultFileName), java);
  }

}
