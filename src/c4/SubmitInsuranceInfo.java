package c4;

import java.io.IOException;
import java.io.PrintWriter;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import c4.beans.InsuranceInfo;
import util.BeanUtilities;
import util.ServletUtilities;

public class SubmitInsuranceInfo extends HttpServlet {
	public void doGet(HttpServletRequest req, HttpServletResponse res) 
			throws ServletException, IOException {
		InsuranceInfo info = new InsuranceInfo();
		BeanUtilities.populateBean(info, req);
		res.setContentType("text/html");
		PrintWriter out = res.getWriter();
		String title = "Insurance Info for " + info.getName();
		out.println(ServletUtilities.headWithTitle(title) + 
				"<body bgcolor=\"#FDF5E6\">\n" +
				"<center>\n" +
				"<h1>" + title + "</h1>\n" +
				"<ul>\n" +
				"  <li>Employee ID: " + info.getEmployeeID() + "\n" +
				"  <li>Number of children: " + info.getNumChildren() +
				"  <li>Married?: " + info.isMarried() +
				"</ul></center></body></html>");
	}
}
