{{define "request"}}
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>Baked Bean Values: request-based Sharing</title>
</head>
<body>
<h1>Baked Bean Values: request-based Sharing</h1>
<h2>Bean level: {{.Level}}</h2>
<h2>Dish bean goes with: {{.GoesWith}}</h2>
{{template "snippet" .}}
</body>
</html>
{{end}}