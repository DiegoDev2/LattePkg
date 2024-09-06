package main

// CloudflareWrangler2 - CLI tool for Cloudflare Workers
// Homepage: https://github.com/cloudflare/workers-sdk

import (
	"fmt"
	
	"os/exec"
)

func installCloudflareWrangler2() {
	// Método 1: Descargar y extraer .tar.gz
	cloudflarewrangler2_tar_url := "https://registry.npmjs.org/wrangler/-/wrangler-3.74.0.tgz"
	cloudflarewrangler2_cmd_tar := exec.Command("curl", "-L", cloudflarewrangler2_tar_url, "-o", "package.tar.gz")
	err := cloudflarewrangler2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cloudflarewrangler2_zip_url := "https://registry.npmjs.org/wrangler/-/wrangler-3.74.0.tgz"
	cloudflarewrangler2_cmd_zip := exec.Command("curl", "-L", cloudflarewrangler2_zip_url, "-o", "package.zip")
	err = cloudflarewrangler2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cloudflarewrangler2_bin_url := "https://registry.npmjs.org/wrangler/-/wrangler-3.74.0.tgz"
	cloudflarewrangler2_cmd_bin := exec.Command("curl", "-L", cloudflarewrangler2_bin_url, "-o", "binary.bin")
	err = cloudflarewrangler2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cloudflarewrangler2_src_url := "https://registry.npmjs.org/wrangler/-/wrangler-3.74.0.tgz"
	cloudflarewrangler2_cmd_src := exec.Command("curl", "-L", cloudflarewrangler2_src_url, "-o", "source.tar.gz")
	err = cloudflarewrangler2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cloudflarewrangler2_cmd_direct := exec.Command("./binary")
	err = cloudflarewrangler2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
