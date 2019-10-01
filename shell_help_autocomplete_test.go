package shell_help_autocomplete_test

import (
	"github.com/kusabashira/acgen"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"regexp"
	s "strings"
)

var _ = Describe("ShellHelpAutocomplete", func() {
	Describe("parseArgs", func() {
		It("returns empty array if input is empty", func() {
			Expect(parseArgs("")).To(HaveLen(0))
		})
		It("returns empty array if input has no args", func() {
			Expect(parseArgs(`irrelevant text
				- some other line
			another-line`)).To(HaveLen(0))
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
				expectLineToEqual("", nil)

			})
			It("if string does not contain a valid arg", func() {
				expectLineToEqual("some irrelevant - text", nil)
			})
		})

		Context("in case of valid input which", func() {
			It("has string containing short arg and description", func() {
				expectLineToEqual("-s    a short arg", &Arg{shortArg: "-s", description: "a short arg"})
			})
			It("has string containing long arg and description", func() {
				expectLineToEqual("--long a long arg", &Arg{longArg: "--long", description: "a long arg"})
			})
			It("has string containing both short and long arg and description", func() {
				expectLineToEqual("-s, --long short and long arg", &Arg{
					shortArg: "-s", longArg: "--long", description: "short and long arg",
				})
			})
			It("has string containing both short and long arg and description with hyphen", func() {
				expectLineToEqual("-s, --long - 1 short and 1 long - the arg", &Arg{
					shortArg: "-s", longArg: "--long", description: "1 short and 1 long - the arg",
				})
			})
		})

	})
	Describe("mapToFlag", func() {
		It("should map Arg to acgen.Flag", func() {
			arg := Arg{
				description: "desc",
				longArg:     "--desc",
				shortArg:    "-d",
			}
			flag := mapToFlag(arg)
			Expect(flag.Description).To(Equal(arg.description))
			Expect(flag.Long[0]).To(Equal(arg.longArg))
			Expect(flag.Short[0]).To(Equal(arg.shortArg))
		})
	})

})

func mapToFlag(arg Arg) acgen.Flag {
	flag := acgen.Flag{Description: arg.description, Arg: arg.description}
	if arg.shortArg != "" {
		flag.Short = append(flag.Short, arg.shortArg)
	}
	if arg.longArg != "" {
		flag.Long = append(flag.Long, arg.longArg)
	}
	return flag
}

func expectLineToEqual(line string, arg *Arg) {
	Expect(parseArgLine(line)).To(Equal(arg))
}

func parseArgLine(line string) *Arg {
	line = s.TrimSpace(line)
	if !s.HasPrefix(line, "-") {
		return nil
	}
	shortArg := regexp.MustCompile(`^-\w+`).FindString(line)
	longArg := regexp.MustCompile(`--\w+`).FindString(line)
	if shortArg == "" && longArg == "" {
		return nil
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
		return &arg
	}

	return nil
}

func parseArgs(input string) []Arg {
	var res []Arg
	lines := s.Split(input, "\n")
	for _, line := range lines {
		arg := parseArgLine(line)
		if arg != nil {
			res = append(res, *arg)
		}
	}
	return res
}

type Arg struct {
	description string
	longArg     string
	shortArg    string
}
