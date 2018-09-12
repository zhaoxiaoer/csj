<%@ page language="java" contentType="text/html; charset=UTF-8"
    pageEncoding="UTF-8"%>
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>Reading Three Request Parameters</title>
</head>
<body>
<h1>Reading Three Request Parameters</h1>
<ul>
  <li><b>param1</b>: <%= request.getParameter("param1") %></li>
  <li><b>param2</b>: <%= request.getParameter("param2") %></li>
  <li><b>param3</b>: <%= request.getParameter("param3") %></li>
</ul>
</body>
</html>