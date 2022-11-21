# html_templates

This includes a template example which includes other templates into one final html file.

Important:

    {{ template "Header" . }}

The . in there means you give the Header template your current map of key/values. Without this . the Header template could not use {{.Title}}.

