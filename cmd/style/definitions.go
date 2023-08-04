package style

import (
	"fmt"

	"github.com/pterm/pterm"
	"github.com/thatmattlove/addr/addr"
)

var Box = &pterm.BoxPrinter{
	VerticalString:          "│",
	TopRightCornerString:    "╰",
	TopLeftCornerString:     "╯",
	BottomLeftCornerString:  "╮",
	BottomRightCornerString: "╭",
	HorizontalString:        "─",
	BoxStyle:                &pterm.Style{pterm.FgGray},
	TextStyle:               &pterm.ThemeDefault.BoxTextStyle,
	RightPadding:            3,
	LeftPadding:             3,
	TopPadding:              1,
	BottomPadding:           1,
	TitleTopLeft:            true,
}

var Wrapper = &pterm.BoxPrinter{
	VerticalString:          "",
	TopRightCornerString:    "",
	TopLeftCornerString:     "",
	BottomLeftCornerString:  "",
	BottomRightCornerString: "",
	HorizontalString:        "",
	BoxStyle:                &pterm.ThemeDefault.BoxStyle,
	TextStyle:               &pterm.ThemeDefault.BoxTextStyle,
	RightPadding:            1,
	LeftPadding:             1,
	TopPadding:              0,
	BottomPadding:           0,
	TitleTopLeft:            false,
}

var Title = pterm.NewStyle(pterm.Bold, pterm.FgLightRed).Sprintf

var Subtitle = pterm.NewStyle(pterm.Italic, pterm.FgLightRed).Sprintf

var Highlight1 = pterm.NewStyle(pterm.Bold, pterm.FgLightCyan).Sprintf

var Highlight2 = pterm.NewStyle(pterm.Bold, pterm.FgLightMagenta).Sprintf

var Subtle = pterm.NewStyle(pterm.FgGray).Sprintf

var Plain = pterm.NewStyle(pterm.FgWhite).Sprintf

func Country(r *addr.Response) string {
	return fmt.Sprintf("%s %s", Plain(r.Name), r.Country.Emoji())
}
