package main

// DiffPdf - Visually compare two PDF files
// Homepage: https://vslavik.github.io/diff-pdf/

import (
	"fmt"
	
	"os/exec"
)

func installDiffPdf() {
	// Método 1: Descargar y extraer .tar.gz
	diffpdf_tar_url := "https://github.com/vslavik/diff-pdf/releases/download/v0.5.2/diff-pdf-0.5.2.tar.gz"
	diffpdf_cmd_tar := exec.Command("curl", "-L", diffpdf_tar_url, "-o", "package.tar.gz")
	err := diffpdf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	diffpdf_zip_url := "https://github.com/vslavik/diff-pdf/releases/download/v0.5.2/diff-pdf-0.5.2.zip"
	diffpdf_cmd_zip := exec.Command("curl", "-L", diffpdf_zip_url, "-o", "package.zip")
	err = diffpdf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	diffpdf_bin_url := "https://github.com/vslavik/diff-pdf/releases/download/v0.5.2/diff-pdf-0.5.2.bin"
	diffpdf_cmd_bin := exec.Command("curl", "-L", diffpdf_bin_url, "-o", "binary.bin")
	err = diffpdf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	diffpdf_src_url := "https://github.com/vslavik/diff-pdf/releases/download/v0.5.2/diff-pdf-0.5.2.src.tar.gz"
	diffpdf_cmd_src := exec.Command("curl", "-L", diffpdf_src_url, "-o", "source.tar.gz")
	err = diffpdf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	diffpdf_cmd_direct := exec.Command("./binary")
	err = diffpdf_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: poppler")
exec.Command("latte", "install", "poppler")
	fmt.Println("Instalando dependencia: wxwidgets")
exec.Command("latte", "install", "wxwidgets")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
