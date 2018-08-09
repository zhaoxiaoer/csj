package c4;

import java.io.IOException;
import java.io.PrintWriter;
import java.util.Enumeration;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import util.ServletUtilities;

public class ShowParameters extends HttpServlet {
	public void init() throws ServletException {
		
	}
	
	public void doGet(HttpServletRequest req, HttpServletResponse res) 
			throws ServletException, IOException {
		res.setContentType("text/html");
		PrintWriter out = res.getWriter();
		String title = "Reading All Request Parameters";
		out.println(ServletUtilities.headWithTitle(title) +
				"<body bgcolor=\"#FDF5E6\">\n" +
				"<h1 align=center>" + title + "</h1>\n" +
				"<table border=1 align=center>\n" +
				"<tr bgcolor=\"#FFAD00\">" +
				"<th>Parameter Name</th><th>Parameter Value(s)</th></tr>");
		Enumeration parameters = req.getParameterNames();
		while (parameters.hasMoreElements()) {
			String paramName = (String) parameters.nextElement();
			out.print("<tr><td>" + paramName + "</td><td>");
			String[] paramValues = req.getParameterValues(paramName);
			if (paramValues.length == 1) {
				String paramValue = paramValues[0];
				if (paramValue.length() == 0) {
					out.println("<i>No Value</i>");
				} else {
					out.println(paramValue);
				}
			} else {
				out.println("<ul>");
				for (int i = 0; i < paramValues.length; i++) {
					out.println("<li>" + paramValues[i] + "</li>");
				}
				out.println("</ul>");
			}
			out.println("</td>");
		}
		out.println("</table>\n</body>\n</html>");
	}
	
	public void doPost(HttpServletRequest req, HttpServletResponse res) 
			throws ServletException, IOException {
		doGet(req, res);
	}
	
	public long getLastModified(HttpServletRequest req) {
		return -1;
	}
	
	public void destroy() {
		
	}
}
