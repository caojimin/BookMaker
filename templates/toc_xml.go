package templates

const TocXml = `<?xml version="1.0" encoding="UTF-8" standalone="no" ?>
<!DOCTYPE ncx PUBLIC "-//NISO//DTD ncx 2005-1//EN" "http://www.daisy.org/z3986/2005/ncx-2005-1.dtd">
<ncx xmlns="http://www.daisy.org/z3986/2005/ncx/" version="2005-1" xml:lang="en">
    <!-- Metadata Section -->
    <!-- Title and Author Section -->
    <!-- Navigation Map Section -->
    <head>
        <!-- Must be exactly the same as dc:identifier in the content.opf file -->
        <meta name="dtb:uid" content="{{.BookId}}"/>
        <!-- Set for 2 if you want a sub-level. It can go up to 4 -->
        <meta name="dtb:depth" content="{{.MaxLevel}}"/>
        <meta name="dtb:totalPageCount" content="0"/> <!-- Do Not change -->
        <meta name="dtb:maxPageNumber" content="0"/> <!-- Do Not change -->
    </head>
    <docTitle>
        <text>{{.BookName}}</text>
    </docTitle>
    {{- range .Authors}}
    <docAuthor>
        <text>{{.}}</text>
    </docAuthor>
    {{- end}}
    <docAuthor>
        <text>Make by CTKindle</text>
    </docAuthor>

    <navMap>
        <navPoint id="toc" playOrder="1">
            <navLabel>
                <text>Table of Contents</text>
            </navLabel>
            <content src="toc.xhtml"/>
        </navPoint>

        {{- range $index , $value := .Chapters}}
        {{- template "NavPoint" $value}}
        {{- end}}
    </navMap>

</ncx>`
