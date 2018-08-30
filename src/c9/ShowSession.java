package c9;

import java.io.*;
import java.util.Date;

import javax.servlet.*;
import javax.servlet.http.*;

import util.ServletUtilities;

public class ShowSession extends HttpServlet {
	public void doGet(HttpServletRequest req, HttpServletResponse res) 
			throws ServletException, IOException {
		res.setContentType("text/html");
		HttpSession session = req.getSession();
		String heading;
		Integer accessCount = (Integer) session.getAttribute("accessCount");
		if (accessCount == null) {
			accessCount = new Integer(0);
			heading = "Welcome, Newcomer";
		} else {
			heading = "Welcome Back";
			accessCount = new Integer(accessCount.intValue() + 1);
		}
//		System.out.println(session.getMaxInactiveInterval());
		session.setAttribute("accessCount", accessCount);
		PrintWriter out = res.getWriter();
		String title = "Session Tracking Example";
		out.println(ServletUtilities.headWithTitle(title) +
				"<body bgcolor=\"#FDF5E6\">\n" +
				"<center>\n" +
				"<h1>" + heading + "</h1>\n" + 
				"<h2>Information on Your Session:</h2>\n" +
				"<table border=1>\n" +
				"<tr bgcolor=\"#FFAD00\">\n" +
				"  <th>Info Type</th><th>Value</th>\n" +
				"</tr>\n" +
				"<tr>\n" +
				"  <td>ID</td><td>" + session.getId() + "</td>\n" +
				"</tr>\n" +
				"<tr>\n" +
				"  <td>Creation Time</td><td>" + new Date(session.getCreationTime()) + "</td>\n" +
				"</tr>\n" +
				"<tr>\n" +
				"  <td>Time of Last Access</td><td>" + new Date(session.getLastAccessedTime()) + "</td>\n" +
				"</tr>\n" +
				"<tr>\n" +
				"  <td>Number of Previous Accesses</td><td>" + accessCount + "</td>\n" +
				"</tr>\n" +
				"</table>\n" +
				"</center></body></html>");
	}
}
