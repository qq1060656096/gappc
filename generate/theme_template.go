package generate

const themeTemplate = `{{ define "{{themeName}}/{{subDirName}}/{{fileName}}" }}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8" />
    <title>{{ .title }}</title>
    <style>
    body{
        text-align: center;
        color: red
    }
    </style>
</head>
<body>
    <h1>{{ .content }}</h1>
    <p>时间: {{ .nowDate }}</p>
    <p class="theme">主题: default</p>
    <p>模板: {{themeName}}/{{subDirName}}/{{fileName}}</p>
</body>
</html>
{{end}}}
`