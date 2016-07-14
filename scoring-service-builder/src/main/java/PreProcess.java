//import java.util.Map;
//import java.util.HashMap;

import java.io.BufferedReader;
import java.io.FileReader;

import hex.genmodel.easy.*;

//import org.deeplearning4j.vectorizer.TfidfVectorizer;
import org.canova.nd4j.nlp.vectorizer.TfidfVectorizer;

public class PreProcess implements Transform {

	class TfiDf {
		private Map<String, Int> df = new HashMap<String, Int>();

		public void add(String[] terms) {
			for (String term: terms) {
				if (!df.containsKey(term))
					df[term] = 0;
				df[term] += 1;
			}
		}

		public void add(String[] terms) {
			for (String term: terms) {
				if (!df.containsKey(term))
					df[term] = 0;
				df[term] += 1;
			}
		}

		public Float[] tfidf(String[] terms) {
			private Map<String, Int> tf = new HashMap<String, Int>();

			String[] keys = df.getKeys();
			for (String term: terms) {
				if (!tf.containsKey(term))
					tf[term] = 0;
				tf[term] += 1;
			}


		}


	}

    public RowData fit(byte[] input) {
	//Map<String, Object> res = new HashMap<String, Object>();
	RowData res = new RowData();
	String s = new String(input);
	String[] a = s.split(" ");
	

	res.put("Dest",a[0]);
	// res.put("dist", 3.14);
	return res;
    }

	public static void main(String[] args) throws Exception {
		String trainingFile = args[0];

		TfidfVectorizer vectorizer = new TfidfVectorizer();

		BufferedReader br = new BufferedReader(new FileReader(trainingFile));
		try {
			String line = br.readLine();

			while (line != null) {
				String a[] = line.split("\t");
				String label = a[0];
				String text = a[1];



				line = br.readLine();
			}
		} finally {
			br.close();
		}
	}

}
