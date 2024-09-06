package main

// Wallpaper - Manage the desktop wallpaper
// Homepage: https://github.com/sindresorhus/macos-wallpaper

import (
	"fmt"
	
	"os/exec"
)

func installWallpaper() {
	// Método 1: Descargar y extraer .tar.gz
	wallpaper_tar_url := "https://github.com/sindresorhus/macos-wallpaper/archive/refs/tags/v2.3.1.tar.gz"
	wallpaper_cmd_tar := exec.Command("curl", "-L", wallpaper_tar_url, "-o", "package.tar.gz")
	err := wallpaper_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wallpaper_zip_url := "https://github.com/sindresorhus/macos-wallpaper/archive/refs/tags/v2.3.1.zip"
	wallpaper_cmd_zip := exec.Command("curl", "-L", wallpaper_zip_url, "-o", "package.zip")
	err = wallpaper_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wallpaper_bin_url := "https://github.com/sindresorhus/macos-wallpaper/archive/refs/tags/v2.3.1.bin"
	wallpaper_cmd_bin := exec.Command("curl", "-L", wallpaper_bin_url, "-o", "binary.bin")
	err = wallpaper_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wallpaper_src_url := "https://github.com/sindresorhus/macos-wallpaper/archive/refs/tags/v2.3.1.src.tar.gz"
	wallpaper_cmd_src := exec.Command("curl", "-L", wallpaper_src_url, "-o", "source.tar.gz")
	err = wallpaper_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wallpaper_cmd_direct := exec.Command("./binary")
	err = wallpaper_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
