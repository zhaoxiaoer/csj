package c8;

import java.io.*;

import javax.servlet.*;
import javax.servlet.http.*;

import util.ServletUtilities;

public class CookieTest extends HttpServlet {
	public void doGet(HttpServletRequest req, HttpServletResponse res) 
			throws ServletException, IOException {
		for (int i = 0; i < 3; i++) {
			Cookie cookie = new Cookie("Session-Cookie-" + i, "Cookie-Valie-S" + i);
			res.addCookie(cookie);
			cookie = new Cookie("Persistent-Cookie-" + i, "Cookie-Value-P" + i);
			cookie.setMaxAge(3600);
			res.addCookie(cookie);
		}
		res.setContentType("text/html");
		PrintWriter out = res.getWriter();
		String title = "Active Cookies";
		out.println(ServletUtilities.headWithTitle(title) +
				"<body bgcolor=\"#FDF5E6\">\n" +
				"<h1 align=\"center\">" + title + "</h1>" +
				"<table border=1 align=\"center\">\n" +
				"<tr bgcolor=\"#FFAD00\">\n" +
				"  <th>Cookie Name</th>\n" + 
				"  <th>Cookie Value</th>");
		Cookie[] cookies = req.getCookies();
		if (cookies == null) {
			out.println("<tr><th colspan=2>No cookies</th></tr>");
		} else {
			Cookie cookie;
			for (int i = 0; i < cookies.length; i++) {
				cookie = cookies[i];
				out.println("<tr>\n" +
						"  <td>" + cookie.getName() + "</td>\n" +
						"  <td>" + cookie.getValue() + "</td>\n" +
						"</tr>");
			}
		}
		out.println("</table></body></html>");
	}
}
