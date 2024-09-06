package main

// Gperftools - Multi-threaded malloc() and performance analysis tools
// Homepage: https://github.com/gperftools/gperftools

import (
	"fmt"
	
	"os/exec"
)

func installGperftools() {
	// Método 1: Descargar y extraer .tar.gz
	gperftools_tar_url := "https://github.com/gperftools/gperftools/releases/download/gperftools-2.15/gperftools-2.15.tar.gz"
	gperftools_cmd_tar := exec.Command("curl", "-L", gperftools_tar_url, "-o", "package.tar.gz")
	err := gperftools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gperftools_zip_url := "https://github.com/gperftools/gperftools/releases/download/gperftools-2.15/gperftools-2.15.zip"
	gperftools_cmd_zip := exec.Command("curl", "-L", gperftools_zip_url, "-o", "package.zip")
	err = gperftools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gperftools_bin_url := "https://github.com/gperftools/gperftools/releases/download/gperftools-2.15/gperftools-2.15.bin"
	gperftools_cmd_bin := exec.Command("curl", "-L", gperftools_bin_url, "-o", "binary.bin")
	err = gperftools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gperftools_src_url := "https://github.com/gperftools/gperftools/releases/download/gperftools-2.15/gperftools-2.15.src.tar.gz"
	gperftools_cmd_src := exec.Command("curl", "-L", gperftools_src_url, "-o", "source.tar.gz")
	err = gperftools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gperftools_cmd_direct := exec.Command("./binary")
	err = gperftools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: libunwind")
exec.Command("latte", "install", "libunwind")
}
