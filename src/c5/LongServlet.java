package c5;

import java.io.IOException;
import java.io.PrintWriter;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import util.GzipUtilities;
import util.ServletUtilities;

public class LongServlet extends HttpServlet {
	public void doGet(HttpServletRequest req, HttpServletResponse res) 
			throws ServletException, IOException {
		res.setContentType("text/html");
		PrintWriter out;
		if (GzipUtilities.isGzipSupported(req) &&
				!GzipUtilities.isGzipDisabled(req)) {
			out = GzipUtilities.getGzipWriter(res);
			res.setHeader("Content-Encoding", "gzip");
		} else {
			out = res.getWriter();
		}
		String title = "Long Page";
		out.println(ServletUtilities.headWithTitle(title) +
				"<body bgcolor=\"#FDF5E6\">\n" +
				"<h1 align=\"center\">" + title + "</h1>\n");
		String line = "Blah, blah, blah, blah, blah. Yadda, yadda, yadda, yadda.";
		for (int i = 0; i < 10000; i++) {
			out.println(line);
		}
		out.println("</body></html>");
		out.close(); // Needed for gzip; optional otherwise.
	}
}
