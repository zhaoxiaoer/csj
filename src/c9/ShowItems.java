package c9;

import java.io.*;
import java.util.*;
import javax.servlet.*;
import javax.servlet.http.*;

import util.ServletUtilities;

public class ShowItems extends HttpServlet {
	public void doGet(HttpServletRequest req, HttpServletResponse res) 
			throws ServletException, IOException {
		
//		try {
//			System.out.println("current thread id: " + Thread.currentThread().getId());
//			Thread.currentThread().sleep(30000);
//			System.out.println("current thread id: " + Thread.currentThread().getId());
//		} catch (InterruptedException e) {
//			// TODO Auto-generated catch block
//			e.printStackTrace();
//		}
		
		HttpSession session = req.getSession();
		ArrayList previousItems = (ArrayList) session.getAttribute("previousItems");
		if (previousItems == null) {
			previousItems = new ArrayList();
			session.setAttribute("previousItems", previousItems);
		}
		String newItem = req.getParameter("newItem");
		PrintWriter out = res.getWriter();
		String title = "Items Purchased";
		
//		try {
//			System.out.println("2 current thread id: " + Thread.currentThread().getId());
//			Thread.currentThread().sleep(30000);
//			System.out.println("2 current thread id: " + Thread.currentThread().getId());
//		} catch (InterruptedException e) {
//			// TODO Auto-generated catch block
//			e.printStackTrace();
//		}
		
		out.println(ServletUtilities.headWithTitle(title) +
				"<body bgcolor=\"#FDF5E6\">\n" +
				"<h1>" + title + "</h1>");
		synchronized(previousItems) {
			if ((newItem != null) && (!newItem.trim().equals(""))) {
				previousItems.add(newItem);
			}
			if (previousItems.size() == 0) {
				out.println("<i>No items</i>");
			} else {
				out.println("<ul>");
				for (int i = 0; i < previousItems.size(); i++) {
					out.println("<li>" + (String) previousItems.get(i) + "</li>");
				}
				out.println("</ul>");
			}
		}
		out.println("</body></html>");
	}
}
