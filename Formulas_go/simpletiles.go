package main

// SimpleTiles - Image generation library for spatial data
// Homepage: https://github.com/propublica/simple-tiles

import (
	"fmt"
	
	"os/exec"
)

func installSimpleTiles() {
	// Método 1: Descargar y extraer .tar.gz
	simpletiles_tar_url := "https://github.com/propublica/simple-tiles/archive/refs/tags/v0.6.2.tar.gz"
	simpletiles_cmd_tar := exec.Command("curl", "-L", simpletiles_tar_url, "-o", "package.tar.gz")
	err := simpletiles_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	simpletiles_zip_url := "https://github.com/propublica/simple-tiles/archive/refs/tags/v0.6.2.zip"
	simpletiles_cmd_zip := exec.Command("curl", "-L", simpletiles_zip_url, "-o", "package.zip")
	err = simpletiles_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	simpletiles_bin_url := "https://github.com/propublica/simple-tiles/archive/refs/tags/v0.6.2.bin"
	simpletiles_cmd_bin := exec.Command("curl", "-L", simpletiles_bin_url, "-o", "binary.bin")
	err = simpletiles_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	simpletiles_src_url := "https://github.com/propublica/simple-tiles/archive/refs/tags/v0.6.2.src.tar.gz"
	simpletiles_cmd_src := exec.Command("curl", "-L", simpletiles_src_url, "-o", "source.tar.gz")
	err = simpletiles_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	simpletiles_cmd_direct := exec.Command("./binary")
	err = simpletiles_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: gdal")
exec.Command("latte", "install", "gdal")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: harfbuzz")
exec.Command("latte", "install", "harfbuzz")
}
