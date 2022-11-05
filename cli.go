package cli

import (
	"os"
	"strconv"
	"strings"
)

type Command struct {
	Name     string
	Function func(args, flags []string)
}

type CmdFunc func(args, flags []string)

/*type Parsable interface {
	float32 | float64 | int | int32 | int64 | complex64 | complex128 | bool | string
}

type Flag[Type Parsable] struct {
	Name           string
	Default, Value Type
}*/

var (
	ARGS, FLAGS []string
)

var (
	commands []Command
)

func init() {
	for _, a := range os.Args[1:] {
		if strings.HasPrefix(a, "--") || strings.HasPrefix(a, "-") {
			trimmed := strings.TrimPrefix(strings.TrimPrefix(a, "--"), "-")
			FLAGS = append(FLAGS, trimmed)
		} else {
			ARGS = append(ARGS, a)
		}
	}
}

func Bool(name string, def bool) bool {
	for _, arg := range FLAGS {
		if strings.HasPrefix(arg, name) {
			eqSplit := strings.Split(arg, "=")
			if len(eqSplit) == 1 {
				return !def
			} else {
				val, err := strconv.ParseBool(eqSplit[len(eqSplit)-1])
				if err != nil {
					return def
				}
				return val
			}
		}
	}
	return def
}

func Int(name string, def int) int {
	for _, arg := range FLAGS {
		if strings.HasPrefix(arg, name) {
			eqSplit := strings.Split(arg, "=")
			if len(eqSplit) == 1 {
				return def
			} else {
				val, err := strconv.Atoi(eqSplit[len(eqSplit)-1])
				if err != nil {
					return def
				}
				return val
			}
		}
	}

	return def
}

func Float[T float32 | float64](name string, def float64) float64 {
	for _, arg := range FLAGS {
		if strings.HasPrefix(arg, name) {
			eqSplit := strings.Split(arg, "=")
			if len(eqSplit) == 1 {
				return def
			} else {
				val, err := strconv.ParseFloat(eqSplit[len(eqSplit)-1], 64)
				if err != nil {
					return def
				}
				return val
			}
		}
	}

	return def
}

func Complex(name string, def complex128) complex128 {
	for _, arg := range FLAGS {
		if strings.HasPrefix(arg, name) {
			eqSplit := strings.Split(arg, "=")
			if len(eqSplit) == 1 {
				return def
			} else {
				val, err := strconv.ParseComplex(eqSplit[len(eqSplit)-1], 128)
				if err != nil {
					return def
				}
				return val
			}
		}
	}

	return def
}

func String(name string, def string) string {
	for _, arg := range FLAGS {
		if strings.HasPrefix(arg, name) {
			eqSplit := strings.Split(arg, "=")
			if len(eqSplit) == 1 {
				return def
			} else {
				return strings.ReplaceAll(strings.Join(eqSplit[1:], ""), "\\", "\\")
			}
		}
	}

	return def
}

func RegisterCommand(name string, function CmdFunc) {
	commands = append(commands, Command{name, function})
}

func Default(function CmdFunc) {
	commands = append(commands, Command{"", function})	
}

func Exec() {
	if len(ARGS) == 0 {
		for _, c := range commands {
			if c.Name == "" {
				c.Function([]string{}, FLAGS)
			}
		}
		return
	}

	for _, c := range commands {
		if c.Name == ARGS[0] {
			c.Function(ARGS[1:], FLAGS)
		}
	}
}
