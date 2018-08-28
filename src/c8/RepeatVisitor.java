package c8;

import java.io.IOException;
import java.io.PrintWriter;

import javax.servlet.*;
import javax.servlet.http.*;

import util.ServletUtilities;

public class RepeatVisitor extends HttpServlet {
	public void doGet(HttpServletRequest req, HttpServletResponse res) 
			throws ServletException, IOException {
		boolean newbie = true;
		Cookie[] cookies = req.getCookies();
		if (cookies != null) {
			for (int i = 0; i < cookies.length; i++) {
				Cookie c = cookies[i];
				if ((c.getName().equals("repeatVisitor")) 
						&& (c.getValue().equals("yes"))) {
					newbie = false;
					break;
				}
			}
		}
		String title;
		if (newbie) {
			Cookie returnVisitorCookie = new Cookie("repeatVisitor", "yes");
			returnVisitorCookie.setMaxAge(60 * 60 * 24 * 365);
			res.addCookie(returnVisitorCookie);
			title = "Welcome Aboart";
		} else {
			title = "Welcome Back";
		}
		res.setContentType("text/html");
		PrintWriter out = res.getWriter();
		out.println(ServletUtilities.headWithTitle(title) + 
				"<body bgcolor=\"#FDF5E6\">\n" +
				"<h1 align=\"center\">" + title + "</h1>\n" +
				"</body></html>");
	}
}
