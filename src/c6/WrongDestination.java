package c6;

import javax.servlet.*;
import javax.servlet.http.*;
import java.io.IOException;

public class WrongDestination extends HttpServlet {
	public void doGet(HttpServletRequest req, HttpServletResponse res) 
			throws ServletException, IOException {
		String userAgent = req.getHeader("User-Agent");
		if ((userAgent != null) && (userAgent.indexOf("MSIE") != -1)) {
			res.sendRedirect("http://home.netscape.com");
		} else {
			res.sendRedirect("http://www.microsoft.com");
		}
	}
}
