package templates

const TocXhtml = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.1//EN" "http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
    <title>Table of Contents</title>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8"/>
</head>
<body>
    <h1><b>TABLE OF CONTENTS</b></h1>
    <br />
    {{- range $index , $value := .Chapters}}
    {{template "TOCChapter" $value}}
    {{- end}}
</body>
</html>`
