package main

// VscodeLangserversExtracted - Language servers for HTML, CSS, JavaScript, and JSON extracted from vscode
// Homepage: https://github.com/hrsh7th/vscode-langservers-extracted

import (
	"fmt"
	
	"os/exec"
)

func installVscodeLangserversExtracted() {
	// Método 1: Descargar y extraer .tar.gz
	vscodelangserversextracted_tar_url := "https://registry.npmjs.org/vscode-langservers-extracted/-/vscode-langservers-extracted-4.10.0.tgz"
	vscodelangserversextracted_cmd_tar := exec.Command("curl", "-L", vscodelangserversextracted_tar_url, "-o", "package.tar.gz")
	err := vscodelangserversextracted_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vscodelangserversextracted_zip_url := "https://registry.npmjs.org/vscode-langservers-extracted/-/vscode-langservers-extracted-4.10.0.tgz"
	vscodelangserversextracted_cmd_zip := exec.Command("curl", "-L", vscodelangserversextracted_zip_url, "-o", "package.zip")
	err = vscodelangserversextracted_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vscodelangserversextracted_bin_url := "https://registry.npmjs.org/vscode-langservers-extracted/-/vscode-langservers-extracted-4.10.0.tgz"
	vscodelangserversextracted_cmd_bin := exec.Command("curl", "-L", vscodelangserversextracted_bin_url, "-o", "binary.bin")
	err = vscodelangserversextracted_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vscodelangserversextracted_src_url := "https://registry.npmjs.org/vscode-langservers-extracted/-/vscode-langservers-extracted-4.10.0.tgz"
	vscodelangserversextracted_cmd_src := exec.Command("curl", "-L", vscodelangserversextracted_src_url, "-o", "source.tar.gz")
	err = vscodelangserversextracted_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vscodelangserversextracted_cmd_direct := exec.Command("./binary")
	err = vscodelangserversextracted_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
