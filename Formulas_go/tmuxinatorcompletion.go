package main

// TmuxinatorCompletion - Shell completion for Tmuxinator
// Homepage: https://github.com/tmuxinator/tmuxinator

import (
	"fmt"
	
	"os/exec"
)

func installTmuxinatorCompletion() {
	// Método 1: Descargar y extraer .tar.gz
	tmuxinatorcompletion_tar_url := "https://github.com/tmuxinator/tmuxinator/archive/refs/tags/v3.3.0.tar.gz"
	tmuxinatorcompletion_cmd_tar := exec.Command("curl", "-L", tmuxinatorcompletion_tar_url, "-o", "package.tar.gz")
	err := tmuxinatorcompletion_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tmuxinatorcompletion_zip_url := "https://github.com/tmuxinator/tmuxinator/archive/refs/tags/v3.3.0.zip"
	tmuxinatorcompletion_cmd_zip := exec.Command("curl", "-L", tmuxinatorcompletion_zip_url, "-o", "package.zip")
	err = tmuxinatorcompletion_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tmuxinatorcompletion_bin_url := "https://github.com/tmuxinator/tmuxinator/archive/refs/tags/v3.3.0.bin"
	tmuxinatorcompletion_cmd_bin := exec.Command("curl", "-L", tmuxinatorcompletion_bin_url, "-o", "binary.bin")
	err = tmuxinatorcompletion_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tmuxinatorcompletion_src_url := "https://github.com/tmuxinator/tmuxinator/archive/refs/tags/v3.3.0.src.tar.gz"
	tmuxinatorcompletion_cmd_src := exec.Command("curl", "-L", tmuxinatorcompletion_src_url, "-o", "source.tar.gz")
	err = tmuxinatorcompletion_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tmuxinatorcompletion_cmd_direct := exec.Command("./binary")
	err = tmuxinatorcompletion_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
