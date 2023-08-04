package style

import (
	"time"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

func NewSpinner(cmd *cobra.Command) pterm.SpinnerPrinter {
	printer := pterm.SpinnerPrinter{
		Sequence: []string{
			"▁",
			"▃",
			"▄",
			"▅",
			"▆",
			"▇",
			"▆",
			"▅",
			"▄",
			"▃",
		},
		Style:               &pterm.ThemeDefault.SpinnerStyle,
		Delay:               time.Millisecond * 100,
		ShowTimer:           false,
		TimerRoundingFactor: time.Second,
		TimerStyle:          &pterm.ThemeDefault.TimerStyle,
		MessageStyle:        &pterm.ThemeDefault.SpinnerTextStyle,
		InfoPrinter:         &pterm.PrefixPrinter{},
		SuccessPrinter:      &pterm.PrefixPrinter{},
		FailPrinter:         &pterm.PrefixPrinter{},
		WarningPrinter:      &pterm.PrefixPrinter{},
		Writer:              cmd.OutOrStderr(),
		RemoveWhenDone:      true,
	}
	return printer
}
