package main

// Ptpython - Advanced Python REPL
// Homepage: https://github.com/prompt-toolkit/ptpython

import (
	"fmt"
	
	"os/exec"
)

func installPtpython() {
	// Método 1: Descargar y extraer .tar.gz
	ptpython_tar_url := "https://files.pythonhosted.org/packages/56/61/352792c9f47de98a910526ff8a684466a6217e53ffa6627b3781960e4f0d/ptpython-3.0.29.tar.gz"
	ptpython_cmd_tar := exec.Command("curl", "-L", ptpython_tar_url, "-o", "package.tar.gz")
	err := ptpython_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ptpython_zip_url := "https://files.pythonhosted.org/packages/56/61/352792c9f47de98a910526ff8a684466a6217e53ffa6627b3781960e4f0d/ptpython-3.0.29.zip"
	ptpython_cmd_zip := exec.Command("curl", "-L", ptpython_zip_url, "-o", "package.zip")
	err = ptpython_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ptpython_bin_url := "https://files.pythonhosted.org/packages/56/61/352792c9f47de98a910526ff8a684466a6217e53ffa6627b3781960e4f0d/ptpython-3.0.29.bin"
	ptpython_cmd_bin := exec.Command("curl", "-L", ptpython_bin_url, "-o", "binary.bin")
	err = ptpython_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ptpython_src_url := "https://files.pythonhosted.org/packages/56/61/352792c9f47de98a910526ff8a684466a6217e53ffa6627b3781960e4f0d/ptpython-3.0.29.src.tar.gz"
	ptpython_cmd_src := exec.Command("curl", "-L", ptpython_src_url, "-o", "source.tar.gz")
	err = ptpython_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ptpython_cmd_direct := exec.Command("./binary")
	err = ptpython_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
