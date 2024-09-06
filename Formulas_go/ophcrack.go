package main

// Ophcrack - Microsoft Windows password cracker using rainbow tables
// Homepage: https://ophcrack.sourceforge.io/

import (
	"fmt"
	
	"os/exec"
)

func installOphcrack() {
	// Método 1: Descargar y extraer .tar.gz
	ophcrack_tar_url := "https://downloads.sourceforge.net/project/ophcrack/ophcrack/3.8.0/ophcrack-3.8.0.tar.bz2"
	ophcrack_cmd_tar := exec.Command("curl", "-L", ophcrack_tar_url, "-o", "package.tar.gz")
	err := ophcrack_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ophcrack_zip_url := "https://downloads.sourceforge.net/project/ophcrack/ophcrack/3.8.0/ophcrack-3.8.0.tar.bz2"
	ophcrack_cmd_zip := exec.Command("curl", "-L", ophcrack_zip_url, "-o", "package.zip")
	err = ophcrack_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ophcrack_bin_url := "https://downloads.sourceforge.net/project/ophcrack/ophcrack/3.8.0/ophcrack-3.8.0.tar.bz2"
	ophcrack_cmd_bin := exec.Command("curl", "-L", ophcrack_bin_url, "-o", "binary.bin")
	err = ophcrack_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ophcrack_src_url := "https://downloads.sourceforge.net/project/ophcrack/ophcrack/3.8.0/ophcrack-3.8.0.tar.bz2"
	ophcrack_cmd_src := exec.Command("curl", "-L", ophcrack_src_url, "-o", "source.tar.gz")
	err = ophcrack_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ophcrack_cmd_direct := exec.Command("./binary")
	err = ophcrack_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
