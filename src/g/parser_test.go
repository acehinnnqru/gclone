package g

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParse(t *testing.T) {
	cases := []string{
		"https://github.com/spf13/cobra.git",
		"git@github.com:spf13/cobra.git",
	}
	addr := Address{
		Server:     "github.com",
		Namespace:  "spf13",
		Repository: "cobra",
	}

	for _, c := range cases {
		addrGet, e := Parse(c)

		require.Nil(t, e)
		require.Equal(t, addr, *addrGet)
	}
}
