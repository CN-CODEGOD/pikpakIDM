package main

import (
	_ "embed"

	"github.com/pikpakidm/pikpakidm/cli"
	"github.com/pikpakidm/pikpakidm/cli/setup"
)

//go:embed config_example.yml
var configTemplate []byte

func main() {
	setup.SetConfigTemplate(configTemplate)
	cli.Execute()
}
