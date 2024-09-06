package main

// Ftjam - Build tool that can be used as a replacement for Make
// Homepage: https://www.freetype.org/jam/

import (
	"fmt"
	
	"os/exec"
)

func installFtjam() {
	// Método 1: Descargar y extraer .tar.gz
	ftjam_tar_url := "https://downloads.sourceforge.net/project/freetype/ftjam/2.5.2/ftjam-2.5.2.tar.bz2"
	ftjam_cmd_tar := exec.Command("curl", "-L", ftjam_tar_url, "-o", "package.tar.gz")
	err := ftjam_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ftjam_zip_url := "https://downloads.sourceforge.net/project/freetype/ftjam/2.5.2/ftjam-2.5.2.tar.bz2"
	ftjam_cmd_zip := exec.Command("curl", "-L", ftjam_zip_url, "-o", "package.zip")
	err = ftjam_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ftjam_bin_url := "https://downloads.sourceforge.net/project/freetype/ftjam/2.5.2/ftjam-2.5.2.tar.bz2"
	ftjam_cmd_bin := exec.Command("curl", "-L", ftjam_bin_url, "-o", "binary.bin")
	err = ftjam_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ftjam_src_url := "https://downloads.sourceforge.net/project/freetype/ftjam/2.5.2/ftjam-2.5.2.tar.bz2"
	ftjam_cmd_src := exec.Command("curl", "-L", ftjam_src_url, "-o", "source.tar.gz")
	err = ftjam_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ftjam_cmd_direct := exec.Command("./binary")
	err = ftjam_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
}
