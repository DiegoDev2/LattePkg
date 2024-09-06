package main

// ElanInit - Lean Theorem Prover installer and version manager
// Homepage: https://github.com/leanprover/elan

import (
	"fmt"
	
	"os/exec"
)

func installElanInit() {
	// Método 1: Descargar y extraer .tar.gz
	elaninit_tar_url := "https://github.com/leanprover/elan/archive/refs/tags/v3.1.1.tar.gz"
	elaninit_cmd_tar := exec.Command("curl", "-L", elaninit_tar_url, "-o", "package.tar.gz")
	err := elaninit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	elaninit_zip_url := "https://github.com/leanprover/elan/archive/refs/tags/v3.1.1.zip"
	elaninit_cmd_zip := exec.Command("curl", "-L", elaninit_zip_url, "-o", "package.zip")
	err = elaninit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	elaninit_bin_url := "https://github.com/leanprover/elan/archive/refs/tags/v3.1.1.bin"
	elaninit_cmd_bin := exec.Command("curl", "-L", elaninit_bin_url, "-o", "binary.bin")
	err = elaninit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	elaninit_src_url := "https://github.com/leanprover/elan/archive/refs/tags/v3.1.1.src.tar.gz"
	elaninit_cmd_src := exec.Command("curl", "-L", elaninit_src_url, "-o", "source.tar.gz")
	err = elaninit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	elaninit_cmd_direct := exec.Command("./binary")
	err = elaninit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: coreutils")
exec.Command("latte", "install", "coreutils")
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
