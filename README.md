# gotmpl benthos processor plugin

 Gotmpl is a [Benthos](https://www.benthos.dev/) processor plugin that allows you to use [Go templates](https://golang.org/pkg/text/template/) to render messages.

## Installation

``` sh
go get github.com/wincus/benthos-gotmpl-plugin
```

## Configuration example

``` yaml
input:
  generate:
    mapping: root = {"name":"John"}
    count: 1

pipeline:
  processors:
    - gotmpl:
        template: Hi There {{.name}}
```

```bash
$ go run main.go -c conf/config.yaml
Hi There John
```