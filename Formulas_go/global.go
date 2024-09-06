package main

// Global - Source code tag system
// Homepage: https://www.gnu.org/software/global/

import (
	"fmt"
	
	"os/exec"
)

func installGlobal() {
	// Método 1: Descargar y extraer .tar.gz
	global_tar_url := "https://ftp.gnu.org/gnu/global/global-6.6.13.tar.gz"
	global_cmd_tar := exec.Command("curl", "-L", global_tar_url, "-o", "package.tar.gz")
	err := global_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	global_zip_url := "https://ftp.gnu.org/gnu/global/global-6.6.13.zip"
	global_cmd_zip := exec.Command("curl", "-L", global_zip_url, "-o", "package.zip")
	err = global_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	global_bin_url := "https://ftp.gnu.org/gnu/global/global-6.6.13.bin"
	global_cmd_bin := exec.Command("curl", "-L", global_bin_url, "-o", "binary.bin")
	err = global_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	global_src_url := "https://ftp.gnu.org/gnu/global/global-6.6.13.src.tar.gz"
	global_cmd_src := exec.Command("curl", "-L", global_src_url, "-o", "source.tar.gz")
	err = global_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	global_cmd_direct := exec.Command("./binary")
	err = global_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: bison")
exec.Command("latte", "install", "bison")
	fmt.Println("Instalando dependencia: flex")
exec.Command("latte", "install", "flex")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: ncurses")
exec.Command("latte", "install", "ncurses")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: sqlite")
exec.Command("latte", "install", "sqlite")
	fmt.Println("Instalando dependencia: universal-ctags")
exec.Command("latte", "install", "universal-ctags")
}
