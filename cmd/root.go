package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/thatmattlove/addr/cmd/style"
	"github.com/thatmattlove/addr/internal/util"
	addr "github.com/thatmattlove/addr/pkg"
)

func Init(version string) *cobra.Command {
	root := &cobra.Command{
		Use:     "addr",
		Short:   "addr is a tool to look up IP & ASN ownership and routing information.",
		Args:    cobra.ArbitraryArgs,
		Version: version,
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
						cmd.PrintErr(err.Error() + "\n")
						os.Exit(1)
					}
					ptrs, _ := addr.DNSReverseLookup(r.IP)
					p.Stop()
					cmd.Println(style.IPBox(r, ptrs))
				} else if util.IsASN(arg) {
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
			os.Exit(0)
		},
	}
	root.AddCommand(ASNCmd, IPCmd)
	return root
}
