package pdftk

import "strings"

type Permission string

const (
	Printing         Permission = "Printing"
	DegradedPrinting Permission = "DegradedPrinting"
	Assembly         Permission = "Assembly"
	CopyContents     Permission = "CopyContents"
	ScreenReaders    Permission = "ScreenReaders"
	ModifyContents   Permission = "ModifyContents"
	FillIn           Permission = "FillIn"
)

type Option func(cmd command)

// Set executable name instead of using default "pdftk"
func OptionExecutable(name string) Option {
	return func(cmd command) {
		// replace the command entirely
		c := createCmd(name, cmd.Stdout, cmd.Stdin, cmd.Args...)
		cmd.Cmd = c.Cmd
	}
}

// Flatten the PDF before output
func OptionFlatten() Option {
	return func(cmd command) {
		cmd.Args = append(cmd.Args, "flatten")
	}
}
func OptionAllow(perms ...Permission) Option {
	return func(cmd command) {
		strPer4ms := unwrapPerms(perms)

		cmd.Args = append(cmd.Args, "allow"+strings.Join(strPer4ms, " "))

	}
}

func unwrapPerms(perms []Permission) []string {
	unwrapped := make([]string, len(perms))
	for i, p := range perms {
		unwrapped[i] = string(p)
	}
	return unwrapped
}
