package main

// Nethack - Single-player roguelike video game
// Homepage: https://www.nethack.org/

import (
	"fmt"
	
	"os/exec"
)

func installNethack() {
	// Método 1: Descargar y extraer .tar.gz
	nethack_tar_url := "https://www.nethack.org/download/3.6.7/nethack-367-src.tgz"
	nethack_cmd_tar := exec.Command("curl", "-L", nethack_tar_url, "-o", "package.tar.gz")
	err := nethack_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nethack_zip_url := "https://www.nethack.org/download/3.6.7/nethack-367-src.tgz"
	nethack_cmd_zip := exec.Command("curl", "-L", nethack_zip_url, "-o", "package.zip")
	err = nethack_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nethack_bin_url := "https://www.nethack.org/download/3.6.7/nethack-367-src.tgz"
	nethack_cmd_bin := exec.Command("curl", "-L", nethack_bin_url, "-o", "binary.bin")
	err = nethack_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nethack_src_url := "https://www.nethack.org/download/3.6.7/nethack-367-src.tgz"
	nethack_cmd_src := exec.Command("curl", "-L", nethack_src_url, "-o", "source.tar.gz")
	err = nethack_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nethack_cmd_direct := exec.Command("./binary")
	err = nethack_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
