package main

import (
	"github.com/kusabashira/acgen"
	"log"
	"os"
)

func main() {
	generator, err := acgen.LookGenerator("fish")
	if err != nil {
		log.Fatal(err)
	}
	command := &acgen.Command{
		Name: "sed",
		Flags: []*acgen.Flag{
			&acgen.Flag{
				Short:       []string{"n"},
				Long:        []string{"quiet", "silent"},
				Description: "suppress automatic printing of pattern space",
			},
			&acgen.Flag{
				Short:       []string{"e"},
				Long:        []string{"expression"},
				Arg:         "script",
				Description: "add the script to the commands to be executed",
			},
		},
	}
	err = generator(os.Stdout, command)
	if err != nil {
		log.Fatal(err)
	}
}
