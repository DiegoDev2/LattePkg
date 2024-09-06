package main

// Dmagnetic - Magnetic Scrolls Interpreter
// Homepage: https://www.dettus.net/dMagnetic/

import (
	"fmt"
	
	"os/exec"
)

func installDmagnetic() {
	// Método 1: Descargar y extraer .tar.gz
	dmagnetic_tar_url := "https://www.dettus.net/dMagnetic/dMagnetic_0.37.tar.bz2"
	dmagnetic_cmd_tar := exec.Command("curl", "-L", dmagnetic_tar_url, "-o", "package.tar.gz")
	err := dmagnetic_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dmagnetic_zip_url := "https://www.dettus.net/dMagnetic/dMagnetic_0.37.tar.bz2"
	dmagnetic_cmd_zip := exec.Command("curl", "-L", dmagnetic_zip_url, "-o", "package.zip")
	err = dmagnetic_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dmagnetic_bin_url := "https://www.dettus.net/dMagnetic/dMagnetic_0.37.tar.bz2"
	dmagnetic_cmd_bin := exec.Command("curl", "-L", dmagnetic_bin_url, "-o", "binary.bin")
	err = dmagnetic_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dmagnetic_src_url := "https://www.dettus.net/dMagnetic/dMagnetic_0.37.tar.bz2"
	dmagnetic_cmd_src := exec.Command("curl", "-L", dmagnetic_src_url, "-o", "source.tar.gz")
	err = dmagnetic_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dmagnetic_cmd_direct := exec.Command("./binary")
	err = dmagnetic_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
