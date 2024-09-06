package main

// Softhsm - Cryptographic store accessible through a PKCS#11 interface
// Homepage: https://www.opendnssec.org/softhsm/

import (
	"fmt"
	
	"os/exec"
)

func installSofthsm() {
	// Método 1: Descargar y extraer .tar.gz
	softhsm_tar_url := "https://dist.opendnssec.org/source/softhsm-2.6.1.tar.gz"
	softhsm_cmd_tar := exec.Command("curl", "-L", softhsm_tar_url, "-o", "package.tar.gz")
	err := softhsm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	softhsm_zip_url := "https://dist.opendnssec.org/source/softhsm-2.6.1.zip"
	softhsm_cmd_zip := exec.Command("curl", "-L", softhsm_zip_url, "-o", "package.zip")
	err = softhsm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	softhsm_bin_url := "https://dist.opendnssec.org/source/softhsm-2.6.1.bin"
	softhsm_cmd_bin := exec.Command("curl", "-L", softhsm_bin_url, "-o", "binary.bin")
	err = softhsm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	softhsm_src_url := "https://dist.opendnssec.org/source/softhsm-2.6.1.src.tar.gz"
	softhsm_cmd_src := exec.Command("curl", "-L", softhsm_src_url, "-o", "source.tar.gz")
	err = softhsm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	softhsm_cmd_direct := exec.Command("./binary")
	err = softhsm_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
