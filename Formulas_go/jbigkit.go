package main

// Jbigkit - JBIG1 data compression standard implementation
// Homepage: https://www.cl.cam.ac.uk/~mgk25/jbigkit/

import (
	"fmt"
	
	"os/exec"
)

func installJbigkit() {
	// Método 1: Descargar y extraer .tar.gz
	jbigkit_tar_url := "https://www.cl.cam.ac.uk/~mgk25/jbigkit/download/jbigkit-2.1.tar.gz"
	jbigkit_cmd_tar := exec.Command("curl", "-L", jbigkit_tar_url, "-o", "package.tar.gz")
	err := jbigkit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jbigkit_zip_url := "https://www.cl.cam.ac.uk/~mgk25/jbigkit/download/jbigkit-2.1.zip"
	jbigkit_cmd_zip := exec.Command("curl", "-L", jbigkit_zip_url, "-o", "package.zip")
	err = jbigkit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jbigkit_bin_url := "https://www.cl.cam.ac.uk/~mgk25/jbigkit/download/jbigkit-2.1.bin"
	jbigkit_cmd_bin := exec.Command("curl", "-L", jbigkit_bin_url, "-o", "binary.bin")
	err = jbigkit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jbigkit_src_url := "https://www.cl.cam.ac.uk/~mgk25/jbigkit/download/jbigkit-2.1.src.tar.gz"
	jbigkit_cmd_src := exec.Command("curl", "-L", jbigkit_src_url, "-o", "source.tar.gz")
	err = jbigkit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jbigkit_cmd_direct := exec.Command("./binary")
	err = jbigkit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
