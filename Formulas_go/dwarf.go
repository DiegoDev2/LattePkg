package main

// Dwarf - Object file manipulation tool
// Homepage: https://github.com/elboza/dwarf-ng/

import (
	"fmt"
	
	"os/exec"
)

func installDwarf() {
	// Método 1: Descargar y extraer .tar.gz
	dwarf_tar_url := "https://github.com/elboza/dwarf-ng/archive/refs/tags/dwarf-0.4.0.tar.gz"
	dwarf_cmd_tar := exec.Command("curl", "-L", dwarf_tar_url, "-o", "package.tar.gz")
	err := dwarf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dwarf_zip_url := "https://github.com/elboza/dwarf-ng/archive/refs/tags/dwarf-0.4.0.zip"
	dwarf_cmd_zip := exec.Command("curl", "-L", dwarf_zip_url, "-o", "package.zip")
	err = dwarf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dwarf_bin_url := "https://github.com/elboza/dwarf-ng/archive/refs/tags/dwarf-0.4.0.bin"
	dwarf_cmd_bin := exec.Command("curl", "-L", dwarf_bin_url, "-o", "binary.bin")
	err = dwarf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dwarf_src_url := "https://github.com/elboza/dwarf-ng/archive/refs/tags/dwarf-0.4.0.src.tar.gz"
	dwarf_cmd_src := exec.Command("curl", "-L", dwarf_src_url, "-o", "source.tar.gz")
	err = dwarf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dwarf_cmd_direct := exec.Command("./binary")
	err = dwarf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: flex")
exec.Command("latte", "install", "flex")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
