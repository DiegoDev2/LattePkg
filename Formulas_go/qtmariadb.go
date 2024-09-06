package main

// QtMariadb - Qt SQL Database Driver
// Homepage: https://www.qt.io/

import (
	"fmt"
	
	"os/exec"
)

func installQtMariadb() {
	// Método 1: Descargar y extraer .tar.gz
	qtmariadb_tar_url := "https://download.qt.io/official_releases/qt/6.7/6.7.0/submodules/qtbase-everywhere-src-6.7.0.tar.xz"
	qtmariadb_cmd_tar := exec.Command("curl", "-L", qtmariadb_tar_url, "-o", "package.tar.gz")
	err := qtmariadb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	qtmariadb_zip_url := "https://download.qt.io/official_releases/qt/6.7/6.7.0/submodules/qtbase-everywhere-src-6.7.0.tar.xz"
	qtmariadb_cmd_zip := exec.Command("curl", "-L", qtmariadb_zip_url, "-o", "package.zip")
	err = qtmariadb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	qtmariadb_bin_url := "https://download.qt.io/official_releases/qt/6.7/6.7.0/submodules/qtbase-everywhere-src-6.7.0.tar.xz"
	qtmariadb_cmd_bin := exec.Command("curl", "-L", qtmariadb_bin_url, "-o", "binary.bin")
	err = qtmariadb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	qtmariadb_src_url := "https://download.qt.io/official_releases/qt/6.7/6.7.0/submodules/qtbase-everywhere-src-6.7.0.tar.xz"
	qtmariadb_cmd_src := exec.Command("curl", "-L", qtmariadb_src_url, "-o", "source.tar.gz")
	err = qtmariadb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	qtmariadb_cmd_direct := exec.Command("./binary")
	err = qtmariadb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: mariadb")
exec.Command("latte", "install", "mariadb")
	fmt.Println("Instalando dependencia: qt")
exec.Command("latte", "install", "qt")
}
