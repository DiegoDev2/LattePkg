
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// nfdumpFormula represents a formula in Go.
type nfdumpFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg nfdumpFormula) Print() {
    fmt.Printf("Name: nfdump\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := nfdumpFormula{
        Description:  "Tools to collect and process netflow data on the command-line",
        Homepage:     "https://github.com/phaag/nfdump",
        URL:          "https://github.com/phaag/nfdump/archive/refs/tags/v1.7.4.tar.gz",
        Sha256:       "64b660aaaa2bea081d802e32c32af4d2869b8a421cefb000467758b106f20847",
        Dependencies: []string{"autoconf", "automake", "libtool", "pkg-config"},
    }

    pkg.Print()

    // Instalar dependencias si no están instaladas
    for _, dep := range pkg.Dependencies {
        if !isDependencyInstalled(dep) {
            fmt.Printf("🛠️ Dependency %s not found. Installing...
", dep)
            cmd := exec.Command("brew", "install", dep)
            if err := cmd.Run(); err != nil {
                log.Fatalf("Error installing dependency %s: %v", dep, err)
            }
        } else {
            fmt.Printf("✅ Dependency %s is already installed.
", dep)
        }
    }

    if err := pkg.Installnfdump(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg nfdumpFormula) Installnfdump() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "v1.7.4.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "v1.7.4.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}

func isDependencyInstalled(dep string) bool {
    cmd := exec.Command("brew", "list", dep)
    output, err := cmd.CombinedOutput()
    return err == nil && strings.TrimSpace(string(output)) != ""
}
