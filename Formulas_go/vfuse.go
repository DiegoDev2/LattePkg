package main

// Vfuse - Convert bootable DMG images for use in VMware Fusion
// Homepage: https://github.com/chilcote/vfuse

import (
	"fmt"
	
	"os/exec"
)

func installVfuse() {
	// Método 1: Descargar y extraer .tar.gz
	vfuse_tar_url := "https://github.com/chilcote/vfuse/archive/refs/tags/2.2.6.tar.gz"
	vfuse_cmd_tar := exec.Command("curl", "-L", vfuse_tar_url, "-o", "package.tar.gz")
	err := vfuse_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vfuse_zip_url := "https://github.com/chilcote/vfuse/archive/refs/tags/2.2.6.zip"
	vfuse_cmd_zip := exec.Command("curl", "-L", vfuse_zip_url, "-o", "package.zip")
	err = vfuse_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vfuse_bin_url := "https://github.com/chilcote/vfuse/archive/refs/tags/2.2.6.bin"
	vfuse_cmd_bin := exec.Command("curl", "-L", vfuse_bin_url, "-o", "binary.bin")
	err = vfuse_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vfuse_src_url := "https://github.com/chilcote/vfuse/archive/refs/tags/2.2.6.src.tar.gz"
	vfuse_cmd_src := exec.Command("curl", "-L", vfuse_src_url, "-o", "source.tar.gz")
	err = vfuse_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vfuse_cmd_direct := exec.Command("./binary")
	err = vfuse_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
