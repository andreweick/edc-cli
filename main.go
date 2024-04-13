package main

import (
	"edc-cli/cmd"
	"fmt"
	"github.com/alecthomas/kong"
	"github.com/caarlos0/env/v10"
)

var CLI struct {
	Photo struct {
		Force     bool `help:"Force update."`
		Recursive bool `help:"Recursively upload photos."`

		Paths []string `arg:"" name:"path" help:"Photos to sync." type:"path"`
	} `cmd:"" help:"Synchronize Photographs to Eick.com."`

	Sync struct {
		Force     bool `help:"Force removal."`
		Recursive bool `help:"Recursively remove files."`

		Paths []string `arg:"" name:"path" help:"Files to sync." type:"path"`
	} `cmd:"" help:"Synchronize files to Eick.com."`

	Cookbook struct {
		Force bool `help:"Force update."`
	} `cmd:"" help:"Synchronize recipes from https://cookbook.eick.com."`

	Bookmark struct {
		Force bool `help:"Force update."`
	} `cmd:"" help:"Synchronize bookmarks."`
}

type config struct {
	ConnStr        string `env:"CONN_STR"`
	CookbookApiKey string `env:"COOKBOOK_API_KEY"`
}

func main() {
	// Get environment variables
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	ctx := kong.Parse(&CLI)
	switch ctx.Command() {
	case "photo <path>":
		// photo
	case "cookbook":
		err := cmd.SyncCookbook(cfg.CookbookApiKey, cfg.ConnStr)

		if err != nil {
			fmt.Printf("Problems syncing cookbook %s\n", err)
		}

	default:
		panic(ctx.Command())
	}
}
