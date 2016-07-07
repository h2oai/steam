import java.util.Map;

import hex.genmodel.easy.*;

public interface Transform {
    /**
     *
     * @param input is the original data to be transformed
     * @return an array of
     */
    //    Object[] fit(byte[] input);
    //Map<String, Object> fit(byte[] input);
    RowData fit(byte[] input);
}

