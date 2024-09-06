package main

// Reminiscence - Flashback engine reimplementation
// Homepage: http://cyxdown.free.fr/reminiscence/

import (
	"fmt"
	
	"os/exec"
)

func installReminiscence() {
	// Método 1: Descargar y extraer .tar.gz
	reminiscence_tar_url := "https://pkg.freebsd.org/ports-distfiles/REminiscence-0.5.1.tar.bz2"
	reminiscence_cmd_tar := exec.Command("curl", "-L", reminiscence_tar_url, "-o", "package.tar.gz")
	err := reminiscence_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	reminiscence_zip_url := "https://pkg.freebsd.org/ports-distfiles/REminiscence-0.5.1.tar.bz2"
	reminiscence_cmd_zip := exec.Command("curl", "-L", reminiscence_zip_url, "-o", "package.zip")
	err = reminiscence_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	reminiscence_bin_url := "https://pkg.freebsd.org/ports-distfiles/REminiscence-0.5.1.tar.bz2"
	reminiscence_cmd_bin := exec.Command("curl", "-L", reminiscence_bin_url, "-o", "binary.bin")
	err = reminiscence_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	reminiscence_src_url := "https://pkg.freebsd.org/ports-distfiles/REminiscence-0.5.1.tar.bz2"
	reminiscence_cmd_src := exec.Command("curl", "-L", reminiscence_src_url, "-o", "source.tar.gz")
	err = reminiscence_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	reminiscence_cmd_direct := exec.Command("./binary")
	err = reminiscence_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libmodplug")
exec.Command("latte", "install", "libmodplug")
	fmt.Println("Instalando dependencia: libogg")
exec.Command("latte", "install", "libogg")
	fmt.Println("Instalando dependencia: sdl2")
exec.Command("latte", "install", "sdl2")
}
