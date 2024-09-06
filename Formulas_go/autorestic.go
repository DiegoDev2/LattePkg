package main

// Autorestic - High level CLI utility for restic
// Homepage: https://autorestic.vercel.app/

import (
	"fmt"
	
	"os/exec"
)

func installAutorestic() {
	// Método 1: Descargar y extraer .tar.gz
	autorestic_tar_url := "https://github.com/cupcakearmy/autorestic/archive/refs/tags/v1.8.3.tar.gz"
	autorestic_cmd_tar := exec.Command("curl", "-L", autorestic_tar_url, "-o", "package.tar.gz")
	err := autorestic_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	autorestic_zip_url := "https://github.com/cupcakearmy/autorestic/archive/refs/tags/v1.8.3.zip"
	autorestic_cmd_zip := exec.Command("curl", "-L", autorestic_zip_url, "-o", "package.zip")
	err = autorestic_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	autorestic_bin_url := "https://github.com/cupcakearmy/autorestic/archive/refs/tags/v1.8.3.bin"
	autorestic_cmd_bin := exec.Command("curl", "-L", autorestic_bin_url, "-o", "binary.bin")
	err = autorestic_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	autorestic_src_url := "https://github.com/cupcakearmy/autorestic/archive/refs/tags/v1.8.3.src.tar.gz"
	autorestic_cmd_src := exec.Command("curl", "-L", autorestic_src_url, "-o", "source.tar.gz")
	err = autorestic_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	autorestic_cmd_direct := exec.Command("./binary")
	err = autorestic_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: restic")
exec.Command("latte", "install", "restic")
}
