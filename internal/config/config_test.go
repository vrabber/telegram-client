package config

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestLoad(t *testing.T) {
	cases := []struct {
		name        string
		token       string
		vrabberHost string
		vrabberPort string
		mustPanic   bool
	}{
		{
			name:        "correct",
			token:       "test_token",
			vrabberHost: "localhost",
			vrabberPort: "9090",
			mustPanic:   false,
		},
		{
			name:        "empty token",
			token:       "",
			vrabberHost: "localhost",
			vrabberPort: "9090",
			mustPanic:   true,
		},
		{
			name:        "incorrect port",
			token:       "test_token",
			vrabberHost: "localhost",
			vrabberPort: "-10",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Setenv("TG_TOKEN", c.token)
			t.Setenv("VRABBER_HOST", c.vrabberHost)
			t.Setenv("VRABBER_PORT", c.vrabberPort)

			defer func() {
				err := recover()

				if c.mustPanic {
					assert.NotNil(t, err)
				} else {
					assert.Nil(t, err)
				}
			}()

			conf := MustLoad()
			assert.NotNil(t, c)

			assert.Equal(t, c.token, conf.TgToken)
			assert.Equal(t, c.vrabberHost, conf.ServerHost)

			p, err := strconv.Atoi(c.vrabberPort)
			assert.Nil(t, err)
			assert.Equal(t, p, conf.ServerPort)
		})
	}
}
