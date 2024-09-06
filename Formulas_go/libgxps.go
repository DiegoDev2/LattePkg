package main

// Libgxps - GObject based library for handling and rendering XPS documents
// Homepage: https://wiki.gnome.org/Projects/libgxps

import (
	"fmt"
	
	"os/exec"
)

func installLibgxps() {
	// Método 1: Descargar y extraer .tar.gz
	libgxps_tar_url := "https://download.gnome.org/sources/libgxps/0.3/libgxps-0.3.2.tar.xz"
	libgxps_cmd_tar := exec.Command("curl", "-L", libgxps_tar_url, "-o", "package.tar.gz")
	err := libgxps_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libgxps_zip_url := "https://download.gnome.org/sources/libgxps/0.3/libgxps-0.3.2.tar.xz"
	libgxps_cmd_zip := exec.Command("curl", "-L", libgxps_zip_url, "-o", "package.zip")
	err = libgxps_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libgxps_bin_url := "https://download.gnome.org/sources/libgxps/0.3/libgxps-0.3.2.tar.xz"
	libgxps_cmd_bin := exec.Command("curl", "-L", libgxps_bin_url, "-o", "binary.bin")
	err = libgxps_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libgxps_src_url := "https://download.gnome.org/sources/libgxps/0.3/libgxps-0.3.2.tar.xz"
	libgxps_cmd_src := exec.Command("curl", "-L", libgxps_src_url, "-o", "source.tar.gz")
	err = libgxps_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libgxps_cmd_direct := exec.Command("./binary")
	err = libgxps_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gobject-introspection")
exec.Command("latte", "install", "gobject-introspection")
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: libarchive")
exec.Command("latte", "install", "libarchive")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libtiff")
exec.Command("latte", "install", "libtiff")
	fmt.Println("Instalando dependencia: little-cms2")
exec.Command("latte", "install", "little-cms2")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
