package main

// Reprepro - Debian package repository manager
// Homepage: https://salsa.debian.org/debian/reprepro

import (
	"fmt"
	
	"os/exec"
)

func installReprepro() {
	// Método 1: Descargar y extraer .tar.gz
	reprepro_tar_url := "https://deb.debian.org/debian/pool/main/r/reprepro/reprepro_5.3.1.orig.tar.xz"
	reprepro_cmd_tar := exec.Command("curl", "-L", reprepro_tar_url, "-o", "package.tar.gz")
	err := reprepro_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	reprepro_zip_url := "https://deb.debian.org/debian/pool/main/r/reprepro/reprepro_5.3.1.orig.tar.xz"
	reprepro_cmd_zip := exec.Command("curl", "-L", reprepro_zip_url, "-o", "package.zip")
	err = reprepro_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	reprepro_bin_url := "https://deb.debian.org/debian/pool/main/r/reprepro/reprepro_5.3.1.orig.tar.xz"
	reprepro_cmd_bin := exec.Command("curl", "-L", reprepro_bin_url, "-o", "binary.bin")
	err = reprepro_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	reprepro_src_url := "https://deb.debian.org/debian/pool/main/r/reprepro/reprepro_5.3.1.orig.tar.xz"
	reprepro_cmd_src := exec.Command("curl", "-L", reprepro_src_url, "-o", "source.tar.gz")
	err = reprepro_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	reprepro_cmd_direct := exec.Command("./binary")
	err = reprepro_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: berkeley-db@5")
exec.Command("latte", "install", "berkeley-db@5")
	fmt.Println("Instalando dependencia: gpgme")
exec.Command("latte", "install", "gpgme")
	fmt.Println("Instalando dependencia: libarchive")
exec.Command("latte", "install", "libarchive")
	fmt.Println("Instalando dependencia: libgpg-error")
exec.Command("latte", "install", "libgpg-error")
	fmt.Println("Instalando dependencia: xz")
exec.Command("latte", "install", "xz")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
	fmt.Println("Instalando dependencia: gcc")
exec.Command("latte", "install", "gcc")
}
