package c9;

import java.io.*;
import javax.servlet.*;
import javax.servlet.http.*;

import c9.shoppingcart.Catalog;
import c9.shoppingcart.CatalogItem;
import util.ServletUtilities;

public abstract class CatalogPage extends HttpServlet {
	private CatalogItem[] items;
	private String[] itemIDs;
	private String title;
	
	protected void setItems(String[] itemIDs) {
		this.itemIDs = itemIDs;
		items = new CatalogItem[itemIDs.length];
		for (int i = 0; i < items.length; i++) {
			items[i] = Catalog.getItem(itemIDs[i]);
		}
	}
	
	protected void setTitle(String title) {
		this.title = title;
	}
	
	public void doGet(HttpServletRequest req, HttpServletResponse res) 
			throws ServletException, IOException {
		if (items == null) {
			res.sendError(res.SC_NOT_FOUND, "Missing Items.");
			return;
		}
		res.setContentType("text/html");
		PrintWriter out = res.getWriter();
		out.println(ServletUtilities.headWithTitle(title) +
				"<body bgcolor=\"#FDF5E6\">\n" +
				"<h1 align=\"center\">" + title + "</h1>");
		CatalogItem item;
		for (int i = 0; i < items.length; i++) {
			out.println("<hr>");
			item = items[i];
			if (item == null) {
				out.println("<font color=\"red\">" + 
						"Unknown item ID " + itemIDs[i] +
						"</font>");
			} else {
				out.println();
				String formURL = "/csj/c9/orderpage";
				formURL = res.encodeURL(formURL);
				out.println("<form action=\"" + formURL + "\">\n" +
						"<input type=\"hidden\" name=\"itemID\" " +
						"  value=\"" + item.getItemID() + "\">\n" +
						"<h2>" + item.getShortDescription() + 
						" ($" + item.getCost() + ")</h2>\n" +
						item.getLongDescription() + "\n" +
						"<p>\n<center>\n" +
						"<input type=\"submit\" " +
						"value=\"Add to Shopping Cart\">\n" +
						"</center>\n</p>\n</form>");
			}
		}
		out.println("<hr>\n</body></html>");
	}
}
