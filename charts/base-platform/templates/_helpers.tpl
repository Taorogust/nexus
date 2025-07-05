{{- /*
Define el fullname usado por todos los recursos:
<release-name>-<chart-name>, p.ej. base-platform-my-release
*/ -}}
{{- define "base-platform.fullname" -}}
{{- printf "%s-%s" .Release.Name .Chart.Name | trunc 63 | trimSuffix "-" }}
{{- end -}}

{{- /*
Labels comunes para todos los recursos de base-platform
*/ -}}
{{- define "base-platform.labels" -}}
app.kubernetes.io/name: {{ .Chart.Name }}
app.kubernetes.io/instance: {{ .Release.Name }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end -}}
