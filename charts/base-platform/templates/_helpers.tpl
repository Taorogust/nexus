{{/* 
Define el fullname usado por todos los recursos:
  release-name-chart-name, p.ej. base-platform-my-release
*/}}
{{- define "base-platform.fullname" -}}
{{- printf "%s-%s" .Release.Name .Chart.Name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Define las labels estándar para Kubernetes.
Todas las claves y valores deben ser strings para evitar errores de deserialización.
*/}}
{{- define "base-platform.labels" -}}
app.kubernetes.io/name: {{ .Chart.Name | quote }}
app.kubernetes.io/instance: {{ .Release.Name | quote }}
app.kubernetes.io/version: {{ .Chart.Version | quote }}
app.kubernetes.io/managed-by: {{ .Release.Service | quote }}
{{- end -}}
