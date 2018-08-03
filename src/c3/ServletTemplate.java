package c3;

import java.io.IOException;
import java.io.PrintWriter;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

public class ServletTemplate extends HttpServlet {
	public void doGet(HttpServletRequest req, HttpServletResponse res) 
			throws ServletException, IOException{
		
		// Use "request" to read incoming HTTP headers
		// (e.g., cookies) and query data from HTML forms.
		System.out.println(req);
		System.out.println("user-Agent: " + req.getHeader("user-Agent"));
		
		// Use "response" to specify the HTTP response status
		// code and headers (e.g., the content type, cookies).
		
		PrintWriter out = res.getWriter();
		// Use "out" to send content to browser
		
		System.out.println("11111");
	}
}
