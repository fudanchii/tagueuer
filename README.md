# tagueuer

build status: [![CircleCI](https://circleci.com/gh/fudanchii/tagueuer.svg?style=svg)](https://circleci.com/gh/fudanchii/tagueuer)

(/tag'ju:/ _tagyew_)

Populate struct with declarative struct tags.

Tagueuer provides infrastructure to populate struct from other data source, such as environment variables etc.
You can define how the struct get populated by specifying callback which returns string value, the string value
will then get parsed into its proper field data type.

examples:

```go

// Let's have a struct populated wiht value from environment variables

type AppConfig struct {
    Host string `env:"HOST" required:"true"`
    Port int    `env:"PORT" requierd:"true"`
}

func newConfig() *AppConfig {
    var config AppConfig

    envParser := tagueuer.New()
    envParser.On("env", func (c *tagueuer.Context) (string, error) {
	    return os.Getenv(strings.ToUpper(c.TagValue("env"))), nil
    })

    envParser.ParseInto(&config)
    return &config
}

func main() {
    config := newConfig()

    fmt.Printf("address: %s:%d", config.Host, config.Port)
}

```
