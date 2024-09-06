package main

// Bee - Tool for managing database changes
// Homepage: https://github.com/bluesoft/bee

import (
	"fmt"
	
	"os/exec"
)

func installBee() {
	// Método 1: Descargar y extraer .tar.gz
	bee_tar_url := "https://github.com/bluesoft/bee/releases/download/1.103/bee-1.103.zip"
	bee_cmd_tar := exec.Command("curl", "-L", bee_tar_url, "-o", "package.tar.gz")
	err := bee_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bee_zip_url := "https://github.com/bluesoft/bee/releases/download/1.103/bee-1.103.zip"
	bee_cmd_zip := exec.Command("curl", "-L", bee_zip_url, "-o", "package.zip")
	err = bee_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bee_bin_url := "https://github.com/bluesoft/bee/releases/download/1.103/bee-1.103.zip"
	bee_cmd_bin := exec.Command("curl", "-L", bee_bin_url, "-o", "binary.bin")
	err = bee_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bee_src_url := "https://github.com/bluesoft/bee/releases/download/1.103/bee-1.103.zip"
	bee_cmd_src := exec.Command("curl", "-L", bee_src_url, "-o", "source.tar.gz")
	err = bee_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bee_cmd_direct := exec.Command("./binary")
	err = bee_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
