

# SilverSaphire / The Arch Install Alternative

> The Arch Linux Alternative Installer — Go TUI + Nim hooks

[![Typing SVG](https://readme-typing-svg.demolab.com?font=Fira+Code&weight=100&size=17&pause=1000&color=868686&vCenter=true&width=435&lines=Just+A+Installer+For+Arch+And+Based-Distors;Archinstall+Alternative;CachyOS+Kernel+And+linux-rt+...)](https://git.io/typing-svg)

**Author:** Silvercore (s1lverarch) <s1lv3rcore@proton.me>

**Website:** https://s1lverarch.github.io/linuxdistro/

**License:** GPL-v3.0

---
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue?style=for-the-badge&logo=gnu&logoColor=white)](LICENSE)
[![Platform](https://img.shields.io/badge/Platform-Linux-FCC624?style=for-the-badge&logo=linux&logoColor=black)](#)
[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev)

---
<img width="950" height="329" alt="image" src="https://github.com/user-attachments/assets/d4ada8d0-6ab9-4b92-abe2-bd51c93aae7d" />
<img width="944" height="354" alt="image" src="https://github.com/user-attachments/assets/7c7289bb-8f60-4d40-a47c-18e2b276be76" />
<img width="947" height="403" alt="image" src="https://github.com/user-attachments/assets/e53b25f1-6e56-42e0-b08a-435ce7865a2e" />



## Features

- Boot animation → interactive TUI menu
- Full installer: Language, Keyboard, Timezone, Disk, Filesystem, Bootloader, Kernel, Network+VPN, User, Desktop/WM, Packages, AUR, Security Wizard, Swap/Zram, GPU, Flatpak, SSH, Dotfiles
- All DEs/WMs supported (GNOME, KDE, Hyprland, i3, Sway...)
- All kernels (linux, lts, zen, hardened, cachyos, tkg, custom)
- Reproducible TOML profile system
- Post-install Nim hooks (firewall, zram, flatpak, dotfiles, SilverNebula repo)
- Written in **Go** (TUI) + **Nim** (system hooks)

## Install

```bash
# From SilverNebula repo
sudo pacman -S silversaphire

# Run as root
sudo ss
```

## Build from source

```bash
git clone https://github.com/s1lverarch/silversaphire
cd silversaphire
go build -o ss ./cmd/
nim compile --opt:size -o:ss-hooks nim/hooks/post_install.nim
sudo ./ss
```

## Profile (reproducible installs)

```toml
# silverarch-profile.toml
[system]
language   = "English"
kernel     = "linux-zen"
filesystem = "btrfs"
bootloader = "systemd-boot"

[desktop]
environment = "Hyprland"
```

```bash
sudo ss --profile silverarch-profile.toml
```
## Contributing

Issues and PRs welcome. This project follows the SilverArch Linux
open-source philosophy — build it minimal, ship it open.

```bash
git clone https://github.com/s1lverarch/SilverSaphire.git
cd SilverSaphire
go build ./...
go vet ./...
```

---
## 󰚰  New Update : SilverSaphire v0.2.6 / 2026-10-25

What's New :

- memory maps + CPU init + ACPI + security mitigations + systemd
- About screen — opens with Enter on "About", goes back with any key
- Skip — now s + Tab + Ctrl+S all work
- Theme chooser — the first step in the installer (StepTheme)
- Fix Some Bugs In Install Steps
- Update GoLang From 1.22 To 1.27.2 

<div align="center">

<img width="949" height="485" alt="image" src="https://github.com/user-attachments/assets/41aba533-3478-4c57-a206-059d11798d5f" />
<img width="950" height="387" alt="image" src="https://github.com/user-attachments/assets/b92fa4db-ce39-4291-8e6f-e13993b83a55" />
<img width="952" height="444" alt="image" src="https://github.com/user-attachments/assets/3ab3519c-9e4a-478d-8740-f2956d65eee8" />

---

### Part of the SilverArch Linux Project

[![Wiki](https://img.shields.io/badge/Wiki-4DA6FF?style=for-the-badge&logo=wikidot&logoColor=white)](https://silverarchlinux.miraheze.org)
[![Discord](https://img.shields.io/badge/Discord-5865F2?style=for-the-badge&logo=discord&logoColor=white)](https://discord.gg/eBgVHSPru9)
[![GitHub](https://img.shields.io/badge/GitHub-181717?style=for-the-badge&logo=github&logoColor=white)](https://github.com/s1lverarch)

<br>
