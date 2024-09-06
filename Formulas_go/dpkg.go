package main

// Dpkg - Debian package management system
// Homepage: https://wiki.debian.org/Teams/Dpkg

import (
	"fmt"
	
	"os/exec"
)

func installDpkg() {
	// Método 1: Descargar y extraer .tar.gz
	dpkg_tar_url := "https://deb.debian.org/debian/pool/main/d/dpkg/dpkg_1.22.11.tar.xz"
	dpkg_cmd_tar := exec.Command("curl", "-L", dpkg_tar_url, "-o", "package.tar.gz")
	err := dpkg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dpkg_zip_url := "https://deb.debian.org/debian/pool/main/d/dpkg/dpkg_1.22.11.tar.xz"
	dpkg_cmd_zip := exec.Command("curl", "-L", dpkg_zip_url, "-o", "package.zip")
	err = dpkg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dpkg_bin_url := "https://deb.debian.org/debian/pool/main/d/dpkg/dpkg_1.22.11.tar.xz"
	dpkg_cmd_bin := exec.Command("curl", "-L", dpkg_bin_url, "-o", "binary.bin")
	err = dpkg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dpkg_src_url := "https://deb.debian.org/debian/pool/main/d/dpkg/dpkg_1.22.11.tar.xz"
	dpkg_cmd_src := exec.Command("curl", "-L", dpkg_src_url, "-o", "source.tar.gz")
	err = dpkg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dpkg_cmd_direct := exec.Command("./binary")
	err = dpkg_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: po4a")
exec.Command("latte", "install", "po4a")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: gnu-tar")
exec.Command("latte", "install", "gnu-tar")
	fmt.Println("Instalando dependencia: gpatch")
exec.Command("latte", "install", "gpatch")
	fmt.Println("Instalando dependencia: libmd")
exec.Command("latte", "install", "libmd")
	fmt.Println("Instalando dependencia: perl")
exec.Command("latte", "install", "perl")
	fmt.Println("Instalando dependencia: xz")
exec.Command("latte", "install", "xz")
}
