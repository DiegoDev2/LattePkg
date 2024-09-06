package main

// Sigi - Organizing tool for terminal lovers that hate organizing
// Homepage: https://sigi-cli.org

import (
	"fmt"
	
	"os/exec"
)

func installSigi() {
	// Método 1: Descargar y extraer .tar.gz
	sigi_tar_url := "https://github.com/sigi-cli/sigi/archive/refs/tags/v3.7.1.tar.gz"
	sigi_cmd_tar := exec.Command("curl", "-L", sigi_tar_url, "-o", "package.tar.gz")
	err := sigi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sigi_zip_url := "https://github.com/sigi-cli/sigi/archive/refs/tags/v3.7.1.zip"
	sigi_cmd_zip := exec.Command("curl", "-L", sigi_zip_url, "-o", "package.zip")
	err = sigi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sigi_bin_url := "https://github.com/sigi-cli/sigi/archive/refs/tags/v3.7.1.bin"
	sigi_cmd_bin := exec.Command("curl", "-L", sigi_bin_url, "-o", "binary.bin")
	err = sigi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sigi_src_url := "https://github.com/sigi-cli/sigi/archive/refs/tags/v3.7.1.src.tar.gz"
	sigi_cmd_src := exec.Command("curl", "-L", sigi_src_url, "-o", "source.tar.gz")
	err = sigi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sigi_cmd_direct := exec.Command("./binary")
	err = sigi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
