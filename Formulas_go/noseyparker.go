package main

// Noseyparker - Finds secrets and sensitive information in textual data and Git history
// Homepage: https://github.com/praetorian-inc/noseyparker

import (
	"fmt"
	
	"os/exec"
)

func installNoseyparker() {
	// Método 1: Descargar y extraer .tar.gz
	noseyparker_tar_url := "https://github.com/praetorian-inc/noseyparker/archive/refs/tags/v0.19.0.tar.gz"
	noseyparker_cmd_tar := exec.Command("curl", "-L", noseyparker_tar_url, "-o", "package.tar.gz")
	err := noseyparker_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	noseyparker_zip_url := "https://github.com/praetorian-inc/noseyparker/archive/refs/tags/v0.19.0.zip"
	noseyparker_cmd_zip := exec.Command("curl", "-L", noseyparker_zip_url, "-o", "package.zip")
	err = noseyparker_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	noseyparker_bin_url := "https://github.com/praetorian-inc/noseyparker/archive/refs/tags/v0.19.0.bin"
	noseyparker_cmd_bin := exec.Command("curl", "-L", noseyparker_bin_url, "-o", "binary.bin")
	err = noseyparker_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	noseyparker_src_url := "https://github.com/praetorian-inc/noseyparker/archive/refs/tags/v0.19.0.src.tar.gz"
	noseyparker_cmd_src := exec.Command("curl", "-L", noseyparker_src_url, "-o", "source.tar.gz")
	err = noseyparker_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	noseyparker_cmd_direct := exec.Command("./binary")
	err = noseyparker_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
