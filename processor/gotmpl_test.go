package processor

import (
	"context"
	"testing"

	"github.com/benthosdev/benthos/v4/public/service"
	"github.com/stretchr/testify/require"
)

func TestGoTmplProcessor(t *testing.T) {

	proc, err := getGoTmplProcessor("Hello {{.name}}", nil)

	require.NoError(t, err)

	res, err := proc.Process(context.TODO(), service.NewMessage([]byte(`{"name":"Bob"}`)))

	require.NoError(t, err)

	msg, err := res[0].AsBytes()

	require.NoError(t, err)

	require.Equal(t, []byte("Hello Bob"), msg)

}
