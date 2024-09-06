package main

// Gitbatch - Manage your git repositories in one place
// Homepage: https://github.com/isacikgoz/gitbatch

import (
	"fmt"
	
	"os/exec"
)

func installGitbatch() {
	// Método 1: Descargar y extraer .tar.gz
	gitbatch_tar_url := "https://github.com/isacikgoz/gitbatch/archive/refs/tags/v0.6.1.tar.gz"
	gitbatch_cmd_tar := exec.Command("curl", "-L", gitbatch_tar_url, "-o", "package.tar.gz")
	err := gitbatch_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitbatch_zip_url := "https://github.com/isacikgoz/gitbatch/archive/refs/tags/v0.6.1.zip"
	gitbatch_cmd_zip := exec.Command("curl", "-L", gitbatch_zip_url, "-o", "package.zip")
	err = gitbatch_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitbatch_bin_url := "https://github.com/isacikgoz/gitbatch/archive/refs/tags/v0.6.1.bin"
	gitbatch_cmd_bin := exec.Command("curl", "-L", gitbatch_bin_url, "-o", "binary.bin")
	err = gitbatch_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitbatch_src_url := "https://github.com/isacikgoz/gitbatch/archive/refs/tags/v0.6.1.src.tar.gz"
	gitbatch_cmd_src := exec.Command("curl", "-L", gitbatch_src_url, "-o", "source.tar.gz")
	err = gitbatch_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitbatch_cmd_direct := exec.Command("./binary")
	err = gitbatch_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
