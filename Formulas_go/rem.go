package main

// Rem - Command-line tool to access OSX Reminders.app database
// Homepage: https://github.com/kykim/rem

import (
	"fmt"
	
	"os/exec"
)

func installRem() {
	// Método 1: Descargar y extraer .tar.gz
	rem_tar_url := "https://github.com/kykim/rem/archive/refs/tags/20150618.tar.gz"
	rem_cmd_tar := exec.Command("curl", "-L", rem_tar_url, "-o", "package.tar.gz")
	err := rem_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rem_zip_url := "https://github.com/kykim/rem/archive/refs/tags/20150618.zip"
	rem_cmd_zip := exec.Command("curl", "-L", rem_zip_url, "-o", "package.zip")
	err = rem_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rem_bin_url := "https://github.com/kykim/rem/archive/refs/tags/20150618.bin"
	rem_cmd_bin := exec.Command("curl", "-L", rem_bin_url, "-o", "binary.bin")
	err = rem_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rem_src_url := "https://github.com/kykim/rem/archive/refs/tags/20150618.src.tar.gz"
	rem_cmd_src := exec.Command("curl", "-L", rem_src_url, "-o", "source.tar.gz")
	err = rem_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rem_cmd_direct := exec.Command("./binary")
	err = rem_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
