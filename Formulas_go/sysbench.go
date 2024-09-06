package main

// Sysbench - System performance benchmark tool
// Homepage: https://github.com/akopytov/sysbench

import (
	"fmt"
	
	"os/exec"
)

func installSysbench() {
	// Método 1: Descargar y extraer .tar.gz
	sysbench_tar_url := "https://github.com/akopytov/sysbench/archive/refs/tags/1.0.20.tar.gz"
	sysbench_cmd_tar := exec.Command("curl", "-L", sysbench_tar_url, "-o", "package.tar.gz")
	err := sysbench_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sysbench_zip_url := "https://github.com/akopytov/sysbench/archive/refs/tags/1.0.20.zip"
	sysbench_cmd_zip := exec.Command("curl", "-L", sysbench_zip_url, "-o", "package.zip")
	err = sysbench_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sysbench_bin_url := "https://github.com/akopytov/sysbench/archive/refs/tags/1.0.20.bin"
	sysbench_cmd_bin := exec.Command("curl", "-L", sysbench_bin_url, "-o", "binary.bin")
	err = sysbench_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sysbench_src_url := "https://github.com/akopytov/sysbench/archive/refs/tags/1.0.20.src.tar.gz"
	sysbench_cmd_src := exec.Command("curl", "-L", sysbench_src_url, "-o", "source.tar.gz")
	err = sysbench_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sysbench_cmd_direct := exec.Command("./binary")
	err = sysbench_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libpq")
exec.Command("latte", "install", "libpq")
	fmt.Println("Instalando dependencia: luajit")
exec.Command("latte", "install", "luajit")
	fmt.Println("Instalando dependencia: mysql-client@8.0")
exec.Command("latte", "install", "mysql-client@8.0")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
