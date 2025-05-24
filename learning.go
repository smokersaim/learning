package learning

import (
	"os"

	"github.com/droqsic/probe"
)

func IsTerminal() bool {
	return probe.IsTerminal(os.Stdout.Fd())
}

func IsCygwinTerminal() bool {
	return probe.IsCygwinTerminal(os.Stdout.Fd())
}
