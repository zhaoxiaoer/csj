<%@ page language="java" contentType="text/html; charset=UTF-8"
    pageEncoding="UTF-8"%>
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>Order Confirmation</title>
</head>
<body>
<h2>Order Confirmation</h2>
Thanks for ordering <i><%= request.getParameter("title") %></i>!
</body>
</html>