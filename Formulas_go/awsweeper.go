package main

// Awsweeper - CLI tool for cleaning your AWS account
// Homepage: https://github.com/jckuester/awsweeper/

import (
	"fmt"
	
	"os/exec"
)

func installAwsweeper() {
	// Método 1: Descargar y extraer .tar.gz
	awsweeper_tar_url := "https://github.com/jckuester/awsweeper/archive/refs/tags/v0.12.0.tar.gz"
	awsweeper_cmd_tar := exec.Command("curl", "-L", awsweeper_tar_url, "-o", "package.tar.gz")
	err := awsweeper_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	awsweeper_zip_url := "https://github.com/jckuester/awsweeper/archive/refs/tags/v0.12.0.zip"
	awsweeper_cmd_zip := exec.Command("curl", "-L", awsweeper_zip_url, "-o", "package.zip")
	err = awsweeper_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	awsweeper_bin_url := "https://github.com/jckuester/awsweeper/archive/refs/tags/v0.12.0.bin"
	awsweeper_cmd_bin := exec.Command("curl", "-L", awsweeper_bin_url, "-o", "binary.bin")
	err = awsweeper_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	awsweeper_src_url := "https://github.com/jckuester/awsweeper/archive/refs/tags/v0.12.0.src.tar.gz"
	awsweeper_cmd_src := exec.Command("curl", "-L", awsweeper_src_url, "-o", "source.tar.gz")
	err = awsweeper_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	awsweeper_cmd_direct := exec.Command("./binary")
	err = awsweeper_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
