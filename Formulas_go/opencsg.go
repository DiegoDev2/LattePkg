package main

// Opencsg - Constructive solid geometry rendering library
// Homepage: https://www.opencsg.org/

import (
	"fmt"
	
	"os/exec"
)

func installOpencsg() {
	// Método 1: Descargar y extraer .tar.gz
	opencsg_tar_url := "https://www.opencsg.org/OpenCSG-1.6.0.tar.gz"
	opencsg_cmd_tar := exec.Command("curl", "-L", opencsg_tar_url, "-o", "package.tar.gz")
	err := opencsg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	opencsg_zip_url := "https://www.opencsg.org/OpenCSG-1.6.0.zip"
	opencsg_cmd_zip := exec.Command("curl", "-L", opencsg_zip_url, "-o", "package.zip")
	err = opencsg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	opencsg_bin_url := "https://www.opencsg.org/OpenCSG-1.6.0.bin"
	opencsg_cmd_bin := exec.Command("curl", "-L", opencsg_bin_url, "-o", "binary.bin")
	err = opencsg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	opencsg_src_url := "https://www.opencsg.org/OpenCSG-1.6.0.src.tar.gz"
	opencsg_cmd_src := exec.Command("curl", "-L", opencsg_src_url, "-o", "source.tar.gz")
	err = opencsg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	opencsg_cmd_direct := exec.Command("./binary")
	err = opencsg_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: qt")
exec.Command("latte", "install", "qt")
	fmt.Println("Instalando dependencia: glew")
exec.Command("latte", "install", "glew")
}
