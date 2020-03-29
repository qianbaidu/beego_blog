<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>

<h1>default</h1>

<ul>
    <li>{{if .TrueCond}} true condition {{ end }}</li>
    <li>{{.user}} </li>
    <li>{{with .user}}<li>{{.Name}}</li><li>{{.Age}} </li>{{end }}</li>
    <li></li>
</ul>


<ul>
    {{range .nums}}
    <li>{{.}}</li>
    {{end}}
</ul>


<ul>
    <li>{{$tplVal := .TplVar}}</li>
    <li>{{$tplVal}}</li>
    <li>{{.html}}</li>
    <li>{{str2html .html}}</li>
    <li>{{.html | htmlquote}}</li>
</ul>


{{template "footer"}}


</body>
</html>

