package main

// SourceHighlight - Source-code syntax highlighter
// Homepage: https://www.gnu.org/software/src-highlite/

import (
	"fmt"
	
	"os/exec"
)

func installSourceHighlight() {
	// Método 1: Descargar y extraer .tar.gz
	sourcehighlight_tar_url := "https://ftp.gnu.org/gnu/src-highlite/source-highlight-3.1.9.tar.gz"
	sourcehighlight_cmd_tar := exec.Command("curl", "-L", sourcehighlight_tar_url, "-o", "package.tar.gz")
	err := sourcehighlight_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sourcehighlight_zip_url := "https://ftp.gnu.org/gnu/src-highlite/source-highlight-3.1.9.zip"
	sourcehighlight_cmd_zip := exec.Command("curl", "-L", sourcehighlight_zip_url, "-o", "package.zip")
	err = sourcehighlight_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sourcehighlight_bin_url := "https://ftp.gnu.org/gnu/src-highlite/source-highlight-3.1.9.bin"
	sourcehighlight_cmd_bin := exec.Command("curl", "-L", sourcehighlight_bin_url, "-o", "binary.bin")
	err = sourcehighlight_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sourcehighlight_src_url := "https://ftp.gnu.org/gnu/src-highlite/source-highlight-3.1.9.src.tar.gz"
	sourcehighlight_cmd_src := exec.Command("curl", "-L", sourcehighlight_src_url, "-o", "source.tar.gz")
	err = sourcehighlight_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sourcehighlight_cmd_direct := exec.Command("./binary")
	err = sourcehighlight_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
}
