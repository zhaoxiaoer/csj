package c9;

import java.io.*;
import java.text.NumberFormat;
import java.util.List;

import javax.servlet.*;
import javax.servlet.http.*;

import c9.shoppingcart.*;
import util.ServletUtilities;

public class OrderPage extends HttpServlet {
	public void doGet(HttpServletRequest req, HttpServletResponse res) 
			throws ServletException, IOException {
		HttpSession session = req.getSession();
		ShoppingCart cart;
		synchronized(session) {
			cart = (ShoppingCart) session.getAttribute("shoppingCart");
			if (cart == null) {
				cart = new ShoppingCart();
				session.setAttribute("shoppingCart", cart);
			}
			String itemID = req.getParameter("itemID");
			if (itemID != null) {
				String numItemsString = req.getParameter("numItems");
				if (numItemsString == null) {
					cart.addItem(itemID);
				} else {
					int numItems;
					try {
						numItems = Integer.parseInt(numItemsString);
					} catch (NumberFormatException nfe) {
						numItems = -1;
					}
					cart.setNumOrdered(itemID, numItems);
				}
			}
		}
		res.setContentType("text/html");
		PrintWriter out = res.getWriter();
		String title = "Status of Your Order";
		out.println(ServletUtilities.headWithTitle(title) +
				"<body bgcolor=\"#FDF5E6\">\n" +
				"<h1 align=\"center\">" + title + "</h1>");
		synchronized(session) {
			List itemsOrdered = cart.getItemsOrdered();
			if (itemsOrdered.size() == 0) {
				out.println("<h2><i>No items in your cart...</i></h2>");
			} else {
				out.println("<table border=1 align=\"center\">\n" +
						"<tr bgcolor=\"#FFAD00\">\n" +
						"  <th>Item ID</th><th>Description\n</th>\n" +
						"  <th>Unit Cost</th><th>Number</th><th>Total Cost</th>" +
						"</tr>");
				ItemOrder order;
				NumberFormat formatter = NumberFormat.getCurrencyInstance();
				for (int i = 0; i < itemsOrdered.size(); i++) {
					order = (ItemOrder) itemsOrdered.get(i);
					out.println("<tr>\n" +
							"  <td>" + order.getItemID() + "</td>\n" +
							"  <td>" + order.getShortDescription() + "</td>\n" +
							"  <td>" + formatter.format(order.getUnitCost()) + "</td>\n" +
							"  <td>" +
							"<form>\n" + 
							"<input type=\"hidden\" name=\"itemID\" "+ 
							"value=\"" + order.getItemID() + "\">\n" +
							"<input type=\"text\" name=\"numItems\" size=3 value=\"" +
							order.getNumItems() + "\">\n" +
							"<small>\n" +
							"<input type=\"submit\" value=\"Update Order\">\n" +
							"</small>\n" +
							"</form>\n" +
							"</td>" +
							"<td>" +
							formatter.format(order.getTotalCost()) + "</td>");
				}
				String checkoutURL = res.encodeURL("/csj/c9/Checkout.html");
				out.println("</table>\n" +
						"<form action=\"" + checkoutURL + "\">\n" +
						"<big><center>\n" + 
						"<input type=\"submit\" value=\"Proceed to Checkout\">\n" +
						"</center></big></form>");
			}
			out.println("</body></html>");
		}
	}
}
