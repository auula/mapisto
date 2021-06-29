package cmd

import "github.com/fatih/color"

var (
	version = "0.0.1 Alpha"
	banner  = color.MagentaString("\n" +
		"    __  ___            _      __      \n" +
		"   /  |/  /___ _____  (_)____/ /_____ \n" +
		"  / /|_/ / __ `/ __ \\/ / ___/ __/ __ \\\n" +
		" / /  / / /_/ / /_/ / (__  ) /_/ /_/ /\n" +
		"/_/  /_/\\__,_/ .___/_/____/\\__/\\____/ \n" +
		"            /_/   " + version + "\n")
	description = color.GreenString(`Mapisto is a command-line database tool that can generate tables
with the type of Golang, Python,Rust, Java, TypeScript language,
with the command line.`)
)
