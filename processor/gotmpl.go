package processor

import (
	"bytes"
	"context"
	"text/template"

	"github.com/benthosdev/benthos/v4/public/service"
)

func init() {

	configSpec := service.NewConfigSpec().
		Beta().
		Summary("Renders a go template").
		Field(service.NewStringField("template").Description("Template string").Example(`Hello {{.name}}`))

	constructor := func(conf *service.ParsedConfig, mgr *service.Resources) (service.Processor, error) {

		tmpl, err := conf.FieldString("template")

		if err != nil {
			return nil, err
		}

		return getGoTmplProcessor(tmpl, mgr.Logger())

	}

	err := service.RegisterProcessor("gotmpl", configSpec, constructor)
	if err != nil {
		panic(err)
	}
}

type gotmplProcessor struct {
	template *template.Template
	logger   *service.Logger
}

func getGoTmplProcessor(text string, logger *service.Logger) (*gotmplProcessor, error) {

	t, err := template.New("").Parse(text)

	if err != nil {
		return nil, err
	}

	return &gotmplProcessor{
		template: t,
		logger:   logger,
	}, nil

}

func (r *gotmplProcessor) Process(ctx context.Context, m *service.Message) (service.MessageBatch, error) {

	var buf bytes.Buffer

	doc, err := m.AsStructured()

	if err != nil {
		return nil, err
	}

	err = r.template.Execute(&buf, doc)

	if err != nil {
		return nil, err
	}

	return service.MessageBatch{service.NewMessage(buf.Bytes())}, nil
}

func (r *gotmplProcessor) Close(ctx context.Context) error {
	return nil
}
