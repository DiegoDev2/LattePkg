package main

// Geometry - Minimal, fully customizable and composable zsh prompt theme
// Homepage: https://github.com/geometry-zsh/geometry

import (
	"fmt"
	
	"os/exec"
)

func installGeometry() {
	// Método 1: Descargar y extraer .tar.gz
	geometry_tar_url := "https://github.com/geometry-zsh/geometry/archive/refs/tags/v2.2.0.tar.gz"
	geometry_cmd_tar := exec.Command("curl", "-L", geometry_tar_url, "-o", "package.tar.gz")
	err := geometry_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	geometry_zip_url := "https://github.com/geometry-zsh/geometry/archive/refs/tags/v2.2.0.zip"
	geometry_cmd_zip := exec.Command("curl", "-L", geometry_zip_url, "-o", "package.zip")
	err = geometry_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	geometry_bin_url := "https://github.com/geometry-zsh/geometry/archive/refs/tags/v2.2.0.bin"
	geometry_cmd_bin := exec.Command("curl", "-L", geometry_bin_url, "-o", "binary.bin")
	err = geometry_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	geometry_src_url := "https://github.com/geometry-zsh/geometry/archive/refs/tags/v2.2.0.src.tar.gz"
	geometry_cmd_src := exec.Command("curl", "-L", geometry_src_url, "-o", "source.tar.gz")
	err = geometry_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	geometry_cmd_direct := exec.Command("./binary")
	err = geometry_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: zsh-async")
exec.Command("latte", "install", "zsh-async")
}
