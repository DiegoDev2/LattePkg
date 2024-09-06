package main

// Demumble - More powerful symbol demangler (a la c++filt)
// Homepage: https://github.com/nico/demumble

import (
	"fmt"
	
	"os/exec"
)

func installDemumble() {
	// Método 1: Descargar y extraer .tar.gz
	demumble_tar_url := "https://github.com/nico/demumble/archive/refs/tags/v1.3.0.tar.gz"
	demumble_cmd_tar := exec.Command("curl", "-L", demumble_tar_url, "-o", "package.tar.gz")
	err := demumble_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	demumble_zip_url := "https://github.com/nico/demumble/archive/refs/tags/v1.3.0.zip"
	demumble_cmd_zip := exec.Command("curl", "-L", demumble_zip_url, "-o", "package.zip")
	err = demumble_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	demumble_bin_url := "https://github.com/nico/demumble/archive/refs/tags/v1.3.0.bin"
	demumble_cmd_bin := exec.Command("curl", "-L", demumble_bin_url, "-o", "binary.bin")
	err = demumble_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	demumble_src_url := "https://github.com/nico/demumble/archive/refs/tags/v1.3.0.src.tar.gz"
	demumble_cmd_src := exec.Command("curl", "-L", demumble_src_url, "-o", "source.tar.gz")
	err = demumble_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	demumble_cmd_direct := exec.Command("./binary")
	err = demumble_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
