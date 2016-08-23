package main

const template_rest_routes = `package main

import (
    "github.hpe.com/christophe-larsonneur/goforjj"
)

var routes = Routes{
    Route{ "Index",    "GET",  "/",         Index               },
    Route{ "Quit",     "GET",  "/quit",     Quit                },
    Route{ "Ping",     "GET",  "/ping",     goforjj.PingHandler },
    Route{ "Create",   "POST", "/create",   Create              },
    Route{ "Update",   "POST", "/update",   Update              },
    Route{ "Maintain", "POST", "/maintain", Maintain            },
}
`