<%@ page language="java" contentType="text/html; charset=UTF-8"
    pageEncoding="UTF-8"%>
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>The import Attribute</title>
</head>
<body>
<h2>The import Attribute</h2>
<%-- JSP page Derective --%>
<%@ page import="java.util.*,util.*" %>
<%-- JSP Declaration --%>
<%!
private String randomID() {
	int num = (int) (Math.random() * 10000000.0);
	return "id" + num;
}
private final String NO_VALUE = "<i>No Value</i>";
%>
<%-- JSP Scriptlet --%>
<%
String oldID = CookieUtilities.getCookieValue(request, "userID", NO_VALUE);
if (oldID.equals(NO_VALUE)) {
	String newID = randomID();
	Cookie cookie = new Cookie("userID", newID);
	cookie.setMaxAge(60 * 60 * 24 * 365);
	response.addCookie(cookie);
}
%>
<%-- JSP Expressions --%>
This page was accessed on <%= new Date() %> with a userID cookie of <%= oldID %>
</body>
</html>