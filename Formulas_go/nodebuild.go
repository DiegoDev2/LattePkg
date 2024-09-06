package main

// NodeBuild - Install NodeJS versions
// Homepage: https://github.com/nodenv/node-build

import (
	"fmt"
	
	"os/exec"
)

func installNodeBuild() {
	// Método 1: Descargar y extraer .tar.gz
	nodebuild_tar_url := "https://github.com/nodenv/node-build/archive/refs/tags/v5.3.11.tar.gz"
	nodebuild_cmd_tar := exec.Command("curl", "-L", nodebuild_tar_url, "-o", "package.tar.gz")
	err := nodebuild_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nodebuild_zip_url := "https://github.com/nodenv/node-build/archive/refs/tags/v5.3.11.zip"
	nodebuild_cmd_zip := exec.Command("curl", "-L", nodebuild_zip_url, "-o", "package.zip")
	err = nodebuild_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nodebuild_bin_url := "https://github.com/nodenv/node-build/archive/refs/tags/v5.3.11.bin"
	nodebuild_cmd_bin := exec.Command("curl", "-L", nodebuild_bin_url, "-o", "binary.bin")
	err = nodebuild_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nodebuild_src_url := "https://github.com/nodenv/node-build/archive/refs/tags/v5.3.11.src.tar.gz"
	nodebuild_cmd_src := exec.Command("curl", "-L", nodebuild_src_url, "-o", "source.tar.gz")
	err = nodebuild_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nodebuild_cmd_direct := exec.Command("./binary")
	err = nodebuild_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
