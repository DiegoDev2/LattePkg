package main

// Solana - Web-Scale Blockchain for decentralized apps and marketplaces
// Homepage: https://solana.com

import (
	"fmt"
	
	"os/exec"
)

func installSolana() {
	// Método 1: Descargar y extraer .tar.gz
	solana_tar_url := "https://github.com/solana-labs/solana/archive/refs/tags/v1.18.20.tar.gz"
	solana_cmd_tar := exec.Command("curl", "-L", solana_tar_url, "-o", "package.tar.gz")
	err := solana_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	solana_zip_url := "https://github.com/solana-labs/solana/archive/refs/tags/v1.18.20.zip"
	solana_cmd_zip := exec.Command("curl", "-L", solana_zip_url, "-o", "package.zip")
	err = solana_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	solana_bin_url := "https://github.com/solana-labs/solana/archive/refs/tags/v1.18.20.bin"
	solana_cmd_bin := exec.Command("curl", "-L", solana_bin_url, "-o", "binary.bin")
	err = solana_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	solana_src_url := "https://github.com/solana-labs/solana/archive/refs/tags/v1.18.20.src.tar.gz"
	solana_cmd_src := exec.Command("curl", "-L", solana_src_url, "-o", "source.tar.gz")
	err = solana_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	solana_cmd_direct := exec.Command("./binary")
	err = solana_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: protobuf")
exec.Command("latte", "install", "protobuf")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: systemd")
exec.Command("latte", "install", "systemd")
}
