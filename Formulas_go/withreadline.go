package main

// WithReadline - Allow GNU Readline to be used with arbitrary programs
// Homepage: https://www.greenend.org.uk/rjk/sw/withreadline.html

import (
	"fmt"
	
	"os/exec"
)

func installWithReadline() {
	// Método 1: Descargar y extraer .tar.gz
	withreadline_tar_url := "https://www.greenend.org.uk/rjk/sw/with-readline-0.1.1.tar.gz"
	withreadline_cmd_tar := exec.Command("curl", "-L", withreadline_tar_url, "-o", "package.tar.gz")
	err := withreadline_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	withreadline_zip_url := "https://www.greenend.org.uk/rjk/sw/with-readline-0.1.1.zip"
	withreadline_cmd_zip := exec.Command("curl", "-L", withreadline_zip_url, "-o", "package.zip")
	err = withreadline_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	withreadline_bin_url := "https://www.greenend.org.uk/rjk/sw/with-readline-0.1.1.bin"
	withreadline_cmd_bin := exec.Command("curl", "-L", withreadline_bin_url, "-o", "binary.bin")
	err = withreadline_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	withreadline_src_url := "https://www.greenend.org.uk/rjk/sw/with-readline-0.1.1.src.tar.gz"
	withreadline_cmd_src := exec.Command("curl", "-L", withreadline_src_url, "-o", "source.tar.gz")
	err = withreadline_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	withreadline_cmd_direct := exec.Command("./binary")
	err = withreadline_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
