package main

const template_rest_main = `package main

import (
    "gopkg.in/alecthomas/kingpin.v2"
    "os"
)

var cliApp {{.Yaml.Name}}App

func main() {
    cliApp.init()

    switch kingpin.MustParse(cliApp.App.Parse(os.Args[1:])) {
    case "service start":
        cliApp.start_server()
    default:
        kingpin.Usage()
    }
}
`
