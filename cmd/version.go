package cmd

import (
	"github.com/spf13/cobra"
	"runtime"
	"os"
	"text/template"
	"fmt"
)

var (
	Version = "0.2"
	Codename = "Happy Gabian"
	BuildDate = "I don't remember exactly"
)

var versionTemplate = `
Version:         {{.Version}}
Codename:        {{.Codename}}
Go version:      {{.GoVersion}}
BuildTime:    	 {{.BuildTime}}
OS/Arch:      	 {{.Os}}/{{.Arch}}
Compiled on:     {{.Hostname}}
`

var BuildTime string

func InitVersion() *cobra.Command {
	//version Command init
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "version",
		Long:  `dockit version`,
		Run: func(cmd *cobra.Command, args []string) {
			tmpl, err := template.New("").Parse(versionTemplate)
			if err == nil {
				v := struct {
					Version   string
					Codename  string
					GoVersion string
					BuildTime string
					Os        string
					Arch      string
					Hostname  string
				}{
					Version:   Version,
					Codename:  Codename,
					GoVersion: runtime.Version(),
					BuildTime: BuildTime,
					Os:        runtime.GOOS,
					Arch:      runtime.GOARCH,
					Hostname:  getHostname(os.Hostname()),
				}
				if err := tmpl.Execute(os.Stdout, v); err != nil {
				}
				fmt.Printf("\n")
			}
		},
	}
	return versionCmd
}


func getHostname(name string, err error) string {
	if err != nil {
		return "undef hostname"
	}
	return name
}