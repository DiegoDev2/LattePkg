package main

// Alure - Manage common tasks with OpenAL applications
// Homepage: https://kcat.tomasu.net/alure.html

import (
	"fmt"
	
	"os/exec"
)

func installAlure() {
	// Método 1: Descargar y extraer .tar.gz
	alure_tar_url := "https://kcat.tomasu.net/alure-releases/alure-1.2.tar.bz2"
	alure_cmd_tar := exec.Command("curl", "-L", alure_tar_url, "-o", "package.tar.gz")
	err := alure_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	alure_zip_url := "https://kcat.tomasu.net/alure-releases/alure-1.2.tar.bz2"
	alure_cmd_zip := exec.Command("curl", "-L", alure_zip_url, "-o", "package.zip")
	err = alure_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	alure_bin_url := "https://kcat.tomasu.net/alure-releases/alure-1.2.tar.bz2"
	alure_cmd_bin := exec.Command("curl", "-L", alure_bin_url, "-o", "binary.bin")
	err = alure_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	alure_src_url := "https://kcat.tomasu.net/alure-releases/alure-1.2.tar.bz2"
	alure_cmd_src := exec.Command("curl", "-L", alure_src_url, "-o", "source.tar.gz")
	err = alure_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	alure_cmd_direct := exec.Command("./binary")
	err = alure_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: openal-soft")
exec.Command("latte", "install", "openal-soft")
}
