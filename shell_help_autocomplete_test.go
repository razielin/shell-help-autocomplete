package shell_help_autocomplete

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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

func expectLineToEqual(line string, arg *Arg) {
	Expect(parseArgLine(line)).To(Equal(arg))
}
