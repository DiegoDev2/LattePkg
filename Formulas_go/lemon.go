package main

// Lemon - LALR(1) parser generator like yacc or bison
// Homepage: https://www.hwaci.com/sw/lemon/

import (
	"fmt"
	
	"os/exec"
)

func installLemon() {
	// Método 1: Descargar y extraer .tar.gz
	lemon_tar_url := "https://www.sqlite.org/2024/sqlite-src-3460100.zip"
	lemon_cmd_tar := exec.Command("curl", "-L", lemon_tar_url, "-o", "package.tar.gz")
	err := lemon_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lemon_zip_url := "https://www.sqlite.org/2024/sqlite-src-3460100.zip"
	lemon_cmd_zip := exec.Command("curl", "-L", lemon_zip_url, "-o", "package.zip")
	err = lemon_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lemon_bin_url := "https://www.sqlite.org/2024/sqlite-src-3460100.zip"
	lemon_cmd_bin := exec.Command("curl", "-L", lemon_bin_url, "-o", "binary.bin")
	err = lemon_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lemon_src_url := "https://www.sqlite.org/2024/sqlite-src-3460100.zip"
	lemon_cmd_src := exec.Command("curl", "-L", lemon_src_url, "-o", "source.tar.gz")
	err = lemon_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lemon_cmd_direct := exec.Command("./binary")
	err = lemon_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
