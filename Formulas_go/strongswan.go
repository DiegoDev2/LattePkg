package main

// Strongswan - VPN based on IPsec
// Homepage: https://www.strongswan.org

import (
	"fmt"
	
	"os/exec"
)

func installStrongswan() {
	// Método 1: Descargar y extraer .tar.gz
	strongswan_tar_url := "https://download.strongswan.org/strongswan-5.9.14.tar.bz2"
	strongswan_cmd_tar := exec.Command("curl", "-L", strongswan_tar_url, "-o", "package.tar.gz")
	err := strongswan_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	strongswan_zip_url := "https://download.strongswan.org/strongswan-5.9.14.tar.bz2"
	strongswan_cmd_zip := exec.Command("curl", "-L", strongswan_zip_url, "-o", "package.zip")
	err = strongswan_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	strongswan_bin_url := "https://download.strongswan.org/strongswan-5.9.14.tar.bz2"
	strongswan_cmd_bin := exec.Command("curl", "-L", strongswan_bin_url, "-o", "binary.bin")
	err = strongswan_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	strongswan_src_url := "https://download.strongswan.org/strongswan-5.9.14.tar.bz2"
	strongswan_cmd_src := exec.Command("curl", "-L", strongswan_src_url, "-o", "source.tar.gz")
	err = strongswan_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	strongswan_cmd_direct := exec.Command("./binary")
	err = strongswan_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: bison")
exec.Command("latte", "install", "bison")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
