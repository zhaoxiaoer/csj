package util;

import javax.servlet.http.*;

public class CookieUtilities {
	public static String getCookieValue(HttpServletRequest req, 
			String cookieName, String defaultValue) {
		Cookie[] cookies = req.getCookies();
		if (cookies != null) {
			for (int i = 0; i < cookies.length; i++) {
				Cookie cookie = cookies[i];
				if (cookieName.equals(cookie.getName())) {
					return cookie.getValue();
				}
			}
		}
		return defaultValue;
	}
	
	public static Cookie getCookie(HttpServletRequest req, String cookieName) {
		Cookie[] cookies = req.getCookies();
		if (cookies != null) {
			for (int i = 0; i < cookies.length; i++) {
				Cookie cookie = cookies[i];
				if (cookieName.equals(cookie.getName())) {
					return cookie;
				}
			}
		}
		return null;
	}
}
