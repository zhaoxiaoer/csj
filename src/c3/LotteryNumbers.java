package c3;

import java.io.IOException;
import java.io.PrintWriter;

import javax.servlet.ServletConfig;
import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import util.ServletUtilities;

/**
 * 
 * Example using servlet initialization and the
 * getLastModified method.
 *
 */

public class LotteryNumbers extends HttpServlet {
	private long modTime;
	private int[] numbers = new int[10];
	
	/**
	 * The init method is called only when the servlet
	 * is first loaded, before the first request is processed.
	 */
	public void init() throws ServletException {
		ServletConfig cfg = getServletConfig();
		String addr = cfg.getInitParameter("address");
		System.out.println("addr: " + addr);
		
		// Round to nearest second (i.e, 1000 milliseconds)
		modTime = System.currentTimeMillis() / 1000 * 1000;
		for (int i = 0; i < numbers.length; i++) {
			numbers[i] = randomNum();
		}
	}
	
	// Return the list of numbers that init computed.
	public void doGet(HttpServletRequest req, HttpServletResponse res) throws ServletException, IOException {
		System.out.println("11111111");
		res.setContentType("text/html");
		PrintWriter out = res.getWriter();
		String title = "Your Lottery Numbers";
		out.println(ServletUtilities.DOCTYPE +
				ServletUtilities.headWithTitle(title) +
				"<body bgcolor=\"#FDF5E6\"\n" +
				"<h1 align=center>" + title + "</h1><br />\n" +
				"<b>Based upon extensive research of " +
				"astro-illogical trends, psychic farces, " +
				"and detailed statistical claptrap, " +
				"we have chosen the " + numbers.length +
				" best lottery numbers for you.</b>\n" +
				"<ol>");
		for (int i = 0; i < numbers.length; i++) {
			out.println("  <li>" + numbers[i] + "</li>");
		}
		out.println("</ol>\n</body>\n</html>\n");
	}
	
	public long getLastModified(HttpServletRequest req) {
		System.out.println("222222");
		return modTime;
	}
	
	public void destroy() {
		System.out.println("LotteryNumber destroy");
	}
	
	// A random int from 0 to 99
	private int randomNum() {
		return (int) (Math.random() * 100);
	}
}
