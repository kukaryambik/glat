package arg

import (
	"errors"
	"strings"
)

var (
	t, s string
	err  error
)

// Parse - parse args for subcommand
func Parse(a []string) (string, string, error) {
	if len(a) == 0 {
		t, s, err = "", "", errors.New("Not enough args")
	} else if len(a) >= 1 {
		t = a[0]
		if len(a) > 1 {
			s = a[1]
		}
	}
	return t, s, err
}

// Split subject parts
func Split(subj string, sep string) (string, string) {
	name := subj[strings.LastIndex(subj, sep)+1:]
	tmp := strings.Split(subj, sep)
	namespace := strings.Join(tmp[:len(tmp)-1], sep)
	return namespace, name
}

// Cook arg for use
func Cook(raw string) string {
	cropped := Crop(raw, "/")
	cooked := strings.ReplaceAll(cropped, "/", "%2F")
	return cooked
}

// Crop arg
func Crop(arg string, sep string) string {
	first := string(arg[0])
	last := string(arg[len(arg)-1])

	if first == sep {
		arg = strings.TrimLeft(arg, string(arg[0]))
	}
	if last == sep {
		arg = string(arg[:len(arg)-1])
	}
	return arg
}
