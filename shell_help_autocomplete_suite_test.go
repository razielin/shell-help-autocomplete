package shell_help_autocomplete_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestShellHelpAutocomplete(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ShellHelpAutocomplete Suite")
}
