package main

// LibomemoC - Implementation of Signal's ratcheting forward secrecy protocol
// Homepage: https://github.com/dino/libomemo-c

import (
	"fmt"
	
	"os/exec"
)

func installLibomemoC() {
	// Método 1: Descargar y extraer .tar.gz
	libomemoc_tar_url := "https://github.com/dino/libomemo-c/archive/refs/tags/v0.5.0.tar.gz"
	libomemoc_cmd_tar := exec.Command("curl", "-L", libomemoc_tar_url, "-o", "package.tar.gz")
	err := libomemoc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libomemoc_zip_url := "https://github.com/dino/libomemo-c/archive/refs/tags/v0.5.0.zip"
	libomemoc_cmd_zip := exec.Command("curl", "-L", libomemoc_zip_url, "-o", "package.zip")
	err = libomemoc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libomemoc_bin_url := "https://github.com/dino/libomemo-c/archive/refs/tags/v0.5.0.bin"
	libomemoc_cmd_bin := exec.Command("curl", "-L", libomemoc_bin_url, "-o", "binary.bin")
	err = libomemoc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libomemoc_src_url := "https://github.com/dino/libomemo-c/archive/refs/tags/v0.5.0.src.tar.gz"
	libomemoc_cmd_src := exec.Command("curl", "-L", libomemoc_src_url, "-o", "source.tar.gz")
	err = libomemoc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libomemoc_cmd_direct := exec.Command("./binary")
	err = libomemoc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: protobuf-c")
exec.Command("latte", "install", "protobuf-c")
}
