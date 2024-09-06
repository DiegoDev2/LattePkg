package main

// Tivodecode - Convert .tivo to .mpeg
// Homepage: https://tivodecode.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installTivodecode() {
	// Método 1: Descargar y extraer .tar.gz
	tivodecode_tar_url := "https://downloads.sourceforge.net/project/tivodecode/tivodecode/0.2pre4/tivodecode-0.2pre4.tar.gz"
	tivodecode_cmd_tar := exec.Command("curl", "-L", tivodecode_tar_url, "-o", "package.tar.gz")
	err := tivodecode_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tivodecode_zip_url := "https://downloads.sourceforge.net/project/tivodecode/tivodecode/0.2pre4/tivodecode-0.2pre4.zip"
	tivodecode_cmd_zip := exec.Command("curl", "-L", tivodecode_zip_url, "-o", "package.zip")
	err = tivodecode_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tivodecode_bin_url := "https://downloads.sourceforge.net/project/tivodecode/tivodecode/0.2pre4/tivodecode-0.2pre4.bin"
	tivodecode_cmd_bin := exec.Command("curl", "-L", tivodecode_bin_url, "-o", "binary.bin")
	err = tivodecode_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tivodecode_src_url := "https://downloads.sourceforge.net/project/tivodecode/tivodecode/0.2pre4/tivodecode-0.2pre4.src.tar.gz"
	tivodecode_cmd_src := exec.Command("curl", "-L", tivodecode_src_url, "-o", "source.tar.gz")
	err = tivodecode_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tivodecode_cmd_direct := exec.Command("./binary")
	err = tivodecode_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
