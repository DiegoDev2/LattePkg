package main

// Optipng - PNG file optimizer
// Homepage: https://optipng.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installOptipng() {
	// Método 1: Descargar y extraer .tar.gz
	optipng_tar_url := "https://downloads.sourceforge.net/project/optipng/OptiPNG/optipng-0.7.8/optipng-0.7.8.tar.gz"
	optipng_cmd_tar := exec.Command("curl", "-L", optipng_tar_url, "-o", "package.tar.gz")
	err := optipng_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	optipng_zip_url := "https://downloads.sourceforge.net/project/optipng/OptiPNG/optipng-0.7.8/optipng-0.7.8.zip"
	optipng_cmd_zip := exec.Command("curl", "-L", optipng_zip_url, "-o", "package.zip")
	err = optipng_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	optipng_bin_url := "https://downloads.sourceforge.net/project/optipng/OptiPNG/optipng-0.7.8/optipng-0.7.8.bin"
	optipng_cmd_bin := exec.Command("curl", "-L", optipng_bin_url, "-o", "binary.bin")
	err = optipng_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	optipng_src_url := "https://downloads.sourceforge.net/project/optipng/OptiPNG/optipng-0.7.8/optipng-0.7.8.src.tar.gz"
	optipng_cmd_src := exec.Command("curl", "-L", optipng_src_url, "-o", "source.tar.gz")
	err = optipng_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	optipng_cmd_direct := exec.Command("./binary")
	err = optipng_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
}
