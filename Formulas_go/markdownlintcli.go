package main

// MarkdownlintCli - CLI for Node.js style checker and lint tool for Markdown files
// Homepage: https://github.com/igorshubovych/markdownlint-cli

import (
	"fmt"
	
	"os/exec"
)

func installMarkdownlintCli() {
	// Método 1: Descargar y extraer .tar.gz
	markdownlintcli_tar_url := "https://registry.npmjs.org/markdownlint-cli/-/markdownlint-cli-0.41.0.tgz"
	markdownlintcli_cmd_tar := exec.Command("curl", "-L", markdownlintcli_tar_url, "-o", "package.tar.gz")
	err := markdownlintcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	markdownlintcli_zip_url := "https://registry.npmjs.org/markdownlint-cli/-/markdownlint-cli-0.41.0.tgz"
	markdownlintcli_cmd_zip := exec.Command("curl", "-L", markdownlintcli_zip_url, "-o", "package.zip")
	err = markdownlintcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	markdownlintcli_bin_url := "https://registry.npmjs.org/markdownlint-cli/-/markdownlint-cli-0.41.0.tgz"
	markdownlintcli_cmd_bin := exec.Command("curl", "-L", markdownlintcli_bin_url, "-o", "binary.bin")
	err = markdownlintcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	markdownlintcli_src_url := "https://registry.npmjs.org/markdownlint-cli/-/markdownlint-cli-0.41.0.tgz"
	markdownlintcli_cmd_src := exec.Command("curl", "-L", markdownlintcli_src_url, "-o", "source.tar.gz")
	err = markdownlintcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	markdownlintcli_cmd_direct := exec.Command("./binary")
	err = markdownlintcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
