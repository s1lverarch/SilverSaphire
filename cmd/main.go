// SilverSaphire (ss) - The SilverArch Linux Installer
// Author: Silvercore (s1lverarch) <s1lv3rcore@proton.me>
// Website: https://s1lverarch.github.io/linuxdistro/
// License: GPL-v3.0
// Copyright 2026 Silvercore

package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/s1lverarch/SilverSaphire/internal/ui"
)

const asciiLogo = `
  ███████╗███████╗
  ██╔════╝██╔════╝
  ███████╗███████╗
  ╚════██║╚════██║
  ███████║███████║
  ╚══════╝╚══════╝`

func main() {
	if os.Geteuid() != 0 {
		fmt.Fprintf(os.Stderr, "SilverSaphire must be run as root\n")
		os.Exit(1)
	}

	p := tea.NewProgram(
		ui.NewBootAnimation(asciiLogo),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
