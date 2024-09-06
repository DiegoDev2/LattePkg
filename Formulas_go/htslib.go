package main

// Htslib - C library for high-throughput sequencing data formats
// Homepage: https://www.htslib.org/

import (
	"fmt"
	
	"os/exec"
)

func installHtslib() {
	// Método 1: Descargar y extraer .tar.gz
	htslib_tar_url := "https://github.com/samtools/htslib/releases/download/1.20/htslib-1.20.tar.bz2"
	htslib_cmd_tar := exec.Command("curl", "-L", htslib_tar_url, "-o", "package.tar.gz")
	err := htslib_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	htslib_zip_url := "https://github.com/samtools/htslib/releases/download/1.20/htslib-1.20.tar.bz2"
	htslib_cmd_zip := exec.Command("curl", "-L", htslib_zip_url, "-o", "package.zip")
	err = htslib_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	htslib_bin_url := "https://github.com/samtools/htslib/releases/download/1.20/htslib-1.20.tar.bz2"
	htslib_cmd_bin := exec.Command("curl", "-L", htslib_bin_url, "-o", "binary.bin")
	err = htslib_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	htslib_src_url := "https://github.com/samtools/htslib/releases/download/1.20/htslib-1.20.tar.bz2"
	htslib_cmd_src := exec.Command("curl", "-L", htslib_src_url, "-o", "source.tar.gz")
	err = htslib_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	htslib_cmd_direct := exec.Command("./binary")
	err = htslib_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libdeflate")
exec.Command("latte", "install", "libdeflate")
	fmt.Println("Instalando dependencia: xz")
exec.Command("latte", "install", "xz")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
