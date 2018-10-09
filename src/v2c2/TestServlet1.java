package v2c2;

import java.io.*;
import javax.servlet.*;
import javax.servlet.http.*;

public class TestServlet1 extends HttpServlet {
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
				"</body>\n" +
				"</html>");
	}
}
