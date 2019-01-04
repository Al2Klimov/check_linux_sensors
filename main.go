//go:generate go run vendor/github.com/Al2Klimov/go-gen-source-repos/main.go github.com/Al2Klimov/check_linux_sensors

package main

import (
	"fmt"
	_ "github.com/Al2Klimov/go-gen-source-repos"
	. "github.com/Al2Klimov/go-monplug-utils"
	"os"
	"strings"
)

func main() {
	os.Exit(ExecuteCheck(onTerminal, checkLinuxSensors))
}

func onTerminal() (output string) {
	return fmt.Sprintf(
		"For the terms of use, the source code and the authors\n"+
			"see the projects this program is assembled from:\n\n  %s\n",
		strings.Join(GithubcomAl2klimovGo_gen_source_repos, "\n  "),
	)
}

func checkLinuxSensors() (output string, perfdata PerfdataCollection, errs map[string]error) {
	return "", nil, nil
}
