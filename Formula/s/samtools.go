
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// samtoolsFormula represents a formula in Go.
type samtoolsFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg samtoolsFormula) Print() {
    fmt.Printf("Name: samtools\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := samtoolsFormula{
        Description:  "Tools for manipulating next-generation sequencing data",
        Homepage:     "https://www.htslib.org/",
        URL:          "https://github.com/samtools/samtools/releases/download/1.20/samtools-1.20.tar.bz2",
        Sha256:       "1390d215456d06f7de7aa10f32d34eedb1dd41d86e47309450e7cca1facb8d2a",
        Dependencies: []string{"htslib"},
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

    if err := pkg.Installsamtools(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg samtoolsFormula) Installsamtools() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "samtools-1.20.tar.bz2"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "samtools-1.20.tar"
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
