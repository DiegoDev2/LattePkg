package main

// Gmssl - Toolkit for Chinese national cryptographic standards
// Homepage: https://github.com/guanzhi/GmSSL

import (
	"fmt"
	
	"os/exec"
)

func installGmssl() {
	// Método 1: Descargar y extraer .tar.gz
	gmssl_tar_url := "https://github.com/guanzhi/GmSSL/archive/refs/tags/v3.1.1.tar.gz"
	gmssl_cmd_tar := exec.Command("curl", "-L", gmssl_tar_url, "-o", "package.tar.gz")
	err := gmssl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gmssl_zip_url := "https://github.com/guanzhi/GmSSL/archive/refs/tags/v3.1.1.zip"
	gmssl_cmd_zip := exec.Command("curl", "-L", gmssl_zip_url, "-o", "package.zip")
	err = gmssl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gmssl_bin_url := "https://github.com/guanzhi/GmSSL/archive/refs/tags/v3.1.1.bin"
	gmssl_cmd_bin := exec.Command("curl", "-L", gmssl_bin_url, "-o", "binary.bin")
	err = gmssl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gmssl_src_url := "https://github.com/guanzhi/GmSSL/archive/refs/tags/v3.1.1.src.tar.gz"
	gmssl_cmd_src := exec.Command("curl", "-L", gmssl_src_url, "-o", "source.tar.gz")
	err = gmssl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gmssl_cmd_direct := exec.Command("./binary")
	err = gmssl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
