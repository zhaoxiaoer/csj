{{ define "layout" }}
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
  <li>{{ template "item1" }}</li>
  <li>{{ template "item2" }}</li>
  <li>{{ template "item3" }}</li>
</ol>
</body>
</html>
{{ end }}