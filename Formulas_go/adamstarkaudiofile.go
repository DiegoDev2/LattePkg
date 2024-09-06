package main

// AdamstarkAudiofile - C++ Audio File Library by Adam Stark
// Homepage: https://github.com/adamstark/AudioFile

import (
	"fmt"
	
	"os/exec"
)

func installAdamstarkAudiofile() {
	// Método 1: Descargar y extraer .tar.gz
	adamstarkaudiofile_tar_url := "https://github.com/adamstark/AudioFile/archive/refs/tags/1.1.1.tar.gz"
	adamstarkaudiofile_cmd_tar := exec.Command("curl", "-L", adamstarkaudiofile_tar_url, "-o", "package.tar.gz")
	err := adamstarkaudiofile_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	adamstarkaudiofile_zip_url := "https://github.com/adamstark/AudioFile/archive/refs/tags/1.1.1.zip"
	adamstarkaudiofile_cmd_zip := exec.Command("curl", "-L", adamstarkaudiofile_zip_url, "-o", "package.zip")
	err = adamstarkaudiofile_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	adamstarkaudiofile_bin_url := "https://github.com/adamstark/AudioFile/archive/refs/tags/1.1.1.bin"
	adamstarkaudiofile_cmd_bin := exec.Command("curl", "-L", adamstarkaudiofile_bin_url, "-o", "binary.bin")
	err = adamstarkaudiofile_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	adamstarkaudiofile_src_url := "https://github.com/adamstark/AudioFile/archive/refs/tags/1.1.1.src.tar.gz"
	adamstarkaudiofile_cmd_src := exec.Command("curl", "-L", adamstarkaudiofile_src_url, "-o", "source.tar.gz")
	err = adamstarkaudiofile_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	adamstarkaudiofile_cmd_direct := exec.Command("./binary")
	err = adamstarkaudiofile_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
