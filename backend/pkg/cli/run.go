package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
)

func Run(cmd *cobra.Command) int {
	if initialized, err := run(cmd); err != nil {
		if !initialized {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		} else {
			klog.ErrorS(err, "command failed")
		}
		return 1
	}
	return 0
}

func run(cmd *cobra.Command) (initialized bool, err error) {
	defer klog.Flush()
	switch {
	case cmd.PersistentPreRun != nil:
		pre := cmd.PersistentPreRun
		cmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
			initialized = true
			pre(cmd, args)
		}
	case cmd.PersistentPreRunE != nil:
		pre := cmd.PersistentPreRunE
		cmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
			initialized = true
			return pre(cmd, args)
		}
	default:
		cmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
			initialized = true
		}
	}
	err = cmd.Execute()
	return
}
