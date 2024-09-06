package main

// GitQuickStats - Simple and efficient way to access statistics in git
// Homepage: https://github.com/arzzen/git-quick-stats

import (
	"fmt"
	
	"os/exec"
)

func installGitQuickStats() {
	// Método 1: Descargar y extraer .tar.gz
	gitquickstats_tar_url := "https://github.com/arzzen/git-quick-stats/archive/refs/tags/2.5.6.tar.gz"
	gitquickstats_cmd_tar := exec.Command("curl", "-L", gitquickstats_tar_url, "-o", "package.tar.gz")
	err := gitquickstats_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitquickstats_zip_url := "https://github.com/arzzen/git-quick-stats/archive/refs/tags/2.5.6.zip"
	gitquickstats_cmd_zip := exec.Command("curl", "-L", gitquickstats_zip_url, "-o", "package.zip")
	err = gitquickstats_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitquickstats_bin_url := "https://github.com/arzzen/git-quick-stats/archive/refs/tags/2.5.6.bin"
	gitquickstats_cmd_bin := exec.Command("curl", "-L", gitquickstats_bin_url, "-o", "binary.bin")
	err = gitquickstats_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitquickstats_src_url := "https://github.com/arzzen/git-quick-stats/archive/refs/tags/2.5.6.src.tar.gz"
	gitquickstats_cmd_src := exec.Command("curl", "-L", gitquickstats_src_url, "-o", "source.tar.gz")
	err = gitquickstats_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitquickstats_cmd_direct := exec.Command("./binary")
	err = gitquickstats_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: coreutils")
exec.Command("latte", "install", "coreutils")
	fmt.Println("Instalando dependencia: util-linux")
exec.Command("latte", "install", "util-linux")
}
