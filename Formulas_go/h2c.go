package main

// H2c - Headers 2 curl
// Homepage: https://curl.se/h2c/

import (
	"fmt"
	
	"os/exec"
)

func installH2c() {
	// Método 1: Descargar y extraer .tar.gz
	h2c_tar_url := "https://github.com/curl/h2c/archive/refs/tags/1.0.tar.gz"
	h2c_cmd_tar := exec.Command("curl", "-L", h2c_tar_url, "-o", "package.tar.gz")
	err := h2c_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	h2c_zip_url := "https://github.com/curl/h2c/archive/refs/tags/1.0.zip"
	h2c_cmd_zip := exec.Command("curl", "-L", h2c_zip_url, "-o", "package.zip")
	err = h2c_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	h2c_bin_url := "https://github.com/curl/h2c/archive/refs/tags/1.0.bin"
	h2c_cmd_bin := exec.Command("curl", "-L", h2c_bin_url, "-o", "binary.bin")
	err = h2c_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	h2c_src_url := "https://github.com/curl/h2c/archive/refs/tags/1.0.src.tar.gz"
	h2c_cmd_src := exec.Command("curl", "-L", h2c_src_url, "-o", "source.tar.gz")
	err = h2c_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	h2c_cmd_direct := exec.Command("./binary")
	err = h2c_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
