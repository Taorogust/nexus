{{- /*
Define el fullname usado por todos los recursos:
release-name-chart-name, p.ej. base-platform-my-release
*/ -}}
{{- define "base-platform.fullname" -}}
{{ printf "%s-%s" .Release.Name .Chart.Name | trunc 63 | trimSuffix "-" }}
{{- end -}}
