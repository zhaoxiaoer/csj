package c7;

import java.io.IOException;
import java.io.PrintWriter;
import java.util.ArrayList;

import javax.servlet.ServletException;
import javax.servlet.http.*;

import util.ServletUtilities;

public class PrimeNumberServlet extends HttpServlet {
	private ArrayList primeListCollection = new ArrayList();
	private int maxPrimeLists = 30;
	
	public void doGet(HttpServletRequest req, HttpServletResponse res)
			throws ServletException, IOException {
		int numPrimes = ServletUtilities.getIntParameter(req, "numPrimes", 50);
		int numDigits = ServletUtilities.getIntParameter(req, "numDigits", 120);
		
		PrimeList primeList = findPrimeList(primeListCollection, numPrimes, numDigits);
		if (primeList == null) {
			primeList = new PrimeList(numPrimes, numDigits, true);
			synchronized(primeListCollection) {
				if (primeListCollection.size() >= maxPrimeLists)
					primeListCollection.remove(0);
				primeListCollection.add(primeList);
			}
		}
		ArrayList currentPrimes = primeList.getPrimes();
		int numCurrentPrimes = currentPrimes.size();
		int numPrimesRemaining = numPrimes - numCurrentPrimes;
		boolean isLastResult = (numPrimesRemaining == 0);
		if (!isLastResult) {
			res.setIntHeader("Refresh",	5);
		}
		res.setContentType("text/html");
		PrintWriter out = res.getWriter();
		String title = "Some " + numDigits + "-Digit Prime Numbers";
		out.println(ServletUtilities.headWithTitle(title) +
				"<body bgcolor=\"#FDF5E6\">\n" + 
				"<h2 align=\"center\">" + title + "</h2>\n" +
				"<h3>Primes found with " + numDigits +
				" or more digits: " + numCurrentPrimes + ".</h3>");
		if (isLastResult)
			out.println("<b>Done searching.</b>");
		else
			out.println("<b>Still looking for " + numPrimesRemaining +
					" more<blink>...</blink></b>");
		out.println("<ol>");
		for (int i = 0; i < numCurrentPrimes; i++) {
			out.println("  <li>" + currentPrimes.get(i));
		}
		out.println("</ol>");
		out.println("</body></html>");
	}
	
	private PrimeList findPrimeList(ArrayList primeListCollection, 
			int numPrimes, int numDigits) {
		synchronized(primeListCollection) {
			for (int i = 0; i < primeListCollection.size(); i++) {
				PrimeList primes = (PrimeList) primeListCollection.get(i);
				if ((numPrimes == primes.numPrimes()) && (numDigits == primes.numDigits())) {
					return primes;
				}
			}
			return null;
		}
	}
}
