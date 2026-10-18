// SilverSaphire (ss) - Boot Animation & Main Menu
// Author: Silvercore (s1lverarch) <s1lv3rcore@proton.me>
// Website: https://s1lverarch.github.io/linuxdistro/
// License: GPL-v3.0

package ui

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// ── Colors ────────────────────────────────────────────────────────────────────
var (
	colorSilver = lipgloss.Color("#C0C0C0")
	colorBlue   = lipgloss.Color("#00BFFF")
	colorCyan   = lipgloss.Color("#00FFD0")
	colorGreen  = lipgloss.Color("#00FF7F")
	colorDim    = lipgloss.Color("#333333")
	colorDim2   = lipgloss.Color("#666666")
	colorWhite  = lipgloss.Color("#EEEEEE")
	colorGray   = lipgloss.Color("#888888")

	styleLogo = lipgloss.NewStyle().Foreground(colorSilver).Bold(true)

	styleOKBracket = lipgloss.NewStyle().Foreground(colorGreen).Bold(true)
	styleOKText    = lipgloss.NewStyle().Foreground(colorDim2)
	styleDim       = lipgloss.NewStyle().Foreground(colorDim)
	styleDivider   = lipgloss.NewStyle().Foreground(colorDim)
)

// ── Boot log ──────────────────────────────────────────────────────────────────
var bootLines = []string{
	"QNX Neutrino RTOS microkernel initializing...",
	"QNX kernel 8.0.0 SMP loaded",
	"CPU: Intel x86 SMP detected — 4 cores online",
	"Memory manager procnto-smp-instr OK",
	"Physical memory: 2048 MB mapped",
	"SilverArch Hardware Abstraction Layer ready",
	"devb-sdmmc block device driver OK",
	"eMMC 16GB internal storage detected",
	"USB subsystem io-usb-otg initialized",
	"Network stack io-pkt-v6-hc loaded",
	"SilverNebula repository connected",
	"Loading SilverSaphire installer engine...",
	"Disk subsystem ready",
	"Welcome to SilverArch Linux / QNX Utility",
}

// ── Phases ────────────────────────────────────────────────────────────────────
type phase int

const (
	phaseBoot phase = iota
	phaseMenu
)

type tickMsg time.Time
type nextLineMsg struct{}

// ── Model ─────────────────────────────────────────────────────────────────────
type BootModel struct {
	logo      string
	phase     phase
	logLines  []string
	logIdx    int
	cursor    int
	width     int
	height    int
	blinkTick int
	blink     bool
}

type menuEntry struct {
	icon  string
	label string
	desc  string
}

var menuItems = []menuEntry{
	{"", "Install SilverArch Linux", "start guided installation"},
	{"", "Load profile (TOML)", "reproducible install from file"},
	{"", "About SilverSaphire", "version and project info"},
	{"", "Exit", "quit the installer"},
}

func NewBootAnimation(logo string) BootModel {
	return BootModel{logo: logo, phase: phaseBoot}
}

func (m BootModel) Init() tea.Cmd {
	return tea.Tick(70*time.Millisecond, func(t time.Time) tea.Msg { return nextLineMsg{} })
}

func (m BootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case nextLineMsg:
		if m.phase == phaseBoot {
			if m.logIdx < len(bootLines) {
				m.logLines = append(m.logLines, bootLines[m.logIdx])
				m.logIdx++
				return m, tea.Tick(65*time.Millisecond, func(t time.Time) tea.Msg { return nextLineMsg{} })
			}
			return m, tea.Tick(400*time.Millisecond, func(t time.Time) tea.Msg { return tickMsg(t) })
		}

	case tickMsg:
		if m.phase == phaseBoot {
			m.phase = phaseMenu
		}
		m.blinkTick++
		m.blink = m.blinkTick%8 < 4
		return m, tea.Tick(100*time.Millisecond, func(t time.Time) tea.Msg { return tickMsg(t) })

	case tea.KeyMsg:
		if m.phase == phaseMenu {
			switch msg.String() {
			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}
			case "down", "j":
				if m.cursor < len(menuItems)-1 {
					m.cursor++
				}
			case "enter", " ":
				switch m.cursor {
				case 0:
					next := NewInstallerModel(m.width, m.height)
					return next, next.Init()
				case 3:
					return m, tea.Quit
				}
			case "q", "ctrl+c":
				return m, tea.Quit
			}
		}
	}
	return m, nil
}

func (m BootModel) View() string {
	if m.phase == phaseBoot {
		return m.viewBoot()
	}
	return m.viewMenu()
}

func (m BootModel) viewBoot() string {
	var b strings.Builder
	b.WriteString(styleLogo.Render(m.logo) + "\n")
	b.WriteString(styleDivider.Render("  "+strings.Repeat("─", 52)) + "\n\n")
	for _, line := range m.logLines {
		b.WriteString(styleOKBracket.Render("  [ OK ] ") + styleOKText.Render(line) + "\n")
	}
	if len(m.logLines) > 0 {
		cur := " "
		if m.blink {
			cur = "█"
		}
		b.WriteString(styleDim.Render("         " + cur))
	}
	return b.String()
}

func (m BootModel) viewMenu() string {
	w := m.width
	if w < 80 {
		w = 80
	}
	h := m.height
	if h < 24 {
		h = 24
	}

	// ── left panel: logo + info ───────────────────────────────────────────────
	leftW := w / 2

	logoStyled := styleLogo.Render(m.logo)

	infoBlock := strings.Join([]string{
		lipgloss.NewStyle().Foreground(colorCyan).Bold(true).Render("  SilverSaphire Installer"),
		lipgloss.NewStyle().Foreground(colorGray).Render("  Part Of SilverArch Linux Project  ·  v0.1.0  ·  GPL-v3.0"),
		"",
		lipgloss.NewStyle().Foreground(colorDim2).Render("  ") +
			lipgloss.NewStyle().Foreground(colorGray).Render("Author :  ") +
			lipgloss.NewStyle().Foreground(colorWhite).Render("Silvercore (s1lverarch)"),
		lipgloss.NewStyle().Foreground(colorDim2).Render("  ") +
			lipgloss.NewStyle().Foreground(colorGray).Render("E-mail :   ") +
			lipgloss.NewStyle().Foreground(colorWhite).Render("s1lv3rcore@proton.me"),
		lipgloss.NewStyle().Foreground(colorDim2).Render("  ") +
			lipgloss.NewStyle().Foreground(colorGray).Render("WebSite     ") +
			lipgloss.NewStyle().Foreground(colorBlue).Render("https://s1lverarch.github.io/linuxdistro"),
		"",
		lipgloss.NewStyle().Foreground(colorDim2).Render("  " + strings.Repeat("─", 30)),
		"",
		lipgloss.NewStyle().Foreground(colorDim2).Render("  ") +
			lipgloss.NewStyle().Foreground(colorGray).Render("kernel  ") +
			lipgloss.NewStyle().Foreground(colorGreen).Render("Linux"),
		lipgloss.NewStyle().Foreground(colorDim2).Render("  ") +
			lipgloss.NewStyle().Foreground(colorGray).Render("arch    ") +
			lipgloss.NewStyle().Foreground(colorGreen).Render("x86_64"),
		lipgloss.NewStyle().Foreground(colorDim2).Render("  ") +
			lipgloss.NewStyle().Foreground(colorGray).Render("repo    ") +
			lipgloss.NewStyle().Foreground(colorGreen).Render("SilverNebula  ✓"),
	}, "\n")

	leftPanel := lipgloss.NewStyle().
		Width(leftW).
		Height(h - 4).
		PaddingTop(1).
		BorderStyle(lipgloss.NormalBorder()).
		BorderRight(true).
		BorderForeground(colorDim).
		Render(logoStyled + "\n\n" + infoBlock)

	// ── right panel: menu ────────────────────────────────────────────────────
	rightW := w - leftW - 2

	menuTitle := lipgloss.NewStyle().
		Foreground(colorCyan).Bold(true).
		PaddingLeft(2).PaddingBottom(1).
		Render(" Main Menu")

	menuDivider := lipgloss.NewStyle().
		Foreground(colorDim).PaddingLeft(2).
		Render(strings.Repeat("─", rightW-6))

	menuContent := menuTitle + "\n" + menuDivider + "\n\n"

	for i, item := range menuItems {
		icon := lipgloss.NewStyle().Foreground(colorCyan).Render(item.icon)
		label := item.label
		desc := item.desc

		if i == m.cursor {
			// selected row — full width highlight
			row := lipgloss.NewStyle().
				Foreground(colorBlue).Bold(true).
				Background(lipgloss.Color("#051525")).
				Width(rightW - 4).
				PaddingLeft(2).
				Render(fmt.Sprintf("▶  %s  %s    %s",
					icon, label,
					lipgloss.NewStyle().
						Foreground(colorDim2).
						Background(lipgloss.Color("#051525")).
						Render(desc),
				))
			menuContent += row + "\n\n"
		} else {
			row := lipgloss.NewStyle().
				Foreground(colorGray).
				Width(rightW - 4).
				PaddingLeft(2).
				Render(fmt.Sprintf("   %s  %s    %s",
					icon, label,
					lipgloss.NewStyle().Foreground(colorDim).Render(desc),
				))
			menuContent += row + "\n\n"
		}
	}

	// bottom hint inside right panel
	hint := lipgloss.NewStyle().
		Foreground(colorDim2).Italic(true).
		PaddingLeft(2).PaddingTop(1).
		Render("  ↑↓ navigate    enter select    q quit")

	rightPanel := lipgloss.NewStyle().
		Width(rightW).
		Height(h - 4).
		PaddingTop(2).
		Render(menuContent + "\n\n" + hint)

	// ── header ───────────────────────────────────────────────────────────────
	header := lipgloss.NewStyle().
		Foreground(colorSilver).Bold(true).
		Background(lipgloss.Color("#0a0a0a")).
		Width(w).PaddingLeft(2).
		BorderStyle(lipgloss.NormalBorder()).
		BorderBottom(true).BorderForeground(colorDim).
		Render("  ss — SilverSaphire  ·  SilverArch Linux Installer  ·  v0.1.0")

	// ── footer ───────────────────────────────────────────────────────────────
	footer := lipgloss.NewStyle().
		Foreground(colorDim2).
		Width(w).PaddingLeft(2).
		BorderStyle(lipgloss.NormalBorder()).
		BorderTop(true).BorderForeground(colorDim).
		Render(" SilverNebula ·  Linux  ·  SilverArch Linux  · (C) Silvercore 2026")

	body := lipgloss.JoinHorizontal(lipgloss.Top, leftPanel, rightPanel)

	return header + "\n" + body + "\n" + footer
}

func init() { _ = fmt.Sprintf }
