package shell_help_autocomplete_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"regexp"
	s "strings"
)

var _ = Describe("ShellHelpAutocomplete", func() {
	emptyArg := Arg{}
	Describe("parseArgs", func() {
		var emptyArgArray []Arg
		It("returns empty array if input is empty", func() {
			Expect(parseArgs("")).To(Equal(emptyArgArray))
		})
		It("returns empty array if input has no args", func() {
			Expect(parseArgs(`irrelevant text
				- some other line
			another-line`)).To(Equal(emptyArgArray))
		})
		It("should parse a single arg in a line", func() {
			Expect(parseArgs("-v    Print version")).To(HaveLen(1))
		})
		It("should parse multiple arg in multiple lines", func() {
			Expect(parseArgs(`-v    Print version

-g  -g opt
irrelevant - text`)).To(HaveLen(2))
		})
	})
	Describe("parseArgLine", func() {
		Context("in case of invalid input without args", func() {
			It("if string is empty", func() {
				Expect(parseArgLine("")).To(Equal(emptyArg))
			})
			It("if string does not contain a valid arg", func() {
				Expect(parseArgLine("some irrelevant - text")).To(Equal(emptyArg))
			})
		})

		Context("in case of valid input which", func() {
			It("has string containing short arg and description", func() {
				Expect(parseArgLine("-s    a short arg")).To(Equal(Arg{shortArg: "-s", description: "a short arg"}))
			})
			It("has string containing long arg and description", func() {
				Expect(parseArgLine("--long a long arg")).To(Equal(Arg{longArg: "--long", description: "a long arg"}))
			})
			It("has string containing both short and long arg and description", func() {
				Expect(parseArgLine("-s, --long short and long arg")).To(Equal(Arg{
					shortArg: "-s", longArg: "--long", description: "short and long arg",
				}))
			})
			It("has string containing both short and long arg and description with hyphen", func() {
				Expect(parseArgLine("-s, --long - 1 short and 1 long - the arg")).To(Equal(Arg{
					shortArg: "-s", longArg: "--long", description: "1 short and 1 long - the arg",
				}))
			})
		})

	})

})

func parseArgLine(line string) Arg {
	line = s.TrimSpace(line)
	if !s.HasPrefix(line, "-") {
		return Arg{}
	}
	shortArg := regexp.MustCompile(`^-\w+`).FindString(line)
	longArg := regexp.MustCompile(`--\w+`).FindString(line)
	if shortArg == "" && longArg == "" {
		return Arg{}
	}

	description := line
	description = s.Replace(description, shortArg, "", -1)
	description = s.Replace(description, longArg, "", -1)
	description = regexp.MustCompile(`^\W+\s`).ReplaceAllString(description, "")
	description = s.TrimSpace(description)

	if description != "" {
		arg := Arg{description: description}
		if shortArg != "" {
			arg.shortArg = shortArg
		}
		if longArg != "" {
			arg.longArg = longArg
		}
		return arg
	}

	return Arg{}
}

func parseArgs(input string) []Arg {
	var res []Arg
	emptyArg := Arg{}
	lines := s.Split(input, "\n")
	for _, line := range lines {
		arg := parseArgLine(line)
		if arg != emptyArg {
			res = append(res, arg)
		}
	}
	return res
}

type Arg struct {
	description string
	longArg     string
	shortArg    string
}
