package main

// Fastjar - Implementation of Sun's jar tool
// Homepage: https://savannah.nongnu.org/projects/fastjar

import (
	"fmt"
	
	"os/exec"
)

func installFastjar() {
	// Método 1: Descargar y extraer .tar.gz
	fastjar_tar_url := "https://download.savannah.nongnu.org/releases/fastjar/fastjar-0.98.tar.gz"
	fastjar_cmd_tar := exec.Command("curl", "-L", fastjar_tar_url, "-o", "package.tar.gz")
	err := fastjar_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fastjar_zip_url := "https://download.savannah.nongnu.org/releases/fastjar/fastjar-0.98.zip"
	fastjar_cmd_zip := exec.Command("curl", "-L", fastjar_zip_url, "-o", "package.zip")
	err = fastjar_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fastjar_bin_url := "https://download.savannah.nongnu.org/releases/fastjar/fastjar-0.98.bin"
	fastjar_cmd_bin := exec.Command("curl", "-L", fastjar_bin_url, "-o", "binary.bin")
	err = fastjar_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fastjar_src_url := "https://download.savannah.nongnu.org/releases/fastjar/fastjar-0.98.src.tar.gz"
	fastjar_cmd_src := exec.Command("curl", "-L", fastjar_src_url, "-o", "source.tar.gz")
	err = fastjar_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fastjar_cmd_direct := exec.Command("./binary")
	err = fastjar_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
