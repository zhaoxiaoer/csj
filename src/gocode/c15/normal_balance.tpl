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
<ul>
  <li>First name: {{.FirstName}}</li>
  <li>Last name: {{.LastName}}</li>
  <li>ID: {{.Id}}</li>
  <li>Balance: ${{.Balance}}</li>
</ul>
</body>
</html>