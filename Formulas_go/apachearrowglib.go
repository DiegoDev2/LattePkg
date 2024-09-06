package main

// ApacheArrowGlib - GLib bindings for Apache Arrow
// Homepage: https://arrow.apache.org/

import (
	"fmt"
	
	"os/exec"
)

func installApacheArrowGlib() {
	// Método 1: Descargar y extraer .tar.gz
	apachearrowglib_tar_url := "https://www.apache.org/dyn/closer.lua?path=arrow/arrow-17.0.0/apache-arrow-17.0.0.tar.gz"
	apachearrowglib_cmd_tar := exec.Command("curl", "-L", apachearrowglib_tar_url, "-o", "package.tar.gz")
	err := apachearrowglib_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	apachearrowglib_zip_url := "https://www.apache.org/dyn/closer.lua?path=arrow/arrow-17.0.0/apache-arrow-17.0.0.zip"
	apachearrowglib_cmd_zip := exec.Command("curl", "-L", apachearrowglib_zip_url, "-o", "package.zip")
	err = apachearrowglib_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	apachearrowglib_bin_url := "https://www.apache.org/dyn/closer.lua?path=arrow/arrow-17.0.0/apache-arrow-17.0.0.bin"
	apachearrowglib_cmd_bin := exec.Command("curl", "-L", apachearrowglib_bin_url, "-o", "binary.bin")
	err = apachearrowglib_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	apachearrowglib_src_url := "https://www.apache.org/dyn/closer.lua?path=arrow/arrow-17.0.0/apache-arrow-17.0.0.src.tar.gz"
	apachearrowglib_cmd_src := exec.Command("curl", "-L", apachearrowglib_src_url, "-o", "source.tar.gz")
	err = apachearrowglib_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	apachearrowglib_cmd_direct := exec.Command("./binary")
	err = apachearrowglib_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gobject-introspection")
exec.Command("latte", "install", "gobject-introspection")
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: apache-arrow")
exec.Command("latte", "install", "apache-arrow")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
}
