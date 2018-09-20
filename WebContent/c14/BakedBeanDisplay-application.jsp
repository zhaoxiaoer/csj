<%@ page language="java" contentType="text/html; charset=UTF-8"
    pageEncoding="UTF-8"%>
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>Baked Bean Values: application-based Sharing</title>
</head>
<body>
<h1>Baked Bean Values: application-based Sharing</h1>
<jsp:useBean id="applicationBean" class="c14.BakedBean" scope="application" />
<jsp:setProperty name="applicationBean" property="*" />
<h2>Bean level: <jsp:getProperty name="applicationBean" property="level" /></h2>
<h2>Dish bean goes with: <jsp:getProperty name="applicationBean" property="goesWith" /></h2>
</body>
</html>