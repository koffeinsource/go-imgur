package imgur

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClientCreationWithoutClientID(t *testing.T) {
	client, err := NewClient(new(http.Client), "", "")
	require.Error(t, err)
	require.Nil(t, client)

	client, err = NewClient(new(http.Client), "some client id", "")
	require.NoError(t, err)
	require.NotNil(t, client)
}
