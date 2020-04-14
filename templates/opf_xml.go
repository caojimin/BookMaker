package templates

const OpfXml = `<?xml version="1.0" encoding="utf-8"?>
<package xmlns="http://www.idpf.org/2007/opf" version="2.0" unique-identifier="{{.BookId}}">

    <metadata xmlns:dc="http://purl.org/dc/elements/1.1/">
        <dc:title>{{ .BookName }}</dc:title>
        <dc:language>zh-cn</dc:language>
        {{- range $index , $value := .Authors}}
        <dc:creator>{{ $value }}</dc:creator>
        {{- end}}
        <meta name="cover" content="cover_image"/>
    </metadata>

    <manifest>
        <item id="cover_image" href="cover.jpg" media-type="image/jpeg"/>
        <item id="toc" media-type="application/x-dtbncx+xml" href="toc.ncx"/>
        <item id="toc_html" media-type="application/xhtml+xml" href="toc.xhtml"/>
        {{- range $index , $value := .Chapters}}
        {{template "ManifestItem" $value}}
        {{- end}}
    </manifest>

    <spine toc="toc">
        <itemref idref="toc_html"/>
        {{- range $index , $value := .Chapters}}
        {{template "SpineItemRef" $value}}
        {{- end}}
    </spine>
</package>`
