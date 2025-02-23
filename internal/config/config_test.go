package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	KeyTgToken         = "TG_TOKEN"
	KeyVrabberHost     = "VRABBER_HOST"
	KeyVrabberPort     = "VRABBER_PORT"
	KeyVrabberTimout   = "VRABBER_TIMEOUT"
	KeyMessagesBuffer  = "MESSAGES_BUFFER"
	KeyResponsesBuffer = "RESPONSES_BUFFER"
)

func TestLoad(t *testing.T) {
	cases := []struct {
		name      string
		mustPanic bool
		vars      map[string]string
	}{
		{
			name:      "all filled",
			mustPanic: false,
			vars: map[string]string{
				KeyTgToken:         "TG_TOKEN",
				KeyVrabberHost:     "testhost",
				KeyVrabberPort:     "8080",
				KeyMessagesBuffer:  "200",
				KeyResponsesBuffer: "200",
				KeyVrabberTimout:   "1",
			},
		},
		{
			name:      "minimum fields filled",
			mustPanic: false,
			vars: map[string]string{
				KeyTgToken:     "TG_TOKEN",
				KeyVrabberPort: "8080",
			},
		},
		{
			name:      "empty token",
			mustPanic: true,
			vars: map[string]string{
				KeyTgToken:         "",
				KeyVrabberHost:     "localhost",
				KeyVrabberPort:     "8080",
				KeyMessagesBuffer:  "200",
				KeyResponsesBuffer: "200",
			},
		},
		{
			name:      "no token",
			mustPanic: true,
			vars: map[string]string{
				KeyVrabberHost:     "localhost",
				KeyVrabberPort:     "8080",
				KeyMessagesBuffer:  "200",
				KeyResponsesBuffer: "200",
			},
		},
		{
			name: "no vrabber host",
			vars: map[string]string{
				KeyTgToken:         "TG_TOKEN",
				KeyVrabberPort:     "8080",
				KeyMessagesBuffer:  "200",
				KeyResponsesBuffer: "200",
			},
		},
		{
			name:      "empty vrabber host",
			mustPanic: true,
			vars: map[string]string{
				KeyTgToken:         "TG_TOKEN",
				KeyVrabberHost:     "",
				KeyVrabberPort:     "8080",
				KeyMessagesBuffer:  "200",
				KeyResponsesBuffer: "200",
			},
		},
		{
			name:      "incorrect vrabber port",
			mustPanic: true,
			vars: map[string]string{
				KeyTgToken:         "TG_TOKEN",
				KeyVrabberHost:     "localhost",
				KeyVrabberPort:     "test",
				KeyMessagesBuffer:  "200",
				KeyResponsesBuffer: "200",
			},
		},
		{
			name:      "empty vrabber port",
			mustPanic: true,
			vars: map[string]string{
				KeyTgToken:         "TG_TOKEN",
				KeyVrabberHost:     "localhost",
				KeyVrabberPort:     "",
				KeyMessagesBuffer:  "200",
				KeyResponsesBuffer: "200",
			},
		},
		{
			name:      "vrabber port too big",
			mustPanic: true,
			vars: map[string]string{
				KeyTgToken:         "TG_TOKEN",
				KeyVrabberHost:     "localhost",
				KeyVrabberPort:     "65536",
				KeyMessagesBuffer:  "200",
				KeyResponsesBuffer: "200",
			},
		},
		{
			name:      "vrabber port too small",
			mustPanic: true,
			vars: map[string]string{
				KeyTgToken:         "TG_TOKEN",
				KeyVrabberHost:     "localhost",
				KeyVrabberPort:     "0",
				KeyMessagesBuffer:  "200",
				KeyResponsesBuffer: "200",
			},
		},
		{
			name:      "vrabber port negative",
			mustPanic: true,
			vars: map[string]string{
				KeyTgToken:         "TG_TOKEN",
				KeyVrabberHost:     "localhost",
				KeyVrabberPort:     "-1",
				KeyMessagesBuffer:  "200",
				KeyResponsesBuffer: "200",
			},
		},
		{
			name:      "incorrect messages buffer",
			mustPanic: true,
			vars: map[string]string{
				KeyTgToken:         "TG_TOKEN",
				KeyVrabberHost:     "localhost",
				KeyVrabberPort:     "8080",
				KeyMessagesBuffer:  "test",
				KeyResponsesBuffer: "200",
			},
		},
		{
			name:      "empty messages buffer",
			mustPanic: true,
			vars: map[string]string{
				KeyTgToken:         "TG_TOKEN",
				KeyVrabberHost:     "localhost",
				KeyVrabberPort:     "8080",
				KeyMessagesBuffer:  "",
				KeyResponsesBuffer: "200",
			},
		},
		{
			name:      "messages buffer too big",
			mustPanic: true,
			vars: map[string]string{
				KeyTgToken:         "TG_TOKEN",
				KeyVrabberHost:     "localhost",
				KeyVrabberPort:     "8080",
				KeyMessagesBuffer:  "20001",
				KeyResponsesBuffer: "200",
			},
		},
		{
			name:      "messages buffer too small",
			mustPanic: true,
			vars: map[string]string{
				KeyTgToken:         "TG_TOKEN",
				KeyVrabberHost:     "localhost",
				KeyVrabberPort:     "8080",
				KeyMessagesBuffer:  "0",
				KeyResponsesBuffer: "200",
			},
		},
		{
			name:      "messages buffer negative",
			mustPanic: true,
			vars: map[string]string{
				KeyTgToken:         "TG_TOKEN",
				KeyVrabberHost:     "localhost",
				KeyVrabberPort:     "8080",
				KeyMessagesBuffer:  "-1",
				KeyResponsesBuffer: "200",
			},
		},
		{
			name:      "incorrect responses buffer",
			mustPanic: true,
			vars: map[string]string{
				KeyTgToken:         "TG_TOKEN",
				KeyVrabberHost:     "localhost",
				KeyVrabberPort:     "8080",
				KeyMessagesBuffer:  "200",
				KeyResponsesBuffer: "test",
			},
		},
		{
			name:      "empty responses buffer",
			mustPanic: true,
			vars: map[string]string{
				KeyTgToken:         "TG_TOKEN",
				KeyVrabberHost:     "localhost",
				KeyVrabberPort:     "8080",
				KeyMessagesBuffer:  "200",
				KeyResponsesBuffer: "",
			},
		},
		{
			name:      "responses buffer too big",
			mustPanic: true,
			vars: map[string]string{
				KeyTgToken:         "TG_TOKEN",
				KeyVrabberHost:     "localhost",
				KeyVrabberPort:     "8080",
				KeyMessagesBuffer:  "200",
				KeyResponsesBuffer: "2001",
			},
		},
		{
			name:      "responses buffer too small",
			mustPanic: true,
			vars: map[string]string{
				KeyTgToken:         "TG_TOKEN",
				KeyVrabberHost:     "localhost",
				KeyVrabberPort:     "8080",
				KeyMessagesBuffer:  "200",
				KeyResponsesBuffer: "0",
			},
		},
		{
			name:      "responses buffer negative",
			mustPanic: true,
			vars: map[string]string{
				KeyTgToken:         "TG_TOKEN",
				KeyVrabberHost:     "localhost",
				KeyVrabberPort:     "8080",
				KeyMessagesBuffer:  "200",
				KeyResponsesBuffer: "-1",
			},
		},
		{
			name:      "incorrect server timeout",
			mustPanic: true,
			vars: map[string]string{
				KeyTgToken:       "TG_TOKEN",
				KeyVrabberPort:   "8080",
				KeyVrabberTimout: "test",
			},
		},
		{
			name:      "empty server timeout",
			mustPanic: true,
			vars: map[string]string{
				KeyTgToken:       "TG_TOKEN",
				KeyVrabberPort:   "8080",
				KeyVrabberTimout: "",
			},
		},
		{
			name:      "server timeout too big",
			mustPanic: true,
			vars: map[string]string{
				KeyTgToken:       "TG_TOKEN",
				KeyVrabberPort:   "8080",
				KeyVrabberTimout: "2001",
			},
		},
		{
			name:      "server timeout too small",
			mustPanic: true,
			vars: map[string]string{
				KeyTgToken:       "TG_TOKEN",
				KeyVrabberPort:   "8080",
				KeyVrabberTimout: "0",
			},
		},
		{
			name:      "server timeout negative",
			mustPanic: true,
			vars: map[string]string{
				KeyTgToken:       "TG_TOKEN",
				KeyVrabberPort:   "8080",
				KeyVrabberTimout: "-1",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			vars := []string{
				KeyTgToken,
				KeyVrabberHost,
				KeyVrabberPort,
				KeyMessagesBuffer,
				KeyResponsesBuffer,
				KeyVrabberTimout,
			}

			for _, v := range vars {
				if val, exists := c.vars[v]; exists {
					t.Setenv(v, val)
				}
			}

			defer func() {
				err := recover()

				if c.mustPanic {
					assert.NotNil(t, err)
				} else {
					assert.Nil(t, err)
				}
			}()

			conf := MustLoad()
			assert.NotNil(t, conf)

			if _, exists := c.vars[KeyVrabberHost]; !exists {
				assert.Equal(t, conf.ServerHost, "localhost")
			}
			if _, exists := c.vars[KeyMessagesBuffer]; !exists {
				assert.Equal(t, conf.MessagesBuffer, 100)
			}
			if _, exists := c.vars[KeyResponsesBuffer]; !exists {
				assert.Equal(t, conf.ResponsesBuffer, 100)
			}
		})
	}
}
