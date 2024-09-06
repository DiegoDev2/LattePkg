package main

// PhpAT81 - General-purpose scripting language
// Homepage: https://www.php.net/

import (
	"fmt"
	
	"os/exec"
)

func installPhpAT81() {
	// Método 1: Descargar y extraer .tar.gz
	phpat81_tar_url := "https://www.php.net/distributions/php-8.1.29.tar.xz"
	phpat81_cmd_tar := exec.Command("curl", "-L", phpat81_tar_url, "-o", "package.tar.gz")
	err := phpat81_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	phpat81_zip_url := "https://www.php.net/distributions/php-8.1.29.tar.xz"
	phpat81_cmd_zip := exec.Command("curl", "-L", phpat81_zip_url, "-o", "package.zip")
	err = phpat81_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	phpat81_bin_url := "https://www.php.net/distributions/php-8.1.29.tar.xz"
	phpat81_cmd_bin := exec.Command("curl", "-L", phpat81_bin_url, "-o", "binary.bin")
	err = phpat81_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	phpat81_src_url := "https://www.php.net/distributions/php-8.1.29.tar.xz"
	phpat81_cmd_src := exec.Command("curl", "-L", phpat81_src_url, "-o", "source.tar.gz")
	err = phpat81_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	phpat81_cmd_direct := exec.Command("./binary")
	err = phpat81_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: httpd")
exec.Command("latte", "install", "httpd")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: apr")
exec.Command("latte", "install", "apr")
	fmt.Println("Instalando dependencia: apr-util")
exec.Command("latte", "install", "apr-util")
	fmt.Println("Instalando dependencia: argon2")
exec.Command("latte", "install", "argon2")
	fmt.Println("Instalando dependencia: aspell")
exec.Command("latte", "install", "aspell")
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: curl")
exec.Command("latte", "install", "curl")
	fmt.Println("Instalando dependencia: freetds")
exec.Command("latte", "install", "freetds")
	fmt.Println("Instalando dependencia: gd")
exec.Command("latte", "install", "gd")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
	fmt.Println("Instalando dependencia: icu4c")
exec.Command("latte", "install", "icu4c")
	fmt.Println("Instalando dependencia: krb5")
exec.Command("latte", "install", "krb5")
	fmt.Println("Instalando dependencia: libpq")
exec.Command("latte", "install", "libpq")
	fmt.Println("Instalando dependencia: libsodium")
exec.Command("latte", "install", "libsodium")
	fmt.Println("Instalando dependencia: libzip")
exec.Command("latte", "install", "libzip")
	fmt.Println("Instalando dependencia: oniguruma")
exec.Command("latte", "install", "oniguruma")
	fmt.Println("Instalando dependencia: openldap")
exec.Command("latte", "install", "openldap")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: pcre2")
exec.Command("latte", "install", "pcre2")
	fmt.Println("Instalando dependencia: sqlite")
exec.Command("latte", "install", "sqlite")
	fmt.Println("Instalando dependencia: tidy-html5")
exec.Command("latte", "install", "tidy-html5")
	fmt.Println("Instalando dependencia: unixodbc")
exec.Command("latte", "install", "unixodbc")
}
