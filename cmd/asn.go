package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/thatmattlove/addr/addr"
	"github.com/thatmattlove/addr/cmd/style"
	"github.com/thatmattlove/addr/internal/util"
)

var ASNCmd *cobra.Command = &cobra.Command{
	Use:   "asn",
	Short: "Look up an ASN",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
		s := style.NewSpinner(cmd)
		for _, arg := range args {
			if util.IsASN(arg) {
				p, _ := s.Start()
				r, err := addr.QueryASN(arg)
				p.Stop()
				if err != nil {
					cmd.PrintErr(err.Error() + "\n")
					os.Exit(1)
				}
				cmd.Println(style.ASNBox(r))
			} else {
				cmd.PrintErrf("invalid argument '%s'\n", arg)
				os.Exit(1)
			}
		}
	},
}
