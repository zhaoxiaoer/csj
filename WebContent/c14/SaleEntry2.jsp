<%@ page language="java" contentType="text/html; charset=UTF-8"
    pageEncoding="UTF-8"%>
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>Using jsp:setProperty</title>
</head>
<body>
<center>
  <table border=5>
    <tr>
      <th>Using jsp:setProperty</th>
    </tr>
  </table>
  <jsp:useBean id="entry" class="c14.SaleEntry" />
  <jsp:setProperty 
    name="entry" 
    property="itemID"
    param="itemID" />
  <jsp:setProperty 
    name="entry"
    property="numItems"
    param="numItems" />
  <jsp:setProperty 
    name="entry"
    property="discountCode"
    param="discountCode" />
  <br>
  <table border=1>
    <tr>
      <th>Item ID</th>
      <th>Unit Price</th>
      <th>Number Ordered</th>
      <th>Total Price</th>
    </tr>
    <tr align="right">
      <td><jsp:getProperty name="entry" property="itemID" /></td>
      <td><jsp:getProperty name="entry" property="itemCost"/></td>
      <td><jsp:getProperty name="entry" property="numItems" /></td>
      <td><jsp:getProperty name="entry" property="totalCost" /></td>
    </tr>
  </table>
</center>
</body>
</html>