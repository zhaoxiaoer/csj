<%@ page language="java" contentType="text/html; charset=UTF-8"
    pageEncoding="UTF-8"%>
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>Baked Bean Values: request-based Sharing</title>
</head>
<body>
<h1>Baked Bean Values: request-based Sharing</h1>
<jsp:useBean id="requestBean" class="c14.BakedBean" scope="request" />
<jsp:setProperty name="requestBean" property="*" />
<h2>Bean level: <jsp:getProperty name="requestBean" property="level" /></h2>
<h2>Dish bean goes with: <jsp:getProperty name="requestBean" property="goesWith" /></h2>
<jsp:include page="BakedBeanDisplay-snippet.jsp" />
</body>
</html>