package templates

const SpineItemRef = `{{- define "SpineItemRef"}}
<itemref idref="{{.ChapterID}}"/>
{{if .SubChapter }}
{{- range $index , $value := .SubChapter}}
{{template "SpineItemRef" $value}}
{{- end}}
{{- end}}
{{- end}}`
