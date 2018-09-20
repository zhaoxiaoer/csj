<%@ page language="java" contentType="text/html; charset=UTF-8"
    pageEncoding="UTF-8"%>
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>Baked Bean Values: session-based Sharing</title>
</head>
<body>
<h1>Baked Bean Values: session-based Sharing</h1>
<jsp:useBean id="sessionBean" class="c14.BakedBean" scope="session" />
<jsp:setProperty name="sessionBean" property="*" />
<h2>Bean level: <jsp:getProperty name="sessionBean" property="level" /></h2>
<h2>Dish bean goes with: <jsp:getProperty name="sessionBean" property="goesWith" /></h2>
</body>
</html>