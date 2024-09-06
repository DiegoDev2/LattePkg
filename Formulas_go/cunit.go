package main

// Cunit - Lightweight unit testing framework for C
// Homepage: https://cunit.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installCunit() {
	// Método 1: Descargar y extraer .tar.gz
	cunit_tar_url := "https://downloads.sourceforge.net/project/cunit/CUnit/2.1-3/CUnit-2.1-3.tar.bz2"
	cunit_cmd_tar := exec.Command("curl", "-L", cunit_tar_url, "-o", "package.tar.gz")
	err := cunit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cunit_zip_url := "https://downloads.sourceforge.net/project/cunit/CUnit/2.1-3/CUnit-2.1-3.tar.bz2"
	cunit_cmd_zip := exec.Command("curl", "-L", cunit_zip_url, "-o", "package.zip")
	err = cunit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cunit_bin_url := "https://downloads.sourceforge.net/project/cunit/CUnit/2.1-3/CUnit-2.1-3.tar.bz2"
	cunit_cmd_bin := exec.Command("curl", "-L", cunit_bin_url, "-o", "binary.bin")
	err = cunit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cunit_src_url := "https://downloads.sourceforge.net/project/cunit/CUnit/2.1-3/CUnit-2.1-3.tar.bz2"
	cunit_cmd_src := exec.Command("curl", "-L", cunit_src_url, "-o", "source.tar.gz")
	err = cunit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cunit_cmd_direct := exec.Command("./binary")
	err = cunit_cmd_direct.Run()
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
