
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// libgedit-gtksourceviewFormula represents a formula in Go.
type libgedit-gtksourceviewFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg libgedit-gtksourceviewFormula) Print() {
    fmt.Printf("Name: libgedit-gtksourceview\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := libgedit-gtksourceviewFormula{
        Description:  "Text editor widget for code editing",
        Homepage:     "https://gitlab.gnome.org/World/gedit/libgedit-gtksourceview",
        URL:          "https://gitlab.gnome.org/World/gedit/libgedit-gtksourceview/-/archive/299.2.1/libgedit-gtksourceview-299.2.1.tar.bz2",
        Sha256:       "528ce89f65e7be3f6184df2d3e64ec16c9e733ac55a5e3722b55088d386c4d80",
        Dependencies: []string{"gettext", "gobject-introspection", "meson", "ninja", "pkg-config", "cairo", "gdk-pixbuf", "glib", "gtk+3", "libxml2", "pango", "gettext"},
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

    if err := pkg.Installlibgedit-gtksourceview(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg libgedit-gtksourceviewFormula) Installlibgedit-gtksourceview() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "libgedit-gtksourceview-299.2.1.tar.bz2"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "libgedit-gtksourceview-299.2.1.tar"
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
