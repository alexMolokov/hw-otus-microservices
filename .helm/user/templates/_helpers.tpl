{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "app.secret" }}
{{-  printf "%s-%s" .Release.Name  "secret" | trunc 63 -}}
{{- end }}

{{- define "app.configmap" }}
{{-  printf "%s-%s" .Release.Name  "configmap" | trunc 63 -}}
{{- end }}

{{- define "app.migration" }}
{{-  printf "%s-%s" .Release.Name  "migration" | trunc 63 -}}
{{- end }}

{{- define "app.service" }}
{{-  printf "%s-%s" .Release.Name  "service" | trunc 63 -}}
{{- end }}

{{- define "app.ingress" }}
{{-  printf "%s-%s" .Release.Name  "ingress" | trunc 63 -}}
{{- end }}

{{- define "app.servicemonitor" }}
{{-  printf "%s-%s" .Release.Name  "servicemonitor" | trunc 63 -}}
{{- end }}