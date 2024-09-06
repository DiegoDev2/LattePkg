package main

// Libxscrnsaver - X.Org: X11 Screen Saver extension client library
// Homepage: https://www.x.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibxscrnsaver() {
	// Método 1: Descargar y extraer .tar.gz
	libxscrnsaver_tar_url := "https://www.x.org/archive/individual/lib/libXScrnSaver-1.2.4.tar.xz"
	libxscrnsaver_cmd_tar := exec.Command("curl", "-L", libxscrnsaver_tar_url, "-o", "package.tar.gz")
	err := libxscrnsaver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libxscrnsaver_zip_url := "https://www.x.org/archive/individual/lib/libXScrnSaver-1.2.4.tar.xz"
	libxscrnsaver_cmd_zip := exec.Command("curl", "-L", libxscrnsaver_zip_url, "-o", "package.zip")
	err = libxscrnsaver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libxscrnsaver_bin_url := "https://www.x.org/archive/individual/lib/libXScrnSaver-1.2.4.tar.xz"
	libxscrnsaver_cmd_bin := exec.Command("curl", "-L", libxscrnsaver_bin_url, "-o", "binary.bin")
	err = libxscrnsaver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libxscrnsaver_src_url := "https://www.x.org/archive/individual/lib/libXScrnSaver-1.2.4.tar.xz"
	libxscrnsaver_cmd_src := exec.Command("curl", "-L", libxscrnsaver_src_url, "-o", "source.tar.gz")
	err = libxscrnsaver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libxscrnsaver_cmd_direct := exec.Command("./binary")
	err = libxscrnsaver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
	fmt.Println("Instalando dependencia: libxext")
exec.Command("latte", "install", "libxext")
	fmt.Println("Instalando dependencia: xorgproto")
exec.Command("latte", "install", "xorgproto")
}
