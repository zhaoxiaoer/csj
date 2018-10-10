package v2c2;

import java.io.*;
import javax.servlet.*;
import javax.servlet.http.*;

public class TestServlet1 extends HttpServlet {
	private String firstName = "First name is missing.";
	private String supportEmail = "Support Email is missing";
	
	public void init() {
		ServletConfig config = getServletConfig();
		if (config.getInitParameter("firstName") != null) {
			firstName = config.getInitParameter("firstName");
//			System.out.println("firstName: " + firstName);
		}
		
		ServletContext context = getServletContext();
		if (context.getInitParameter("support-email") != null) {
			supportEmail = context.getInitParameter("support-email");
		}
	}
	
	public void doGet(HttpServletRequest request, 
					HttpServletResponse response) 
					throws ServletException, IOException {
		response.setContentType("text/html");
		PrintWriter out = response.getWriter();
		String uri = request.getRequestURI();
		out.println("<!DOCTYPE html>\n" +
				"<html>\n" +
				"<head>\n" +
				"<title>" + "Test Servlet 1" + "</title>\n" +
				"</head>\n" +
				"<body bgcolor=\"#FDF5E6\">\n" +
				"<h2>" + "Servlet Name: Test1" + "</h2>\n" +
				"<h2>URI: " + uri + "</h2>\n" +
				"<h2>firstName: " + firstName + "</h2>\n" +
				"<h2>Support Email: " + supportEmail + "</h2>\n" +
				"</body>\n" +
				"</html>");
	}
}
