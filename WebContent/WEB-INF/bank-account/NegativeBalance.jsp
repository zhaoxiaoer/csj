<%@ page language="java" contentType="text/html; charset=UTF-8"
    pageEncoding="UTF-8"%>
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>You Owe Us Money!</title>
</head>
<body>
<table border=5 align="center">
  <tr>
    <th>We Know Where You Live!</th>
  </tr>
</table>
<jsp:useBean id="badCustomer" type="c15.BankCustomer" scope="request" />
Watch out,
<jsp:getProperty name="badCustomer" property="firstName" />, 
we know where you live.
Pay us the
$<jsp:getProperty name="badCustomer" property="balanceNoSign" />
you owe us before it is too late!
</body>
</html>