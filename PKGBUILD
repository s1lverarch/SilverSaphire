# Maintainer: Silvercore (s1lverarch) <s1lv3rcore@proton.me>
# Website: https://s1lverarch.github.io/linuxdistro/

pkgname=silversaphire
pkgver=0.2.6
pkgrel=1
pkgdesc="SilverArch Linux installer — Go TUI + Nim hooks"
arch=('x86_64')
url="https://github.com/s1lverarch/silversaphire"
license=('GPL3')
depends=('glibc' 'nim' 'git')
makedepends=('go')
source=("$pkgname-$pkgver.tar.gz::https://github.com/s1lverarch/$pkgname/archive/v$pkgver.tar.gz")
sha256sums=('SKIP')

build() {
  cd "$srcdir/$pkgname-$pkgver"
  export CGO_ENABLED=0
  export GOFLAGS="-buildmode=pie -trimpath -mod=readonly -modcacherw"
  go build -o ss ./cmd/
  nim compile --opt:size -o:ss-hooks nim/hooks/post_install.nim
}

package() {
  cd "$srcdir/$pkgname-$pkgver"
  install -Dm755 ss          "$pkgdir/usr/bin/ss"
  install -Dm755 ss-hooks    "$pkgdir/usr/lib/silversaphire/hooks/post_install"
  install -Dm644 LICENSE     "$pkgdir/usr/share/licenses/$pkgname/LICENSE"
}
