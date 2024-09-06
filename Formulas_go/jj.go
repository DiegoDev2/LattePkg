package main

// Jj - Git-compatible distributed version control system
// Homepage: https://github.com/martinvonz/jj

import (
	"fmt"
	
	"os/exec"
)

func installJj() {
	// Método 1: Descargar y extraer .tar.gz
	jj_tar_url := "https://github.com/martinvonz/jj/archive/refs/tags/v0.21.0.tar.gz"
	jj_cmd_tar := exec.Command("curl", "-L", jj_tar_url, "-o", "package.tar.gz")
	err := jj_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jj_zip_url := "https://github.com/martinvonz/jj/archive/refs/tags/v0.21.0.zip"
	jj_cmd_zip := exec.Command("curl", "-L", jj_zip_url, "-o", "package.zip")
	err = jj_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jj_bin_url := "https://github.com/martinvonz/jj/archive/refs/tags/v0.21.0.bin"
	jj_cmd_bin := exec.Command("curl", "-L", jj_bin_url, "-o", "binary.bin")
	err = jj_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jj_src_url := "https://github.com/martinvonz/jj/archive/refs/tags/v0.21.0.src.tar.gz"
	jj_cmd_src := exec.Command("curl", "-L", jj_src_url, "-o", "source.tar.gz")
	err = jj_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jj_cmd_direct := exec.Command("./binary")
	err = jj_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: libgit2")
exec.Command("latte", "install", "libgit2")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
