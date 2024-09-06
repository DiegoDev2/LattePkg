package main

// Edbrowse - Command-line editor and web browser
// Homepage: https://edbrowse.org

import (
	"fmt"
	
	"os/exec"
)

func installEdbrowse() {
	// Método 1: Descargar y extraer .tar.gz
	edbrowse_tar_url := "https://github.com/CMB/edbrowse/archive/refs/tags/v3.8.10.tar.gz"
	edbrowse_cmd_tar := exec.Command("curl", "-L", edbrowse_tar_url, "-o", "package.tar.gz")
	err := edbrowse_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	edbrowse_zip_url := "https://github.com/CMB/edbrowse/archive/refs/tags/v3.8.10.zip"
	edbrowse_cmd_zip := exec.Command("curl", "-L", edbrowse_zip_url, "-o", "package.zip")
	err = edbrowse_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	edbrowse_bin_url := "https://github.com/CMB/edbrowse/archive/refs/tags/v3.8.10.bin"
	edbrowse_cmd_bin := exec.Command("curl", "-L", edbrowse_bin_url, "-o", "binary.bin")
	err = edbrowse_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	edbrowse_src_url := "https://github.com/CMB/edbrowse/archive/refs/tags/v3.8.10.src.tar.gz"
	edbrowse_cmd_src := exec.Command("curl", "-L", edbrowse_src_url, "-o", "source.tar.gz")
	err = edbrowse_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	edbrowse_cmd_direct := exec.Command("./binary")
	err = edbrowse_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: quickjs")
exec.Command("latte", "install", "quickjs")
	fmt.Println("Instalando dependencia: curl")
exec.Command("latte", "install", "curl")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: pcre2")
exec.Command("latte", "install", "pcre2")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
	fmt.Println("Instalando dependencia: unixodbc")
exec.Command("latte", "install", "unixodbc")
}
