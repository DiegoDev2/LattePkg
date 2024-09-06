package main

// Gtkextra - Widgets for creating GUIs for GTK+
// Homepage: https://gtkextra.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installGtkextra() {
	// Método 1: Descargar y extraer .tar.gz
	gtkextra_tar_url := "https://downloads.sourceforge.net/project/gtkextra/3.3/gtkextra-3.3.4.tar.gz"
	gtkextra_cmd_tar := exec.Command("curl", "-L", gtkextra_tar_url, "-o", "package.tar.gz")
	err := gtkextra_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gtkextra_zip_url := "https://downloads.sourceforge.net/project/gtkextra/3.3/gtkextra-3.3.4.zip"
	gtkextra_cmd_zip := exec.Command("curl", "-L", gtkextra_zip_url, "-o", "package.zip")
	err = gtkextra_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gtkextra_bin_url := "https://downloads.sourceforge.net/project/gtkextra/3.3/gtkextra-3.3.4.bin"
	gtkextra_cmd_bin := exec.Command("curl", "-L", gtkextra_bin_url, "-o", "binary.bin")
	err = gtkextra_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gtkextra_src_url := "https://downloads.sourceforge.net/project/gtkextra/3.3/gtkextra-3.3.4.src.tar.gz"
	gtkextra_cmd_src := exec.Command("curl", "-L", gtkextra_src_url, "-o", "source.tar.gz")
	err = gtkextra_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gtkextra_cmd_direct := exec.Command("./binary")
	err = gtkextra_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: gtk+")
exec.Command("latte", "install", "gtk+")
}
