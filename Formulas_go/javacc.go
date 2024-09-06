package main

// Javacc - Parser generator for use with Java applications
// Homepage: https://javacc.github.io/javacc/

import (
	"fmt"
	
	"os/exec"
)

func installJavacc() {
	// Método 1: Descargar y extraer .tar.gz
	javacc_tar_url := "https://github.com/javacc/javacc/archive/refs/tags/javacc-7.0.13.tar.gz"
	javacc_cmd_tar := exec.Command("curl", "-L", javacc_tar_url, "-o", "package.tar.gz")
	err := javacc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	javacc_zip_url := "https://github.com/javacc/javacc/archive/refs/tags/javacc-7.0.13.zip"
	javacc_cmd_zip := exec.Command("curl", "-L", javacc_zip_url, "-o", "package.zip")
	err = javacc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	javacc_bin_url := "https://github.com/javacc/javacc/archive/refs/tags/javacc-7.0.13.bin"
	javacc_cmd_bin := exec.Command("curl", "-L", javacc_bin_url, "-o", "binary.bin")
	err = javacc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	javacc_src_url := "https://github.com/javacc/javacc/archive/refs/tags/javacc-7.0.13.src.tar.gz"
	javacc_cmd_src := exec.Command("curl", "-L", javacc_src_url, "-o", "source.tar.gz")
	err = javacc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	javacc_cmd_direct := exec.Command("./binary")
	err = javacc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ant")
exec.Command("latte", "install", "ant")
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
