package main

// Cliclick - Tool for emulating mouse and keyboard events
// Homepage: https://www.bluem.net/jump/cliclick/

import (
	"fmt"
	
	"os/exec"
)

func installCliclick() {
	// Método 1: Descargar y extraer .tar.gz
	cliclick_tar_url := "https://github.com/BlueM/cliclick/archive/refs/tags/5.1.tar.gz"
	cliclick_cmd_tar := exec.Command("curl", "-L", cliclick_tar_url, "-o", "package.tar.gz")
	err := cliclick_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cliclick_zip_url := "https://github.com/BlueM/cliclick/archive/refs/tags/5.1.zip"
	cliclick_cmd_zip := exec.Command("curl", "-L", cliclick_zip_url, "-o", "package.zip")
	err = cliclick_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cliclick_bin_url := "https://github.com/BlueM/cliclick/archive/refs/tags/5.1.bin"
	cliclick_cmd_bin := exec.Command("curl", "-L", cliclick_bin_url, "-o", "binary.bin")
	err = cliclick_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cliclick_src_url := "https://github.com/BlueM/cliclick/archive/refs/tags/5.1.src.tar.gz"
	cliclick_cmd_src := exec.Command("curl", "-L", cliclick_src_url, "-o", "source.tar.gz")
	err = cliclick_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cliclick_cmd_direct := exec.Command("./binary")
	err = cliclick_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
