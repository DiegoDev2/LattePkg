package main

// Dnscontrol - Synchronize your DNS to multiple providers from a simple DSL
// Homepage: https://dnscontrol.org/

import (
	"fmt"
	
	"os/exec"
)

func installDnscontrol() {
	// Método 1: Descargar y extraer .tar.gz
	dnscontrol_tar_url := "https://github.com/StackExchange/dnscontrol/archive/refs/tags/v4.13.0.tar.gz"
	dnscontrol_cmd_tar := exec.Command("curl", "-L", dnscontrol_tar_url, "-o", "package.tar.gz")
	err := dnscontrol_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dnscontrol_zip_url := "https://github.com/StackExchange/dnscontrol/archive/refs/tags/v4.13.0.zip"
	dnscontrol_cmd_zip := exec.Command("curl", "-L", dnscontrol_zip_url, "-o", "package.zip")
	err = dnscontrol_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dnscontrol_bin_url := "https://github.com/StackExchange/dnscontrol/archive/refs/tags/v4.13.0.bin"
	dnscontrol_cmd_bin := exec.Command("curl", "-L", dnscontrol_bin_url, "-o", "binary.bin")
	err = dnscontrol_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dnscontrol_src_url := "https://github.com/StackExchange/dnscontrol/archive/refs/tags/v4.13.0.src.tar.gz"
	dnscontrol_cmd_src := exec.Command("curl", "-L", dnscontrol_src_url, "-o", "source.tar.gz")
	err = dnscontrol_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dnscontrol_cmd_direct := exec.Command("./binary")
	err = dnscontrol_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
