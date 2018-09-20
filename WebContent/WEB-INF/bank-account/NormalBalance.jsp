<%@ page language="java" contentType="text/html; charset=UTF-8"
    pageEncoding="UTF-8"%>
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>Your Balance</title>
</head>
<body>
<table border=5 align="center">
  <tr>
    <th>Your Balance</th>
  </tr>
</table>
<jsp:useBean id="regularCustomer" type="c15.BankCustomer" scope="request" />
<ul>
  <li>First name: <jsp:getProperty name="regularCustomer" property="firstName" />
  <li>Last name: <jsp:getProperty name="regularCustomer" property="lastName" />
  <li>ID: <jsp:getProperty name="regularCustomer" property="id" />
  <li>Balance: $<jsp:getProperty name="regularCustomer" property="balance" />
</ul>
</body>
</html>