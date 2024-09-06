package main

// Wolfmqtt - Small, fast, portable MQTT client C implementation
// Homepage: https://github.com/wolfSSL/wolfMQTT

import (
	"fmt"
	
	"os/exec"
)

func installWolfmqtt() {
	// Método 1: Descargar y extraer .tar.gz
	wolfmqtt_tar_url := "https://github.com/wolfSSL/wolfMQTT/archive/refs/tags/v1.19.0.tar.gz"
	wolfmqtt_cmd_tar := exec.Command("curl", "-L", wolfmqtt_tar_url, "-o", "package.tar.gz")
	err := wolfmqtt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wolfmqtt_zip_url := "https://github.com/wolfSSL/wolfMQTT/archive/refs/tags/v1.19.0.zip"
	wolfmqtt_cmd_zip := exec.Command("curl", "-L", wolfmqtt_zip_url, "-o", "package.zip")
	err = wolfmqtt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wolfmqtt_bin_url := "https://github.com/wolfSSL/wolfMQTT/archive/refs/tags/v1.19.0.bin"
	wolfmqtt_cmd_bin := exec.Command("curl", "-L", wolfmqtt_bin_url, "-o", "binary.bin")
	err = wolfmqtt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wolfmqtt_src_url := "https://github.com/wolfSSL/wolfMQTT/archive/refs/tags/v1.19.0.src.tar.gz"
	wolfmqtt_cmd_src := exec.Command("curl", "-L", wolfmqtt_src_url, "-o", "source.tar.gz")
	err = wolfmqtt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wolfmqtt_cmd_direct := exec.Command("./binary")
	err = wolfmqtt_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: wolfssl")
exec.Command("latte", "install", "wolfssl")
}
