package main

const template_shell_main = `package main

import (
    "fmt"
)

// Adapt this source to create the code for each actions


func (*{{.Name}}App) create() {
    fmt.Printf("create")
}

func (*{{.Name}}App) update() {
    fmt.Printf("update")
}

func (*{{.Name}}App) maintain() {
    fmt.Printf("maintain")
}

`

const template_shell_app=`package main

import (
  "gopkg.in/alecthomas/kingpin.v2"
  "github.hpe.com/christophe-larsonneur/goforjj"
  "os"
  "fmt"
)

type {{.Yaml.Name}}App struct {
    goforjj.ForjjPluginApp
}

var cliApp {{.Yaml.Name}}App

func (a *{{.Yaml.Name}}App)init() {
  a.App = kingpin.New("{{.Yaml.Name}}", "{{.Yaml.Description}}")
{{ if .Yaml.Version }}\
  a.App.Version("{{ .Yaml.Version }}")
{{ end }}\

  // true to create the Infra
  a.IsInfra = a.App.Flag("infra", "Used by upstream plugins to create initial repos.").Hidden().Bool()
  a.Flags = make(map[string]*string)
  a.Tasks = make(map[string]goforjj.PluginTask)
{{ range $Flagname, $Opts := .Yaml.Actions.common.Flags }}\
  a.Flags["{{ $Flagname }}"] = a.App.Flag("{{ $Flagname }}", "{{ $Opts.help }}")\
{{   if $Opts.Required }}.Required(){{ end }}\
{{   if $Opts.Default  }}.Default("{{$Opts.Default}}"){{ end }}\
{{   if $Opts.Hidden"  }}.Hidden(){{ end }}.String()
{{ end }}

{{ range $cmd, $flags := (filter_cmds .Yaml.Actions) }}\
  a.Tasks["{{$cmd}}"] = goforjj.PluginTask {
    Flags: make(map[string]*string),
    Cmd  : a.App.Command("{{ $cmd }}", "{{ $flags.Help }}"),
  }
{{   range $Flagname, $Opts := $flags.Flags }}\
  a.Tasks["{{$cmd}}"].Flags["{{ $Flagname }}"] = a.Tasks["{{$cmd}}"].Cmd.Flag("{{ $Flagname }}", "{{ $Opts.Help }}")\
{{     if $Opts.Required" }}.Required(){{ end }}\
{{     if $Opts.Default   }}.Default("{{$Opts.Default}}"){{ end }}\
{{     if $Opts.Hidden    }}.Hidden(){{ end }}.String()
{{   end }}
{{ end }}\
}

func main() {
  cliApp.init()

  switch kingpin.MustParse(cliApp.App.Parse(os.Args[1:])) {
{{ range $cmd, $flags := (filter_cmds .Yaml.Actions) }}\
    case "{{$cmd}}":
       cliApp.{{$cmd}}()
{{ end }}\
  }
}

const YamlDesc="{{ escape .YamlData}}"

`