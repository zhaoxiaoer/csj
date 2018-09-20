<%@ page language="java" contentType="text/html; charset=UTF-8"
    pageEncoding="UTF-8"%>
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>Baked Bean Values: page-based Sharing</title>
</head>
<body>
<h1>Baked Bean Values: page-based Sharing</h1>
<jsp:useBean id="pageBean" class="c14.BakedBean" />
<jsp:setProperty name="pageBean" property="*" />
<h2>Bean level: <jsp:getProperty name="pageBean" property="level" /></h2>
<h2>Dish bean goes with: <jsp:getProperty name="pageBean" property="goesWith" /></h2>
</body>
</html>