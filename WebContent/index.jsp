<%@ page language="java" contentType="text/html; charset=UTF-8"
    pageEncoding="UTF-8"%>
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>111111111</title>
</head>
<body>
<h1>
Time: <%= new java.util.Date() %>
</h1>
<p>
<%
	for (int i = 0; i < 3; i++) {
		out.print(i);
	}
%>
</p>
</body>
</html>