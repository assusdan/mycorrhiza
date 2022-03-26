//go:generate go run github.com/valyala/quicktemplate/qtc -dir=views
//go:generate go run github.com/valyala/quicktemplate/qtc -dir=tree
//go:generate go run github.com/valyala/quicktemplate/qtc -dir=history
// Command mycorrhiza is a program that runs a mycorrhiza wiki.
package main

import (
	"github.com/bouncepaw/mycorrhiza/hyphae/categories"
	"github.com/bouncepaw/mycorrhiza/migration"
	"log"
	"os"

	"github.com/bouncepaw/mycorrhiza/hyphae/backlinks"

	"github.com/bouncepaw/mycorrhiza/cfg"
	"github.com/bouncepaw/mycorrhiza/files"
	"github.com/bouncepaw/mycorrhiza/history"
	"github.com/bouncepaw/mycorrhiza/hyphae"
	"github.com/bouncepaw/mycorrhiza/shroom"
	"github.com/bouncepaw/mycorrhiza/static"
	"github.com/bouncepaw/mycorrhiza/user"
	"github.com/bouncepaw/mycorrhiza/web"
)

func main() {
	parseCliArgs()

	if err := files.PrepareWikiRoot(); err != nil {
		log.Fatal(err)
	}

	if err := cfg.ReadConfigFile(files.ConfigPath()); err != nil {
		log.Fatal(err)
	}

	log.Println("Running Mycorrhiza Wiki 1.9.0")
	if err := os.Chdir(files.HyphaeDir()); err != nil {
		log.Fatal(err)
	}
	log.Println("Wiki directory is", cfg.WikiDir)
	log.Println("Using Git storage at", files.HyphaeDir())

	// Init the subsystems:
	hyphae.Index(files.HyphaeDir())
	backlinks.IndexBacklinks()
	go backlinks.RunBacklinksConveyor()
	user.InitUserDatabase()
	history.Start()
	history.InitGitRepo()
	migration.MigrateRocketsMaybe()
	shroom.SetHeaderLinks()
	categories.InitCategories()

	// Static files:
	static.InitFS(files.StaticFiles())

	serveHTTP(web.Handler())
}
