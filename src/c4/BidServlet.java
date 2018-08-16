package c4;

import java.io.IOException;
import java.io.PrintWriter;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import c4.beans.BidInfo;
import util.BeanUtilities;
import util.ServletUtilities;

public class BidServlet extends HttpServlet {
	public void doGet(HttpServletRequest req, HttpServletResponse res) 
			throws ServletException, IOException {
		BidInfo bid = new BidInfo();
		BeanUtilities.populateBean(bid, req);
		if (bid.isComplete()) {
			showBid(req, res, bid);
		} else {
			showEntryForm(req, res, bid);
		}
	}
	
	private void showBid(HttpServletRequest req, HttpServletResponse res, BidInfo bid) 
			throws ServletException, IOException {
		submitBid(bid);
		res.setContentType("text/html");
		PrintWriter out = res.getWriter();
		String title = "Bid Submitted";
		out.println(ServletUtilities.headWithTitle(title) +
				"<body bgcolor=\"#FDF5E6\"><center>\n" +
				"<h1>" + title + "</h1>\n" +
				"Your bid is now active. If your bid is successful,\n" +
				"you will be notified within 24 hours of the close\n" + 
				"of bidding.\n" +
				"<p>\n" +
				"<table border=1>\n" +
				"  <tr><th bgcolor=\"black\"><font color=\"white\">" +
				bid.getItemName() + "</font></th></tr>\n" +
				"  <tr><td>Item ID: " +
				bid.getItemID() + "</td></tr>\n" +
				"  <tr><td>Name: " +
				bid.getBidderName() + "</td></tr>\n" +
				"  <tr><td>Email address: " +
				bid.getEmailAddress() + "</td></tr>\n" +
				"  <tr><td>Bid price: $" +
				bid.getBidPrice() + "</td></tr>\n" +
				"  <tr><td>Auto-increment price: " +
				bid.isAutoIncrement() + "</td></tr>\n" +
				"</table></center></body></html>");
	}
	
	private void showEntryForm(HttpServletRequest req, HttpServletResponse res, BidInfo bid) 
			throws ServletException, IOException {
		boolean isPartlyComplete = bid.isPartlyComplete();
		res.setContentType("text/html");
		PrintWriter out = res.getWriter();
		String title = "Welcome to Auctions-R-Us. Please Enter Bid.";
		out.println(ServletUtilities.headWithTitle(title) +
				"<body bgcolor=\"#FDF5E6\"><center>\n" +
				"<h1>" + title + "</h1>\n" +
				warning(isPartlyComplete) +
				"<form>\n" +
				inputElement("Item ID", "itemID", bid.getItemID(), isPartlyComplete) +
				inputElement("Item Name", "itemName", bid.getItemName(), isPartlyComplete) +
				inputElement("Your Name", "bidderName", bid.getBidderName(), isPartlyComplete) +
				inputElement("Your Email Address", "emailAddress", bid.getEmailAddress(), isPartlyComplete) +
				inputElement("Amount Bid", "bidPrice", bid.getBidPrice(), isPartlyComplete) +
				checkbox("Auto-increment bid to match other bidders?", "autoIncrement", bid.isAutoIncrement()) +
				"<input type=\"submit\" value=\"Submit Bid\">\n" +
				"</form></center></body></html>");
	}
	
	private void submitBid(BidInfo bid) {
		
	}
	
	private String warning(boolean isFormPartlyComplete) {
		if (isFormPartlyComplete) {
			return "<h2>Required Data Missing! " +
					"Enter and Resubmit.</h2>";
		} else {
			return "";
		}
	}
	
	private String inputElement(String prompt, String name, String value, boolean shouldPrompt) {
		String message = "";
		if (shouldPrompt && ((value == null) || value.equals(""))) {
			message = "<b>Required field!</b> ";
		}
		return (message + prompt + ": " +
				"<input type=\"text\" name=\"" + name + "\"" +
				" value=\"" + value + "\"><br>\n");
	}
	
	private String inputElement(String prompt, String name, double value, boolean shouldPrompt) {
		String num;
		if (value == 0.0) {
			num = "";
		} else {
			num = String.valueOf(value);
		}
		return inputElement(prompt, name, num, shouldPrompt);
	}
	
	private String checkbox(String prompt, String name, boolean isChecked) {
		String result = prompt + ": " +
				"<input type=\"checkbox\" name=\"" + name + "\"";
		if (isChecked) {
			result = result + " checked";
		}
		result += "><br>\n";
		return result;
	}
}
