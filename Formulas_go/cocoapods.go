package main

// Cocoapods - Dependency manager for Cocoa projects
// Homepage: https://cocoapods.org/

import (
	"fmt"
	
	"os/exec"
)

func installCocoapods() {
	// Método 1: Descargar y extraer .tar.gz
	cocoapods_tar_url := "https://github.com/CocoaPods/CocoaPods/archive/refs/tags/1.15.2.tar.gz"
	cocoapods_cmd_tar := exec.Command("curl", "-L", cocoapods_tar_url, "-o", "package.tar.gz")
	err := cocoapods_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cocoapods_zip_url := "https://github.com/CocoaPods/CocoaPods/archive/refs/tags/1.15.2.zip"
	cocoapods_cmd_zip := exec.Command("curl", "-L", cocoapods_zip_url, "-o", "package.zip")
	err = cocoapods_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cocoapods_bin_url := "https://github.com/CocoaPods/CocoaPods/archive/refs/tags/1.15.2.bin"
	cocoapods_cmd_bin := exec.Command("curl", "-L", cocoapods_bin_url, "-o", "binary.bin")
	err = cocoapods_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cocoapods_src_url := "https://github.com/CocoaPods/CocoaPods/archive/refs/tags/1.15.2.src.tar.gz"
	cocoapods_cmd_src := exec.Command("curl", "-L", cocoapods_src_url, "-o", "source.tar.gz")
	err = cocoapods_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cocoapods_cmd_direct := exec.Command("./binary")
	err = cocoapods_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: ruby")
exec.Command("latte", "install", "ruby")
}
