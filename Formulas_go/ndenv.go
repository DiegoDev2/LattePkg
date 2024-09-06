package main

// Ndenv - Node version manager
// Homepage: https://github.com/riywo/ndenv

import (
	"fmt"
	
	"os/exec"
)

func installNdenv() {
	// Método 1: Descargar y extraer .tar.gz
	ndenv_tar_url := "https://github.com/riywo/ndenv/archive/refs/tags/v0.4.0.tar.gz"
	ndenv_cmd_tar := exec.Command("curl", "-L", ndenv_tar_url, "-o", "package.tar.gz")
	err := ndenv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ndenv_zip_url := "https://github.com/riywo/ndenv/archive/refs/tags/v0.4.0.zip"
	ndenv_cmd_zip := exec.Command("curl", "-L", ndenv_zip_url, "-o", "package.zip")
	err = ndenv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ndenv_bin_url := "https://github.com/riywo/ndenv/archive/refs/tags/v0.4.0.bin"
	ndenv_cmd_bin := exec.Command("curl", "-L", ndenv_bin_url, "-o", "binary.bin")
	err = ndenv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ndenv_src_url := "https://github.com/riywo/ndenv/archive/refs/tags/v0.4.0.src.tar.gz"
	ndenv_cmd_src := exec.Command("curl", "-L", ndenv_src_url, "-o", "source.tar.gz")
	err = ndenv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ndenv_cmd_direct := exec.Command("./binary")
	err = ndenv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node-build")
exec.Command("latte", "install", "node-build")
}
