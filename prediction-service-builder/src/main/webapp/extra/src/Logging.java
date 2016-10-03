import org.slf4j.LoggerFactory;

public class Logging {

  public static void initLogger() {
    // show date and time
    if (System.getProperty("org.slf4j.simpleLogger.showDateTime") == null)
      System.setProperty("org.slf4j.simpleLogger.showDateTime", "true");

    // date and time format
    if (System.getProperty("org.slf4j.simpleLogger.dateTimeFormat") == null)
      System.setProperty("org.slf4j.simpleLogger.dateTimeFormat", "yyyy-MM-dd HH:mm:ss.SSS Z");
  }

  public static org.slf4j.Logger getLogger(Class theClass) {
    initLogger();
    String className = theClass.getSimpleName();
    return LoggerFactory.getLogger(className);
  }

}
