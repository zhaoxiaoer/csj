package c8;

import java.io.*;
import javax.servlet.*;
import javax.servlet.http.*;

import util.CookieUtilities;
import util.ServletUtilities;

public class RegistrationForm extends HttpServlet {
	public void doGet(HttpServletRequest req, HttpServletResponse res) 
			throws ServletException, IOException {
		res.setContentType("text/html");
		PrintWriter out = res.getWriter();
		String actionURL = "/csj/c8/registrationservlet";
		String firstName = CookieUtilities.getCookieValue(req, "firstName", "");
		String lastName = CookieUtilities.getCookieValue(req, "lastName", "");
		String emailAddress = CookieUtilities.getCookieValue(req, "emailAddress", "");
		String title = "Please Register";
		out.println(ServletUtilities.headWithTitle(title) +
				"<body bgcolor=\"#FDF5E6\">\n" +
				"<center>\n" +
				"<h1>" + title + "</h1>" +
				"<form action=\"" + actionURL + "\">\n" +
				"First Name:\n" +
				"  <input type=\"text\" name=\"firstName\" " +
						"value=\"" + firstName + "\"><br>\n" +
				"Last Name:\n" +
				"  <input type=\"text\" name=\"lastName\" " +
						"value=\"" + lastName + "\"><br>\n" +
				"Email Address: \n" +
				"  <input type=\"text\" name=\"emailAddress\" " +
						"value=\"" + emailAddress + "\"><br>\n" +
				"<input type=\"submit\" value=\"Register\">\n" +
				"</form></center></body></html>");
	}
}
