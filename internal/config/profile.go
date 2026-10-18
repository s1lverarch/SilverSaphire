// SilverSaphire (ss) - TOML Profile (reproducible installs)
// Author: Silvercore (s1lverarch) <s1lv3rcore@proton.me>
// License: GPL-v3.0

package config

import (
	"fmt"
	"os"
	"time"
)

// Profile mirrors InstallConfig as TOML-serializable struct
type Profile struct {
	Meta       MetaSection       `toml:"meta"`
	System     SystemSection     `toml:"system"`
	User       UserSection       `toml:"user"`
	Desktop    DesktopSection    `toml:"desktop"`
	Packages   PackagesSection   `toml:"packages"`
	Security   SecuritySection   `toml:"security"`
	PostInstall PostInstallSection `toml:"post_install"`
}

type MetaSection struct {
	CreatedBy string `toml:"created_by"`
	Version   string `toml:"version"`
	Date      string `toml:"date"`
	Website   string `toml:"website"`
}

type SystemSection struct {
	Language   string `toml:"language"`
	Keyboard   string `toml:"keyboard"`
	Timezone   string `toml:"timezone"`
	Hostname   string `toml:"hostname"`
	Disk       string `toml:"disk"`
	Filesystem string `toml:"filesystem"`
	Bootloader string `toml:"bootloader"`
	Kernel     string `toml:"kernel"`
	Swap       string `toml:"swap"`
	GPU        string `toml:"gpu_driver"`
	Network    string `toml:"network"`
	SSH        string `toml:"ssh"`
}

type UserSection struct {
	Username string `toml:"username"`
	Shell    string `toml:"shell"`
	AUR      string `toml:"aur_helper"`
}

type DesktopSection struct {
	Environment string `toml:"environment"`
	Flatpak     string `toml:"flatpak"`
}

type PackagesSection struct {
	Extra []string `toml:"extra"`
}

type SecuritySection struct {
	Profile  string `toml:"profile"`
	AppArmor bool   `toml:"apparmor"`
	SELinux  bool   `toml:"selinux"`
	Firewall string `toml:"firewall"`
}

type PostInstallSection struct {
	Dotfiles string   `toml:"dotfiles"`
	Hooks    []string `toml:"hooks"`
}

func DefaultProfile() Profile {
	return Profile{
		Meta: MetaSection{
			CreatedBy: "SilverSaphire (ss) v0.1.0",
			Version:   "0.1.0",
			Date:      time.Now().Format("2006-01-02"),
			Website:   "https://s1lverarch.github.io/linuxdistro/",
		},
		System: SystemSection{
			Language:   "English",
			Keyboard:   "us",
			Timezone:   "UTC",
			Hostname:   "silverarch",
			Disk:       "Auto partition",
			Filesystem: "ext4",
			Bootloader: "GRUB",
			Kernel:     "linux",
			Swap:       "zram",
			GPU:        "Auto-detect",
			Network:    "NetworkManager",
			SSH:        "Disable",
		},
		User: UserSection{
			Username: "user",
			Shell:    "/bin/bash",
			AUR:      "paru",
		},
		Desktop: DesktopSection{
			Environment: "Hyprland",
			Flatpak:     "Enable Flatpak + Flathub",
		},
		Security: SecuritySection{
			Profile:  "normal",
			AppArmor: true,
			SELinux:  false,
			Firewall: "nftables",
		},
		PostInstall: PostInstallSection{
			Dotfiles: "",
			Hooks:    []string{},
		},
	}
}

func SaveProfile(p Profile, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	// manual TOML write (no external dep needed)
	fmt.Fprintf(f, "# SilverSaphire Install Profile\n")
	fmt.Fprintf(f, "# Author: Silvercore (s1lverarch) <s1lv3rcore@proton.me>\n")
	fmt.Fprintf(f, "# License: GPL-v3.0\n\n")

	fmt.Fprintf(f, "[meta]\n")
	fmt.Fprintf(f, "created_by = %q\n", p.Meta.CreatedBy)
	fmt.Fprintf(f, "version    = %q\n", p.Meta.Version)
	fmt.Fprintf(f, "date       = %q\n", p.Meta.Date)
	fmt.Fprintf(f, "website    = %q\n\n", p.Meta.Website)

	fmt.Fprintf(f, "[system]\n")
	fmt.Fprintf(f, "language   = %q\n", p.System.Language)
	fmt.Fprintf(f, "keyboard   = %q\n", p.System.Keyboard)
	fmt.Fprintf(f, "timezone   = %q\n", p.System.Timezone)
	fmt.Fprintf(f, "hostname   = %q\n", p.System.Hostname)
	fmt.Fprintf(f, "disk       = %q\n", p.System.Disk)
	fmt.Fprintf(f, "filesystem = %q\n", p.System.Filesystem)
	fmt.Fprintf(f, "bootloader = %q\n", p.System.Bootloader)
	fmt.Fprintf(f, "kernel     = %q\n", p.System.Kernel)
	fmt.Fprintf(f, "swap       = %q\n", p.System.Swap)
	fmt.Fprintf(f, "gpu_driver = %q\n", p.System.GPU)
	fmt.Fprintf(f, "network    = %q\n", p.System.Network)
	fmt.Fprintf(f, "ssh        = %q\n\n", p.System.SSH)

	fmt.Fprintf(f, "[user]\n")
	fmt.Fprintf(f, "username   = %q\n", p.User.Username)
	fmt.Fprintf(f, "shell      = %q\n", p.User.Shell)
	fmt.Fprintf(f, "aur_helper = %q\n\n", p.User.AUR)

	fmt.Fprintf(f, "[desktop]\n")
	fmt.Fprintf(f, "environment = %q\n", p.Desktop.Environment)
	fmt.Fprintf(f, "flatpak     = %q\n\n", p.Desktop.Flatpak)

	fmt.Fprintf(f, "[security]\n")
	fmt.Fprintf(f, "profile  = %q\n", p.Security.Profile)
	fmt.Fprintf(f, "apparmor = %v\n", p.Security.AppArmor)
	fmt.Fprintf(f, "selinux  = %v\n", p.Security.SELinux)
	fmt.Fprintf(f, "firewall = %q\n\n", p.Security.Firewall)

	fmt.Fprintf(f, "[post_install]\n")
	fmt.Fprintf(f, "dotfiles = %q\n", p.PostInstall.Dotfiles)
	fmt.Fprintf(f, "hooks    = []\n")

	return nil
}
