{{- /*
Define el fullname usado por todos los recursos:
  release-name-chart-name, p.ej. base-platform-my-release
*/ -}}
{{- define "base-platform.fullname" -}}
{{- printf "%s-%s" .Release.Name .Chart.Name | trunc 63 | trimSuffix "-" }}
{{- end }}

{{- /*
Labels comunes para todos los recursos del chart.
Se aseguran como cadenas con la función `quote`.
*/ -}}
{{- define "base-platform.labels" -}}
app.kubernetes.io/name:       {{ .Chart.Name | quote }}
app.kubernetes.io/instance:   {{ .Release.Name | quote }}
app.kubernetes.io/version:    {{ .Chart.AppVersion | default .Chart.AppVersion | quote }}
app.kubernetes.io/managed-by: {{ "Helm" | quote }}
{{- end }}
