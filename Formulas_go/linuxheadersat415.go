package main

// LinuxHeadersAT415 - Header files of the Linux kernel
// Homepage: https://kernel.org/

import (
	"fmt"
	
	"os/exec"
)

func installLinuxHeadersAT415() {
	// Método 1: Descargar y extraer .tar.gz
	linuxheadersat415_tar_url := "https://cdn.kernel.org/pub/linux/kernel/v4.x/linux-4.15.18.tar.gz"
	linuxheadersat415_cmd_tar := exec.Command("curl", "-L", linuxheadersat415_tar_url, "-o", "package.tar.gz")
	err := linuxheadersat415_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	linuxheadersat415_zip_url := "https://cdn.kernel.org/pub/linux/kernel/v4.x/linux-4.15.18.zip"
	linuxheadersat415_cmd_zip := exec.Command("curl", "-L", linuxheadersat415_zip_url, "-o", "package.zip")
	err = linuxheadersat415_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	linuxheadersat415_bin_url := "https://cdn.kernel.org/pub/linux/kernel/v4.x/linux-4.15.18.bin"
	linuxheadersat415_cmd_bin := exec.Command("curl", "-L", linuxheadersat415_bin_url, "-o", "binary.bin")
	err = linuxheadersat415_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	linuxheadersat415_src_url := "https://cdn.kernel.org/pub/linux/kernel/v4.x/linux-4.15.18.src.tar.gz"
	linuxheadersat415_cmd_src := exec.Command("curl", "-L", linuxheadersat415_src_url, "-o", "source.tar.gz")
	err = linuxheadersat415_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	linuxheadersat415_cmd_direct := exec.Command("./binary")
	err = linuxheadersat415_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
