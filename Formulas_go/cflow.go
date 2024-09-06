package main

// Cflow - Generate call graphs from C code
// Homepage: https://www.gnu.org/software/cflow/

import (
	"fmt"
	
	"os/exec"
)

func installCflow() {
	// Método 1: Descargar y extraer .tar.gz
	cflow_tar_url := "https://ftp.gnu.org/gnu/cflow/cflow-1.7.tar.bz2"
	cflow_cmd_tar := exec.Command("curl", "-L", cflow_tar_url, "-o", "package.tar.gz")
	err := cflow_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cflow_zip_url := "https://ftp.gnu.org/gnu/cflow/cflow-1.7.tar.bz2"
	cflow_cmd_zip := exec.Command("curl", "-L", cflow_zip_url, "-o", "package.zip")
	err = cflow_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cflow_bin_url := "https://ftp.gnu.org/gnu/cflow/cflow-1.7.tar.bz2"
	cflow_cmd_bin := exec.Command("curl", "-L", cflow_bin_url, "-o", "binary.bin")
	err = cflow_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cflow_src_url := "https://ftp.gnu.org/gnu/cflow/cflow-1.7.tar.bz2"
	cflow_cmd_src := exec.Command("curl", "-L", cflow_src_url, "-o", "source.tar.gz")
	err = cflow_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cflow_cmd_direct := exec.Command("./binary")
	err = cflow_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
