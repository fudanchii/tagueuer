package envconfig

import (
	"os"
	"strings"

	"github.com/fudanchii/tagueuer"
)

var envParser = tagueuer.New()

func init() {
	envParser.On("envconfig", readEnv)
}

func readEnv(c *tagueuer.Context) (string, error) {
	return os.Getenv(strings.ToUpper(c.TagValue("envconfig"))), nil
}

func ParseInto(obj interface{}) error {
	return envParser.ParseInto(obj)
}
