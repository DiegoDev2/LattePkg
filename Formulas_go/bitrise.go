package main

// Bitrise - Command-line automation tool
// Homepage: https://github.com/bitrise-io/bitrise

import (
	"fmt"
	
	"os/exec"
)

func installBitrise() {
	// Método 1: Descargar y extraer .tar.gz
	bitrise_tar_url := "https://github.com/bitrise-io/bitrise/archive/refs/tags/2.20.1.tar.gz"
	bitrise_cmd_tar := exec.Command("curl", "-L", bitrise_tar_url, "-o", "package.tar.gz")
	err := bitrise_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bitrise_zip_url := "https://github.com/bitrise-io/bitrise/archive/refs/tags/2.20.1.zip"
	bitrise_cmd_zip := exec.Command("curl", "-L", bitrise_zip_url, "-o", "package.zip")
	err = bitrise_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bitrise_bin_url := "https://github.com/bitrise-io/bitrise/archive/refs/tags/2.20.1.bin"
	bitrise_cmd_bin := exec.Command("curl", "-L", bitrise_bin_url, "-o", "binary.bin")
	err = bitrise_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bitrise_src_url := "https://github.com/bitrise-io/bitrise/archive/refs/tags/2.20.1.src.tar.gz"
	bitrise_cmd_src := exec.Command("curl", "-L", bitrise_src_url, "-o", "source.tar.gz")
	err = bitrise_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bitrise_cmd_direct := exec.Command("./binary")
	err = bitrise_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
