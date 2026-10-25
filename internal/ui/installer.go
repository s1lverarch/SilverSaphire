// SilverSaphire (ss) - Installer Steps
// Author: Silvercore (s1lverarch) <s1lv3rcore@proton.me>
// License: GPL-v3.0

package ui

import (
	"fmt"
	"strings"
	"unicode"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// ── Local styles ──────────────────────────────────────────────────────────────
var (
	iStyleTitle = lipgloss.NewStyle().Foreground(colorCyan).Bold(true)
	iStyleOK    = lipgloss.NewStyle().Foreground(colorGreen).Bold(true)
	iStyleSel   = lipgloss.NewStyle().Foreground(colorBlue).Bold(true)
	iStyleNorm  = lipgloss.NewStyle().Foreground(colorWhite)
	iStyleDim   = lipgloss.NewStyle().Foreground(colorDim2)
	iStyleGray  = lipgloss.NewStyle().Foreground(colorGray)
	iStyleWarn  = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFD700")).Bold(true)
	iStyleSkip  = lipgloss.NewStyle().Foreground(lipgloss.Color("#666666")).Italic(true)
	iStyleInput = lipgloss.NewStyle().Foreground(colorWhite).Bold(true)

	iStyleHeader = lipgloss.NewStyle().
			Foreground(colorSilver).Bold(true).
			BorderStyle(lipgloss.NormalBorder()).
			BorderBottom(true).BorderForeground(colorDim)

	iStyleFooter = lipgloss.NewStyle().
			Foreground(colorDim2).Italic(true).
			BorderStyle(lipgloss.NormalBorder()).
			BorderTop(true).BorderForeground(colorDim)

	iStyleOptSel = lipgloss.NewStyle().
			Foreground(colorBlue).Bold(true).
			Background(lipgloss.Color("#051525")).
			PaddingLeft(2).PaddingRight(2)

	iStyleOptNorm = lipgloss.NewStyle().
			Foreground(colorWhite).PaddingLeft(2).PaddingRight(2)

	iStyleOptSkip = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#666666")).Italic(true).
			PaddingLeft(2).PaddingRight(2)
)

// ── Arabic transliteration ────────────────────────────────────────────────────
var arabicTranslit = map[rune]string{
	'ا': "a", 'ب': "b", 'ت': "t", 'ث': "th", 'ج': "j", 'ح': "h",
	'خ': "kh", 'د': "d", 'ذ': "dh", 'ر': "r", 'ز': "z", 'س': "s",
	'ش': "sh", 'ص': "s", 'ض': "d", 'ط': "t", 'ظ': "z",
	'غ': "gh", 'ف': "f", 'ق': "q", 'ك': "k", 'ل': "l", 'م': "m",
	'ن': "n", 'ه': "h", 'و': "w", 'ي': "y", 'ة': "a", 'ى': "a",
	'ء': "'", 'أ': "a", 'إ': "i", 'آ': "aa", 'ئ': "y", 'ؤ': "w",
}

func safeText(s string) string {
	var out strings.Builder
	for _, r := range s {
		if r < 128 {
			out.WriteRune(r)
			continue
		}
		if r >= 0x0600 && r <= 0x06FF {
			if t, ok := arabicTranslit[r]; ok {
				out.WriteString(t)
			}
			continue
		}
		if r >= 0x3040 && r <= 0x30FF {
			out.WriteString("[ja]")
			continue
		}
		if r >= 0x4E00 && r <= 0x9FFF {
			out.WriteString("[zh]")
			continue
		}
		if unicode.IsPrint(r) {
			out.WriteRune(r)
		}
	}
	return out.String()
}

// ── Steps ─────────────────────────────────────────────────────────────────────
type Step int

const (
	StepTheme Step = iota
	StepLanguage
	StepKeyboard
	StepTimezone
	StepDisk
	StepFilesystem
	StepBootloader
	StepKernel
	StepNetwork
	StepUser
	StepHostname
	StepDesktop
	StepFonts
	StepPackages
	StepAUR
	StepSecurity
	StepSwap
	StepGPU
	StepFlatpak
	StepSSH
	StepDotfiles
	StepReview
	StepDone
)

var stepNames = map[Step]string{
	StepTheme:      "  Theme",
	StepLanguage:   "  Language",
	StepKeyboard:   "  Keyboard",
	StepTimezone:   "  Timezone",
	StepDisk:       "  Disk",
	StepFilesystem: "  Filesystem",
	StepBootloader: "  Bootloader",
	StepKernel:     "  Kernel",
	StepNetwork:    "󰤨  Network & VPN",
	StepUser:       "  User",
	StepHostname:   "  Hostname",
	StepDesktop:    "  Desktop / WM",
	StepFonts:      "  Fonts",
	StepPackages:   "󰏖  Packages",
	StepAUR:        "  AUR Helper",
	StepSecurity:   "  Security",
	StepSwap:       "  Swap / Zram",
	StepGPU:        "󰍹  GPU Driver",
	StepFlatpak:    "  Flatpak",
	StepSSH:        "  SSH",
	StepDotfiles:   "  Dotfiles",
	StepReview:     "  Review & Install",
}

type langOption struct {
	display  string
	fallback string
	value    string
}

var languageOptions = []langOption{
	{"English", "English", "en"},
	{"العربية  (Arabic)", "al-Arabiyya  (Arabic)", "ar"},
	{"Français  (French)", "Francais  (French)", "fr"},
	{"Deutsch  (German)", "Deutsch  (German)", "de"},
	{"Español  (Spanish)", "Espanol  (Spanish)", "es"},
	{"Italiano  (Italian)", "Italiano  (Italian)", "it"},
	{"Português  (Portuguese)", "Portugues  (Portuguese)", "pt"},
	{"Türkçe  (Turkish)", "Turkce  (Turkish)", "tr"},
	{"日本語  (Japanese)", "Nihongo  (Japanese)", "ja"},
	{"中文  (Chinese)", "Zhongwen  (Chinese)", "zh"},
}

var stepOptions = map[Step][]string{
	StepKeyboard: {
		"us    English US",
		"uk    English UK",
		"fr    French AZERTY",
		"de    German QWERTZ",
		"ar    Arabic",
		"es    Spanish",
		"it    Italian",
		"jp    Japanese",
		"tr    Turkish",
	},
	StepTimezone: {
		"Africa/Tunis",
		"UTC",
		"Europe/Paris",
		"Europe/London",
		"Europe/Berlin",
		"Asia/Riyadh",
		"Asia/Dubai",
		"Asia/Tokyo",
		"Asia/Shanghai",
		"America/New_York",
		"America/Los_Angeles",
		" Custom  (type your own)",
	},
	StepDisk: {
		"󰋊  Auto partition       recommended",
		"   Guided               step by step",
		"   Manual               cfdisk style",
	},
	StepFilesystem: {
		"  ext4        stable, default",
		"  btrfs       snapshots + compression",
		"  xfs         high performance",
		"  zfs         enterprise, advanced",
		"  f2fs        flash/SSD optimized",
		"  bcachefs    new CoW filesystem",
		"   Custom      type manually",
	},
	StepBootloader: {
		"  GRUB            universal",
		"  systemd-boot    fast, UEFI only",
		"  rEFInd          graphical",
		"  EFISTUB         minimal, no bootloader",
		"   Custom",
	},
	StepKernel: {
		"  linux           vanilla upstream",
		"  linux-lts       long term stable",
		"  linux-zen       desktop performance",
		"  linux-hardened  security focused",
		"  linux-rt        realtime",
		"  linux-cachyos   optimized patches",
		"  linux-tkg       custom scheduler",
		"   Custom          enter manually",
	},
	StepNetwork: {
		"󰤨  NetworkManager            recommended",
		"  iwd + dhcpcd",
		"  dhcpcd only               minimal",
		"󰒃  NetworkManager + WireGuard",
		"󰒃  NetworkManager + OpenVPN",
	},
	StepDesktop: {
		"  GNOME",
		"  KDE Plasma",
		"  XFCE          lightweight",
		"  MATE",
		"  Cinnamon",
		"  Hyprland      Wayland, modern",
		"  Sway          Wayland, i3-like",
		"  i3            X11, tiling",
		"  Openbox       X11, minimal",
		"  bspwm",
		"  dwm           suckless",
		"  TTY only      no DE",
	},
	StepFonts: {
		"  JetBrains Mono NF      coding, NF icons",
		"  Noto Fonts             full unicode + Arabic/CJK",
		"  JB Mono NF + Noto      best combo",
		"  FiraCode NF            ligatures",
		"  Hack NF",
		"  Cascadia Code NF",
		"  Source Code Pro",
		"  IBM Plex Mono NF",
		"  DejaVu + Liberation    minimal",
		"  Skip                   use system default",
	},
	StepAUR: {
		"  paru     recommended, written in Rust",
		"  yay      popular, written in Go",
		"  pikaur",
		"  trizen",
		"  None     skip AUR helper",
	},
	StepSecurity: {
		"  Run Security Wizard",
		"  Skip",
	},
	StepSwap: {
		"  zram            RAM compression, recommended",
		"  swapfile        traditional",
		"  swap partition",
		"  None            skip",
	},
	StepGPU: {
		"󰍹  Auto-detect    recommended",
		"  Intel          i915 open source",
		"  AMD            amdgpu open source",
		"  Nvidia         proprietary",
		"  Nvidia         nouveau open source",
		"  None",
	},
	StepFlatpak: {
		"  Enable Flatpak + Flathub",
		"  Flatpak only   no Flathub",
		"  Skip",
	},
	StepSSH: {
		"  Enable SSH     port 22",
		"  Enable SSH     custom port",
		"  Disable        skip",
	},
	StepDotfiles: {
		"  Import from Git URL",
		"  Import from local path",
		"  Skip",
	},
	StepReview: {
		"  Confirm & Install",
		"  Go back",
		"  Save profile  TOML",
		"  Abort",
	},
}

var skippable = map[Step]bool{
	StepPackages: true,
	StepSecurity: true,
	StepFlatpak:  true,
	StepSSH:      true,
	StepDotfiles: true,
	StepFonts:    true,
}

// ── Config ────────────────────────────────────────────────────────────────────
type InstallConfig struct {
	Theme, Language, Keyboard, Timezone, Disk, Filesystem string
	Bootloader, Kernel, Network, Username          string
	Hostname, Desktop, Fonts, AUR, Security        string
	Swap, GPU, Flatpak, SSH, Dotfiles              string
}

// ── Model ─────────────────────────────────────────────────────────────────────
type InstallerModel struct {
	step         Step
	cursor       int
	config       InstallConfig
	width        int
	height       int
	input        string
	typing       bool
	typingCustom bool
	inputPrompt  string
	inputHint    string
}

func NewInstallerModel(w, h int) InstallerModel {
	return InstallerModel{step: StepTheme, width: w, height: h}
}

func (m InstallerModel) Init() tea.Cmd { return nil }

func (m InstallerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.typing || m.typingCustom {
			switch msg.String() {
			case "enter":
				if m.typingCustom {
					m.config.Timezone = m.input
					m.typingCustom = false
				} else {
					m.applyInput()
					m.typing = false
				}
				m.input = ""
				m.step++
				m.cursor = 0
			case "backspace":
				if len(m.input) > 0 {
					m.input = m.input[:len(m.input)-1]
				}
			case "esc":
				m.typing = false
				m.typingCustom = false
				m.input = ""
			default:
				if len(msg.String()) == 1 {
					m.input += msg.String()
				}
			}
			return m, nil
		}
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < m.maxCursor() {
				m.cursor++
			}
		case "tab", "s", "ctrl+s":
			if skippable[m.step] {
				m.step++
				m.cursor = 0
			}
		case "enter", " ":
			switch m.step {
			case StepUser:
				m.typing = true
				m.inputPrompt = "  Enter username"
				m.inputHint   = "  lowercase, no spaces   e.g. silvercore"
			case StepHostname:
				m.typing = true
				m.inputPrompt = "  Enter hostname"
				m.inputHint   = "  device name on network   e.g. silverarch"
			case StepTimezone:
				opts := stepOptions[StepTimezone]
				if m.cursor == len(opts)-1 {
					m.typingCustom = true
					m.input = ""
					return m, nil
				}
				m.selectCurrent()
				m.step++
				m.cursor = 0
			default:
				m.selectCurrent()
				m.step++
				m.cursor = 0
			}
		case "backspace", "b":
			if m.step > StepLanguage {
				m.step--
				m.cursor = 0
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}
	return m, nil
}

func (m InstallerModel) maxCursor() int {
	switch m.step {
	case StepTheme:
		return len(ThemeOrder) - 1
	case StepLanguage:
		return len(languageOptions) - 1
	default:
		return len(stepOptions[m.step]) - 1
	}
}

func (m *InstallerModel) selectCurrent() {
	if m.step == StepTheme {
		if m.cursor < len(ThemeOrder) {
			name := ThemeOrder[m.cursor]
			m.config.Theme = name
			ApplyTheme(name)
		}
		return
	}
	if m.step == StepLanguage {
		if m.cursor < len(languageOptions) {
			m.config.Language = languageOptions[m.cursor].value
		}
		return
	}
	opts := stepOptions[m.step]
	if m.cursor >= len(opts) {
		return
	}
	v := strings.TrimSpace(opts[m.cursor])
	switch m.step {
	case StepKeyboard:   m.config.Keyboard = v
	case StepTimezone:   m.config.Timezone = v
	case StepDisk:       m.config.Disk = v
	case StepFilesystem: m.config.Filesystem = v
	case StepBootloader: m.config.Bootloader = v
	case StepKernel:     m.config.Kernel = v
	case StepNetwork:    m.config.Network = v
	case StepDesktop:    m.config.Desktop = v
	case StepFonts:      m.config.Fonts = v
	case StepAUR:        m.config.AUR = v
	case StepSecurity:   m.config.Security = v
	case StepSwap:       m.config.Swap = v
	case StepGPU:        m.config.GPU = v
	case StepFlatpak:    m.config.Flatpak = v
	case StepSSH:        m.config.SSH = v
	case StepDotfiles:   m.config.Dotfiles = v
	}
}

func (m *InstallerModel) applyInput() {
	switch m.step {
	case StepUser:     m.config.Username = m.input
	case StepHostname: m.config.Hostname = m.input
	}
}

// ── View ──────────────────────────────────────────────────────────────────────
func (m InstallerModel) View() string {
	w := m.width
	if w < 100 {
		w = 100
	}

	// progress bar
	total  := int(StepReview) + 1
	cur    := int(m.step) + 1
	filled := cur * 36 / total
	progress := iStyleOK.Render(strings.Repeat("█", filled)) +
		iStyleGray.Render(strings.Repeat("░", 36-filled))

	skipHint := ""
	if skippable[m.step] {
		skipHint = iStyleSkip.Render("    s/tab = skip")
	}

	header := iStyleHeader.Width(w-2).Render(
		fmt.Sprintf("   ss — SilverSaphire — Utility For SilverArch Linux    %s  %d/%d%s",
			progress, cur, total, skipHint),
	)

	// ── sidebar (fixed 24 wide) ───────────────────────────────────────────────
	sidebarW := 24
	sidebar  := ""
	for s := StepLanguage; s <= StepReview; s++ {
		name := stepNames[s]
		switch {
		case s == m.step:
			sidebar += iStyleSel.Render("▶"+name) + "\n"
		case s < m.step:
			sidebar += iStyleOK.Render(" ✓"+name) + "\n"
		default:
			sidebar += iStyleGray.Render("  "+name) + "\n"
		}
	}
	sidebarPanel := lipgloss.NewStyle().
		Width(sidebarW).
		BorderStyle(lipgloss.NormalBorder()).
		BorderRight(true).BorderForeground(colorDim).
		PaddingLeft(1).PaddingTop(1).
		Render(sidebar)

	// ── content (remaining width) ─────────────────────────────────────────────
	contentW := w - sidebarW - 6
	content  := m.renderStep(contentW)
	contentPanel := lipgloss.NewStyle().
		Width(contentW).
		PaddingLeft(4).PaddingTop(1).
		Render(content)

	body := lipgloss.JoinHorizontal(lipgloss.Top, sidebarPanel, contentPanel)

	footer := iStyleFooter.Width(w-2).Render(
		"   ↑↓ navigate    enter select    b back    s skip    q quit",
	)

	return header + "\n" + body + "\n" + footer
}

// ── Step renderers ────────────────────────────────────────────────────────────
func (m InstallerModel) renderStep(w int) string {
	// custom timezone
	if m.typingCustom {
		return m.renderInput(w,
			"  Custom Timezone",
			"  format: Region/City    e.g. Africa/Algiers",
		)
	}
	// text input steps
	if m.typing {
		return m.renderInput(w, m.inputPrompt, m.inputHint)
	}

	switch m.step {
	case StepTheme:
		return m.renderTheme(w)
	case StepLanguage:
		return m.renderLanguage(w)
	case StepReview:
		return m.renderReview(w)
	case StepDone:
		return iStyleOK.Render("\n\n   ✓ Installation complete!\n\n") +
			iStyleDim.Render("   Rebooting in 5 seconds...")
	}

	name := stepNames[m.step]
	var b strings.Builder
	b.WriteString(iStyleTitle.Render(name) + "\n")
	b.WriteString(iStyleGray.Render(strings.Repeat("─", w-4)) + "\n\n")
	if skippable[m.step] {
		b.WriteString(iStyleSkip.Render("  press s or tab to skip this step\n\n"))
	}
	for i, opt := range stepOptions[m.step] {
		isSkip := strings.Contains(strings.ToLower(opt), "skip") ||
			strings.Contains(strings.ToLower(opt), "none")
		isCustom := strings.Contains(strings.ToLower(opt), "custom")
		switch {
		case i == m.cursor:
			b.WriteString(iStyleOptSel.Width(w-4).Render("▶  "+opt) + "\n")
		case isCustom:
			b.WriteString(iStyleWarn.PaddingLeft(2).Render("   "+opt) + "\n")
		case isSkip:
			b.WriteString(iStyleOptSkip.Width(w-4).Render("   "+opt) + "\n")
		default:
			b.WriteString(iStyleOptNorm.Width(w-4).Render("   "+opt) + "\n")
		}
	}
	return b.String()
}

func (m InstallerModel) renderInput(w int, prompt, hint string) string {
	var b strings.Builder
	b.WriteString(iStyleTitle.Render(prompt) + "\n")
	b.WriteString(iStyleGray.Render(strings.Repeat("─", w-4)) + "\n\n")
	b.WriteString(iStyleDim.Render(hint) + "\n\n")
	inputLine := "  ❯ " + m.input + "█"
	b.WriteString(
		lipgloss.NewStyle().
			Width(w-4).
			Border(lipgloss.NormalBorder()).
			BorderForeground(colorBlue).
			PaddingLeft(2).PaddingRight(2).
			Render(iStyleInput.Render(inputLine)),
	)
	b.WriteString("\n\n" + iStyleGray.Render("  enter to confirm    esc to cancel"))
	return b.String()
}

func (m InstallerModel) renderTheme(w int) string {
	var b strings.Builder
	b.WriteString(iStyleTitle.Render("  Choose Installer Theme") + "\n")
	b.WriteString(iStyleGray.Render(strings.Repeat("─", w-4)) + "\n\n")
	b.WriteString(iStyleDim.Render("  theme applies instantly — change anytime\n\n"))

	for i, name := range ThemeOrder {
		t := Themes[name]
		// colored preview blocks
		preview := lipgloss.NewStyle().Foreground(t.Primary).Render("██") +
			lipgloss.NewStyle().Foreground(t.Secondary).Render("██") +
			lipgloss.NewStyle().Foreground(t.Accent).Render("██") +
			lipgloss.NewStyle().Foreground(t.OK).Render("██")

		label := fmt.Sprintf("%s  %-14s  %s", t.Icon, t.Name, preview)

		if i == m.cursor {
			// apply theme preview live
			ApplyTheme(name)
			b.WriteString(iStyleOptSel.Width(w-4).Render("▶  "+label) + "\n")
		} else {
			b.WriteString(iStyleOptNorm.Width(w-4).Render("   "+label) + "\n")
		}
	}
	return b.String()
}

func (m InstallerModel) renderLanguage(w int) string {
	var b strings.Builder
	b.WriteString(iStyleTitle.Render("  Select Language / اختر اللغة") + "\n")
	b.WriteString(iStyleGray.Render(strings.Repeat("─", w-4)) + "\n\n")
	b.WriteString(iStyleWarn.Render("  󰍪  boxes □□□ = font fallback shown") + "\n\n")
	for i, opt := range languageOptions {
		label := fmt.Sprintf("%-30s  %s", opt.display,
			iStyleGray.Render(safeText(opt.fallback)))
		if i == m.cursor {
			b.WriteString(iStyleOptSel.Width(w-4).Render("▶  "+label) + "\n")
		} else {
			b.WriteString(iStyleOptNorm.Width(w-4).Render("   "+label) + "\n")
		}
	}
	return b.String()
}

func (m InstallerModel) renderReview(w int) string {
	c := m.config
	var b strings.Builder
	b.WriteString(iStyleTitle.Render("  Review Configuration") + "\n")
	b.WriteString(iStyleGray.Render(strings.Repeat("─", w-4)) + "\n\n")
	rows := [][2]string{
		{"  Theme", c.Theme}, {"  Language", c.Language}, {"  Keyboard", c.Keyboard},
		{"  Timezone", c.Timezone}, {"  Disk", c.Disk},
		{"  Filesystem", c.Filesystem}, {"  Bootloader", c.Bootloader},
		{"  Kernel", c.Kernel}, {"󰤨  Network", c.Network},
		{"  Username", c.Username}, {"  Hostname", c.Hostname},
		{"  Desktop", c.Desktop}, {"  Fonts", c.Fonts},
		{"  AUR", c.AUR}, {"  Security", c.Security},
		{"  Swap", c.Swap}, {"󰍹  GPU", c.GPU},
		{"  Flatpak", c.Flatpak}, {"  SSH", c.SSH},
		{"  Dotfiles", c.Dotfiles},
	}
	for _, row := range rows {
		val := row[1]
		var v string
		if val == "" {
			v = iStyleSkip.Render("(skipped)")
		} else {
			v = iStyleOK.Render(val)
		}
		b.WriteString(iStyleDim.Render(fmt.Sprintf("  %-16s", row[0])) + v + "\n")
	}
	b.WriteString("\n")
	for i, opt := range stepOptions[StepReview] {
		if i == m.cursor {
			b.WriteString(iStyleOptSel.Width(w-4).Render("▶  "+opt) + "\n")
		} else {
			b.WriteString(iStyleOptNorm.Width(w-4).Render("   "+opt) + "\n")
		}
	}
	return b.String()
}
