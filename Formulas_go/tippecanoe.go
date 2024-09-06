package main

// Tippecanoe - Build vector tilesets from collections of GeoJSON features
// Homepage: https://github.com/felt/tippecanoe

import (
	"fmt"
	
	"os/exec"
)

func installTippecanoe() {
	// Método 1: Descargar y extraer .tar.gz
	tippecanoe_tar_url := "https://github.com/felt/tippecanoe/archive/refs/tags/2.62.0.tar.gz"
	tippecanoe_cmd_tar := exec.Command("curl", "-L", tippecanoe_tar_url, "-o", "package.tar.gz")
	err := tippecanoe_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tippecanoe_zip_url := "https://github.com/felt/tippecanoe/archive/refs/tags/2.62.0.zip"
	tippecanoe_cmd_zip := exec.Command("curl", "-L", tippecanoe_zip_url, "-o", "package.zip")
	err = tippecanoe_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tippecanoe_bin_url := "https://github.com/felt/tippecanoe/archive/refs/tags/2.62.0.bin"
	tippecanoe_cmd_bin := exec.Command("curl", "-L", tippecanoe_bin_url, "-o", "binary.bin")
	err = tippecanoe_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tippecanoe_src_url := "https://github.com/felt/tippecanoe/archive/refs/tags/2.62.0.src.tar.gz"
	tippecanoe_cmd_src := exec.Command("curl", "-L", tippecanoe_src_url, "-o", "source.tar.gz")
	err = tippecanoe_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tippecanoe_cmd_direct := exec.Command("./binary")
	err = tippecanoe_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
