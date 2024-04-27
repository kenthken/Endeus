package main

import (
	"endeus/api/config"
	"endeus/api/route"
	"endeus/generate"
	"os"
)

func main() {
	args := os.Args
	env := ""
	if len(args) > 1 {
		env = args[1]
	}

	if env == "migrate" {
		generate.Migrate()

	} else {
		config.InitEnvConfigs()
		db := config.InitDB()
		config.InitLogger(db)
		route.StartRouting(db)
	}
}
