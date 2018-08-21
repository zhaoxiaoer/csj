package c6;

import java.io.IOException;
import java.io.PrintWriter;

import javax.servlet.ServletException;
import javax.servlet.http.*;

import util.SearchSpec;
import util.SearchUtilities;
import util.ServletUtilities;

public class SearchEngineForm extends HttpServlet {
	public void doGet(HttpServletRequest req, HttpServletResponse res) 
			throws ServletException, IOException {
		res.setContentType("text/html");
		PrintWriter out = res.getWriter();
		String title = "One-Stop Web Search!";
		String actionURL = "/csj/c6/searchengines";
		out.println(ServletUtilities.headWithTitle(title) +
				"<body bgcolor=\"#FDF5E6\">\n" +
				"<center>\n" +
				"<h1>" + title + "</h1>\n" +
				"<form action=\"" + actionURL + "\">\n" +
				"  Search keywords: \n" +
				"  <input type=\"text\" name=\"searchString\"><p>\n");
		SearchSpec[] specs = SearchUtilities.getCommonSpecs();
		for (int i = 0; i < specs.length; i++) {
			String searchEngineName = specs[i].getName();
			out.println("<input type=\"radio\" " + 
					"name=\"searchEngine\" " +
					"value=\"" + searchEngineName + "\">\n");
			out.println(searchEngineName + "<br>\n");
		}
		out.println("<br>  <input type=\"submit\">\n" +
				"</form>\n" +
				"</center></body></html>");
	}
}
