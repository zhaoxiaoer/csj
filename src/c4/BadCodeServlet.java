package c4;

import java.io.IOException;
import java.io.PrintWriter;

import javax.servlet.ServletConfig;
import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import util.ServletUtilities;

public class BadCodeServlet extends HttpServlet {
//	public void init() throws ServletException {
//		ServletConfig cfg = getServletConfig();
//	}
	
	public void doGet(HttpServletRequest req, HttpServletResponse res) 
			throws ServletException, IOException {
		res.setContentType("text/html");
		PrintWriter out = res.getWriter();
		String title = "Code Sample";
		out.println(ServletUtilities.headWithTitle(title) +
				"<body bgcolor=\"#FDF5E6\">\n" +
				"<h1 align=\"center\">" + title + "</h1>\n" +
				"<pre>\n" +
				getCode(req) +
				"</pre>\n" +
				"Now, wasn't that an interesting sample\n" +
				"of code?\n" +
				"</body></html>");
	}
	
//	public long getLastModified(HttpServletRequest req) {
//		return -1;
//	}
//	
//	public void destroy() {
//		
//	}
	
	private String getCode(HttpServletRequest req) {
		return ServletUtilities.filter(req.getParameter("code"));
	}
}
