package main

// Valkey - High-performance data structure server that primarily serves key/value workloads
// Homepage: https://valkey.io

import (
	"fmt"
	
	"os/exec"
)

func installValkey() {
	// Método 1: Descargar y extraer .tar.gz
	valkey_tar_url := "https://github.com/valkey-io/valkey/archive/refs/tags/7.2.6.tar.gz"
	valkey_cmd_tar := exec.Command("curl", "-L", valkey_tar_url, "-o", "package.tar.gz")
	err := valkey_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	valkey_zip_url := "https://github.com/valkey-io/valkey/archive/refs/tags/7.2.6.zip"
	valkey_cmd_zip := exec.Command("curl", "-L", valkey_zip_url, "-o", "package.zip")
	err = valkey_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	valkey_bin_url := "https://github.com/valkey-io/valkey/archive/refs/tags/7.2.6.bin"
	valkey_cmd_bin := exec.Command("curl", "-L", valkey_bin_url, "-o", "binary.bin")
	err = valkey_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	valkey_src_url := "https://github.com/valkey-io/valkey/archive/refs/tags/7.2.6.src.tar.gz"
	valkey_cmd_src := exec.Command("curl", "-L", valkey_src_url, "-o", "source.tar.gz")
	err = valkey_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	valkey_cmd_direct := exec.Command("./binary")
	err = valkey_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
