package main

// Dssim - RGBA Structural Similarity Rust implementation
// Homepage: https://github.com/kornelski/dssim

import (
	"fmt"
	
	"os/exec"
)

func installDssim() {
	// Método 1: Descargar y extraer .tar.gz
	dssim_tar_url := "https://github.com/kornelski/dssim/archive/refs/tags/3.3.1.tar.gz"
	dssim_cmd_tar := exec.Command("curl", "-L", dssim_tar_url, "-o", "package.tar.gz")
	err := dssim_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dssim_zip_url := "https://github.com/kornelski/dssim/archive/refs/tags/3.3.1.zip"
	dssim_cmd_zip := exec.Command("curl", "-L", dssim_zip_url, "-o", "package.zip")
	err = dssim_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dssim_bin_url := "https://github.com/kornelski/dssim/archive/refs/tags/3.3.1.bin"
	dssim_cmd_bin := exec.Command("curl", "-L", dssim_bin_url, "-o", "binary.bin")
	err = dssim_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dssim_src_url := "https://github.com/kornelski/dssim/archive/refs/tags/3.3.1.src.tar.gz"
	dssim_cmd_src := exec.Command("curl", "-L", dssim_src_url, "-o", "source.tar.gz")
	err = dssim_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dssim_cmd_direct := exec.Command("./binary")
	err = dssim_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: nasm")
exec.Command("latte", "install", "nasm")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
