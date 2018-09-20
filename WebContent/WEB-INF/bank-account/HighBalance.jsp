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
<jsp:useBean id="eliteCustomer" type="c15.BankCustomer" scope="request" />
It is an honor to serve you,
<jsp:getProperty name="eliteCustomer" property="firstName" />
<jsp:getProperty name="eliteCustomer" property="lastName" />!
Since you are one of our most valued customers, we would like 
to offer you the opportunity to spend a mere fraction of your
<jsp:getProperty name="eliteCustomer" property="balance" />
on a boat worthy of your status. Please visit our boat store for
more information.
</body>
</html>