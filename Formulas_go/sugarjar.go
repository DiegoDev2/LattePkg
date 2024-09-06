package main

// Sugarjar - Helper utility for a better Git/GitHub experience
// Homepage: https://github.com/jaymzh/sugarjar/

import (
	"fmt"
	
	"os/exec"
)

func installSugarjar() {
	// Método 1: Descargar y extraer .tar.gz
	sugarjar_tar_url := "https://github.com/jaymzh/sugarjar/archive/refs/tags/v1.1.2.tar.gz"
	sugarjar_cmd_tar := exec.Command("curl", "-L", sugarjar_tar_url, "-o", "package.tar.gz")
	err := sugarjar_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sugarjar_zip_url := "https://github.com/jaymzh/sugarjar/archive/refs/tags/v1.1.2.zip"
	sugarjar_cmd_zip := exec.Command("curl", "-L", sugarjar_zip_url, "-o", "package.zip")
	err = sugarjar_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sugarjar_bin_url := "https://github.com/jaymzh/sugarjar/archive/refs/tags/v1.1.2.bin"
	sugarjar_cmd_bin := exec.Command("curl", "-L", sugarjar_bin_url, "-o", "binary.bin")
	err = sugarjar_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sugarjar_src_url := "https://github.com/jaymzh/sugarjar/archive/refs/tags/v1.1.2.src.tar.gz"
	sugarjar_cmd_src := exec.Command("curl", "-L", sugarjar_src_url, "-o", "source.tar.gz")
	err = sugarjar_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sugarjar_cmd_direct := exec.Command("./binary")
	err = sugarjar_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gh")
exec.Command("latte", "install", "gh")
	fmt.Println("Instalando dependencia: ruby")
exec.Command("latte", "install", "ruby")
}
