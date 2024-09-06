package main

// CargoAuditable - Make production Rust binaries auditable
// Homepage: https://github.com/rust-secure-code/cargo-auditable

import (
	"fmt"
	
	"os/exec"
)

func installCargoAuditable() {
	// Método 1: Descargar y extraer .tar.gz
	cargoauditable_tar_url := "https://github.com/rust-secure-code/cargo-auditable/archive/refs/tags/v0.6.4.tar.gz"
	cargoauditable_cmd_tar := exec.Command("curl", "-L", cargoauditable_tar_url, "-o", "package.tar.gz")
	err := cargoauditable_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cargoauditable_zip_url := "https://github.com/rust-secure-code/cargo-auditable/archive/refs/tags/v0.6.4.zip"
	cargoauditable_cmd_zip := exec.Command("curl", "-L", cargoauditable_zip_url, "-o", "package.zip")
	err = cargoauditable_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cargoauditable_bin_url := "https://github.com/rust-secure-code/cargo-auditable/archive/refs/tags/v0.6.4.bin"
	cargoauditable_cmd_bin := exec.Command("curl", "-L", cargoauditable_bin_url, "-o", "binary.bin")
	err = cargoauditable_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cargoauditable_src_url := "https://github.com/rust-secure-code/cargo-auditable/archive/refs/tags/v0.6.4.src.tar.gz"
	cargoauditable_cmd_src := exec.Command("curl", "-L", cargoauditable_src_url, "-o", "source.tar.gz")
	err = cargoauditable_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cargoauditable_cmd_direct := exec.Command("./binary")
	err = cargoauditable_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: rustup")
exec.Command("latte", "install", "rustup")
}
