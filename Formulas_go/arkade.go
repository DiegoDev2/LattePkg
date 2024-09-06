package main

// Arkade - Open Source Kubernetes Marketplace
// Homepage: https://blog.alexellis.io/kubernetes-marketplace-two-year-update/

import (
	"fmt"
	
	"os/exec"
)

func installArkade() {
	// Método 1: Descargar y extraer .tar.gz
	arkade_tar_url := "https://github.com/alexellis/arkade/archive/refs/tags/0.11.23.tar.gz"
	arkade_cmd_tar := exec.Command("curl", "-L", arkade_tar_url, "-o", "package.tar.gz")
	err := arkade_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	arkade_zip_url := "https://github.com/alexellis/arkade/archive/refs/tags/0.11.23.zip"
	arkade_cmd_zip := exec.Command("curl", "-L", arkade_zip_url, "-o", "package.zip")
	err = arkade_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	arkade_bin_url := "https://github.com/alexellis/arkade/archive/refs/tags/0.11.23.bin"
	arkade_cmd_bin := exec.Command("curl", "-L", arkade_bin_url, "-o", "binary.bin")
	err = arkade_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	arkade_src_url := "https://github.com/alexellis/arkade/archive/refs/tags/0.11.23.src.tar.gz"
	arkade_cmd_src := exec.Command("curl", "-L", arkade_src_url, "-o", "source.tar.gz")
	err = arkade_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	arkade_cmd_direct := exec.Command("./binary")
	err = arkade_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
