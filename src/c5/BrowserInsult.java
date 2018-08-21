package c5;

import java.io.IOException;
import java.io.PrintWriter;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import util.ServletUtilities;

public class BrowserInsult extends HttpServlet {
	public void doGet(HttpServletRequest req, HttpServletResponse res) 
			throws ServletException, IOException {
		res.setContentType("text/html");
		PrintWriter out = res.getWriter();
		String title, message;
		String userAgent = req.getHeader("User-Agent");
		if ((userAgent != null) && (userAgent.indexOf("MSIE") != -1)) {
			title = "Microsoft Minion";
			message = "Welcome, O spineless slave to the mighty empire.";
		} else {
			title = "Hopeless Netscape Rebel";
			message = "Enjoy it while you can. You <i>will</i> be assimilated!";
		}
		out.println(ServletUtilities.headWithTitle(title) +
				"<body bgcolor=\"#FDF5E6\">\n" +
				"<h1 align=\"center\">" + title + "</h1>\n" +
				message + "\n" +
				"</body></html>");
	}
}
