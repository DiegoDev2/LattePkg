package main

// Texlive - Free software distribution for the TeX typesetting system
// Homepage: https://www.tug.org/texlive/

import (
	"fmt"
	
	"os/exec"
)

func installTexlive() {
	// Método 1: Descargar y extraer .tar.gz
	texlive_tar_url := "https://ftp.math.utah.edu/pub/tex/historic/systems/texlive/2024/texlive-20240312-source.tar.xz"
	texlive_cmd_tar := exec.Command("curl", "-L", texlive_tar_url, "-o", "package.tar.gz")
	err := texlive_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	texlive_zip_url := "https://ftp.math.utah.edu/pub/tex/historic/systems/texlive/2024/texlive-20240312-source.tar.xz"
	texlive_cmd_zip := exec.Command("curl", "-L", texlive_zip_url, "-o", "package.zip")
	err = texlive_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	texlive_bin_url := "https://ftp.math.utah.edu/pub/tex/historic/systems/texlive/2024/texlive-20240312-source.tar.xz"
	texlive_cmd_bin := exec.Command("curl", "-L", texlive_bin_url, "-o", "binary.bin")
	err = texlive_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	texlive_src_url := "https://ftp.math.utah.edu/pub/tex/historic/systems/texlive/2024/texlive-20240312-source.tar.xz"
	texlive_cmd_src := exec.Command("curl", "-L", texlive_src_url, "-o", "source.tar.gz")
	err = texlive_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	texlive_cmd_direct := exec.Command("./binary")
	err = texlive_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: clisp")
exec.Command("latte", "install", "clisp")
	fmt.Println("Instalando dependencia: fontconfig")
exec.Command("latte", "install", "fontconfig")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: gd")
exec.Command("latte", "install", "gd")
	fmt.Println("Instalando dependencia: ghostscript")
exec.Command("latte", "install", "ghostscript")
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
	fmt.Println("Instalando dependencia: graphite2")
exec.Command("latte", "install", "graphite2")
	fmt.Println("Instalando dependencia: harfbuzz")
exec.Command("latte", "install", "harfbuzz")
	fmt.Println("Instalando dependencia: icu4c")
exec.Command("latte", "install", "icu4c")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
	fmt.Println("Instalando dependencia: libxft")
exec.Command("latte", "install", "libxft")
	fmt.Println("Instalando dependencia: lua")
exec.Command("latte", "install", "lua")
	fmt.Println("Instalando dependencia: luajit")
exec.Command("latte", "install", "luajit")
	fmt.Println("Instalando dependencia: mpfr")
exec.Command("latte", "install", "mpfr")
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: perl")
exec.Command("latte", "install", "perl")
	fmt.Println("Instalando dependencia: pixman")
exec.Command("latte", "install", "pixman")
	fmt.Println("Instalando dependencia: potrace")
exec.Command("latte", "install", "potrace")
	fmt.Println("Instalando dependencia: pstoedit")
exec.Command("latte", "install", "pstoedit")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: libice")
exec.Command("latte", "install", "libice")
	fmt.Println("Instalando dependencia: libnsl")
exec.Command("latte", "install", "libnsl")
	fmt.Println("Instalando dependencia: libsm")
exec.Command("latte", "install", "libsm")
	fmt.Println("Instalando dependencia: libxaw")
exec.Command("latte", "install", "libxaw")
	fmt.Println("Instalando dependencia: libxext")
exec.Command("latte", "install", "libxext")
	fmt.Println("Instalando dependencia: libxi")
exec.Command("latte", "install", "libxi")
	fmt.Println("Instalando dependencia: libxmu")
exec.Command("latte", "install", "libxmu")
	fmt.Println("Instalando dependencia: libxpm")
exec.Command("latte", "install", "libxpm")
	fmt.Println("Instalando dependencia: libxt")
exec.Command("latte", "install", "libxt")
	fmt.Println("Instalando dependencia: mesa")
exec.Command("latte", "install", "mesa")
}
