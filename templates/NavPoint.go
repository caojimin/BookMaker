package templates

const NavPoint = `{{- define "NavPoint"}}
<navPoint id="{{.ChapterID}}" playOrder="{{.PlayOrder}}">
    <navLabel>
        <text>{{.Title}}</text>
    </navLabel>
    <content src="{{.FileName}}"/>
    {{- range $index , $value := .SubChapter}}
    {{- template "NavPoint" $value}}
    {{- end}}
</navPoint>
{{- end}}`
