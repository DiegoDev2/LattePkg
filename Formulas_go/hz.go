package main

// Hz - Golang HTTP framework for microservices
// Homepage: https://github.com/cloudwego/hertz

import (
	"fmt"
	
	"os/exec"
)

func installHz() {
	// Método 1: Descargar y extraer .tar.gz
	hz_tar_url := "https://github.com/cloudwego/hertz/archive/refs/tags/cmd/hz/v0.9.1.tar.gz"
	hz_cmd_tar := exec.Command("curl", "-L", hz_tar_url, "-o", "package.tar.gz")
	err := hz_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hz_zip_url := "https://github.com/cloudwego/hertz/archive/refs/tags/cmd/hz/v0.9.1.zip"
	hz_cmd_zip := exec.Command("curl", "-L", hz_zip_url, "-o", "package.zip")
	err = hz_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hz_bin_url := "https://github.com/cloudwego/hertz/archive/refs/tags/cmd/hz/v0.9.1.bin"
	hz_cmd_bin := exec.Command("curl", "-L", hz_bin_url, "-o", "binary.bin")
	err = hz_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hz_src_url := "https://github.com/cloudwego/hertz/archive/refs/tags/cmd/hz/v0.9.1.src.tar.gz"
	hz_cmd_src := exec.Command("curl", "-L", hz_src_url, "-o", "source.tar.gz")
	err = hz_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hz_cmd_direct := exec.Command("./binary")
	err = hz_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
