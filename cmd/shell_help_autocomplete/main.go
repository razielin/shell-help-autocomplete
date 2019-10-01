package main

import (
	"github.com/kusabashira/acgen"
	shell_help_autocomplete "github.com/razielin/shell-help-autocomplete"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	input := string(bytes)
	flags := shell_help_autocomplete.ParseArgsFromString(input)

	generator, err := acgen.LookGenerator("fish")
	if err != nil {
		log.Fatal(err)
	}
	command := &acgen.Command{
		Name:  "sed",
		Flags: flags,
	}
	err = generator(os.Stdout, command)
	if err != nil {
		log.Fatal(err)
	}
}
