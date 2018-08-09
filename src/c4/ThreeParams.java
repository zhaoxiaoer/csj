package c4;

import java.io.IOException;
import java.io.PrintWriter;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import util.ServletUtilities;

public class ThreeParams extends HttpServlet {
	public void init() throws ServletException {
		
	}
	
	// curl -v "http://127.0.0.1:8080/csj/c4/threeparams?param1=aaa&param2=bbb&param3=ccc"
	public void doGet(HttpServletRequest req, HttpServletResponse res) 
			throws ServletException, IOException {
		res.setContentType("text/html");
		PrintWriter out = res.getWriter();
		String title = "Reading Three Request Parameters";
		out.println(ServletUtilities.headWithTitle(title) +
				"<body bgcolor=\"#FDF5E6\">\n" +
				"<h1 align=center>" + title + "</h1>\n" +
				"<ul>\n" +
				"  <li><b>param1</b>: " +
				req.getParameter("param1") + "\n" +
				"  <li><b>param2</b>: " +
				req.getParameter("param2") + "\n" +
				"  <li><b>param3</b>: " +
				req.getParameter("param3") + "\n" +
				"</ul>\n" +
				"</body></html>");
	}
	
	// curl -v -X POST --data "param1=aaa&param2=bbb&param3=ccc" "http://127.0.0.1:8080/csj/c4/threeparams"
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
