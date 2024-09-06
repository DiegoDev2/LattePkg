package main

// Slurm - Yet another network load monitor
// Homepage: https://github.com/mattthias/slurm/wiki/

import (
	"fmt"
	
	"os/exec"
)

func installSlurm() {
	// Método 1: Descargar y extraer .tar.gz
	slurm_tar_url := "https://github.com/mattthias/slurm/archive/refs/tags/upstream/0.4.4.tar.gz"
	slurm_cmd_tar := exec.Command("curl", "-L", slurm_tar_url, "-o", "package.tar.gz")
	err := slurm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	slurm_zip_url := "https://github.com/mattthias/slurm/archive/refs/tags/upstream/0.4.4.zip"
	slurm_cmd_zip := exec.Command("curl", "-L", slurm_zip_url, "-o", "package.zip")
	err = slurm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	slurm_bin_url := "https://github.com/mattthias/slurm/archive/refs/tags/upstream/0.4.4.bin"
	slurm_cmd_bin := exec.Command("curl", "-L", slurm_bin_url, "-o", "binary.bin")
	err = slurm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	slurm_src_url := "https://github.com/mattthias/slurm/archive/refs/tags/upstream/0.4.4.src.tar.gz"
	slurm_cmd_src := exec.Command("curl", "-L", slurm_src_url, "-o", "source.tar.gz")
	err = slurm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	slurm_cmd_direct := exec.Command("./binary")
	err = slurm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
