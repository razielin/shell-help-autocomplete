package shell_help_autocomplete

import (
	"github.com/kusabashira/acgen"
	"regexp"
	s "strings"
)

func ParseArgsFromString(input string) []acgen.Flag {
	var res []acgen.Flag
	args := parseArgs(input)
	for _, arg := range args {
		res = append(res, mapToFlag(arg))
	}
	return res
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

type Arg struct {
	description string
	longArg     string
	shortArg    string
}

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
