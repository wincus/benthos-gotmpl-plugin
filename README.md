# gotmpl benthos processor plugin

 Gotmpl is a [Benthos](https://www.benthos.dev/) processor plugin that allows you to use [Go templates](https://golang.org/pkg/text/template/) to render messages.

## Why?

Sometimes is easier to express content as a rendered go template.

## Installation

``` sh
go install github.com/wincus/benthos-gotmpl-plugin@latest
```

## Configuration example

``` yaml
input:
  generate:
    mapping: |
      root = {"name":"John", "toys": ["car", "phone", "cards"], "good": true}
    count: 1

pipeline:
  processors:
    - gotmpl:
        template: >
          Hi there {{ .name }}! Your toys are:
          {{ range .toys }}
            - {{ . -}}
          {{ end }}
          {{ if .good }}
          
          You have been a good boy!
          {{ end }}
```

```bash
$ go run main.go -c conf/config.yaml
Hi there John! Your toys are: 
  - car
  - phone
  - cards 
You have been a good boy! 
```