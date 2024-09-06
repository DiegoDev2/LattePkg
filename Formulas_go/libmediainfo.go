package main

// Libmediainfo - Shared library for mediainfo
// Homepage: https://mediaarea.net/en/MediaInfo

import (
	"fmt"
	
	"os/exec"
)

func installLibmediainfo() {
	// Método 1: Descargar y extraer .tar.gz
	libmediainfo_tar_url := "https://mediaarea.net/download/source/libmediainfo/24.06/libmediainfo_24.06.tar.xz"
	libmediainfo_cmd_tar := exec.Command("curl", "-L", libmediainfo_tar_url, "-o", "package.tar.gz")
	err := libmediainfo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libmediainfo_zip_url := "https://mediaarea.net/download/source/libmediainfo/24.06/libmediainfo_24.06.tar.xz"
	libmediainfo_cmd_zip := exec.Command("curl", "-L", libmediainfo_zip_url, "-o", "package.zip")
	err = libmediainfo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libmediainfo_bin_url := "https://mediaarea.net/download/source/libmediainfo/24.06/libmediainfo_24.06.tar.xz"
	libmediainfo_cmd_bin := exec.Command("curl", "-L", libmediainfo_bin_url, "-o", "binary.bin")
	err = libmediainfo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libmediainfo_src_url := "https://mediaarea.net/download/source/libmediainfo/24.06/libmediainfo_24.06.tar.xz"
	libmediainfo_cmd_src := exec.Command("curl", "-L", libmediainfo_src_url, "-o", "source.tar.gz")
	err = libmediainfo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libmediainfo_cmd_direct := exec.Command("./binary")
	err = libmediainfo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libmms")
exec.Command("latte", "install", "libmms")
	fmt.Println("Instalando dependencia: libzen")
exec.Command("latte", "install", "libzen")
}
