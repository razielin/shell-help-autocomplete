package shell_help_autocomplete

import (
	"github.com/kusabashira/acgen"
	"log"
	"os"
)

func main() {
	flag := acgen.Flag{Arg: "-t - arg name", Description: "-t arg description", Long: []string{"--t"}, Short: []string{"-t"}}
	command := acgen.Command{Name: "test command", Flags: []*acgen.Flag{&flag}}
	generator, err := acgen.LookGenerator("bash")
	if err != nil {
		log.Fatal("Can't find generator", err)
	}
	err = generator(os.Stdout, &command)
	if err != nil {
		log.Fatal("Error during generation", err)
	}
}
