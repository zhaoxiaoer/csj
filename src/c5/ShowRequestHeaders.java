package c5;

import java.io.IOException;
import java.io.PrintWriter;
import java.util.Collection;
import java.util.Enumeration;
import java.util.Iterator;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import util.ServletUtilities;

public class ShowRequestHeaders extends HttpServlet {
	public void doGet(HttpServletRequest req, HttpServletResponse res) 
			throws ServletException, IOException {
		res.setContentType("text/html");
		PrintWriter out = res.getWriter();
		String title = "Servlet Example: Showing Request Headers";
		out.println(ServletUtilities.headWithTitle(title) +
				"<body bgcolor=\"#FDF5E6\">\n" +
				"<h1 align=\"center\">" + title + "</h1>\n" +
				"<b>Request Method: </b>" + req.getMethod() + "<br>\n" +
				"<b>Request URI: </b>" + req.getRequestURI() + "<br>\n" +
				"<b>Request Protocol: </b>" + req.getProtocol() + "</br>\n" +
				"<table border=1 align=\"center\">\n" +
				"<tr bgcolor=\"#FFAD00\">\n" +
				"  <th>Header Name</th>\n" +
				"  <th>Header Value</th>\n" +
				"</tr>");
		Enumeration headerNames = req.getHeaderNames();
		while (headerNames.hasMoreElements()) {
			String headerName = (String) headerNames.nextElement();
			out.println("<tr>\n  <td>" + headerName + "</td>");
			out.println("  <td>" + req.getHeader(headerName) + "</td>\n</tr>");
		}
		out.println("</table>\n");
		
		out.println("<br><br><br><table border=1 align=\"center\">\n" +
				"<tr bgcolor=\"#FFAD00\">\n" +
				"  <th>Header Name</th>\n" +
				"  <th>Header Value</th>\n" +
				"</tr>");
		res.setHeader("Name", "zhaoxiaoer");
		Collection headerNames2 = res.getHeaderNames();
		System.out.println("" + headerNames2.size());
		Iterator it = headerNames2.iterator();
		while (it.hasNext()) {
			String headerName = (String) it.next();
			out.println("<tr>\n  <td>" + headerName + "</td>");
			out.println("  <td>" + res.getHeader(headerName) + "</td>\n</tr>");
		}
		out.println("</table>\n");
		System.out.println("Content-Type: " + res.getContentType());
		
		out.println("</body></html>");
	}
	
	public void doPost(HttpServletRequest req, HttpServletResponse res) 
			throws ServletException, IOException {
		doGet(req, res);
	}
}
