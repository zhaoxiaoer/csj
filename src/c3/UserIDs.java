package c3;

import java.io.IOException;
import java.io.PrintWriter;

import javax.servlet.ServletConfig;
import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import util.ServletUtilities;

public class UserIDs extends HttpServlet {
	private int nextID = 0;
	
	public void init() throws ServletException {
//		ServletConfig cfg = getServletConfig();
	}
	
	public void doGet(HttpServletRequest req, HttpServletResponse res)
			throws ServletException, IOException {
		res.setContentType("text/html");
		PrintWriter out = res.getWriter();
		String title = "Your ID";
		out.println(ServletUtilities.DOCTYPE +
				ServletUtilities.headWithTitle(title) +
				"<body bgcolor=\"#FDF5E6\">\n" +
				"<h1 align=center>" + title + "</h1>\n");
		synchronized(this) {
			String id = "User-ID-" + nextID;
			out.println("<h2>" + id + "</h2>\n");
			nextID++;
		}
		out.println("</body></html>");
	}
	
//	public long getLastModified(HttpServletRequest req) {
//		return 0;
//	}
	
	public void destroy() {
		
	}
}
