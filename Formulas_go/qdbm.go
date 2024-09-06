package main

// Qdbm - Library of routines for managing a database
// Homepage: https://dbmx.net/qdbm/

import (
	"fmt"
	
	"os/exec"
)

func installQdbm() {
	// Método 1: Descargar y extraer .tar.gz
	qdbm_tar_url := "https://dbmx.net/qdbm/qdbm-1.8.78.tar.gz"
	qdbm_cmd_tar := exec.Command("curl", "-L", qdbm_tar_url, "-o", "package.tar.gz")
	err := qdbm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	qdbm_zip_url := "https://dbmx.net/qdbm/qdbm-1.8.78.zip"
	qdbm_cmd_zip := exec.Command("curl", "-L", qdbm_zip_url, "-o", "package.zip")
	err = qdbm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	qdbm_bin_url := "https://dbmx.net/qdbm/qdbm-1.8.78.bin"
	qdbm_cmd_bin := exec.Command("curl", "-L", qdbm_bin_url, "-o", "binary.bin")
	err = qdbm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	qdbm_src_url := "https://dbmx.net/qdbm/qdbm-1.8.78.src.tar.gz"
	qdbm_cmd_src := exec.Command("curl", "-L", qdbm_src_url, "-o", "source.tar.gz")
	err = qdbm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	qdbm_cmd_direct := exec.Command("./binary")
	err = qdbm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
