# SilverSaphire (ss) - Post-install hooks (Nim)
# Author: Silvercore (s1lverarch) <s1lv3rcore@proton.me>
# Website: https://s1lverarch.github.io/linuxdistro/
# License: GPL-v3.0

import os, osproc, strutils, strformat, terminal

# ── Colors ───────────────────────────────────────────────────────────────────
proc ok(msg: string) =
  styledEcho fgGreen, "[ OK ] ", resetStyle, msg

proc info(msg: string) =
  styledEcho fgCyan, "[ >> ] ", resetStyle, msg

proc err(msg: string) =
  styledEcho fgRed, "[ !! ] ", resetStyle, msg

proc run(cmd: string): bool =
  info &"Running: {cmd}"
  let (output, code) = execCmdEx(cmd)
  if code != 0:
    err &"Failed: {output}"
    return false
  ok cmd
  return true

# ── Hooks ────────────────────────────────────────────────────────────────────

proc hookEnableServices() =
  info "Enabling system services..."
  discard run("systemctl enable NetworkManager")
  discard run("systemctl enable bluetooth")
  discard run("systemctl enable sshd")
  ok "Services enabled"

proc hookSetupZram() =
  info "Setting up zram..."
  discard run("modprobe zram")
  writeFile("/etc/systemd/zram-generator.conf",
    "[zram0]\nzram-size = ram / 2\ncompression-algorithm = zstd\n")
  ok "Zram configured"

proc hookSetupFlatpak() =
  info "Setting up Flatpak + Flathub..."
  discard run("flatpak remote-add --if-not-exists flathub https://dl.flathub.org/repo/flathub.flatpakrepo")
  ok "Flathub added"

proc hookSetupAppArmor() =
  info "Enabling AppArmor..."
  discard run("systemctl enable apparmor")
  ok "AppArmor enabled"

proc hookSetupFirewall() =
  info "Setting up nftables firewall..."
  let rules = """
#!/usr/sbin/nft -f
flush ruleset
table inet filter {
  chain input {
    type filter hook input priority 0; policy drop;
    iif lo accept
    ct state established,related accept
    tcp dport 22 accept
    reject with icmpx admin-prohibited
  }
  chain forward { type filter hook forward priority 0; policy drop; }
  chain output  { type filter hook output  priority 0; policy accept; }
}
"""
  writeFile("/etc/nftables.conf", rules)
  discard run("systemctl enable nftables")
  ok "Firewall configured"

proc hookImportDotfiles(url: string) =
  if url.len == 0:
    return
  info &"Importing dotfiles from {url}..."
  discard run(&"git clone {url} /tmp/dotfiles")
  if fileExists("/tmp/dotfiles/install.sh"):
    discard run("bash /tmp/dotfiles/install.sh")
  ok "Dotfiles imported"

proc hookSilverNebulaRepo() =
  info "Adding SilverNebula repo to pacman..."
  let conf = readFile("/etc/pacman.conf")
  if "[silvernebula]" notin conf:
    let entry = "\n[silvernebula]\nSigLevel = Optional TrustAll\nServer = https://raw.githubusercontent.com/s1lverarch/silverarch-repo/main/x86_64/\n"
    writeFile("/etc/pacman.conf", conf & entry)
  discard run("pacman -Syy")
  ok "SilverNebula repo added"

proc hookWelcomeMessage() =
  info "Setting up welcome message..."
  let motd = """



  Welcome to Installtion Page
  Web  : https://s1lverarch.github.io/linuxdistro/
  Wiki : https://silverarchlinux.miraheze.org/wiki/
  By   : Silvercore (s1lverarch)

"""
  writeFile("/etc/motd", motd)
  ok "Welcome message set"

# ── Main ─────────────────────────────────────────────────────────────────────
when isMainModule:
  let dotfilesUrl = if paramCount() > 0: paramStr(1) else: ""

  info "SilverSaphire post-install hooks starting..."
  hookEnableServices()
  hookSetupZram()
  hookSetupFlatpak()
  hookSetupAppArmor()
  hookSetupFirewall()
  hookSilverNebulaRepo()
  hookImportDotfiles(dotfilesUrl)
  hookWelcomeMessage()
  ok "All post-install hooks complete!"
