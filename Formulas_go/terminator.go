package main

// Terminator - Multiple GNOME terminals in one window
// Homepage: https://gnome-terminator.org

import (
	"fmt"
	
	"os/exec"
)

func installTerminator() {
	// Método 1: Descargar y extraer .tar.gz
	terminator_tar_url := "https://github.com/gnome-terminator/terminator/archive/refs/tags/v2.1.4.tar.gz"
	terminator_cmd_tar := exec.Command("curl", "-L", terminator_tar_url, "-o", "package.tar.gz")
	err := terminator_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	terminator_zip_url := "https://github.com/gnome-terminator/terminator/archive/refs/tags/v2.1.4.zip"
	terminator_cmd_zip := exec.Command("curl", "-L", terminator_zip_url, "-o", "package.zip")
	err = terminator_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	terminator_bin_url := "https://github.com/gnome-terminator/terminator/archive/refs/tags/v2.1.4.bin"
	terminator_cmd_bin := exec.Command("curl", "-L", terminator_bin_url, "-o", "binary.bin")
	err = terminator_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	terminator_src_url := "https://github.com/gnome-terminator/terminator/archive/refs/tags/v2.1.4.src.tar.gz"
	terminator_cmd_src := exec.Command("curl", "-L", terminator_src_url, "-o", "source.tar.gz")
	err = terminator_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	terminator_cmd_direct := exec.Command("./binary")
	err = terminator_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pygobject3")
exec.Command("latte", "install", "pygobject3")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: vte3")
exec.Command("latte", "install", "vte3")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
