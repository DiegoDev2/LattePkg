package main

// Pktanon - Packet trace anonymization
// Homepage: https://www.tm.kit.edu/software/pktanon/index.html

import (
	"fmt"
	
	"os/exec"
)

func installPktanon() {
	// Método 1: Descargar y extraer .tar.gz
	pktanon_tar_url := "https://www.tm.kit.edu/software/pktanon/download/pktanon-1.4.0-dev.tar.gz"
	pktanon_cmd_tar := exec.Command("curl", "-L", pktanon_tar_url, "-o", "package.tar.gz")
	err := pktanon_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pktanon_zip_url := "https://www.tm.kit.edu/software/pktanon/download/pktanon-1.4.0-dev.zip"
	pktanon_cmd_zip := exec.Command("curl", "-L", pktanon_zip_url, "-o", "package.zip")
	err = pktanon_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pktanon_bin_url := "https://www.tm.kit.edu/software/pktanon/download/pktanon-1.4.0-dev.bin"
	pktanon_cmd_bin := exec.Command("curl", "-L", pktanon_bin_url, "-o", "binary.bin")
	err = pktanon_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pktanon_src_url := "https://www.tm.kit.edu/software/pktanon/download/pktanon-1.4.0-dev.src.tar.gz"
	pktanon_cmd_src := exec.Command("curl", "-L", pktanon_src_url, "-o", "source.tar.gz")
	err = pktanon_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pktanon_cmd_direct := exec.Command("./binary")
	err = pktanon_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: xerces-c")
exec.Command("latte", "install", "xerces-c")
}
