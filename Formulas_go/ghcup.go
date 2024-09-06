package main

// Ghcup - Installer for the general purpose language Haskell
// Homepage: https://www.haskell.org/ghcup/

import (
	"fmt"
	
	"os/exec"
)

func installGhcup() {
	// Método 1: Descargar y extraer .tar.gz
	ghcup_tar_url := "https://github.com/haskell/ghcup-hs/archive/refs/tags/v0.1.30.0.tar.gz"
	ghcup_cmd_tar := exec.Command("curl", "-L", ghcup_tar_url, "-o", "package.tar.gz")
	err := ghcup_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ghcup_zip_url := "https://github.com/haskell/ghcup-hs/archive/refs/tags/v0.1.30.0.zip"
	ghcup_cmd_zip := exec.Command("curl", "-L", ghcup_zip_url, "-o", "package.zip")
	err = ghcup_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ghcup_bin_url := "https://github.com/haskell/ghcup-hs/archive/refs/tags/v0.1.30.0.bin"
	ghcup_cmd_bin := exec.Command("curl", "-L", ghcup_bin_url, "-o", "binary.bin")
	err = ghcup_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ghcup_src_url := "https://github.com/haskell/ghcup-hs/archive/refs/tags/v0.1.30.0.src.tar.gz"
	ghcup_cmd_src := exec.Command("curl", "-L", ghcup_src_url, "-o", "source.tar.gz")
	err = ghcup_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ghcup_cmd_direct := exec.Command("./binary")
	err = ghcup_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cabal-install")
exec.Command("latte", "install", "cabal-install")
	fmt.Println("Instalando dependencia: ghc@9.4")
exec.Command("latte", "install", "ghc@9.4")
}
