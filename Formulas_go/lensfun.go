package main

// Lensfun - Remove defects from digital images
// Homepage: https://lensfun.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installLensfun() {
	// Método 1: Descargar y extraer .tar.gz
	lensfun_tar_url := "https://github.com/lensfun/lensfun/archive/refs/tags/v0.3.4.tar.gz"
	lensfun_cmd_tar := exec.Command("curl", "-L", lensfun_tar_url, "-o", "package.tar.gz")
	err := lensfun_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lensfun_zip_url := "https://github.com/lensfun/lensfun/archive/refs/tags/v0.3.4.zip"
	lensfun_cmd_zip := exec.Command("curl", "-L", lensfun_zip_url, "-o", "package.zip")
	err = lensfun_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lensfun_bin_url := "https://github.com/lensfun/lensfun/archive/refs/tags/v0.3.4.bin"
	lensfun_cmd_bin := exec.Command("curl", "-L", lensfun_bin_url, "-o", "binary.bin")
	err = lensfun_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lensfun_src_url := "https://github.com/lensfun/lensfun/archive/refs/tags/v0.3.4.src.tar.gz"
	lensfun_cmd_src := exec.Command("curl", "-L", lensfun_src_url, "-o", "source.tar.gz")
	err = lensfun_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lensfun_cmd_direct := exec.Command("./binary")
	err = lensfun_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: python-setuptools")
exec.Command("latte", "install", "python-setuptools")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
