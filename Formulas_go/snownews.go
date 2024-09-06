package main

// Snownews - Text mode RSS newsreader
// Homepage: https://sourceforge.net/projects/snownews/

import (
	"fmt"
	
	"os/exec"
)

func installSnownews() {
	// Método 1: Descargar y extraer .tar.gz
	snownews_tar_url := "https://downloads.sourceforge.net/project/snownews/snownews-1.11.tar.gz"
	snownews_cmd_tar := exec.Command("curl", "-L", snownews_tar_url, "-o", "package.tar.gz")
	err := snownews_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	snownews_zip_url := "https://downloads.sourceforge.net/project/snownews/snownews-1.11.zip"
	snownews_cmd_zip := exec.Command("curl", "-L", snownews_zip_url, "-o", "package.zip")
	err = snownews_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	snownews_bin_url := "https://downloads.sourceforge.net/project/snownews/snownews-1.11.bin"
	snownews_cmd_bin := exec.Command("curl", "-L", snownews_bin_url, "-o", "binary.bin")
	err = snownews_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	snownews_src_url := "https://downloads.sourceforge.net/project/snownews/snownews-1.11.src.tar.gz"
	snownews_cmd_src := exec.Command("curl", "-L", snownews_src_url, "-o", "source.tar.gz")
	err = snownews_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	snownews_cmd_direct := exec.Command("./binary")
	err = snownews_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: coreutils")
exec.Command("latte", "install", "coreutils")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: ncurses")
exec.Command("latte", "install", "ncurses")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
