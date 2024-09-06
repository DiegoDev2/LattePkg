package main

// Freeciv - Free and Open Source empire-building strategy game
// Homepage: https://freeciv.org/

import (
	"fmt"
	
	"os/exec"
)

func installFreeciv() {
	// Método 1: Descargar y extraer .tar.gz
	freeciv_tar_url := "https://downloads.sourceforge.net/project/freeciv/Freeciv%203.1/3.1.2/freeciv-3.1.2.tar.xz"
	freeciv_cmd_tar := exec.Command("curl", "-L", freeciv_tar_url, "-o", "package.tar.gz")
	err := freeciv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	freeciv_zip_url := "https://downloads.sourceforge.net/project/freeciv/Freeciv%203.1/3.1.2/freeciv-3.1.2.tar.xz"
	freeciv_cmd_zip := exec.Command("curl", "-L", freeciv_zip_url, "-o", "package.zip")
	err = freeciv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	freeciv_bin_url := "https://downloads.sourceforge.net/project/freeciv/Freeciv%203.1/3.1.2/freeciv-3.1.2.tar.xz"
	freeciv_cmd_bin := exec.Command("curl", "-L", freeciv_bin_url, "-o", "binary.bin")
	err = freeciv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	freeciv_src_url := "https://downloads.sourceforge.net/project/freeciv/Freeciv%203.1/3.1.2/freeciv-3.1.2.tar.xz"
	freeciv_cmd_src := exec.Command("curl", "-L", freeciv_src_url, "-o", "source.tar.gz")
	err = freeciv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	freeciv_cmd_direct := exec.Command("./binary")
	err = freeciv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: adwaita-icon-theme")
exec.Command("latte", "install", "adwaita-icon-theme")
	fmt.Println("Instalando dependencia: at-spi2-core")
exec.Command("latte", "install", "at-spi2-core")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: gdk-pixbuf")
exec.Command("latte", "install", "gdk-pixbuf")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gtk+3")
exec.Command("latte", "install", "gtk+3")
	fmt.Println("Instalando dependencia: harfbuzz")
exec.Command("latte", "install", "harfbuzz")
	fmt.Println("Instalando dependencia: icu4c")
exec.Command("latte", "install", "icu4c")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
	fmt.Println("Instalando dependencia: sdl2")
exec.Command("latte", "install", "sdl2")
	fmt.Println("Instalando dependencia: sdl2_mixer")
exec.Command("latte", "install", "sdl2_mixer")
	fmt.Println("Instalando dependencia: sqlite")
exec.Command("latte", "install", "sqlite")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
}
