// SilverSaphire (ss) - Boot Animation & Main Menu
// Author: Silvercore (s1lverarch) <s1lv3rcore@proton.me>
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

	styleLogo      = lipgloss.NewStyle().Foreground(colorSilver).Bold(true)
	styleOKBracket = lipgloss.NewStyle().Foreground(colorGreen).Bold(true)
	styleOKText    = lipgloss.NewStyle().Foreground(colorDim2)
	styleDim       = lipgloss.NewStyle().Foreground(colorDim)
	styleDivider   = lipgloss.NewStyle().Foreground(colorDim)
)

// ── Super verbose kernel-style boot log ───────────────────────────────────────
var bootLines = []struct {
	prefix string
	msg    string
}{
	{"[    0.000000]", "Linux version 6.12.0-silverarch (silvercore@silverarch) (gcc 14.2.1) #1 SMP"},
	{"[    0.000000]", "QNX Neutrino RTOS microkernel co-boot initializing..."},
	{"[    0.000000]", "Command line: BOOT_IMAGE=/boot/vmlinuz root=UUID=a1b2c3d4 rw quiet splash"},
	{"[    0.000000]", "BIOS-provided physical RAM map:"},
	{"[    0.000000]", "BIOS-e820: [mem 0x0000000000000000-0x000000000009efff] usable"},
	{"[    0.000000]", "BIOS-e820: [mem 0x0000000000100000-0x000000007fffffff] usable"},
	{"[    0.000000]", "NX (Execute Disable) protection: active"},
	{"[    0.000000]", "SMBIOS 3.0 present."},
	{"[    0.000000]", "DMI: LENOVO ThinkPad/BlackBerry Q20, BIOS 1.42 2023-09-01"},
	{"[    0.000000]", "tsc: Detected 2600.000 MHz processor"},
	{"[    0.000000]", "e820: last_pfn = 0x7fff0 max_arch_pfn = 0x400000000"},
	{"[    0.000000]", "MTRR default type: uncachable"},
	{"[    0.000000]", "x86/PAT: Configuration [0-7]: WB  WC  UC- UC  WB  WP  UC- WT"},
	{"[    0.000000]", "found SMP MP-table at [mem 0x000f05b0-0x000f05bf]"},
	{"[    0.000000]", "ACPI: Early table checksum verification disabled"},
	{"[    0.000000]", "ACPI: RSDP 0x00000000000F0490 000024 (v02 BOCHS )"},
	{"[    0.000000]", "ACPI: IRQ0 used by override."},
	{"[    0.000000]", "ACPI: IRQ9 used by override."},
	{"[    0.000000]", "PCI: Using configuration type 1 for base access"},
	{"[    0.000000]", "clocksource: hpet: mask: 0xffffffff max_cycles: 0xffffffff"},
	{"[    0.000017]", "SilverArch: QNX HAL bridge layer initializing"},
	{"[    0.000042]", "ACPI: PM-Timer IO Port: 0x608"},
	{"[    0.000100]", "ACPI: LAPIC_NMI (acpi_id[0xff] dfl dfl lint[0x1])"},
	{"[    0.000150]", "IOAPIC[0]: apic_id 0, version 17, address 0xfec00000, GSI 0-23"},
	{"[    0.000200]", "proc: Initially using: 4 CPUs"},
	{"[    0.000300]", "Booting paravirtualized kernel on bare hardware"},
	{"[    0.000400]", "setup_percpu: NR_CPUS:8192 nr_cpumask_bits:4 nr_cpu_ids:4 nr_node_ids:1"},
	{"[    0.001200]", "percpu: Embedded 62 pages/cpu s212992 r8192 d28672 u262144"},
	{"[    0.001500]", "Built 1 zonelists, mobility grouping on. Total pages: 515961"},
	{"[    0.001600]", "Kernel command line: BOOT_IMAGE=/boot/vmlinuz-silverarch"},
	{"[    0.001700]", "Dentry cache hash table entries: 262144 (order: 9, 2097152 bytes, linear)"},
	{"[    0.001800]", "Inode-cache hash table entries: 131072 (order: 8, 1048576 bytes, linear)"},
	{"[    0.002000]", "mem auto-init: stack:all(zero), heap alloc:on, heap free:off"},
	{"[    0.002100]", "Memory: 2010956K/2096632K available (18432K kernel code, 4160K rwdata)"},
	{"[    0.002500]", "SLUB: HWalign=64, Order=0-3, MinObjects=0, CPUs=4, Nodes=1"},
	{"[    0.002600]", "ftrace: allocating 48923 entries in 192 pages"},
	{"[    0.003500]", "Dynamic Preempt: voluntary"},
	{"[    0.004000]", "rcu: Hierarchical RCU implementation."},
	{"[    0.004100]", "rcu: RCU restricting CPUs from NR_CPUS=8192 to nr_cpu_ids=4."},
	{"[    0.004200]", "rcu: RCU debug extended checks: RCU_EQS_DEBUG"},
	{"[    0.004500]", "NR_IRQS: 524544, nr_irqs: 1024, preallocated irqs: 16"},
	{"[    0.005000]", "Console: colour dummy device 80x25"},
	{"[    0.005100]", "printk: legacy console [tty0] enabled"},
	{"[    0.005200]", "ACPI: Core revision 20240827"},
	{"[    0.006000]", "clocksource: tsc-early: mask: 0xffffffffffffffff max_cycles: 0x25694cb7b0c"},
	{"[    0.007000]", "Calibrating delay loop (skipped), value calculated using timer frequency.."},
	{"[    0.007100]", "5199.99 BogoMIPS (lpj=10399980)"},
	{"[    0.007200]", "pid_max: default: 32768 minimum: 301"},
	{"[    0.007500]", "LSM: initializing lsm=capability,yama,loadpin,safesetid,selinux,bpf"},
	{"[    0.007600]", "SELinux: Initializing."},
	{"[    0.007700]", "AppArmor: AppArmor initialized"},
	{"[    0.008000]", "Mount-cache hash table entries: 4096 (order: 3, 32768 bytes, linear)"},
	{"[    0.008200]", "Mountpoint-cache hash table entries: 4096 (order: 3, 32768 bytes, linear)"},
	{"[    0.009000]", "CPU0: Thermal monitoring enabled (TM2)"},
	{"[    0.009100]", "process: using mwait in idle threads"},
	{"[    0.009200]", "Last level iTLB entries: 4KB 512, 2MB 8, 4MB 8"},
	{"[    0.009300]", "Last level dTLB entries: 4KB 512, 2MB 32, 4MB 32, 1GB 0"},
	{"[    0.009500]", "Spectre V1: Mitigation: usercopy/swapgs barriers"},
	{"[    0.009600]", "Spectre V2: Mitigation: Retpolines"},
	{"[    0.009700]", "Spectre V2 : Spectre v2 / SpectreRSB mitigation: Filling RSB on context switch"},
	{"[    0.009800]", "Spectre V2 : mitigation: Enabling conditional Indirect Branch Prediction Barrier"},
	{"[    0.009900]", "Spectre V4: Mitigation: Speculative Store Bypass disabled via prctl"},
	{"[    0.010000]", "MDS: Mitigation: Clear CPU buffers"},
	{"[    0.010200]", "Freeing SMP alternatives memory: 44K"},
	{"[    0.011000]", "smpboot: Booting Node 0 Processor 1 APIC 0x1"},
	{"[    0.011100]", "smpboot: Booting Node 0 Processor 2 APIC 0x2"},
	{"[    0.011200]", "smpboot: Booting Node 0 Processor 3 APIC 0x3"},
	{"[    0.011500]", "Brought up 4 CPUs"},
	{"[    0.011600]", "smpboot: Max logical packages: 1"},
	{"[    0.012000]", "net: Registered PF_UNIX/PF_LOCAL protocol family"},
	{"[    0.012100]", "net: Registered PF_INET protocol family"},
	{"[    0.012200]", "net: Registered PF_INET6 protocol family"},
	{"[    0.013000]", "PCI: Using ACPI for IRQ routing"},
	{"[    0.013500]", "PCI: pci_cache_line_size set to 64 bytes"},
	{"[    0.014000]", "e820: reserve RAM buffer [mem 0x0009f000-0x0009ffff]"},
	{"[    0.014500]", "clocksource: Switched to clocksource tsc"},
	{"[    0.015000]", "VFS: Disk quotas dquot_6.6.0"},
	{"[    0.015200]", "VFS: Dquot-cache hash table entries: 512"},
	{"[    0.015500]", "pnp: PnP ACPI init"},
	{"[    0.016000]", "pnp: PnP ACPI: found 9 devices"},
	{"[    0.017000]", "clocksource: acpi_pm: mask: 0xffffff max_cycles: 0xffffff"},
	{"[    0.017500]", "NET: Registered PF_NETLINK/PF_ROUTE protocol family"},
	{"[    0.018000]", "PCI: CLS 64 bytes, default 64"},
	{"[    0.018200]", "Trying to unpack rootfs image as initramfs..."},
	{"[    0.019000]", "Freeing initrd memory: 8192K"},
	{"[    0.020000]", "PCI-DMA: Using software bounce buffering for IO (SWIOTLB)"},
	{"[    0.020100]", "software IO TLB: area num 4, bus addr limit 0x00000000"},
	{"[    0.020200]", "software IO TLB: mapped [mem 0x7b400000-0x7f400000] (64MB)"},
	{"[    0.021000]", "SCSI subsystem initialized"},
	{"[    0.021100]", "libata version 3.00 loaded."},
	{"[    0.021200]", "ACPI: bus type USB registered"},
	{"[    0.021300]", "usbcore: registered new interface driver usbfs"},
	{"[    0.021400]", "usbcore: registered new interface driver hub"},
	{"[    0.021500]", "usbcore: registered new device driver usb"},
	{"[    0.022000]", "pps_core: LinuxPPS API ver. 1 registered"},
	{"[    0.022100]", "pps_core: Software version 5.3.6 by Rodolfo Giometti"},
	{"[    0.022200]", "PTP clock support registered"},
	{"[    0.022300]", "Advanced Linux Sound Architecture Driver Initialized."},
	{"[    0.023000]", "thermal: Intel PowerClamp Driver"},
	{"[    0.023100]", "intel_idle: Intel Nehalem Core C-state driver, 9 C-states available"},
	{"[    0.024000]", "ledtrig-cpu: registered to indicate activity on CPUs"},
	{"[    0.024100]", "hid: raw HID events driver (C) Jiri Kosina"},
	{"[    0.024200]", "usbcore: registered new interface driver usbhid"},
	{"[    0.024300]", "usbhid: USB HID core driver"},
	{"[    0.025000]", "drop_monitor: Initializing network drop monitor service"},
	{"[    0.025100]", "NET: Registered PF_INET6 protocol family"},
	{"[    0.026000]", "Segment Routing with IPv6"},
	{"[    0.026100]", "In-situ OAM (IOAM) with IPv6"},
	{"[    0.026200]", "mip6: Mobile IPv6"},
	{"[    0.026300]", "NET: Registered PF_PACKET protocol family"},
	{"[    0.027000]", "Bridge firewalling registered"},
	{"[    0.027100]", "8021q: 802.1Q VLAN Support v1.8"},
	{"[    0.028000]", "device-mapper: ioctl: 4.48.0-ioctl ready for use"},
	{"[    0.028100]", "device-mapper: multipath: version 1.14.0 loaded"},
	{"[    0.029000]", "Intel(R) Wireless WiFi driver for Linux"},
	{"[    0.029100]", "iwlwifi: driver: iwlwifi"},
	{"[    0.030000]", "Bluetooth: Core ver 2.22"},
	{"[    0.030100]", "NET: Registered PF_BLUETOOTH protocol family"},
	{"[    0.030200]", "Bluetooth: HCI device and connection manager initialized"},
	{"[    0.030300]", "Bluetooth: HCI socket layer initialized"},
	{"[    0.030400]", "Bluetooth: L2CAP socket layer initialized"},
	{"[    0.030500]", "Bluetooth: SCO socket layer initialized"},
	{"[    0.031000]", "EXT4-fs (sda2): mounted filesystem with ordered data mode"},
	{"[    0.031100]", "EXT4-fs (sda2): recovery complete"},
	{"[    0.032000]", "systemd[1]: SilverArch Linux — QNX Utility Mode"},
	{"[    0.032100]", "systemd[1]: Detected architecture x86-64."},
	{"[    0.032200]", "systemd[1]: Hostname set to <silverarch>."},
	{"[    0.033000]", "systemd[1]: Queued start job for default target Graphical Interface"},
	{"[    0.033100]", "[ OK ] Started Journal Service"},
	{"[    0.033200]", "[ OK ] Listening on LVM2 poll daemon socket"},
	{"[    0.033300]", "[ OK ] Listening on D-Bus System Message Bus Socket"},
	{"[    0.034000]", "[ OK ] Reached target Local Encrypted Volumes"},
	{"[    0.034100]", "[ OK ] Mounted Kernel Debug File System"},
	{"[    0.034200]", "[ OK ] Mounted Kernel Trace File System"},
	{"[    0.035000]", "[ OK ] Started SilverNebula Repository Daemon"},
	{"[    0.035100]", "[ OK ] Started SilverArch HAL Service"},
	{"[    0.035200]", "[ OK ] Started SilverSaphire Installer Engine"},
	{"[    0.035300]", "[ OK ] Reached target SilverArch Linux — Ready"},
}

// ── Phases ────────────────────────────────────────────────────────────────────
type phase int

const (
	phaseBoot phase = iota
	phaseMenu
	phaseAbout
)

type tickMsg time.Time
type nextLineMsg struct{}

// ── Menu entries ──────────────────────────────────────────────────────────────
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

// ── Model ─────────────────────────────────────────────────────────────────────
type BootModel struct {
	logo      string
	phase     phase
	logLines  int
	logIdx    int
	cursor    int
	width     int
	height    int
	blinkTick int
	blink     bool
}

func NewBootAnimation(logo string) BootModel {
	return BootModel{logo: logo, phase: phaseBoot}
}

func (m BootModel) Init() tea.Cmd {
	return tea.Tick(18*time.Millisecond, func(t time.Time) tea.Msg { return nextLineMsg{} })
}

func (m BootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case nextLineMsg:
		if m.phase == phaseBoot {
			if m.logIdx < len(bootLines) {
				m.logLines++
				m.logIdx++
				return m, tea.Tick(18*time.Millisecond, func(t time.Time) tea.Msg { return nextLineMsg{} })
			}
			return m, tea.Tick(300*time.Millisecond, func(t time.Time) tea.Msg { return tickMsg(t) })
		}

	case tickMsg:
		if m.phase == phaseBoot {
			m.phase = phaseMenu
		}
		m.blinkTick++
		m.blink = m.blinkTick%8 < 4
		return m, tea.Tick(100*time.Millisecond, func(t time.Time) tea.Msg { return tickMsg(t) })

	case tea.KeyMsg:
		switch m.phase {
		case phaseAbout:
			m.phase = phaseMenu
			return m, nil
		case phaseMenu:
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
				case 2:
					m.phase = phaseAbout
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
	switch m.phase {
	case phaseBoot:
		return m.viewBoot()
	case phaseAbout:
		return m.viewAbout()
	default:
		return m.viewMenu()
	}
}

// ── Boot view ─────────────────────────────────────────────────────────────────
func (m BootModel) viewBoot() string {
	h := m.height
	if h < 10 {
		h = 40
	}
	// show last h-4 lines
	start := 0
	if m.logLines > h-4 {
		start = m.logLines - (h - 4)
	}

	var b strings.Builder
	for i := start; i < m.logLines && i < len(bootLines); i++ {
		entry := bootLines[i]
		ts := lipgloss.NewStyle().Foreground(colorDim2).Render(entry.prefix)
		if strings.Contains(entry.msg, "[ OK ]") {
			parts := strings.SplitN(entry.msg, "]", 2)
			b.WriteString(ts + " " + styleOKBracket.Render(parts[0]+"]") +
				styleOKText.Render(parts[1]) + "\n")
		} else if strings.Contains(entry.msg, "SilverArch") || strings.Contains(entry.msg, "Silver") {
			b.WriteString(ts + " " + lipgloss.NewStyle().Foreground(colorCyan).Render(entry.msg) + "\n")
		} else {
			b.WriteString(ts + " " + styleOKText.Render(entry.msg) + "\n")
		}
	}

	cur := " "
	if m.blink {
		cur = "█"
	}
	b.WriteString(styleDim.Render(cur))
	return b.String()
}

// ── Menu view ─────────────────────────────────────────────────────────────────
func (m BootModel) viewMenu() string {
	w := m.width
	if w < 80 {
		w = 80
	}
	h := m.height
	if h < 24 {
		h = 24
	}
	leftW  := w / 2
	rightW := w - leftW - 3

	// left panel
	logoStr := styleLogo.Render(m.logo)
	sep     := lipgloss.NewStyle().Foreground(colorDim).Render(strings.Repeat("─", leftW-4))
	col     := func(k, v, vc string) string {
		return lipgloss.NewStyle().Foreground(colorGray).Render("  "+k+"  ") +
			lipgloss.NewStyle().Foreground(lipgloss.Color(vc)).Render(v)
	}
	info := strings.Join([]string{
		lipgloss.NewStyle().Foreground(colorCyan).Bold(true).Render("  SilverSaphire Installer"),
		lipgloss.NewStyle().Foreground(colorGray).Render("  ss  ·  v0.2.6  ·  GPL-v3.0"),
		"",
		sep,
		"",
		col("author ", "Silvercore (s1lverarch)", "#EEEEEE"),
		col("email  ", "s1lv3rcore@proton.me", "#EEEEEE"),
		col("web    ", "s1lverarch.github.io", "#00BFFF"),
		col("license", "GPL-v3.0", "#EEEEEE"),
		"",
		sep,
		"",
		col("kernel ", "QNX Neutrino 8.0.0 + Linux 6.12", "#00FF7F"),
		col("arch   ", "x86_64  SMP  4 cores", "#00FF7F"),
		col("repo   ", "SilverNebula  ✓", "#00FF7F"),
		col("mode   ", "installer (ss)", "#00FFD0"),
	}, "\n")

	leftPanel := lipgloss.NewStyle().
		Width(leftW).Height(h-4).PaddingTop(1).
		BorderStyle(lipgloss.NormalBorder()).
		BorderRight(true).BorderForeground(colorDim).
		Render(logoStr + "\n\n" + info)

	// right panel — menu
	title := lipgloss.NewStyle().Foreground(colorCyan).Bold(true).
		PaddingLeft(3).Render(" Main Menu")
	div := lipgloss.NewStyle().Foreground(colorDim).PaddingLeft(3).
		Render(strings.Repeat("─", rightW-4))

	rows := title + "\n" + div + "\n\n"
	for i, item := range menuItems {
		icon  := lipgloss.NewStyle().Foreground(colorCyan).Render(item.icon)
		descS := lipgloss.NewStyle().Foreground(colorDim2)
		if i == m.cursor {
			descS = descS.Background(lipgloss.Color("#051525"))
			rows += lipgloss.NewStyle().
				Foreground(colorBlue).Bold(true).
				Background(lipgloss.Color("#051525")).
				Width(rightW-2).PaddingLeft(3).
				Render(fmt.Sprintf("▶  %s  %-26s %s", icon, item.label,
					descS.Render(item.desc))) + "\n\n"
		} else {
			rows += lipgloss.NewStyle().
				Foreground(colorGray).Width(rightW-2).PaddingLeft(3).
				Render(fmt.Sprintf("   %s  %-26s %s", icon, item.label,
					lipgloss.NewStyle().Foreground(colorDim).Render(item.desc))) + "\n\n"
		}
	}

	hint := lipgloss.NewStyle().Foreground(colorDim2).Italic(true).PaddingLeft(3).
		Render("  ↑↓ navigate    enter select    q quit")

	rightPanel := lipgloss.NewStyle().
		Width(rightW).Height(h-4).PaddingTop(2).
		Render(rows + "\n" + hint)

	header := lipgloss.NewStyle().
		Foreground(colorSilver).Bold(true).Width(w).PaddingLeft(2).
		BorderStyle(lipgloss.NormalBorder()).BorderBottom(true).BorderForeground(colorDim).
		Render("  ss — SilverSaphire  ·  SilverArch Linux Installer  ·  v0.2.6")

	footer := lipgloss.NewStyle().
		Foreground(colorDim2).Width(w).PaddingLeft(2).
		BorderStyle(lipgloss.NormalBorder()).BorderTop(true).BorderForeground(colorDim).
		Render("  SilverNebula ✓    QNX 8.0.0 + Linux 6.12    x86_64    GPL-v3.0")

	body := lipgloss.JoinHorizontal(lipgloss.Top, leftPanel, rightPanel)
	return header + "\n" + body + "\n" + footer
}

// ── About view ────────────────────────────────────────────────────────────────
func (m BootModel) viewAbout() string {
	w := m.width
	if w < 80 { w = 80 }

	header := lipgloss.NewStyle().
		Foreground(colorSilver).Bold(true).Width(w).PaddingLeft(2).
		BorderStyle(lipgloss.NormalBorder()).BorderBottom(true).BorderForeground(colorDim).
		Render("  ss — About SilverSaphire")

	body := lipgloss.NewStyle().PaddingLeft(4).PaddingTop(2).Render(
		styleLogo.Render(m.logo) + "\n\n" +
		lipgloss.NewStyle().Foreground(colorCyan).Bold(true).Render("  SilverSaphire (ss)  —  SilverArch Linux Installer") + "\n\n" +
		row("version  ", "0.2.6") +
		row("author   ", "Silvercore (s1lverarch)") +
		row("email    ", "s1lv3rcore@proton.me") +
		row("website  ", "https://s1lverarch.github.io/linuxdistro/") +
		row("repo     ", "https://github.com/s1lverarch/SilverSaphire") +
		row("license  ", "GPL-v3.0") +
		"\n" +
		lipgloss.NewStyle().Foreground(colorDim).Render("  "+strings.Repeat("─", 50)) + "\n\n" +
		row("distro   ", "SilverArch Linux — first Arch-based Tunisian distro") +
		row("supported verisons of linux kernel   ", "Linux 6.0.2 aarch +") +
		row("arch     ", "x86_64  SMP  4 cores") +
		row("language ", "Go (TUI) + Nim (post-install hooks)") +
		row("repo     ", "SilverNebula  ✓  connected") +
		"\n" +
		lipgloss.NewStyle().Foreground(colorDim).Render("  "+strings.Repeat("─", 50)) + "\n\n" +
		lipgloss.NewStyle().Foreground(colorDim2).Italic(true).Render("  press any key to go back"),
	)

	return header + "\n" + body
}

func row(k, v string) string {
	return lipgloss.NewStyle().Foreground(colorGray).Render(fmt.Sprintf("  %-10s", k)) +
		lipgloss.NewStyle().Foreground(colorWhite).Render(v) + "\n"
}

func init() { _ = fmt.Sprintf }
