{{ define "packages" }}

<head>
<link rel="stylesheet" type="text/css" href="config/bootstrap.min.css">
<link rel="stylesheet" type="text/css" href="config/font-awesome.min.css">
<link rel="stylesheet" type="text/css" href="config/stylesheet.css">

</head>

{{ with .packages}}
<p>Packages:</p>
<ul>
    {{ range . }}
    <li>
        <a href="#{{- packageAnchorID . -}}">{{ packageDisplayName . }}</a>
    </li>
    {{ end }}
</ul>
{{ end}}

{{ range .packages }}
    <h2 id="{{- packageAnchorID . -}}">
        {{- packageDisplayName . -}}
    </h2>

    {{ with (index .GoPackages 0 )}}
        {{ with .DocComments }}
        <p>
            {{ safe (renderComments .) }}
        </p>
        {{ end }}
    {{ end }}

    Resource Types:
    <ul>
    {{- range (visibleTypes (sortedTypes .Types)) -}}
        {{ if isExportedType . -}}
        <li>
            <a href="{{ linkForType . }}">{{ typeDisplayName . }}</a>
        </li>
        {{- end }}
    {{- end -}}
    </ul>

    {{ range (visibleTypes (sortedTypes .Types))}}
        {{ template "type" .  }}
    {{ end }}
    <hr/>
{{ end }}

<p><em>
    Generated with <code>gen-crd-api-reference-docs</code>
    {{ with .gitCommit }} on git commit <code>{{ . }}</code>{{end}}.
</em></p>

{{ end }}
