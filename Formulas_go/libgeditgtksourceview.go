package main

// LibgeditGtksourceview - Text editor widget for code editing
// Homepage: https://gitlab.gnome.org/World/gedit/libgedit-gtksourceview

import (
	"fmt"
	
	"os/exec"
)

func installLibgeditGtksourceview() {
	// Método 1: Descargar y extraer .tar.gz
	libgeditgtksourceview_tar_url := "https://gitlab.gnome.org/World/gedit/libgedit-gtksourceview/-/archive/299.2.1/libgedit-gtksourceview-299.2.1.tar.bz2"
	libgeditgtksourceview_cmd_tar := exec.Command("curl", "-L", libgeditgtksourceview_tar_url, "-o", "package.tar.gz")
	err := libgeditgtksourceview_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libgeditgtksourceview_zip_url := "https://gitlab.gnome.org/World/gedit/libgedit-gtksourceview/-/archive/299.2.1/libgedit-gtksourceview-299.2.1.tar.bz2"
	libgeditgtksourceview_cmd_zip := exec.Command("curl", "-L", libgeditgtksourceview_zip_url, "-o", "package.zip")
	err = libgeditgtksourceview_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libgeditgtksourceview_bin_url := "https://gitlab.gnome.org/World/gedit/libgedit-gtksourceview/-/archive/299.2.1/libgedit-gtksourceview-299.2.1.tar.bz2"
	libgeditgtksourceview_cmd_bin := exec.Command("curl", "-L", libgeditgtksourceview_bin_url, "-o", "binary.bin")
	err = libgeditgtksourceview_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libgeditgtksourceview_src_url := "https://gitlab.gnome.org/World/gedit/libgedit-gtksourceview/-/archive/299.2.1/libgedit-gtksourceview-299.2.1.tar.bz2"
	libgeditgtksourceview_cmd_src := exec.Command("curl", "-L", libgeditgtksourceview_src_url, "-o", "source.tar.gz")
	err = libgeditgtksourceview_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libgeditgtksourceview_cmd_direct := exec.Command("./binary")
	err = libgeditgtksourceview_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
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
	fmt.Println("Instalando dependencia: gdk-pixbuf")
exec.Command("latte", "install", "gdk-pixbuf")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gtk+3")
exec.Command("latte", "install", "gtk+3")
	fmt.Println("Instalando dependencia: libxml2")
exec.Command("latte", "install", "libxml2")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
