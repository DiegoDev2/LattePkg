package main

// CssCrush - Extensible PHP based CSS preprocessor
// Homepage: https://the-echoplex.net/csscrush

import (
	"fmt"
	
	"os/exec"
)

func installCssCrush() {
	// Método 1: Descargar y extraer .tar.gz
	csscrush_tar_url := "https://github.com/peteboere/css-crush/archive/refs/tags/v4.1.3.tar.gz"
	csscrush_cmd_tar := exec.Command("curl", "-L", csscrush_tar_url, "-o", "package.tar.gz")
	err := csscrush_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	csscrush_zip_url := "https://github.com/peteboere/css-crush/archive/refs/tags/v4.1.3.zip"
	csscrush_cmd_zip := exec.Command("curl", "-L", csscrush_zip_url, "-o", "package.zip")
	err = csscrush_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	csscrush_bin_url := "https://github.com/peteboere/css-crush/archive/refs/tags/v4.1.3.bin"
	csscrush_cmd_bin := exec.Command("curl", "-L", csscrush_bin_url, "-o", "binary.bin")
	err = csscrush_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	csscrush_src_url := "https://github.com/peteboere/css-crush/archive/refs/tags/v4.1.3.src.tar.gz"
	csscrush_cmd_src := exec.Command("curl", "-L", csscrush_src_url, "-o", "source.tar.gz")
	err = csscrush_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	csscrush_cmd_direct := exec.Command("./binary")
	err = csscrush_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: php")
exec.Command("latte", "install", "php")
}
