package main

// ApachePulsar - Cloud-native distributed messaging and streaming platform
// Homepage: https://pulsar.apache.org/

import (
	"fmt"
	
	"os/exec"
)

func installApachePulsar() {
	// Método 1: Descargar y extraer .tar.gz
	apachepulsar_tar_url := "https://www.apache.org/dyn/mirrors/mirrors.cgi?action=download&filename=pulsar/pulsar-3.1.2/apache-pulsar-3.1.2-src.tar.gz"
	apachepulsar_cmd_tar := exec.Command("curl", "-L", apachepulsar_tar_url, "-o", "package.tar.gz")
	err := apachepulsar_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	apachepulsar_zip_url := "https://www.apache.org/dyn/mirrors/mirrors.cgi?action=download&filename=pulsar/pulsar-3.1.2/apache-pulsar-3.1.2-src.zip"
	apachepulsar_cmd_zip := exec.Command("curl", "-L", apachepulsar_zip_url, "-o", "package.zip")
	err = apachepulsar_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	apachepulsar_bin_url := "https://www.apache.org/dyn/mirrors/mirrors.cgi?action=download&filename=pulsar/pulsar-3.1.2/apache-pulsar-3.1.2-src.bin"
	apachepulsar_cmd_bin := exec.Command("curl", "-L", apachepulsar_bin_url, "-o", "binary.bin")
	err = apachepulsar_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	apachepulsar_src_url := "https://www.apache.org/dyn/mirrors/mirrors.cgi?action=download&filename=pulsar/pulsar-3.1.2/apache-pulsar-3.1.2-src.src.tar.gz"
	apachepulsar_cmd_src := exec.Command("curl", "-L", apachepulsar_src_url, "-o", "source.tar.gz")
	err = apachepulsar_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	apachepulsar_cmd_direct := exec.Command("./binary")
	err = apachepulsar_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: cppunit")
exec.Command("latte", "install", "cppunit")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: maven")
exec.Command("latte", "install", "maven")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: protobuf")
exec.Command("latte", "install", "protobuf")
	fmt.Println("Instalando dependencia: openjdk@17")
exec.Command("latte", "install", "openjdk@17")
}
