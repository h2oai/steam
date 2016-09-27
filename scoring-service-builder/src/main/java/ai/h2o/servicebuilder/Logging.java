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
