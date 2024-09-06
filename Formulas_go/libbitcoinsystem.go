package main

// LibbitcoinSystem - Bitcoin Cross-Platform C++ Development Toolkit
// Homepage: https://github.com/libbitcoin/libbitcoin-system

import (
	"fmt"
	
	"os/exec"
)

func installLibbitcoinSystem() {
	// Método 1: Descargar y extraer .tar.gz
	libbitcoinsystem_tar_url := "https://github.com/libbitcoin/libbitcoin-system/archive/refs/tags/v3.8.0.tar.gz"
	libbitcoinsystem_cmd_tar := exec.Command("curl", "-L", libbitcoinsystem_tar_url, "-o", "package.tar.gz")
	err := libbitcoinsystem_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libbitcoinsystem_zip_url := "https://github.com/libbitcoin/libbitcoin-system/archive/refs/tags/v3.8.0.zip"
	libbitcoinsystem_cmd_zip := exec.Command("curl", "-L", libbitcoinsystem_zip_url, "-o", "package.zip")
	err = libbitcoinsystem_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libbitcoinsystem_bin_url := "https://github.com/libbitcoin/libbitcoin-system/archive/refs/tags/v3.8.0.bin"
	libbitcoinsystem_cmd_bin := exec.Command("curl", "-L", libbitcoinsystem_bin_url, "-o", "binary.bin")
	err = libbitcoinsystem_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libbitcoinsystem_src_url := "https://github.com/libbitcoin/libbitcoin-system/archive/refs/tags/v3.8.0.src.tar.gz"
	libbitcoinsystem_cmd_src := exec.Command("curl", "-L", libbitcoinsystem_src_url, "-o", "source.tar.gz")
	err = libbitcoinsystem_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libbitcoinsystem_cmd_direct := exec.Command("./binary")
	err = libbitcoinsystem_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: boost@1.76")
exec.Command("latte", "install", "boost@1.76")
}
