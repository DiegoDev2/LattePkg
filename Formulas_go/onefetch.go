package main

// Onefetch - Command-line Git information tool
// Homepage: https://onefetch.dev/

import (
	"fmt"
	
	"os/exec"
)

func installOnefetch() {
	// Método 1: Descargar y extraer .tar.gz
	onefetch_tar_url := "https://github.com/o2sh/onefetch/archive/refs/tags/2.21.0.tar.gz"
	onefetch_cmd_tar := exec.Command("curl", "-L", onefetch_tar_url, "-o", "package.tar.gz")
	err := onefetch_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	onefetch_zip_url := "https://github.com/o2sh/onefetch/archive/refs/tags/2.21.0.zip"
	onefetch_cmd_zip := exec.Command("curl", "-L", onefetch_zip_url, "-o", "package.zip")
	err = onefetch_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	onefetch_bin_url := "https://github.com/o2sh/onefetch/archive/refs/tags/2.21.0.bin"
	onefetch_cmd_bin := exec.Command("curl", "-L", onefetch_bin_url, "-o", "binary.bin")
	err = onefetch_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	onefetch_src_url := "https://github.com/o2sh/onefetch/archive/refs/tags/2.21.0.src.tar.gz"
	onefetch_cmd_src := exec.Command("curl", "-L", onefetch_src_url, "-o", "source.tar.gz")
	err = onefetch_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	onefetch_cmd_direct := exec.Command("./binary")
	err = onefetch_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
}
