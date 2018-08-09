package util;

public class ServletUtilities {
	public static final String DOCTYPE = "<!DOCTYPE html>\n";
	public static String headWithTitle(String title) {
		return DOCTYPE +
				"<html>\n" +
				"<head>\n" +
				"<title>" + title + "</title>\n" +
				"</head>\n";
	}
}
