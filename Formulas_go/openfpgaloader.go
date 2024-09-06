package main

// Openfpgaloader - Universal utility for programming FPGA
// Homepage: https://github.com/trabucayre/openFPGALoader

import (
	"fmt"
	
	"os/exec"
)

func installOpenfpgaloader() {
	// Método 1: Descargar y extraer .tar.gz
	openfpgaloader_tar_url := "https://github.com/trabucayre/openFPGALoader/archive/refs/tags/v0.12.1.tar.gz"
	openfpgaloader_cmd_tar := exec.Command("curl", "-L", openfpgaloader_tar_url, "-o", "package.tar.gz")
	err := openfpgaloader_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openfpgaloader_zip_url := "https://github.com/trabucayre/openFPGALoader/archive/refs/tags/v0.12.1.zip"
	openfpgaloader_cmd_zip := exec.Command("curl", "-L", openfpgaloader_zip_url, "-o", "package.zip")
	err = openfpgaloader_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openfpgaloader_bin_url := "https://github.com/trabucayre/openFPGALoader/archive/refs/tags/v0.12.1.bin"
	openfpgaloader_cmd_bin := exec.Command("curl", "-L", openfpgaloader_bin_url, "-o", "binary.bin")
	err = openfpgaloader_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openfpgaloader_src_url := "https://github.com/trabucayre/openFPGALoader/archive/refs/tags/v0.12.1.src.tar.gz"
	openfpgaloader_cmd_src := exec.Command("curl", "-L", openfpgaloader_src_url, "-o", "source.tar.gz")
	err = openfpgaloader_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openfpgaloader_cmd_direct := exec.Command("./binary")
	err = openfpgaloader_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libftdi")
exec.Command("latte", "install", "libftdi")
	fmt.Println("Instalando dependencia: libusb")
exec.Command("latte", "install", "libusb")
}
