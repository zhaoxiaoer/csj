package util;

import java.io.IOException;
import java.io.PrintWriter;
import java.util.zip.GZIPOutputStream;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

public class GzipUtilities {
	public static boolean isGzipSupported(HttpServletRequest req) {
		String encodings = req.getHeader("Accept-Encoding");
		return ((encodings != null) && (encodings.indexOf("gzip") != -1));
	}
	
	public static boolean isGzipDisabled(HttpServletRequest req) {
		String flag = req.getParameter("disableGzip");
		return ((flag != null) && (!flag.equalsIgnoreCase("false")));
	}
	
	public static PrintWriter getGzipWriter(HttpServletResponse res) 
			throws IOException {
		return (new PrintWriter(new GZIPOutputStream(res.getOutputStream())));
	}
}
