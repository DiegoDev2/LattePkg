package main

// Multitail - Tail multiple files in one terminal simultaneously
// Homepage: https://vanheusden.com/multitail/

import (
	"fmt"
	
	"os/exec"
)

func installMultitail() {
	// Método 1: Descargar y extraer .tar.gz
	multitail_tar_url := "https://github.com/folkertvanheusden/multitail/archive/refs/tags/7.1.5.tar.gz"
	multitail_cmd_tar := exec.Command("curl", "-L", multitail_tar_url, "-o", "package.tar.gz")
	err := multitail_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	multitail_zip_url := "https://github.com/folkertvanheusden/multitail/archive/refs/tags/7.1.5.zip"
	multitail_cmd_zip := exec.Command("curl", "-L", multitail_zip_url, "-o", "package.zip")
	err = multitail_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	multitail_bin_url := "https://github.com/folkertvanheusden/multitail/archive/refs/tags/7.1.5.bin"
	multitail_cmd_bin := exec.Command("curl", "-L", multitail_bin_url, "-o", "binary.bin")
	err = multitail_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	multitail_src_url := "https://github.com/folkertvanheusden/multitail/archive/refs/tags/7.1.5.src.tar.gz"
	multitail_cmd_src := exec.Command("curl", "-L", multitail_src_url, "-o", "source.tar.gz")
	err = multitail_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	multitail_cmd_direct := exec.Command("./binary")
	err = multitail_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: ncurses")
exec.Command("latte", "install", "ncurses")
}
