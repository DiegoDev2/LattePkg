package main

// Parallel - Shell command parallelization utility
// Homepage: https://savannah.gnu.org/projects/parallel/

import (
	"fmt"
	
	"os/exec"
)

func installParallel() {
	// Método 1: Descargar y extraer .tar.gz
	parallel_tar_url := "https://ftp.gnu.org/gnu/parallel/parallel-20240822.tar.bz2"
	parallel_cmd_tar := exec.Command("curl", "-L", parallel_tar_url, "-o", "package.tar.gz")
	err := parallel_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	parallel_zip_url := "https://ftp.gnu.org/gnu/parallel/parallel-20240822.tar.bz2"
	parallel_cmd_zip := exec.Command("curl", "-L", parallel_zip_url, "-o", "package.zip")
	err = parallel_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	parallel_bin_url := "https://ftp.gnu.org/gnu/parallel/parallel-20240822.tar.bz2"
	parallel_cmd_bin := exec.Command("curl", "-L", parallel_bin_url, "-o", "binary.bin")
	err = parallel_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	parallel_src_url := "https://ftp.gnu.org/gnu/parallel/parallel-20240822.tar.bz2"
	parallel_cmd_src := exec.Command("curl", "-L", parallel_src_url, "-o", "source.tar.gz")
	err = parallel_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	parallel_cmd_direct := exec.Command("./binary")
	err = parallel_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
