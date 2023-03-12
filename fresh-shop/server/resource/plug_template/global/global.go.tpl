package global

{{- if .HasGlobal }}

import "fresh-shop/server/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}