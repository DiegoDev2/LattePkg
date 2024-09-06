package main

// Libobjc2 - Objective-C runtime library intended for use with Clang
// Homepage: https://github.com/gnustep/libobjc2

import (
	"fmt"
	
	"os/exec"
)

func installLibobjc2() {
	// Método 1: Descargar y extraer .tar.gz
	libobjc2_tar_url := "https://github.com/gnustep/libobjc2/archive/refs/tags/v2.2.1.tar.gz"
	libobjc2_cmd_tar := exec.Command("curl", "-L", libobjc2_tar_url, "-o", "package.tar.gz")
	err := libobjc2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libobjc2_zip_url := "https://github.com/gnustep/libobjc2/archive/refs/tags/v2.2.1.zip"
	libobjc2_cmd_zip := exec.Command("curl", "-L", libobjc2_zip_url, "-o", "package.zip")
	err = libobjc2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libobjc2_bin_url := "https://github.com/gnustep/libobjc2/archive/refs/tags/v2.2.1.bin"
	libobjc2_cmd_bin := exec.Command("curl", "-L", libobjc2_bin_url, "-o", "binary.bin")
	err = libobjc2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libobjc2_src_url := "https://github.com/gnustep/libobjc2/archive/refs/tags/v2.2.1.src.tar.gz"
	libobjc2_cmd_src := exec.Command("curl", "-L", libobjc2_src_url, "-o", "source.tar.gz")
	err = libobjc2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libobjc2_cmd_direct := exec.Command("./binary")
	err = libobjc2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: llvm")
exec.Command("latte", "install", "llvm")
	fmt.Println("Instalando dependencia: robin-map")
exec.Command("latte", "install", "robin-map")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
