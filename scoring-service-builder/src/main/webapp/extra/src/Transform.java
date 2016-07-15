import java.util.Map;
import hex.genmodel.easy.*;

public interface Transform {
    /**
     *
     * RowData is a Map<String, Object> where
     * String is the column name
     * Object is the data value in that column Integer/Double/String
     * defined in h2o-genmodel.jar
     *
     * @param input is one incoming row of data to be transformed
     * @return the resulting one row of data
     */
    RowData fit(byte[] input);
}

