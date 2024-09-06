package main

// Rustup - Rust toolchain installer
// Homepage: https://github.com/rust-lang/rustup

import (
	"fmt"
	
	"os/exec"
)

func installRustup() {
	// Método 1: Descargar y extraer .tar.gz
	rustup_tar_url := "https://github.com/rust-lang/rustup/archive/refs/tags/1.27.1.tar.gz"
	rustup_cmd_tar := exec.Command("curl", "-L", rustup_tar_url, "-o", "package.tar.gz")
	err := rustup_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rustup_zip_url := "https://github.com/rust-lang/rustup/archive/refs/tags/1.27.1.zip"
	rustup_cmd_zip := exec.Command("curl", "-L", rustup_zip_url, "-o", "package.zip")
	err = rustup_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rustup_bin_url := "https://github.com/rust-lang/rustup/archive/refs/tags/1.27.1.bin"
	rustup_cmd_bin := exec.Command("curl", "-L", rustup_bin_url, "-o", "binary.bin")
	err = rustup_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rustup_src_url := "https://github.com/rust-lang/rustup/archive/refs/tags/1.27.1.src.tar.gz"
	rustup_cmd_src := exec.Command("curl", "-L", rustup_src_url, "-o", "source.tar.gz")
	err = rustup_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rustup_cmd_direct := exec.Command("./binary")
	err = rustup_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
