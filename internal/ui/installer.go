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
	iStyleSkip  = lipgloss.NewStyle().Foreground(lipgloss.Color("#888888")).Italic(true)
	iStyleInput = lipgloss.NewStyle().Foreground(colorWhite).Bold(true)

	iStyleHeader = lipgloss.NewStyle().
			Foreground(colorSilver).Bold(true).
			BorderStyle(lipgloss.NormalBorder()).
			BorderBottom(true).BorderForeground(colorDim)

	iStyleSidebar = lipgloss.NewStyle().
			Width(26).PaddingRight(1).
			BorderStyle(lipgloss.NormalBorder()).
			BorderRight(true).BorderForeground(colorDim)

	iStyleContent = lipgloss.NewStyle().PaddingLeft(3).PaddingTop(1)

	iStyleOptSel = lipgloss.NewStyle().
			Foreground(colorBlue).Bold(true).
			Background(lipgloss.Color("#0a1a2a")).
			PaddingLeft(1).PaddingRight(2)

	iStyleOptNorm = lipgloss.NewStyle().
			Foreground(colorWhite).PaddingLeft(1).PaddingRight(2)

	iStyleOptSkip = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#888888")).Italic(true).
			PaddingLeft(1).PaddingRight(2)

	iStyleFooter = lipgloss.NewStyle().
			Foreground(colorDim).Italic(true).
			BorderStyle(lipgloss.NormalBorder()).
			BorderTop(true).BorderForeground(colorDim)

	iStyleInputBox = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderForeground(colorBlue).
			PaddingLeft(2).PaddingRight(2).
			PaddingTop(0).PaddingBottom(0)
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
	StepLanguage Step = iota
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
	StepLanguage:   "  Language",
	StepKeyboard:   "  Keyboard",
	StepTimezone:   "  Timezone",
	StepDisk:       "  Disk",
	StepFilesystem: "  Filesystem",
	StepBootloader: "  Bootloader",
	StepKernel:     "  Kernel",
	StepNetwork:    "  Network & VPN",
	StepUser:       "  User",
	StepHostname:   "  Hostname",
	StepDesktop:    "  Desktop / WM",
	StepFonts:      "  Fonts",
	StepPackages:   "  Packages",
	StepAUR:        "  AUR Helper",
	StepSecurity:   "  Security",
	StepSwap:       "  Swap / Zram",
	StepGPU:        "  GPU Driver",
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
	{"العربية  (Arabic)    [ar]", "al-Arabiyya  (Arabic)   [ar]", "ar"},
	{"Français  (French)   [fr]", "Francais  (French)      [fr]", "fr"},
	{"Deutsch  (German)    [de]", "Deutsch  (German)       [de]", "de"},
	{"Español  (Spanish)   [es]", "Espanol  (Spanish)      [es]", "es"},
	{"Italiano  (Italian)  [it]", "Italiano  (Italian)     [it]", "it"},
	{"Português  (Portug.) [pt]", "Portugues  (Portug.)    [pt]", "pt"},
	{"Türkçe  (Turkish)   [tr]", "Turkce  (Turkish)       [tr]", "tr"},
	{"日本語  (Japanese)  [ja]", "Nihongo  (Japanese)     [ja]", "ja"},
	{"中文  (Chinese)     [zh]", "Zhongwen  (Chinese)     [zh]", "zh"},
}

var stepOptions = map[Step][]string{
	StepKeyboard: {
		"  us       (English US)",
		"  uk       (English UK)",
		"  fr       (French AZERTY)",
		"  de       (German QWERTZ)",
		"  ar       (Arabic)",
		"  es       (Spanish)",
		"  it       (Italian)",
		"  jp       (Japanese)",
		"  tr       (Turkish)",
	},
	StepTimezone: {
		"  Africa/Tunis",
		"  UTC",
		"  Europe/Paris",
		"  Europe/London",
		"  Europe/Berlin",
		"  Asia/Riyadh",
		"  Asia/Dubai",
		"  Asia/Tokyo",
		"  Asia/Shanghai",
		"  America/New_York",
		"  America/Los_Angeles",
		"  Custom...",
	},
	StepDisk: {
		"󰋊  Auto partition         (recommended)",
		"   Guided                 (step by step)",
		"   Manual                 (cfdisk style)",
	},
	StepFilesystem: {
		"  ext4                   (stable, default)",
		"  btrfs                  (snapshots, compression)",
		"  xfs                    (high performance)",
		"  zfs                    (advanced, enterprise)",
		"  f2fs                   (flash storage)",
		"  bcachefs               (new, CoW)",
		"  Custom                 (type manually)",
	},
	StepBootloader: {
		"  GRUB                   (universal)",
		"  systemd-boot           (fast, simple)",
		"  rEFInd                 (graphical)",
		"  EFISTUB                (minimal)",
		"  Custom",
	},
	StepKernel: {
		"  linux                  (vanilla upstream)",
		"  linux-lts              (long term stable)",
		"  linux-zen              (desktop performance)",
		"  linux-hardened         (security focused)",
		"  linux-rt               (realtime)",
		"  linux-cachyos          (optimized patches)",
		"  linux-tkg              (custom scheduler)",
		"  Custom                 (enter manually)",
	},
	StepNetwork: {
		"󰤨  NetworkManager         (recommended)",
		"  iwd + dhcpcd",
		"  dhcpcd only            (minimal)",
		"󰒃  NetworkManager + WireGuard VPN",
		"󰒃  NetworkManager + OpenVPN",
	},
	StepDesktop: {
		"  GNOME",
		"  KDE Plasma",
		"  XFCE                   (lightweight)",
		"  MATE",
		"  Cinnamon",
		"  Hyprland               (Wayland, modern)",
		"  Sway                   (Wayland, i3-like)",
		"  i3                     (X11, tiling)",
		"  Openbox                (X11, minimal)",
		"  bspwm",
		"  dwm                    (suckless)",
		"  TTY only               (no DE)",
	},
	StepFonts: {
		"  JetBrains Mono NF      (recommended, coding)",
		"  Noto Fonts             (full unicode, Arabic/CJK)",
		"  Noto Fonts + JB Mono NF (best combo)",
		"  FiraCode NF            (ligatures)",
		"  Hack NF",
		"  Cascadia Code NF",
		"  Source Code Pro",
		"  IBM Plex Mono NF",
		"  DejaVu + Liberation    (minimal)",
		"  Skip                   (use system default)",
	},
	StepAUR: {
		"  paru                   (recommended, Rust)",
		"  yay                    (popular, Go)",
		"  pikaur",
		"  trizen",
		"  None                   (skip AUR)",
	},
	StepSecurity: {
		"  Run Security Wizard",
		"  Skip",
	},
	StepSwap: {
		"  zram                   (RAM compression, recommended)",
		"  swapfile               (traditional)",
		"  swap partition",
		"  None                   (skip)",
	},
	StepGPU: {
		"󰍹  Auto-detect            (recommended)",
		"  Intel                  (i915 open)",
		"  AMD                    (amdgpu open)",
		"  Nvidia                 (proprietary)",
		"  Nvidia                 (nouveau open)",
		"  None",
	},
	StepFlatpak: {
		"  Enable Flatpak + Flathub",
		"  Flatpak only           (no Flathub)",
		"  Skip",
	},
	StepSSH: {
		"  Enable SSH             (port 22)",
		"  Enable SSH             (custom port)",
		"  Disable                (skip)",
	},
	StepDotfiles: {
		"  Import from Git URL",
		"  Import from local path",
		"  Skip",
	},
	StepReview: {
		"  Confirm & Install",
		"  Go back",
		"  Save profile           (TOML)",
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

// steps that use text input
var textInputSteps = map[Step]struct {
	prompt string
	hint   string
}{
	StepUser:     {"  Enter username", "  lowercase, no spaces   e.g. silvercore"},
	StepHostname: {"  Enter hostname", "  device name on network   e.g. silverarch"},
}

// ── Config ────────────────────────────────────────────────────────────────────
type InstallConfig struct {
	Language, Keyboard, Timezone, Disk, Filesystem string
	Bootloader, Kernel, Network, Username, Password string
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
	typingCustom bool // for timezone custom entry
}

func NewInstallerModel(w, h int) InstallerModel {
	return InstallerModel{step: StepLanguage, width: w, height: h}
}

func (m InstallerModel) Init() tea.Cmd { return nil }

func (m InstallerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// ── text input mode ──
		if m.typing || m.typingCustom {
			switch msg.String() {
			case "enter":
				if m.typingCustom {
					// custom timezone
					m.config.Timezone = m.input
					m.typingCustom = false
					m.input = ""
					m.step++
					m.cursor = 0
				} else {
					m.applyInput()
					m.typing = false
					m.input = ""
					m.step++
					m.cursor = 0
				}
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
		case "tab", "s":
			if skippable[m.step] {
				m.step++
				m.cursor = 0
			}
		case "enter", " ":
			// check text input steps
			if _, ok := textInputSteps[m.step]; ok {
				m.typing = true
				return m, nil
			}
			// check timezone custom
			if m.step == StepTimezone {
				opts := stepOptions[StepTimezone]
				if m.cursor == len(opts)-1 {
					// last option = Custom
					m.typingCustom = true
					m.input = ""
					return m, nil
				}
			}
			m.selectCurrent()
			m.step++
			m.cursor = 0
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
	if m.step == StepLanguage {
		return len(languageOptions) - 1
	}
	return len(stepOptions[m.step]) - 1
}

func (m *InstallerModel) selectCurrent() {
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
	case StepKeyboard:
		m.config.Keyboard = v
	case StepTimezone:
		m.config.Timezone = v
	case StepDisk:
		m.config.Disk = v
	case StepFilesystem:
		m.config.Filesystem = v
	case StepBootloader:
		m.config.Bootloader = v
	case StepKernel:
		m.config.Kernel = v
	case StepNetwork:
		m.config.Network = v
	case StepDesktop:
		m.config.Desktop = v
	case StepFonts:
		m.config.Fonts = v
	case StepAUR:
		m.config.AUR = v
	case StepSecurity:
		m.config.Security = v
	case StepSwap:
		m.config.Swap = v
	case StepGPU:
		m.config.GPU = v
	case StepFlatpak:
		m.config.Flatpak = v
	case StepSSH:
		m.config.SSH = v
	case StepDotfiles:
		m.config.Dotfiles = v
	}
}

func (m *InstallerModel) applyInput() {
	switch m.step {
	case StepUser:
		m.config.Username = m.input
	case StepHostname:
		m.config.Hostname = m.input
	}
}

// ── View ──────────────────────────────────────────────────────────────────────
func (m InstallerModel) View() string {
	w := m.width
	if w < 80 {
		w = 80
	}
	total := int(StepReview) + 1
	cur := int(m.step) + 1
	filled := cur * 38 / total
	progress := iStyleOK.Render(strings.Repeat("█", filled)) +
		iStyleGray.Render(strings.Repeat("░", 38-filled))
	skipHint := ""
	if skippable[m.step] {
		skipHint = iStyleSkip.Render("  s/tab=skip")
	}
	header := iStyleHeader.Width(w - 2).Render(
		fmt.Sprintf("   SS — SilverSaphire -- Utility For SilverArch Linux   %s  %d/%d%s", progress, cur, total, skipHint),
	)

	sidebar := ""
	for s := StepLanguage; s <= StepReview; s++ {
		name := stepNames[s]
		switch {
		case s == m.step:
			sidebar += iStyleSel.Render("▶ "+name) + "\n"
		case s < m.step:
			sidebar += iStyleOK.Render("  ✓"+name) + "\n"
		default:
			sidebar += iStyleGray.Render("   "+name) + "\n"
		}
	}

	content := iStyleContent.Render(m.renderStep())
	body := lipgloss.JoinHorizontal(lipgloss.Top,
		iStyleSidebar.Render(sidebar),
		content,
	)
	footer := iStyleFooter.Width(w-2).Render(
		"   ↑↓ navigate    enter select    b back    s skip    q quit",
	)
	return header + "\n" + body + "\n" + footer
}

func (m InstallerModel) renderStep() string {
	// custom timezone typing
	if m.typingCustom {
		return iStyleTitle.Render("  Custom Timezone\n\n") +
			iStyleDim.Render("  format: Region/City   e.g. Africa/Algiers\n\n") +
			iStyleInputBox.Render(iStyleInput.Render("  ❯ "+m.input+"█")) +
			"\n\n" + iStyleGray.Render("  enter to confirm    esc to go back")
	}

	// text input steps
	if info, ok := textInputSteps[m.step]; ok {
		return iStyleTitle.Render(info.prompt+"\n\n") +
			iStyleDim.Render(info.hint+"\n\n") +
			iStyleInputBox.Render(iStyleInput.Render("  ❯ "+m.input+"█")) +
			"\n\n" + iStyleGray.Render("  enter to confirm    esc to cancel")
	}

	switch m.step {
	case StepLanguage:
		return m.renderLanguage()
	case StepReview:
		return m.renderReview()
	case StepDone:
		return iStyleOK.Render("\n   ✓ Installation complete!\n\n") +
			iStyleDim.Render("   Rebooting in 5 seconds...")
	}

	name := stepNames[m.step]
	var b strings.Builder
	b.WriteString(iStyleTitle.Render(name) + "\n")
	b.WriteString(iStyleDim.Render("  "+strings.Repeat("─", 50)) + "\n\n")
	if skippable[m.step] {
		b.WriteString(iStyleSkip.Render("   press s or tab to skip") + "\n\n")
	}
	for i, opt := range stepOptions[m.step] {
		isSkipOpt := strings.Contains(strings.ToLower(opt), "skip") ||
			strings.Contains(strings.ToLower(opt), "none")
		isCustom := strings.Contains(strings.ToLower(opt), "custom...")
		if i == m.cursor {
			b.WriteString(iStyleOptSel.Render("▶  "+opt) + "\n")
		} else if isCustom {
			b.WriteString(iStyleWarn.Render("   "+opt) + "\n")
		} else if isSkipOpt {
			b.WriteString(iStyleOptSkip.Render("   "+opt) + "\n")
		} else {
			b.WriteString(iStyleOptNorm.Render("   "+opt) + "\n")
		}
	}
	return b.String()
}

func (m InstallerModel) renderLanguage() string {
	var b strings.Builder
	b.WriteString(iStyleTitle.Render("  Select Language / اختر اللغة / 言語選択") + "\n")
	b.WriteString(iStyleDim.Render("  "+strings.Repeat("─", 50)) + "\n\n")
	b.WriteString(iStyleWarn.Render("  󰍪  boxes □□□ = font fallback shown on right") + "\n\n")
	for i, opt := range languageOptions {
		label := fmt.Sprintf("%-36s %s", opt.display,
			iStyleGray.Render("  "+safeText(opt.fallback)))
		if i == m.cursor {
			b.WriteString(iStyleOptSel.Render("▶  "+label) + "\n")
		} else {
			b.WriteString(iStyleOptNorm.Render("   "+label) + "\n")
		}
	}
	return b.String()
}

func (m InstallerModel) renderReview() string {
	c := m.config
	var b strings.Builder
	b.WriteString(iStyleTitle.Render("  Review Configuration") + "\n")
	b.WriteString(iStyleDim.Render("  "+strings.Repeat("─", 50)) + "\n\n")
	rows := [][2]string{
		{"  Language", c.Language}, {"  Keyboard", c.Keyboard},
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
		var valStr string
		if val == "" {
			valStr = iStyleSkip.Render("(skipped)")
		} else {
			valStr = iStyleOK.Render(val)
		}
		b.WriteString(iStyleDim.Render(fmt.Sprintf("  %-18s", row[0])) + valStr + "\n")
	}
	b.WriteString("\n")
	for i, opt := range stepOptions[StepReview] {
		if i == m.cursor {
			b.WriteString(iStyleOptSel.Render("▶  "+opt) + "\n")
		} else {
			b.WriteString(iStyleOptNorm.Render("   "+opt) + "\n")
		}
	}
	return b.String()
}
