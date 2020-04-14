package templates

const ManifestItem = `{{- define "ManifestItem"}}
<item id="{{.ChapterID}}" media-type="application/xhtml+xml" href="{{.FileName}}"/>
{{if .SubChapter }}
{{- range $index , $value := .SubChapter}}
{{template "ManifestItem" $value}}
{{- end}}
{{- end}}
{{- end}}`
