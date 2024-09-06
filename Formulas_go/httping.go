package main

// Httping - Ping-like tool for HTTP requests
// Homepage: https://github.com/folkertvanheusden/HTTPing

import (
	"fmt"
	
	"os/exec"
)

func installHttping() {
	// Método 1: Descargar y extraer .tar.gz
	httping_tar_url := "https://github.com/folkertvanheusden/HTTPing/archive/refs/tags/v2.9.tar.gz"
	httping_cmd_tar := exec.Command("curl", "-L", httping_tar_url, "-o", "package.tar.gz")
	err := httping_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	httping_zip_url := "https://github.com/folkertvanheusden/HTTPing/archive/refs/tags/v2.9.zip"
	httping_cmd_zip := exec.Command("curl", "-L", httping_zip_url, "-o", "package.zip")
	err = httping_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	httping_bin_url := "https://github.com/folkertvanheusden/HTTPing/archive/refs/tags/v2.9.bin"
	httping_cmd_bin := exec.Command("curl", "-L", httping_bin_url, "-o", "binary.bin")
	err = httping_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	httping_src_url := "https://github.com/folkertvanheusden/HTTPing/archive/refs/tags/v2.9.src.tar.gz"
	httping_cmd_src := exec.Command("curl", "-L", httping_src_url, "-o", "source.tar.gz")
	err = httping_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	httping_cmd_direct := exec.Command("./binary")
	err = httping_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
