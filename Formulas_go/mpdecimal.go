package main

// Mpdecimal - Library for decimal floating point arithmetic
// Homepage: https://www.bytereef.org/mpdecimal/

import (
	"fmt"
	
	"os/exec"
)

func installMpdecimal() {
	// Método 1: Descargar y extraer .tar.gz
	mpdecimal_tar_url := "https://www.bytereef.org/software/mpdecimal/releases/mpdecimal-4.0.0.tar.gz"
	mpdecimal_cmd_tar := exec.Command("curl", "-L", mpdecimal_tar_url, "-o", "package.tar.gz")
	err := mpdecimal_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mpdecimal_zip_url := "https://www.bytereef.org/software/mpdecimal/releases/mpdecimal-4.0.0.zip"
	mpdecimal_cmd_zip := exec.Command("curl", "-L", mpdecimal_zip_url, "-o", "package.zip")
	err = mpdecimal_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mpdecimal_bin_url := "https://www.bytereef.org/software/mpdecimal/releases/mpdecimal-4.0.0.bin"
	mpdecimal_cmd_bin := exec.Command("curl", "-L", mpdecimal_bin_url, "-o", "binary.bin")
	err = mpdecimal_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mpdecimal_src_url := "https://www.bytereef.org/software/mpdecimal/releases/mpdecimal-4.0.0.src.tar.gz"
	mpdecimal_cmd_src := exec.Command("curl", "-L", mpdecimal_src_url, "-o", "source.tar.gz")
	err = mpdecimal_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mpdecimal_cmd_direct := exec.Command("./binary")
	err = mpdecimal_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
