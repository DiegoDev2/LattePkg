package main

// Wtfis - Passive hostname, domain, and IP lookup tool
// Homepage: https://github.com/pirxthepilot/wtfis

import (
	"fmt"
	
	"os/exec"
)

func installWtfis() {
	// Método 1: Descargar y extraer .tar.gz
	wtfis_tar_url := "https://files.pythonhosted.org/packages/c8/8e/d8005d43adcfd263723de2ee643eb2a6a48e645a9338202c7eda800a0df5/wtfis-0.10.0.tar.gz"
	wtfis_cmd_tar := exec.Command("curl", "-L", wtfis_tar_url, "-o", "package.tar.gz")
	err := wtfis_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wtfis_zip_url := "https://files.pythonhosted.org/packages/c8/8e/d8005d43adcfd263723de2ee643eb2a6a48e645a9338202c7eda800a0df5/wtfis-0.10.0.zip"
	wtfis_cmd_zip := exec.Command("curl", "-L", wtfis_zip_url, "-o", "package.zip")
	err = wtfis_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wtfis_bin_url := "https://files.pythonhosted.org/packages/c8/8e/d8005d43adcfd263723de2ee643eb2a6a48e645a9338202c7eda800a0df5/wtfis-0.10.0.bin"
	wtfis_cmd_bin := exec.Command("curl", "-L", wtfis_bin_url, "-o", "binary.bin")
	err = wtfis_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wtfis_src_url := "https://files.pythonhosted.org/packages/c8/8e/d8005d43adcfd263723de2ee643eb2a6a48e645a9338202c7eda800a0df5/wtfis-0.10.0.src.tar.gz"
	wtfis_cmd_src := exec.Command("curl", "-L", wtfis_src_url, "-o", "source.tar.gz")
	err = wtfis_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wtfis_cmd_direct := exec.Command("./binary")
	err = wtfis_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
