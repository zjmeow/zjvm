package main
import "flag"

type Cmd struct {
	helpFlag    bool
	versionFlag bool
	CpOption    string
	class       string
	XjreOption  string
	args        []string
}

func parseCmd() *Cmd {
	cmd := &Cmd{}
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version")
	flag.StringVar(&cmd.CpOption, "CpOption", "", "classpath")
	flag.StringVar(&cmd.CpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.CpOption, "Xjre", "", "classpath")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}
