package web

import (
	"fmt"
	"github.com/bouncepaw/mycomarkup/v4"
	"github.com/bouncepaw/mycorrhiza/shroom"
	"github.com/bouncepaw/mycorrhiza/viewutil"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"

	"github.com/bouncepaw/mycorrhiza/history"
	"github.com/bouncepaw/mycorrhiza/hyphae"
	"github.com/bouncepaw/mycorrhiza/l18n"
	"github.com/bouncepaw/mycorrhiza/mimetype"
	"github.com/bouncepaw/mycorrhiza/user"
	"github.com/bouncepaw/mycorrhiza/util"
	"github.com/bouncepaw/mycorrhiza/views"

	"github.com/bouncepaw/mycomarkup/v4/mycocontext"
	"github.com/bouncepaw/mycomarkup/v4/tools"
)

func initReaders(r *mux.Router) {
	r.PathPrefix("/page/").HandlerFunc(handlerHypha)
	r.PathPrefix("/hypha/").HandlerFunc(handlerHypha)
	r.PathPrefix("/text/").HandlerFunc(handlerText)
	r.PathPrefix("/binary/").HandlerFunc(handlerBinary)
	r.PathPrefix("/rev/").HandlerFunc(handlerRevision)
	r.PathPrefix("/rev-text/").HandlerFunc(handlerRevisionText)
	r.PathPrefix("/media/").HandlerFunc(handlerMedia)
}

func handlerMedia(w http.ResponseWriter, rq *http.Request) {
	util.PrepareRq(rq)
	var (
		hyphaName = util.HyphaNameFromRq(rq, "media")
		h         = hyphae.ByName(hyphaName)
		u         = user.FromRequest(rq)
		lc        = l18n.FromRequest(rq)
	)
	util.HTTP200Page(w,
		views.Base(
			viewutil.MetaFrom(w, rq),
			lc.Get("ui.media_title", &l18n.Replacements{"name": util.BeautifulName(hyphaName)}),
			views.MediaMenu(rq, h, u)))
}

// handlerRevisionText sends Mycomarkup text of the hypha at the given revision. See also: handlerRevision, handlerText.
//
// /rev-text/<revHash>/<hyphaName>
func handlerRevisionText(w http.ResponseWriter, rq *http.Request) {
	util.PrepareRq(rq)
	var (
		shorterURL      = strings.TrimPrefix(rq.URL.Path, "/rev-text/")
		firstSlashIndex = strings.IndexRune(shorterURL, '/')
		revHash         = shorterURL[:firstSlashIndex]
		hyphaName       = util.CanonicalName(shorterURL[firstSlashIndex+1:])
		h               = hyphae.ByName(hyphaName)
	)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	switch h := h.(type) {
	case *hyphae.EmptyHypha:
		log.Printf(`Hypha ‘%s’ does not exist`)
		w.WriteHeader(http.StatusNotFound)
	case hyphae.ExistingHypha:
		if !h.HasTextFile() {
			log.Printf(`Media hypha ‘%s’ has no text`)
			w.WriteHeader(http.StatusNotFound)
		}
		var textContents, err = history.FileAtRevision(h.TextFilePath(), revHash)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			log.Printf("While serving text of ‘%s’ at revision ‘%s’: %s\n", hyphaName, revHash, err.Error())
			_, _ = io.WriteString(w, "Error: "+err.Error())
			return
		}
		log.Printf("Serving text of ‘%s’ from ‘%s’ at revision ‘%s’\n", hyphaName, h.TextFilePath(), revHash)
		w.WriteHeader(http.StatusOK)
		_, _ = io.WriteString(w, textContents)
	}
}

// handlerRevision displays a specific revision of the text part the hypha
func handlerRevision(w http.ResponseWriter, rq *http.Request) {
	util.PrepareRq(rq)
	var (
		lc              = l18n.FromRequest(rq)
		shorterURL      = strings.TrimPrefix(rq.URL.Path, "/rev/")
		firstSlashIndex = strings.IndexRune(shorterURL, '/')
		revHash         = shorterURL[:firstSlashIndex]
		hyphaName       = util.CanonicalName(shorterURL[firstSlashIndex+1:])
		h               = hyphae.ByName(hyphaName)
		contents        = fmt.Sprintf(`<p>%s</p>`, lc.Get("ui.revision_no_text"))
	)
	switch h := h.(type) {
	case hyphae.ExistingHypha:
		var textContents, err = history.FileAtRevision(h.TextFilePath(), revHash)

		if err == nil {
			ctx, _ := mycocontext.ContextFromStringInput(textContents, shroom.MarkupOptions(hyphaName))
			contents = mycomarkup.BlocksToHTML(ctx, mycomarkup.BlockTree(ctx))
		}
	}
	page := views.Revision(
		viewutil.MetaFrom(w, rq),
		h,
		contents,
		revHash,
	)
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(
		w,
		views.Base(
			viewutil.MetaFrom(w, rq),
			lc.Get("ui.revision_title", &l18n.Replacements{"name": util.BeautifulName(hyphaName), "rev": revHash}),
			page,
		),
	)
}

// handlerText serves raw source text of the hypha.
func handlerText(w http.ResponseWriter, rq *http.Request) {
	util.PrepareRq(rq)
	hyphaName := util.HyphaNameFromRq(rq, "text")
	switch h := hyphae.ByName(hyphaName).(type) {
	case *hyphae.TextualHypha:
		log.Println("Serving", h.TextFilePath())
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		http.ServeFile(w, rq, h.TextFilePath())
	}
}

// handlerBinary serves attachment of the hypha.
func handlerBinary(w http.ResponseWriter, rq *http.Request) {
	util.PrepareRq(rq)
	hyphaName := util.HyphaNameFromRq(rq, "binary")
	switch h := hyphae.ByName(hyphaName).(type) {
	case *hyphae.EmptyHypha:
	case *hyphae.TextualHypha:
		w.WriteHeader(http.StatusNotFound)
		log.Printf("Textual hypha ‘%s’ has no media, cannot serve\n", h.CanonicalName())
	case *hyphae.MediaHypha:
		log.Println("Serving", h.MediaFilePath())
		w.Header().Set("Content-Type", mimetype.FromExtension(filepath.Ext(h.MediaFilePath())))
		http.ServeFile(w, rq, h.MediaFilePath())
	}
}

// handlerHypha is the main hypha action that displays the hypha and the binary upload form along with some navigation.
func handlerHypha(w http.ResponseWriter, rq *http.Request) {
	util.PrepareRq(rq)
	var (
		hyphaName = util.HyphaNameFromRq(rq, "page", "hypha")
		h         = hyphae.ByName(hyphaName)
		contents  string
		openGraph string
		lc        = l18n.FromRequest(rq)
	)

	switch h := h.(type) {
	case *hyphae.EmptyHypha:
		util.HTTP404Page(w,
			views.Base(
				viewutil.MetaFrom(w, rq),
				util.BeautifulName(hyphaName),
				views.Hypha(viewutil.MetaFrom(w, rq), h, contents),
				openGraph))
	case hyphae.ExistingHypha:
		fileContentsT, errT := os.ReadFile(h.TextFilePath())
		if errT == nil {
			ctx, _ := mycocontext.ContextFromStringInput(string(fileContentsT), shroom.MarkupOptions(hyphaName))
			getOpenGraph, descVisitor, imgVisitor := tools.OpenGraphVisitors(ctx)
			ast := mycomarkup.BlockTree(ctx, descVisitor, imgVisitor)
			contents = mycomarkup.BlocksToHTML(ctx, ast)
			openGraph = getOpenGraph()
		}
		switch h := h.(type) {
		case *hyphae.MediaHypha:
			contents = views.Media(h, lc) + contents
		}

		util.HTTP200Page(w,
			views.Base(
				viewutil.MetaFrom(w, rq),
				util.BeautifulName(hyphaName),
				views.Hypha(viewutil.MetaFrom(w, rq), h, contents),
				openGraph))
	}
}
