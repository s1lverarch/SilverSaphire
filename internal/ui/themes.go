// SilverSaphire (ss) - Installer Themes
// Author: Silvercore (s1lverarch) <s1lv3rcore@proton.me>
// License: GPL-v3.0

package ui

import "github.com/charmbracelet/lipgloss"

// ── Theme struct ──────────────────────────────────────────────────────────────
type Theme struct {
	Name       string
	Icon       string
	Primary    lipgloss.Color
	Secondary  lipgloss.Color
	Accent     lipgloss.Color
	OK         lipgloss.Color
	Warn       lipgloss.Color
	Dim        lipgloss.Color
	Dim2       lipgloss.Color
	Background lipgloss.Color
	Text       lipgloss.Color
	Border     lipgloss.Color
	SelBG      lipgloss.Color
}

// ── Themes ────────────────────────────────────────────────────────────────────
var Themes = map[string]Theme{

	"silvercore": {
		Name:       "SilverCore",
		Icon:       "󰴨",
		Primary:    lipgloss.Color("#C0C0C0"),
		Secondary:  lipgloss.Color("#00BFFF"),
		Accent:     lipgloss.Color("#00FFD0"),
		OK:         lipgloss.Color("#00FF7F"),
		Warn:       lipgloss.Color("#FFD700"),
		Dim:        lipgloss.Color("#333333"),
		Dim2:       lipgloss.Color("#666666"),
		Background: lipgloss.Color("#0a0a0a"),
		Text:       lipgloss.Color("#EEEEEE"),
		Border:     lipgloss.Color("#444444"),
		SelBG:      lipgloss.Color("#051525"),
	},

	"crimson": {
		Name:       "Crimson",
		Icon:       "",
		Primary:    lipgloss.Color("#FF4466"),
		Secondary:  lipgloss.Color("#FF0033"),
		Accent:     lipgloss.Color("#FF6688"),
		OK:         lipgloss.Color("#FF4466"),
		Warn:       lipgloss.Color("#FF8800"),
		Dim:        lipgloss.Color("#2a0a0a"),
		Dim2:       lipgloss.Color("#663333"),
		Background: lipgloss.Color("#0d0000"),
		Text:       lipgloss.Color("#FFE0E0"),
		Border:     lipgloss.Color("#441111"),
		SelBG:      lipgloss.Color("#1a0505"),
	},

	"aurora": {
		Name:       "Aurora",
		Icon:       "󰔉",
		Primary:    lipgloss.Color("#88FFCC"),
		Secondary:  lipgloss.Color("#AA88FF"),
		Accent:     lipgloss.Color("#FF88CC"),
		OK:         lipgloss.Color("#88FFAA"),
		Warn:       lipgloss.Color("#FFCC44"),
		Dim:        lipgloss.Color("#1a1a2e"),
		Dim2:       lipgloss.Color("#4a4a6e"),
		Background: lipgloss.Color("#0d0d1a"),
		Text:       lipgloss.Color("#E0E0FF"),
		Border:     lipgloss.Color("#2a2a4a"),
		SelBG:      lipgloss.Color("#0a0a20"),
	},

	"monochrome": {
		Name:       "Monochrome",
		Icon:       "",
		Primary:    lipgloss.Color("#FFFFFF"),
		Secondary:  lipgloss.Color("#AAAAAA"),
		Accent:     lipgloss.Color("#CCCCCC"),
		OK:         lipgloss.Color("#FFFFFF"),
		Warn:       lipgloss.Color("#AAAAAA"),
		Dim:        lipgloss.Color("#222222"),
		Dim2:       lipgloss.Color("#555555"),
		Background: lipgloss.Color("#000000"),
		Text:       lipgloss.Color("#DDDDDD"),
		Border:     lipgloss.Color("#444444"),
		SelBG:      lipgloss.Color("#1a1a1a"),
	},

	"nebula": {
		Name:       "Nebula",
		Icon:       "󱁵",
		Primary:    lipgloss.Color("#7B68EE"),
		Secondary:  lipgloss.Color("#9B59B6"),
		Accent:     lipgloss.Color("#00CED1"),
		OK:         lipgloss.Color("#7B68EE"),
		Warn:       lipgloss.Color("#FF6B6B"),
		Dim:        lipgloss.Color("#1a0a2e"),
		Dim2:       lipgloss.Color("#4a3a6e"),
		Background: lipgloss.Color("#0d0a1a"),
		Text:       lipgloss.Color("#E8E0FF"),
		Border:     lipgloss.Color("#2a1a4a"),
		SelBG:      lipgloss.Color("#100a20"),
	},

	"forest": {
		Name:       "Forest",
		Icon:       "",
		Primary:    lipgloss.Color("#00FF88"),
		Secondary:  lipgloss.Color("#00CC66"),
		Accent:     lipgloss.Color("#88FFCC"),
		OK:         lipgloss.Color("#00FF88"),
		Warn:       lipgloss.Color("#FFCC00"),
		Dim:        lipgloss.Color("#0a1a0a"),
		Dim2:       lipgloss.Color("#336633"),
		Background: lipgloss.Color("#050d05"),
		Text:       lipgloss.Color("#E0FFE8"),
		Border:     lipgloss.Color("#1a3a1a"),
		SelBG:      lipgloss.Color("#051005"),
	},

	"ocean": {
		Name:       "Ocean",
		Icon:       "󰞍",
		Primary:    lipgloss.Color("#00BFFF"),
		Secondary:  lipgloss.Color("#0088CC"),
		Accent:     lipgloss.Color("#00FFFF"),
		OK:         lipgloss.Color("#00BFFF"),
		Warn:       lipgloss.Color("#FFD700"),
		Dim:        lipgloss.Color("#0a1020"),
		Dim2:       lipgloss.Color("#334466"),
		Background: lipgloss.Color("#050a15"),
		Text:       lipgloss.Color("#E0F0FF"),
		Border:     lipgloss.Color("#1a2a44"),
		SelBG:      lipgloss.Color("#051020"),
	},

	"void": {
		Name:       "Void",
		Icon:       "󰁔",
		Primary:    lipgloss.Color("#7733FF"),
		Secondary:  lipgloss.Color("#5500CC"),
		Accent:     lipgloss.Color("#AA55FF"),
		OK:         lipgloss.Color("#7733FF"),
		Warn:       lipgloss.Color("#FF3366"),
		Dim:        lipgloss.Color("#0d0020"),
		Dim2:       lipgloss.Color("#330055"),
		Background: lipgloss.Color("#050005"),
		Text:       lipgloss.Color("#EEE0FF"),
		Border:     lipgloss.Color("#220044"),
		SelBG:      lipgloss.Color("#0a0015"),
	},
}

// ── Theme order for picker ────────────────────────────────────────────────────
var ThemeOrder = []string{
	"silvercore",
	"crimson",
	"aurora",
	"monochrome",
	"nebula",
	"forest",
	"ocean",
	"void",
}

// ── Active theme ──────────────────────────────────────────────────────────────
var ActiveTheme = Themes["silvercore"]

func ApplyTheme(name string) {
	t, ok := Themes[name]
	if !ok {
		t = Themes["silvercore"]
	}
	ActiveTheme = t

	// update all color vars used across boot.go + installer.go
	colorSilver = t.Primary
	colorBlue   = t.Secondary
	colorCyan   = t.Accent
	colorGreen  = t.OK
	colorDim    = t.Dim
	colorDim2   = t.Dim2
	colorWhite  = t.Text
	colorGray   = t.Dim2

	// rebuild styles with new colors
	rebuildStyles()
}

func rebuildStyles() {
	// boot.go styles
	styleLogo       = lipgloss.NewStyle().Foreground(colorSilver).Bold(true)
	styleOKBracket  = lipgloss.NewStyle().Foreground(colorGreen).Bold(true)
	styleOKText     = lipgloss.NewStyle().Foreground(colorDim2)
	styleDim        = lipgloss.NewStyle().Foreground(colorDim)
	styleDivider    = lipgloss.NewStyle().Foreground(colorDim)

	// installer.go styles
	iStyleTitle     = lipgloss.NewStyle().Foreground(colorCyan).Bold(true)
	iStyleOK        = lipgloss.NewStyle().Foreground(colorGreen).Bold(true)
	iStyleSel       = lipgloss.NewStyle().Foreground(colorBlue).Bold(true)
	iStyleNorm      = lipgloss.NewStyle().Foreground(colorWhite)
	iStyleDim       = lipgloss.NewStyle().Foreground(colorDim2)
	iStyleGray      = lipgloss.NewStyle().Foreground(colorGray)
	iStyleWarn      = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFD700")).Bold(true)
	iStyleSkip      = lipgloss.NewStyle().Foreground(lipgloss.Color("#666666")).Italic(true)
	iStyleInput     = lipgloss.NewStyle().Foreground(colorWhite).Bold(true)
	iStyleHeader    = lipgloss.NewStyle().Foreground(colorSilver).Bold(true).
				BorderStyle(lipgloss.NormalBorder()).
				BorderBottom(true).BorderForeground(ActiveTheme.Border)
	iStyleFooter    = lipgloss.NewStyle().Foreground(colorDim2).Italic(true).
				BorderStyle(lipgloss.NormalBorder()).
				BorderTop(true).BorderForeground(ActiveTheme.Border)
	iStyleOptSel    = lipgloss.NewStyle().Foreground(colorBlue).Bold(true).
				Background(ActiveTheme.SelBG).
				PaddingLeft(2).PaddingRight(2)
	iStyleOptNorm   = lipgloss.NewStyle().Foreground(colorWhite).PaddingLeft(2).PaddingRight(2)
	iStyleOptSkip   = lipgloss.NewStyle().Foreground(lipgloss.Color("#666666")).Italic(true).
				PaddingLeft(2).PaddingRight(2)
}
