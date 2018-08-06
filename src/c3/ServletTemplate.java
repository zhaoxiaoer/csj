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
		res.setContentType("text/html");
		PrintWriter out = res.getWriter();
		// Use "out" to send content to browser
		out.println("<!DOCTYPE html>\n" + 
				"<html>\n" + 
				"<head>\n" + 
				"<title>HelloWorld</title>\n" + 
				"</head>\n" + 
				"<body bgcolor=\"#FDF5E6\">\n" + 
				"<p>hello world 111 222</p>\n" + 
				"</body>");
	}
}
