package main

import (
    "github.com/turbot/steampipe-plugin-sdk/plugin"
    "steampipe-plugin-uptimerobot/uptimerobot"
)

func main() {
    plugin.Serve(&plugin.ServeOpts{PluginFunc: uptimerobot.Plugin})
}
