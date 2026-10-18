# SilverSaphire

> The SilverArch Linux Installer — Go TUI + Nim hooks

**Author:** Silvercore (s1lverarch) <s1lv3rcore@proton.me>
**Website:** https://s1lverarch.github.io/linuxdistro/
**License:** GPL-v3.0

---

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
