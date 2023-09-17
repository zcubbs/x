// Package k3s
/*
Copyright © 2023 zcubbs https://github.com/zcubbs
*/
package k3s

import (
	"fmt"
	"github.com/zcubbs/x/bash"
	osx "github.com/zcubbs/x/os"
	"os"
	"text/template"
)

const InstallScript = "/tmp/k3s-install.sh"
const UninstallScript = "/usr/local/bin/k3s-uninstall.sh"
const ConfigFileLocation = "/etc/rancher/k3s"

type Config struct {
	Disable                 []string
	TlsSan                  []string
	DataDir                 string
	DefaultLocalStoragePath string
	WriteKubeconfigMode     string
}

var configTmpl = `---
{{- if .Disable }}
disable: 
{{- range $val := $.Disable }}
  - {{ $val }}
{{- end }}
{{- end }}
{{- if .DefaultLocalStoragePath }}
default-local-storage-path: {{ .DefaultLocalStoragePath }}
{{- end }}
{{- if .TlsSan }}
tls-san:
{{- range $val := $.TlsSan }}
  - {{ $val }}
{{- end }}
{{- end }}
{{- if .DataDir }}
data-dir: {{ .DataDir }}
{{- end }}
{{- if .WriteKubeconfigMode }}
write-kubeconfig-mode: {{ .WriteKubeconfigMode }}
{{- end }}
`

func Install(config Config, debug bool) error {
	if debug {
		fmt.Printf("%+v\n", config)
	}

	// prepare config file
	err := osx.CreateDirIfNotExist(ConfigFileLocation)
	if err != nil {
		return err
	}
	targetFile := fmt.Sprintf("%s/%s", ConfigFileLocation, "config.yaml")
	err = WriteTemplateToFile(configTmpl, config, targetFile)
	if err != nil {
		return err
	}

	//err = PrintFileContents(targetFile)
	//if err != nil {
	//	log.Fatal(err)
	//}

	// curl -sfL https://get.k3s.io -o k3s-install.sh
	err = bash.ExecuteCmd(
		"curl",
		debug,
		"https://get.k3s.io",
		"-o",
		InstallScript,
	)
	if err != nil {
		return fmt.Errorf("error while executing %s \n%v",
			"curl https://get.k3s.io -o k3s-install.sh",
			err,
		)
	}

	// sh ./k3s-install.sh -s - --write-kubeconfig-mode 644
	err = os.Chmod("/tmp/k3s-install.sh", 0700)
	if err != nil {
		return fmt.Errorf("error while running %s \n%v",
			"chmod 0700 ./tmp/k3s-install.sh -s - --write-kubeconfig-mode 644",
			err)
	}

	ok, err := bash.ExecuteScript(
		InstallScript,
		debug,
		InstallScript,
		"-s",
		"-",
		"--write-kubeconfig-mode=644",
	)
	if !ok && err != nil {
		return fmt.Errorf("error while running %s \n%v",
			"./k3s-install.sh -s - --write-kubeconfig-mode 644",
			err)
	}

	return nil
}

func WriteTemplateToFile(templateStr string, config Config, outputFilePath string) error {
	// Create a new template and parse the letter into it.
	tmpl, err := template.New("myTemplate").Parse(templateStr)
	if err != nil {
		return err
	}

	// Open the output file
	f, err := os.OpenFile(outputFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)

	// Apply the template to the config data and write to the file
	err = tmpl.Execute(f, config)
	if err != nil {
		return err
	}

	return nil
}

func Uninstall(debug bool) error {
	_, err := bash.ExecuteScript(
		UninstallScript,
		debug,
		UninstallScript,
	)
	if err != nil {
		return err
	}

	return nil
}
