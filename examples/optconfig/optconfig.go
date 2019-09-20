package optconfig

import "github.com/fudanchii/tagueuer"

type OptFlag struct {
	short string
}

type OptConfig struct {
	*tagueuer.Tagueuer
	flags []OptFlag
}

func New() *tagueuer.Tagueuer {
	optParser := tagueuer.New()
}
