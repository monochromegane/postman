package main

import (
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/monochromegane/postman"
)

var opts postman.Option

func main() {
	parser := flags.NewParser(&opts, flags.Default)
	parser.Name = "postman"
	parser.Usage = "[OPTIONS]"

	_, err := parser.Parse()
	if err != nil {
		os.Exit(1)
	}

	postman := postman.NewPostman(opts.Dir)
	postman.Run()
}
