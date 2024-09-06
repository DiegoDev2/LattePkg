package main

// TypescriptLanguageServer - Language Server Protocol implementation for TypeScript wrapping tsserver
// Homepage: https://github.com/typescript-language-server/typescript-language-server

import (
	"fmt"
	
	"os/exec"
)

func installTypescriptLanguageServer() {
	// Método 1: Descargar y extraer .tar.gz
	typescriptlanguageserver_tar_url := "https://registry.npmjs.org/typescript-language-server/-/typescript-language-server-4.3.3.tgz"
	typescriptlanguageserver_cmd_tar := exec.Command("curl", "-L", typescriptlanguageserver_tar_url, "-o", "package.tar.gz")
	err := typescriptlanguageserver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	typescriptlanguageserver_zip_url := "https://registry.npmjs.org/typescript-language-server/-/typescript-language-server-4.3.3.tgz"
	typescriptlanguageserver_cmd_zip := exec.Command("curl", "-L", typescriptlanguageserver_zip_url, "-o", "package.zip")
	err = typescriptlanguageserver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	typescriptlanguageserver_bin_url := "https://registry.npmjs.org/typescript-language-server/-/typescript-language-server-4.3.3.tgz"
	typescriptlanguageserver_cmd_bin := exec.Command("curl", "-L", typescriptlanguageserver_bin_url, "-o", "binary.bin")
	err = typescriptlanguageserver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	typescriptlanguageserver_src_url := "https://registry.npmjs.org/typescript-language-server/-/typescript-language-server-4.3.3.tgz"
	typescriptlanguageserver_cmd_src := exec.Command("curl", "-L", typescriptlanguageserver_src_url, "-o", "source.tar.gz")
	err = typescriptlanguageserver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	typescriptlanguageserver_cmd_direct := exec.Command("./binary")
	err = typescriptlanguageserver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
	fmt.Println("Instalando dependencia: typescript")
exec.Command("latte", "install", "typescript")
}
