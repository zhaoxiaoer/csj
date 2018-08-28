package c8;

import java.io.*;
import javax.servlet.*;
import javax.servlet.http.*;

import util.ServletUtilities;

public class RegistrationServlet extends HttpServlet {
	public void doGet(HttpServletRequest req, HttpServletResponse res) 
			throws ServletException, IOException {
		res.setContentType("text/html");
		boolean isMissingValue = false;
		String firstName = req.getParameter("firstName");
		if (isMissing(firstName)) {
			firstName = "Missing_first_name";
			isMissingValue = true;
		}
		String lastName = req.getParameter("lastName");
		if (isMissing(lastName)) {
			lastName = "Missing_last_name";
			isMissingValue = true;
		}
		String emailAddress = req.getParameter("emailAddress");
		if (isMissing(emailAddress)) {
			emailAddress = "Missing_email_address";
			isMissingValue = true;
		}
		Cookie cookie = new Cookie("firstName", firstName);
		cookie.setMaxAge(60 * 60 * 24 * 365);
		res.addCookie(cookie);
		Cookie cookie2 = new Cookie("lastName", lastName);
		cookie2.setMaxAge(60 * 60 * 24 * 365);
		res.addCookie(cookie2);
		Cookie cookie3 = new Cookie("emailAddress", emailAddress);
		cookie3.setMaxAge(60 * 60 * 24 * 365);
		res.addCookie(cookie3);
		String formAddress = "/csj/c8/registrationform";
		if (isMissingValue) {
			res.sendRedirect(formAddress);
		} else {
			PrintWriter out = res.getWriter();
			String title = "Thanks for Registering";
			out.println(ServletUtilities.headWithTitle(title) +
					"<body bgcolor=\"#FDF5E6\">\n" +
					"<center>\n" +
					"<h1>" + title + "</h1>\n" + 
					"<ul>\n" + 
					"  <li><b>First Name</b>: " + firstName + "</li>\n" + 
					"  <li><b>Last Name</b>: " + lastName + "</li>\n" +
					"  <li><b>Email address</b>: " + emailAddress + "</li>\n" +
					"</ul>\n" + 
					"</center></body></html>");
		}
	}
	
	private boolean isMissing(String param) {
		return ((param == null) || (param.trim().equals("")));
	}
}
