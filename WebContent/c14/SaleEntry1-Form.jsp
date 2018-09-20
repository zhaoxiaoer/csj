<%@ page language="java" contentType="text/html; charset=UTF-8"
    pageEncoding="UTF-8"%>
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>Invoking SaleEntry1.jsp</title>
</head>
<body>
<center>
  <table boder=5>
    <tr>
      <th>Invoking SaleEntry1.jsp</th>
    </tr>
  </table>
  <form action="SaleEntry1.jsp">
    Item ID: <input type="text" name="itemID"><br>
    Number of Items: <input type="text" name="numItems"><br>
    Discount Code: <input type="text" name="discountCode"><br>
    <input type="submit" value="Show Price">
  </form>
</center>
</body>
</html>