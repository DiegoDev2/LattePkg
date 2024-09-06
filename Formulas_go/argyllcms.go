package main

// ArgyllCms - ICC compatible color management system
// Homepage: https://www.argyllcms.com/

import (
	"fmt"
	
	"os/exec"
)

func installArgyllCms() {
	// Método 1: Descargar y extraer .tar.gz
	argyllcms_tar_url := "https://www.argyllcms.com/Argyll_V3.2.0_src.zip"
	argyllcms_cmd_tar := exec.Command("curl", "-L", argyllcms_tar_url, "-o", "package.tar.gz")
	err := argyllcms_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	argyllcms_zip_url := "https://www.argyllcms.com/Argyll_V3.2.0_src.zip"
	argyllcms_cmd_zip := exec.Command("curl", "-L", argyllcms_zip_url, "-o", "package.zip")
	err = argyllcms_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	argyllcms_bin_url := "https://www.argyllcms.com/Argyll_V3.2.0_src.zip"
	argyllcms_cmd_bin := exec.Command("curl", "-L", argyllcms_bin_url, "-o", "binary.bin")
	err = argyllcms_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	argyllcms_src_url := "https://www.argyllcms.com/Argyll_V3.2.0_src.zip"
	argyllcms_cmd_src := exec.Command("curl", "-L", argyllcms_src_url, "-o", "source.tar.gz")
	err = argyllcms_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	argyllcms_cmd_direct := exec.Command("./binary")
	err = argyllcms_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libtiff")
exec.Command("latte", "install", "libtiff")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
	fmt.Println("Instalando dependencia: libxext")
exec.Command("latte", "install", "libxext")
	fmt.Println("Instalando dependencia: libxinerama")
exec.Command("latte", "install", "libxinerama")
	fmt.Println("Instalando dependencia: libxrandr")
exec.Command("latte", "install", "libxrandr")
	fmt.Println("Instalando dependencia: libxscrnsaver")
exec.Command("latte", "install", "libxscrnsaver")
	fmt.Println("Instalando dependencia: libxxf86vm")
exec.Command("latte", "install", "libxxf86vm")
	fmt.Println("Instalando dependencia: xorgproto")
exec.Command("latte", "install", "xorgproto")
}
