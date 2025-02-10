package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoad(t *testing.T) {
	cases := []struct {
		name  string
		value string
	}{
		{
			name:  "empty",
			value: "",
		},
		{
			name:  "normal",
			value: "test_value",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Setenv("TG_TOKEN", c.value)
			c, err := Load()
			assert.Nil(t, err)
			assert.NotNil(t, c)
			assert.Equal(t, c.TgToken, c.TgToken)
		})
	}
}
