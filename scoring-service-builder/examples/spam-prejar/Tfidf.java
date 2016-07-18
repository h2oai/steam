import java.util.Map;
import java.util.HashMap;
import java.util.Arrays;
import java.io.Serializable;
import java.io.ObjectOutputStream;
import java.io.FileOutputStream;
import java.io.ObjectInputStream;
import java.io.FileInputStream;
import java.io.IOException;
import java.io.BufferedReader;
import java.io.PrintWriter;
import java.io.FileReader;

public class Tfidf implements Serializable {
    Map<String, Integer> df = new HashMap<String, Integer>();
    int totalN = 0;

    public void saveModel(String filename) {
	System.err.println("saving to " + filename + "  totalN " + totalN);
	PrintWriter fout = null;
	try{
	    fout = new PrintWriter(filename);
	    fout.println(totalN);
	    fout.println(df.keySet().size());
	    for (String k: df.keySet()) {
		fout.println(k + "\t" + df.get(k));
	    }
	} catch (Exception e) {
	    e.printStackTrace();
	} finally {
	    if (fout != null){
		fout.close();
	    } 
	}
	System.err.println("tfidf saved to " + filename);
    }

    public void loadModel(String filename) {
	System.err.println("tfidf loading from " + filename);
	BufferedReader streamIn = null;
	try {
	    streamIn = new BufferedReader(new FileReader(filename));
	    totalN = Integer.parseInt(streamIn.readLine());
	    int n = Integer.parseInt(streamIn.readLine());
	    System.err.println("totalN " + totalN + "  size " + n);
	    for (int i = 0; i < n; i++) {
		String line = streamIn.readLine();
		String[] a = line.split("\t");
		String k = a[0];
		int v = Integer.parseInt(a[1]);
		df.put(k ,v);
	    }
	} catch (Exception e) {
	    e.printStackTrace();
	} finally {
	    if (streamIn != null){
		try {
		    streamIn.close();
		}
		catch (IOException e) {
		   e.printStackTrace(); 
		}
	    } 
	}
	System.err.println("tfidf loaded from " + filename);
    }

    public void addDoc(String[] terms) {
	Map<String, Integer> tf = termFrequency(terms);
	for (String term: tf.keySet()) {
	    if (term != null && term.length() > 0) {
		Integer c = df.get(term);
		if (c == null)
		    df.put(term, 1);
		else
		    df.put(term, c + 1);
	    }
	}
	totalN += 1;
    }

    public Map<String, Integer> termFrequency(String[] terms) {
	Map<String, Integer> tf = new HashMap<String, Integer>();
	for (String term : terms) {
	    Integer c = tf.get(term);
	    if (c == null)
		tf.put(term, 1);
	    else
		tf.put(term, c + 1);
	}
	//	    System.out.println(tf);//
	return tf;
    }

    public String[] words() {
	String[] index =  new String[df.keySet().size()];
	df.keySet().toArray(index);
	Arrays.sort(index);
	return index;
    }

    public double[] tfidf(String[] terms) {
	double[] vec = new double[df.keySet().size()];
	String[] index = words();
	Map<String, Integer> ix = new HashMap<String, Integer>();
	int p = 0;
	for (String x: index) {
	    ix.put(x, p);
	    p += 1;
	}
	Map<String, Double> tvals = tfidfMap(terms);
	for (String t: tvals.keySet()) {
	    double tfidf = tvals.get(t);
	    if (ix.containsKey(t)) {
		int ind = ix.get(t);
		vec[ind] = tfidf;
	    }
	}
	return vec;
    }

    public Map<String, Double> tfidfMap(String[] terms) {
	Map<String, Integer> tf = termFrequency(terms);
	double[] vec = new double[df.keySet().size()];
	String[] index =  new String[df.keySet().size()];
	df.keySet().toArray(index);
	Map<String, Double> tvals = new HashMap<String, Double>();
	Map<String, Integer> ix = new HashMap<String, Integer>();
	double s = 0.0;
	for (String t: tf.keySet()) {
	    int docf = df.containsKey(t) ? df.get(t) : 0;
	    double termf = (double) tf.get(t) / tf.size();
	    double idf = totalN > 0 ? (double) (totalN + 1)/ (docf + 1) : 0.0;
	    double tfidf = termf * Math.log(1 + idf);
	    tvals.put(t, tfidf);
	    s += tfidf*tfidf;
	}
	// normalize
	s = Math.sqrt(s);
	for (String k: tvals.keySet()) {
	    tvals.put(k, tvals.get(k) / s);
	}
	return tvals;
    }

}
