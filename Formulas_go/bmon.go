package main

// Bmon - Interface bandwidth monitor
// Homepage: https://github.com/tgraf/bmon

import (
	"fmt"
	
	"os/exec"
)

func installBmon() {
	// Método 1: Descargar y extraer .tar.gz
	bmon_tar_url := "https://github.com/tgraf/bmon/releases/download/v4.0/bmon-4.0.tar.gz"
	bmon_cmd_tar := exec.Command("curl", "-L", bmon_tar_url, "-o", "package.tar.gz")
	err := bmon_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bmon_zip_url := "https://github.com/tgraf/bmon/releases/download/v4.0/bmon-4.0.zip"
	bmon_cmd_zip := exec.Command("curl", "-L", bmon_zip_url, "-o", "package.zip")
	err = bmon_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bmon_bin_url := "https://github.com/tgraf/bmon/releases/download/v4.0/bmon-4.0.bin"
	bmon_cmd_bin := exec.Command("curl", "-L", bmon_bin_url, "-o", "binary.bin")
	err = bmon_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bmon_src_url := "https://github.com/tgraf/bmon/releases/download/v4.0/bmon-4.0.src.tar.gz"
	bmon_cmd_src := exec.Command("curl", "-L", bmon_src_url, "-o", "source.tar.gz")
	err = bmon_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bmon_cmd_direct := exec.Command("./binary")
	err = bmon_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: confuse")
exec.Command("latte", "install", "confuse")
	fmt.Println("Instalando dependencia: libnl")
exec.Command("latte", "install", "libnl")
}
