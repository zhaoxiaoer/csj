<%@ page language="java" contentType="text/html; charset=UTF-8"
    pageEncoding="UTF-8"%>
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>What's New at JspNews.com</title>
</head>
<body>
<table border=5 align="center">
  <tr>
    <th class="title">What's New at JspNews.com</th>
  </tr>
 </table>
<p>Here is a summary of our three most recent news stories:</p>
<ol>
  <li><jsp:include page="/WEB-INF/Item1.html"></jsp:include></li>
  <li><jsp:include page="/WEB-INF/Item2.html"></jsp:include></li>
  <li><jsp:include page="/WEB-INF/Item3.html"></jsp:include></li>
</ol>
</body>
</html>