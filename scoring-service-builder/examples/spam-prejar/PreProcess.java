import java.util.Map;
import java.util.Set;
import java.util.HashMap;
import java.util.HashSet;
import java.util.Arrays;
import java.util.ArrayList;

import java.io.BufferedReader;
import java.io.FileReader;
import java.io.ObjectOutputStream;
import java.io.FileOutputStream;
import java.io.ObjectInputStream;
import java.io.FileInputStream;
import java.io.IOException;

import hex.genmodel.easy.*;

public class PreProcess implements Transform {

    private static Tfidf tfidf = null;

    public RowData fit(String input) {
	if (tfidf == null) {
	    tfidf = new Tfidf();
	    tfidf.loadModel("tfidf.bin");
	}
	RowData res = new RowData();
	String line = input;
	Map<String, Double> map = tfidf.tfidfMap(tokenize(line));
	for (String x: map.keySet()) {
	    double val = map.get(x);
	    res.put(x, val);
	}
	return res;
    }

    private static String[] tokenize(String text) {
	return text.toLowerCase().replaceAll("\\d+", " ").replaceAll("[^\\w ]" ," ").replaceAll("\\s+", " ").trim().split(" ");
    }
    
    public static void main(String[] args) throws Exception {
			
	ArrayList<String[]> vecs = new ArrayList<String[]>();
	ArrayList<String> labels = new ArrayList<String>();
			
	String trainingFile = "smsData.txt";//args[0];

	tfidf = new Tfidf();
	
	BufferedReader br = new BufferedReader(new FileReader(trainingFile));
	try {
	    String line = br.readLine();
	    while (line != null) {
		String a[] = line.toLowerCase().split("\t");
		String label = a[0];
		String text = a[1];
		String[] b = tokenize(text);
		if (label != null && b.length > 0 && b != null && b.length > 0) {
		    labels.add(label);
		    tfidf.addDoc(b);
		    vecs.add(b);
		}

		line = br.readLine();
	    }
	}
	finally {
	    br.close();
	}

	tfidf.saveModel("tfidf.bin");
	
	// bulk convert
	// header first
	String[] words = tfidf.words();
	System.out.print("LABEL");
	for (String w: words) {
	    System.out.print("," + w);
	}
	System.out.println();
	// row by row
	int i = 0;
	for (i = 0; i < vecs.size(); i++) {
	    String[] v = vecs.get(i);
	    String label = labels.get(i);
	    System.out.print(label);
	    double[] t = tfidf.tfidf(v);
	    int len = t.length;
	    for (double x : t) {
		System.out.print("," + x);
	    }
	    System.out.println();
	}
	System.err.println("converted vectors " + i);
	
	br.close();
    }

}
