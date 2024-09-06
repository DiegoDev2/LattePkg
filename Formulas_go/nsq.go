package main

// Nsq - Realtime distributed messaging platform
// Homepage: https://nsq.io/

import (
	"fmt"
	
	"os/exec"
)

func installNsq() {
	// Método 1: Descargar y extraer .tar.gz
	nsq_tar_url := "https://github.com/nsqio/nsq/archive/refs/tags/v1.3.0.tar.gz"
	nsq_cmd_tar := exec.Command("curl", "-L", nsq_tar_url, "-o", "package.tar.gz")
	err := nsq_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nsq_zip_url := "https://github.com/nsqio/nsq/archive/refs/tags/v1.3.0.zip"
	nsq_cmd_zip := exec.Command("curl", "-L", nsq_zip_url, "-o", "package.zip")
	err = nsq_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nsq_bin_url := "https://github.com/nsqio/nsq/archive/refs/tags/v1.3.0.bin"
	nsq_cmd_bin := exec.Command("curl", "-L", nsq_bin_url, "-o", "binary.bin")
	err = nsq_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nsq_src_url := "https://github.com/nsqio/nsq/archive/refs/tags/v1.3.0.src.tar.gz"
	nsq_cmd_src := exec.Command("curl", "-L", nsq_src_url, "-o", "source.tar.gz")
	err = nsq_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nsq_cmd_direct := exec.Command("./binary")
	err = nsq_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
