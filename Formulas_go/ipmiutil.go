package main

// Ipmiutil - IPMI server management utility
// Homepage: https://ipmiutil.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installIpmiutil() {
	// Método 1: Descargar y extraer .tar.gz
	ipmiutil_tar_url := "https://downloads.sourceforge.net/project/ipmiutil/ipmiutil-3.1.9.tar.gz"
	ipmiutil_cmd_tar := exec.Command("curl", "-L", ipmiutil_tar_url, "-o", "package.tar.gz")
	err := ipmiutil_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ipmiutil_zip_url := "https://downloads.sourceforge.net/project/ipmiutil/ipmiutil-3.1.9.zip"
	ipmiutil_cmd_zip := exec.Command("curl", "-L", ipmiutil_zip_url, "-o", "package.zip")
	err = ipmiutil_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ipmiutil_bin_url := "https://downloads.sourceforge.net/project/ipmiutil/ipmiutil-3.1.9.bin"
	ipmiutil_cmd_bin := exec.Command("curl", "-L", ipmiutil_bin_url, "-o", "binary.bin")
	err = ipmiutil_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ipmiutil_src_url := "https://downloads.sourceforge.net/project/ipmiutil/ipmiutil-3.1.9.src.tar.gz"
	ipmiutil_cmd_src := exec.Command("curl", "-L", ipmiutil_src_url, "-o", "source.tar.gz")
	err = ipmiutil_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ipmiutil_cmd_direct := exec.Command("./binary")
	err = ipmiutil_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
}
