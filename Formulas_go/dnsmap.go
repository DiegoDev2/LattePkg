package main

// Dnsmap - Passive DNS network mapper (a.k.a. subdomains bruteforcer)
// Homepage: https://github.com/resurrecting-open-source-projects/dnsmap

import (
	"fmt"
	
	"os/exec"
)

func installDnsmap() {
	// Método 1: Descargar y extraer .tar.gz
	dnsmap_tar_url := "https://github.com/resurrecting-open-source-projects/dnsmap/archive/refs/tags/0.36.tar.gz"
	dnsmap_cmd_tar := exec.Command("curl", "-L", dnsmap_tar_url, "-o", "package.tar.gz")
	err := dnsmap_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dnsmap_zip_url := "https://github.com/resurrecting-open-source-projects/dnsmap/archive/refs/tags/0.36.zip"
	dnsmap_cmd_zip := exec.Command("curl", "-L", dnsmap_zip_url, "-o", "package.zip")
	err = dnsmap_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dnsmap_bin_url := "https://github.com/resurrecting-open-source-projects/dnsmap/archive/refs/tags/0.36.bin"
	dnsmap_cmd_bin := exec.Command("curl", "-L", dnsmap_bin_url, "-o", "binary.bin")
	err = dnsmap_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dnsmap_src_url := "https://github.com/resurrecting-open-source-projects/dnsmap/archive/refs/tags/0.36.src.tar.gz"
	dnsmap_cmd_src := exec.Command("curl", "-L", dnsmap_src_url, "-o", "source.tar.gz")
	err = dnsmap_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dnsmap_cmd_direct := exec.Command("./binary")
	err = dnsmap_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
}
