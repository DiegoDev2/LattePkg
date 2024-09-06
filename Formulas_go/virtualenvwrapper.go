package main

// Virtualenvwrapper - Python virtualenv extensions
// Homepage: https://virtualenvwrapper.readthedocs.io/

import (
	"fmt"
	
	"os/exec"
)

func installVirtualenvwrapper() {
	// Método 1: Descargar y extraer .tar.gz
	virtualenvwrapper_tar_url := "https://files.pythonhosted.org/packages/4e/00/71629f631867c434e49fa276d659894c7cb5715f5de2233c10bc47c1fcc6/virtualenvwrapper-6.1.0.tar.gz"
	virtualenvwrapper_cmd_tar := exec.Command("curl", "-L", virtualenvwrapper_tar_url, "-o", "package.tar.gz")
	err := virtualenvwrapper_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	virtualenvwrapper_zip_url := "https://files.pythonhosted.org/packages/4e/00/71629f631867c434e49fa276d659894c7cb5715f5de2233c10bc47c1fcc6/virtualenvwrapper-6.1.0.zip"
	virtualenvwrapper_cmd_zip := exec.Command("curl", "-L", virtualenvwrapper_zip_url, "-o", "package.zip")
	err = virtualenvwrapper_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	virtualenvwrapper_bin_url := "https://files.pythonhosted.org/packages/4e/00/71629f631867c434e49fa276d659894c7cb5715f5de2233c10bc47c1fcc6/virtualenvwrapper-6.1.0.bin"
	virtualenvwrapper_cmd_bin := exec.Command("curl", "-L", virtualenvwrapper_bin_url, "-o", "binary.bin")
	err = virtualenvwrapper_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	virtualenvwrapper_src_url := "https://files.pythonhosted.org/packages/4e/00/71629f631867c434e49fa276d659894c7cb5715f5de2233c10bc47c1fcc6/virtualenvwrapper-6.1.0.src.tar.gz"
	virtualenvwrapper_cmd_src := exec.Command("curl", "-L", virtualenvwrapper_src_url, "-o", "source.tar.gz")
	err = virtualenvwrapper_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	virtualenvwrapper_cmd_direct := exec.Command("./binary")
	err = virtualenvwrapper_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
