package templates

const TOCChapter = `{{- define "TOCChapter"}}
{{- if eq .Level 1}}<h3><a href="{{.FileName}}">{{.Title}}</a></h3>{{end}}
{{- if eq .Level 2}}<h4><a href="{{.FileName}}">{{.Title}}</a></h4>{{end}}
{{- if eq .Level 3}}<h5><a href="{{.FileName}}">{{.Title}}</a></h5>{{end}}
{{- if eq .Level 4}}<h6><a href="{{.FileName}}">{{.Title}}</a></h6>{{end}}
{{if .SubChapter }}
<ul>
{{- range $index , $value := .SubChapter}}
    <li>{{template "TOCChapter" $value}}</li>
{{- end}}
</ul>
{{- end}}
{{- end}}`
