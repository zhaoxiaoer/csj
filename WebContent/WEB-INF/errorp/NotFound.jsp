<%@ page language="java" contentType="text/html; charset=UTF-8"
    pageEncoding="UTF-8"%>
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>404: Not Found</title>
</head>
<body bgcolor="#FDF5E6">
<h2>Error!</h2>
I'm sorry, but I cannot find a page that matches
<%= request.getAttribute("javax.servlet.forward.request_uri") %>
on the system. Maybe you should try one of the following:
<ul>
  <li>Go to the server's <a href="/csj">home page</a></li>
  <li>Search for relevant pages.</li>
</ul>
</body>
</html>