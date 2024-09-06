package main

// Cvsutils - CVS utilities for use in working directories
// Homepage: https://www.red-bean.com/cvsutils/

import (
	"fmt"
	
	"os/exec"
)

func installCvsutils() {
	// Método 1: Descargar y extraer .tar.gz
	cvsutils_tar_url := "https://www.red-bean.com/cvsutils/releases/cvsutils-0.2.6.tar.gz"
	cvsutils_cmd_tar := exec.Command("curl", "-L", cvsutils_tar_url, "-o", "package.tar.gz")
	err := cvsutils_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cvsutils_zip_url := "https://www.red-bean.com/cvsutils/releases/cvsutils-0.2.6.zip"
	cvsutils_cmd_zip := exec.Command("curl", "-L", cvsutils_zip_url, "-o", "package.zip")
	err = cvsutils_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cvsutils_bin_url := "https://www.red-bean.com/cvsutils/releases/cvsutils-0.2.6.bin"
	cvsutils_cmd_bin := exec.Command("curl", "-L", cvsutils_bin_url, "-o", "binary.bin")
	err = cvsutils_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cvsutils_src_url := "https://www.red-bean.com/cvsutils/releases/cvsutils-0.2.6.src.tar.gz"
	cvsutils_cmd_src := exec.Command("curl", "-L", cvsutils_src_url, "-o", "source.tar.gz")
	err = cvsutils_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cvsutils_cmd_direct := exec.Command("./binary")
	err = cvsutils_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
