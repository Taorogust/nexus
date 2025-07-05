{{- /*
Define el nombre completo usado por todos los recursos:
release-name-chart-name, p.ej. base-platform-my-release
*/ -}}
{{- define "base-platform.fullname" -}}
{{- printf "%s-%s" .Release.Name .Chart.Name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- /*
Define el nombre base del chart (sin el release):
p.ej. base-platform
*/ -}}
{{- define "base-platform.name" -}}
{{- .Chart.Name -}}
{{- end -}}

{{- /*
Labels comunes que queremos inyectar en todos los recursos
*/ -}}
{{- define "base-platform.labels" -}}
app.kubernetes.io/name: {{ include "base-platform.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
app.kubernetes.io/version: {{ .Chart.AppVersion }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end -}}
