package c6;

import java.io.IOException;
import java.net.URLEncoder;

import javax.servlet.*;
import javax.servlet.http.*;

import util.SearchUtilities;

public class SearchEngines extends HttpServlet {
	public void doGet(HttpServletRequest req, HttpServletResponse res) 
			throws ServletException, IOException {
		String searchString = req.getParameter("searchString");
		if ((searchString == null) || (searchString.length() == 0)) {
			reportProblem(res, "Missing search string");
			return;
		}
		
		searchString = URLEncoder.encode(searchString);
		String searchEngineName = req.getParameter("searchEngine");
		if ((searchEngineName == null) || (searchEngineName.length() == 0)) {
			reportProblem(res, "Missing search engine name");
			return;
		}
		String searchURL = SearchUtilities.makeURL(searchEngineName, searchString);
		if (searchURL != null) {
			res.sendRedirect(searchURL);
		} else {
			reportProblem(res, "Unrecognized search engine");
		}
	}
	
	private void reportProblem(HttpServletResponse res, String msg) 
			throws IOException {
		res.sendError(res.SC_NOT_FOUND, msg);
	}
}
