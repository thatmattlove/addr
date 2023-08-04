package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/thatmattlove/addr/addr"
	"github.com/thatmattlove/addr/cmd/style"
	"github.com/thatmattlove/addr/internal/util"
)

var IPCmd *cobra.Command = &cobra.Command{
	Use:   "ip",
	Short: "Look up an IP address or prefix",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
		s := style.NewSpinner(cmd)
		for _, arg := range args {
			if util.IsIP(arg) {
				p, _ := s.Start()
				r, err := addr.QueryIPPrefix(arg)
				if err != nil {
					p.Stop()
					cmd.PrintErr(err.Error())
					os.Exit(1)
				}
				ptrs, _ := addr.DNSReverseLookup(r.IP)
				p.Stop()
				cmd.Println(style.IPBox(r, ptrs))
			} else {
				cmd.PrintErrf("invalid argument '%s'", arg)
				os.Exit(1)
			}
		}
	},
}
