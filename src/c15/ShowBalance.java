package c15;

import java.io.*;
import javax.servlet.*;
import javax.servlet.http.*;

public class ShowBalance extends HttpServlet {
	public void doGet(HttpServletRequest request, HttpServletResponse response) 
			throws ServletException, IOException {
		BankCustomer customer = BankCustomer.getCustomer(request.getParameter("id"));
		String address;
		if (customer == null) {
			address = "/WEB-INF/bank-account/UnknownCustomer.jsp";
		} else if (customer.getBalance() < 0) {
			address = "/WEB-INF/bank-account/NegativeBalance.jsp";
			request.setAttribute("badCustomer", customer);
		} else if (customer.getBalance() < 10000) {
			address = "/WEB-INF/bank-account/NormalBalance.jsp";
			request.setAttribute("regularCustomer", customer);
		} else {
			address = "/WEB-INF/bank-account/HighBalance.jsp";
			request.setAttribute("eliteCustomer", customer);
		}
		RequestDispatcher dispatcher = request.getRequestDispatcher(address);
		dispatcher.forward(request, response);
	}
}
