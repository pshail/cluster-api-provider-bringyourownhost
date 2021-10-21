package algo

import (
	"fmt"
	"path/filepath"
	"strings"
)

type AptStep struct {
	ShellStep
}

func (a *AptStep) create(k *BaseK8sInstaller, aptPkg string) Step {
	pkgName := strings.Split(aptPkg, ".")[0]
	pkgAbsolutePath := filepath.Join(k.BundlePath, aptPkg)
	return &ShellStep{
		BaseK8sInstaller: k,
		Desc:             pkgName,
		DoCmd:            fmt.Sprintf("apt install -y '%s'", pkgAbsolutePath),
		UndoCmd:          fmt.Sprintf("apt remove -y %s", pkgName)}
}
