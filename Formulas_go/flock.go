package main

// Flock - Lock file during command
// Homepage: https://github.com/discoteq/flock

import (
	"fmt"
	
	"os/exec"
)

func installFlock() {
	// Método 1: Descargar y extraer .tar.gz
	flock_tar_url := "https://github.com/discoteq/flock/releases/download/v0.4.0/flock-0.4.0.tar.xz"
	flock_cmd_tar := exec.Command("curl", "-L", flock_tar_url, "-o", "package.tar.gz")
	err := flock_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	flock_zip_url := "https://github.com/discoteq/flock/releases/download/v0.4.0/flock-0.4.0.tar.xz"
	flock_cmd_zip := exec.Command("curl", "-L", flock_zip_url, "-o", "package.zip")
	err = flock_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	flock_bin_url := "https://github.com/discoteq/flock/releases/download/v0.4.0/flock-0.4.0.tar.xz"
	flock_cmd_bin := exec.Command("curl", "-L", flock_bin_url, "-o", "binary.bin")
	err = flock_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	flock_src_url := "https://github.com/discoteq/flock/releases/download/v0.4.0/flock-0.4.0.tar.xz"
	flock_cmd_src := exec.Command("curl", "-L", flock_src_url, "-o", "source.tar.gz")
	err = flock_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	flock_cmd_direct := exec.Command("./binary")
	err = flock_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
