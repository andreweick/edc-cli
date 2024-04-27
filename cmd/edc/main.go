package main

import (
	"edc-cli/internal/bookmark"
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
		Force bool `help:"Force removal."`
	} `cmd:"" help:"Synchronize to Eick.com."`

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
		fmt.Printf("Cannot parse config: %+v\n", err)
	}

	ctx := kong.Parse(&CLI)
	switch ctx.Command() {
	case "photo <path>":
		// photo
	case "cookbook":
		//err := cmd.SyncCookbook(cfg.CookbookApiKey, cfg.ConnStr)
		//
		//if err != nil {
		//	fmt.Printf("Problems syncing cookbook %s\n", err)
		//}
	case "sync":
		err := bookmark.Sync()
		if err != nil {
			fmt.Printf("Problems syncing bookmarks: %s\n", err)
		}
	default:
		panic(ctx.Command())
	}
}
