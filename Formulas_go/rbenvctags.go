package main

// RbenvCtags - Automatically generate ctags for rbenv Ruby stdlibs
// Homepage: https://github.com/tpope/rbenv-ctags

import (
	"fmt"
	
	"os/exec"
)

func installRbenvCtags() {
	// Método 1: Descargar y extraer .tar.gz
	rbenvctags_tar_url := "https://github.com/tpope/rbenv-ctags/archive/refs/tags/v1.0.2.tar.gz"
	rbenvctags_cmd_tar := exec.Command("curl", "-L", rbenvctags_tar_url, "-o", "package.tar.gz")
	err := rbenvctags_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rbenvctags_zip_url := "https://github.com/tpope/rbenv-ctags/archive/refs/tags/v1.0.2.zip"
	rbenvctags_cmd_zip := exec.Command("curl", "-L", rbenvctags_zip_url, "-o", "package.zip")
	err = rbenvctags_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rbenvctags_bin_url := "https://github.com/tpope/rbenv-ctags/archive/refs/tags/v1.0.2.bin"
	rbenvctags_cmd_bin := exec.Command("curl", "-L", rbenvctags_bin_url, "-o", "binary.bin")
	err = rbenvctags_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rbenvctags_src_url := "https://github.com/tpope/rbenv-ctags/archive/refs/tags/v1.0.2.src.tar.gz"
	rbenvctags_cmd_src := exec.Command("curl", "-L", rbenvctags_src_url, "-o", "source.tar.gz")
	err = rbenvctags_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rbenvctags_cmd_direct := exec.Command("./binary")
	err = rbenvctags_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ctags")
exec.Command("latte", "install", "ctags")
	fmt.Println("Instalando dependencia: rbenv")
exec.Command("latte", "install", "rbenv")
}
