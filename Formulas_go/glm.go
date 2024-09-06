package main

// Glm - C++ mathematics library for graphics software
// Homepage: https://glm.g-truc.net/

import (
	"fmt"
	
	"os/exec"
)

func installGlm() {
	// Método 1: Descargar y extraer .tar.gz
	glm_tar_url := "https://github.com/g-truc/glm/archive/refs/tags/1.0.1.tar.gz"
	glm_cmd_tar := exec.Command("curl", "-L", glm_tar_url, "-o", "package.tar.gz")
	err := glm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	glm_zip_url := "https://github.com/g-truc/glm/archive/refs/tags/1.0.1.zip"
	glm_cmd_zip := exec.Command("curl", "-L", glm_zip_url, "-o", "package.zip")
	err = glm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	glm_bin_url := "https://github.com/g-truc/glm/archive/refs/tags/1.0.1.bin"
	glm_cmd_bin := exec.Command("curl", "-L", glm_bin_url, "-o", "binary.bin")
	err = glm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	glm_src_url := "https://github.com/g-truc/glm/archive/refs/tags/1.0.1.src.tar.gz"
	glm_cmd_src := exec.Command("curl", "-L", glm_src_url, "-o", "source.tar.gz")
	err = glm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	glm_cmd_direct := exec.Command("./binary")
	err = glm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: doxygen")
exec.Command("latte", "install", "doxygen")
}
