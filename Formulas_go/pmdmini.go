package main

// Pmdmini - Plays music in PC-88/98 PMD chiptune format
// Homepage: https://github.com/mistydemeo/pmdmini

import (
	"fmt"
	
	"os/exec"
)

func installPmdmini() {
	// Método 1: Descargar y extraer .tar.gz
	pmdmini_tar_url := "https://github.com/mistydemeo/pmdmini/archive/refs/tags/v2.0.0.tar.gz"
	pmdmini_cmd_tar := exec.Command("curl", "-L", pmdmini_tar_url, "-o", "package.tar.gz")
	err := pmdmini_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pmdmini_zip_url := "https://github.com/mistydemeo/pmdmini/archive/refs/tags/v2.0.0.zip"
	pmdmini_cmd_zip := exec.Command("curl", "-L", pmdmini_zip_url, "-o", "package.zip")
	err = pmdmini_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pmdmini_bin_url := "https://github.com/mistydemeo/pmdmini/archive/refs/tags/v2.0.0.bin"
	pmdmini_cmd_bin := exec.Command("curl", "-L", pmdmini_bin_url, "-o", "binary.bin")
	err = pmdmini_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pmdmini_src_url := "https://github.com/mistydemeo/pmdmini/archive/refs/tags/v2.0.0.src.tar.gz"
	pmdmini_cmd_src := exec.Command("curl", "-L", pmdmini_src_url, "-o", "source.tar.gz")
	err = pmdmini_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pmdmini_cmd_direct := exec.Command("./binary")
	err = pmdmini_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: sdl2")
exec.Command("latte", "install", "sdl2")
}
