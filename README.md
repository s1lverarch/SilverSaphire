# SilverSaphire

<img width="800" height="450" alt="ss" src="https://github.com/user-attachments/assets/74b34179-2ec4-4861-a84c-4d59b8a56566" />

> The SilverArch Linux Installer — Go TUI + Nim hooks

[![Typing SVG](https://readme-typing-svg.demolab.com?font=Fira+Code&weight=100&size=17&pause=1000&color=868686&vCenter=true&width=435&lines=Just+A+Installer+For+Arch+And+Based-Distors;Archinstall+Alternative;CachyOS+Kernel+And+linux-rt+...)](https://git.io/typing-svg)

**Author:** Silvercore (s1lverarch) <s1lv3rcore@proton.me>

**Website:** https://s1lverarch.github.io/linuxdistro/

**License:** GPL-v3.0

---
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue?style=for-the-badge&logo=gnu&logoColor=white)](LICENSE)
[![Platform](https://img.shields.io/badge/Platform-Linux-FCC624?style=for-the-badge&logo=linux&logoColor=black)](#)
[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev)

---
<img width="947" height="1045" alt="image" src="https://github.com/user-attachments/assets/70a179ab-eb6b-4dc9-90d0-c097e4dc07f2" />
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

<div align="center">

### Part of the SilverArch Linux Project

[![Wiki](https://img.shields.io/badge/Wiki-4DA6FF?style=for-the-badge&logo=wikidot&logoColor=white)](https://silverarchlinux.miraheze.org)
[![Discord](https://img.shields.io/badge/Discord-5865F2?style=for-the-badge&logo=discord&logoColor=white)](https://discord.gg/eBgVHSPru9)
[![GitHub](https://img.shields.io/badge/GitHub-181717?style=for-the-badge&logo=github&logoColor=white)](https://github.com/s1lverarch)

<br>
