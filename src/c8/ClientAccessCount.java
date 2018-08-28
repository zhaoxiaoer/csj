package c8;

import java.io.*;

import javax.servlet.*;
import javax.servlet.http.*;

import util.*;

public class ClientAccessCount extends HttpServlet {
	public void doGet(HttpServletRequest req, HttpServletResponse res) 
			throws ServletException, IOException {
		String countString = CookieUtilities.getCookieValue(req, "accessCount", "1");
		int count = 1;
		try {
			count = Integer.parseInt(countString);
		} catch (NumberFormatException nfe) {
		}
		Cookie cookie = new Cookie("accessCount", String.valueOf(count + 1));
		cookie.setMaxAge(60 * 60 * 24 * 365);
		res.addCookie(cookie);
		res.setContentType("text/html");
		PrintWriter out = res.getWriter();
		String title = "Access Count Servlet";
		out.println(ServletUtilities.headWithTitle(title) +
				"<body bgcolor=\"#FDF5E6\">\n" + 
				"<center>\n" + 
				"<h1>" + title + "</h1>\n" +
				"<h2>This is visit number " + 
				count + " by this browser.</h2>\n" +
				"</center></body></html>");
	}
}
