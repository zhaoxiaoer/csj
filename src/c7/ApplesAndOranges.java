package c7;

import java.io.IOException;
import java.io.PrintWriter;

import javax.servlet.ServletException;
import javax.servlet.http.*;

public class ApplesAndOranges extends HttpServlet {
	public void doGet(HttpServletRequest req, HttpServletResponse res) 
			throws ServletException, IOException {
		res.setContentType("application/vnd.ms-excel");
		PrintWriter out = res.getWriter();
		out.println("\tQ1\tQ2\tQ3\tQ4\tTotal");
		out.println("Apples\t78\t87\t92\t29\t=SUM(B2:E2)");
		out.println("Oranges\t77\t86\t93\t30\t=SUM(B3:E3)");
	}
}
