package envconfig

import (
	"os"
	"strings"

	"github.com/fudanchii/tagueuer"
)

func New() *tagueuer.Tagueuer {
	envParser := tagueuer.New()
	envParser.On("envconfig", readEnv)
	return envParser
}

func readEnv(c *tagueuer.Context) (string, error) {
	return os.Getenv(strings.ToUpper(c.TagValue("envconfig"))), nil
}
