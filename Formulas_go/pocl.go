package main

// Pocl - Portable Computing Language
// Homepage: https://portablecl.org/

import (
	"fmt"
	
	"os/exec"
)

func installPocl() {
	// Método 1: Descargar y extraer .tar.gz
	pocl_tar_url := "https://github.com/pocl/pocl/archive/refs/tags/v6.0.tar.gz"
	pocl_cmd_tar := exec.Command("curl", "-L", pocl_tar_url, "-o", "package.tar.gz")
	err := pocl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pocl_zip_url := "https://github.com/pocl/pocl/archive/refs/tags/v6.0.zip"
	pocl_cmd_zip := exec.Command("curl", "-L", pocl_zip_url, "-o", "package.zip")
	err = pocl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pocl_bin_url := "https://github.com/pocl/pocl/archive/refs/tags/v6.0.bin"
	pocl_cmd_bin := exec.Command("curl", "-L", pocl_bin_url, "-o", "binary.bin")
	err = pocl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pocl_src_url := "https://github.com/pocl/pocl/archive/refs/tags/v6.0.src.tar.gz"
	pocl_cmd_src := exec.Command("curl", "-L", pocl_src_url, "-o", "source.tar.gz")
	err = pocl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pocl_cmd_direct := exec.Command("./binary")
	err = pocl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: opencl-headers")
exec.Command("latte", "install", "opencl-headers")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: hwloc")
exec.Command("latte", "install", "hwloc")
	fmt.Println("Instalando dependencia: llvm")
exec.Command("latte", "install", "llvm")
	fmt.Println("Instalando dependencia: opencl-icd-loader")
exec.Command("latte", "install", "opencl-icd-loader")
}
