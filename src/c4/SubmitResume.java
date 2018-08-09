package c4;

import java.io.IOException;
import java.io.PrintWriter;
import java.util.StringTokenizer;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import util.ServletUtilities;

public class SubmitResume extends HttpServlet {
	public void init() throws ServletException {
		
	}

	public void doPost(HttpServletRequest req, HttpServletResponse res) 
			throws ServletException, IOException {
		res.setContentType("text/html");
		PrintWriter out = res.getWriter();
		if (req.getParameter("previewButton") != null) {
			showPreview(req, out);
		} else {
			storeResume(req);
			showConfirmation(req, out);
		}
	}
	
//	public long getLastModified(HttpServletRequest req) {
//		return -1;
//	}
//	
//	public void destroy() {
//		
//	}
	
	private void showPreview(HttpServletRequest req, PrintWriter out) {
		String headingFont = req.getParameter("headingFont");
		headingFont = replaceIfMissingOrDefault(headingFont, "");
		int headingSize = getSize(req.getParameter("headingSize"), 32);
		String bodyFont = req.getParameter("bodyFont");
		bodyFont = replaceIfMissingOrDefault(bodyFont, "");
		int bodySize = getSize(req.getParameter("bodySize"), 18);
		String fgColor = req.getParameter("fgColor");
		fgColor = replaceIfMissing(fgColor, "black");
		String bgColor = req.getParameter("bgColor");
		bgColor = replaceIfMissing(bgColor, "white");
		
		String name = req.getParameter("name");
//		name = ServletUtilities.filter(name);
		name = replaceIfMissing(name, "Lou Zer");
		String title = req.getParameter("title");
		title = replaceIfMissing(title, "Loser");
		String email = req.getParameter("email");
		email = replaceIfMissing(email, "contact@hot-computer-jobs.com");
		String languages = req.getParameter("languages");
		languages = replaceIfMissing(languages, "<i>None</i>");
		String languageList = makeList(languages);
		
		String skills = req.getParameter("skills");
		skills = replaceIfMissing(skills, "Not many, obviously.");
		out.println(ServletUtilities.DOCTYPE +
				"<html><head><title>Resume for" + name + "</title>\n" +
				makeStyleSheet(headingFont, headingSize, 
						bodyFont, bodySize, 
						fgColor, bgColor) + "\n" +
				"</head>\n" +
				"<body>\n" +
				"<center>\n" +
				"<span class=\"heading1\">" + name + "</span><br>\n" +
				"<span class=\"heading2\">" + title + "<br>\n" +
				"<a href=\"mailto:" + email + "\">" + email +"</a></span>\n" +
				"</center><br><br>\n" +
				"<span class=\"header3\">Programming Languages</span>\n" +
				makeList(languages) + "<br></br>\n" +
				"<span class=\"heading3\">Skills and Experience" + "</span><br><br>\n" +
				skills + "\n" +
				"</body></html>");
	}
	
	private String makeStyleSheet(String headingFont, int heading1Size,
			String bodyFont, int bodySize, String fgColor, String bgColor) {
		int heading2Size = heading1Size*7/10;
		int heading3Size = heading1Size*6/10;
		String styleSheet = "<style type=\"text/css\">\n" +
				"<!--\n" +
				".heading1 { font-size: " + heading1Size + "px;\n" +
				"			font-weight: bold;\n" +
				"			font-family: " + headingFont + "Arial, Helvetica, sans-serif;\n" +
				"}\n" +
				".heading2 { font-size: " + heading2Size + "px;\n" +
				"			font-weight: bold;\n" +
				"			font-family: " + headingFont + "Arial, Helvetica, sans-serif;\n" +
				"}\n" +
				".heading3 { font-size: " + heading3Size + "px;\n" +
				"			font-weight: bold;\n" +
				"			font-family: " + headingFont + "Arial, Helvetica, sans-serif;\n" +
				"}\n" +
				"body { color: " + fgColor + ";\n" +
				"		background-color: " + bgColor + ";\n" +
				"		font-size: " + bodySize + "px;\n" +
				"		font-family: " + bodyFont + "Times New Roman, Times, serif;\n" +
				"}\n" +
				"a: hover { color: red; }\n" +
				"-->\n" +
				"</style>";
		return styleSheet;
	}
	
	private String replaceIfMissing(String orig, String replacement) {
		if ((orig == null) || (orig.trim().equals(""))) {
			return replacement;
		} else {
			return orig;
		}
	}
	
	private String replaceIfMissingOrDefault(String orig, String replacement) {
		if ((orig == null) || (orig.trim().equals("")) || orig.equals("default")) {
			return replacement;
		} else {
			return orig;
		}
	}
	
	private int getSize(String sizeString, int defaultSize) {
		try {
			return Integer.parseInt(sizeString);
		} catch (NumberFormatException nfe) {
			return defaultSize;
		}
	}
	
	private String makeList(String listItems) {
		StringTokenizer tokenizer = new StringTokenizer(listItems, ", ");
		String list = "<ul>\n";
		while (tokenizer.hasMoreTokens()) {
			list = list + "  <li>" + tokenizer.nextToken() + "\n";
		}
		list += "</ul>";
		return list;
	}
	
	private void showConfirmation(HttpServletRequest req, PrintWriter out) {
		String title = "Submission Confirmed.";
		out.println(ServletUtilities.headWithTitle(title) +
				"<body>\n" +
				"<h1>" + title + "</h1>\n" +
				"Your resume should appear online within\n" +
				"24 hours. If it doesn't, try submitting\n" +
				"again with a different email address.\n" +
				"</body></html>");
	}
	
	private void storeResume(HttpServletRequest req) {
		String email = req.getParameter("email");
		putInSpamList(email);
	}
	
	private void putInSpamList(String emailAddress) {
		// Code removed to protect the guilty.
	}
}
