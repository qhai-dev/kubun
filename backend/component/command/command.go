package command

import (
	"fmt"
	"os"

	"github.com/qhai-dev/kaka/component/log"
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
)

func Run(cmd *cobra.Command) int {
	if initialized, err := preRun(cmd); err != nil {
		if !initialized {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		} else {
			klog.ErrorS(err, "command failed")
		}
		return 1
	}
	return 0
}

func preRun(cmd *cobra.Command) (initialized bool, err error) {
	defer log.FlushLog()

	switch {
	case cmd.PersistentPreRun != nil:
		pre := cmd.PersistentPreRun
		cmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
			log.InitLog()
			initialized = true
			pre(cmd, args)
		}
	default:
		cmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
			log.InitLog()
			initialized = true
		}
	}
	err = cmd.Execute()
	return
}
