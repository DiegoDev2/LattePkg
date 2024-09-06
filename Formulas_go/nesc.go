package main

// Nesc - Programming language for deeply networked systems
// Homepage: https://github.com/tinyos/nesc

import (
	"fmt"
	
	"os/exec"
)

func installNesc() {
	// Método 1: Descargar y extraer .tar.gz
	nesc_tar_url := "https://github.com/tinyos/nesc/archive/refs/tags/v1.4.0.tar.gz"
	nesc_cmd_tar := exec.Command("curl", "-L", nesc_tar_url, "-o", "package.tar.gz")
	err := nesc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nesc_zip_url := "https://github.com/tinyos/nesc/archive/refs/tags/v1.4.0.zip"
	nesc_cmd_zip := exec.Command("curl", "-L", nesc_zip_url, "-o", "package.zip")
	err = nesc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nesc_bin_url := "https://github.com/tinyos/nesc/archive/refs/tags/v1.4.0.bin"
	nesc_cmd_bin := exec.Command("curl", "-L", nesc_bin_url, "-o", "binary.bin")
	err = nesc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nesc_src_url := "https://github.com/tinyos/nesc/archive/refs/tags/v1.4.0.src.tar.gz"
	nesc_cmd_src := exec.Command("curl", "-L", nesc_src_url, "-o", "source.tar.gz")
	err = nesc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nesc_cmd_direct := exec.Command("./binary")
	err = nesc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
	fmt.Println("Instalando dependencia: emacs")
exec.Command("latte", "install", "emacs")
}
