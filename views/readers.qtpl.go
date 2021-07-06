// Code generated by qtc from "readers.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/readers.qtpl:1
package views

//line views/readers.qtpl:1
import "net/http"

//line views/readers.qtpl:2
import "strings"

//line views/readers.qtpl:3
import "path"

//line views/readers.qtpl:4
import "os"

//line views/readers.qtpl:6
import "github.com/bouncepaw/mycorrhiza/cfg"

//line views/readers.qtpl:7
import "github.com/bouncepaw/mycorrhiza/hyphae"

//line views/readers.qtpl:8
import "github.com/bouncepaw/mycorrhiza/mimetype"

//line views/readers.qtpl:9
import "github.com/bouncepaw/mycorrhiza/tree"

//line views/readers.qtpl:10
import "github.com/bouncepaw/mycorrhiza/user"

//line views/readers.qtpl:11
import "github.com/bouncepaw/mycorrhiza/util"

//line views/readers.qtpl:13
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/readers.qtpl:13
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/readers.qtpl:13
func StreamAttachmentMenuHTML(qw422016 *qt422016.Writer, rq *http.Request, h *hyphae.Hypha, u *user.User) {
//line views/readers.qtpl:13
	qw422016.N().S(`
`)
//line views/readers.qtpl:14
	StreamNavHTML(qw422016, rq, h.Name, "attachment")
//line views/readers.qtpl:14
	qw422016.N().S(`
<div class="layout">
<main class="main-width">
	<h1>Attachment of `)
//line views/readers.qtpl:17
	qw422016.E().S(util.BeautifulName(h.Name))
//line views/readers.qtpl:17
	qw422016.N().S(`</h1>
	`)
//line views/readers.qtpl:18
	if h.BinaryPath == "" {
//line views/readers.qtpl:18
		qw422016.N().S(`
	<p class="warning">This hypha has no attachment, you can upload it here.</p>
	`)
//line views/readers.qtpl:20
	} else {
//line views/readers.qtpl:20
		qw422016.N().S(`
	<p class="warning">You can manage the hypha's attachment on this page.</p>
	`)
//line views/readers.qtpl:22
	}
//line views/readers.qtpl:22
	qw422016.N().S(`

	<section class="amnt-grid">

	`)
//line views/readers.qtpl:26
	if h.BinaryPath != "" {
//line views/readers.qtpl:26
		qw422016.N().S(`
		`)
//line views/readers.qtpl:28
		mime := mimetype.FromExtension(path.Ext(h.BinaryPath))
		fileinfo, err := os.Stat(h.BinaryPath)

//line views/readers.qtpl:29
		qw422016.N().S(`
		`)
//line views/readers.qtpl:30
		if err == nil {
//line views/readers.qtpl:30
			qw422016.N().S(`
		<fieldset class="amnt-menu-block">
			<legend class="modal__title modal__title_small">Stat</legend>
			<p class="modal__confirmation-msg"><b>File size:</b> `)
//line views/readers.qtpl:33
			qw422016.N().DL(fileinfo.Size())
//line views/readers.qtpl:33
			qw422016.N().S(` bytes</p>
			<p><b>MIME type:</b> `)
//line views/readers.qtpl:34
			qw422016.E().S(mime)
//line views/readers.qtpl:34
			qw422016.N().S(`</p>
		</fieldset>
		`)
//line views/readers.qtpl:36
		}
//line views/readers.qtpl:36
		qw422016.N().S(`

		`)
//line views/readers.qtpl:38
		if strings.HasPrefix(mime, "image/") {
//line views/readers.qtpl:38
			qw422016.N().S(`
		<fieldset class="amnt-menu-block">
			<legend class="modal__title modal__title_small">Include</legend>
			<p class="modal__confirmation-msg">This attachment is an image. To include it n a hypha, use a syntax like this:</p>
			<pre class="codebleck"><code>img { `)
//line views/readers.qtpl:42
			qw422016.E().S(h.Name)
//line views/readers.qtpl:42
			qw422016.N().S(` }</code></pre>
		</fieldset>
		`)
//line views/readers.qtpl:44
		}
//line views/readers.qtpl:44
		qw422016.N().S(`
	`)
//line views/readers.qtpl:45
	}
//line views/readers.qtpl:45
	qw422016.N().S(`

	`)
//line views/readers.qtpl:47
	if u.CanProceed("upload-binary") {
//line views/readers.qtpl:47
		qw422016.N().S(`
	<form action="/upload-binary/`)
//line views/readers.qtpl:48
		qw422016.E().S(h.Name)
//line views/readers.qtpl:48
		qw422016.N().S(`"
			method="post" enctype="multipart/form-data"
			class="upload-binary modal amnt-menu-block">
		<fieldset class="modal__fieldset">
			<legend class="modal__title modal__title_small">Attach</legend>
			<p class="modal__confirmation-msg">You can upload a new attachment. Please do not upload too big pictures unless you need to because may not want to wait for big pictures to load.</p>
			<label for="upload-binary__input"></label>
			<input type="file" id="upload-binary__input" name="binary">

			<input type="submit" class="btn stick-to-bottom" value="Upload">
		</fieldset>
	</form>
	`)
//line views/readers.qtpl:60
	}
//line views/readers.qtpl:60
	qw422016.N().S(`

	`)
//line views/readers.qtpl:62
	if h.BinaryPath != "" && u.CanProceed("unattach-confirm") {
//line views/readers.qtpl:62
		qw422016.N().S(`
	<form action="/unattach-confirm/`)
//line views/readers.qtpl:63
		qw422016.E().S(h.Name)
//line views/readers.qtpl:63
		qw422016.N().S(`" method="post" class="modal amnt-menu-block">
		<fieldset class="modal__fieldset">
			<legend class="modal__title modal__title_small">Unattach</legend>
			<p class="modal__confirmation-msg">Please note that you don't have to unattach before uploading a new attachment.</p>
			<input type="submit" class="btn" value="Unattach">
		</fieldset>
	</form>
	`)
//line views/readers.qtpl:70
	}
//line views/readers.qtpl:70
	qw422016.N().S(`

	</section>
</main>
</div>
`)
//line views/readers.qtpl:75
}

//line views/readers.qtpl:75
func WriteAttachmentMenuHTML(qq422016 qtio422016.Writer, rq *http.Request, h *hyphae.Hypha, u *user.User) {
//line views/readers.qtpl:75
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/readers.qtpl:75
	StreamAttachmentMenuHTML(qw422016, rq, h, u)
//line views/readers.qtpl:75
	qt422016.ReleaseWriter(qw422016)
//line views/readers.qtpl:75
}

//line views/readers.qtpl:75
func AttachmentMenuHTML(rq *http.Request, h *hyphae.Hypha, u *user.User) string {
//line views/readers.qtpl:75
	qb422016 := qt422016.AcquireByteBuffer()
//line views/readers.qtpl:75
	WriteAttachmentMenuHTML(qb422016, rq, h, u)
//line views/readers.qtpl:75
	qs422016 := string(qb422016.B)
//line views/readers.qtpl:75
	qt422016.ReleaseByteBuffer(qb422016)
//line views/readers.qtpl:75
	return qs422016
//line views/readers.qtpl:75
}

// If `contents` == "", a helpful message is shown instead.

//line views/readers.qtpl:78
func StreamHyphaHTML(qw422016 *qt422016.Writer, rq *http.Request, h *hyphae.Hypha, contents string) {
//line views/readers.qtpl:78
	qw422016.N().S(`
`)
//line views/readers.qtpl:80
	sisters, subhyphae, prevHyphaName, nextHyphaName := tree.Tree(h.Name)
	u := user.FromRequest(rq)

//line views/readers.qtpl:82
	qw422016.N().S(`
`)
//line views/readers.qtpl:83
	StreamNavHTML(qw422016, rq, h.Name, "hypha")
//line views/readers.qtpl:83
	qw422016.N().S(`
<div class="layout">
<main class="main-width">
	<article>
		`)
//line views/readers.qtpl:87
	qw422016.N().S(NaviTitleHTML(h))
//line views/readers.qtpl:87
	qw422016.N().S(`
		`)
//line views/readers.qtpl:88
	if h.Exists {
//line views/readers.qtpl:88
		qw422016.N().S(`
			`)
//line views/readers.qtpl:89
		qw422016.N().S(contents)
//line views/readers.qtpl:89
		qw422016.N().S(`
		`)
//line views/readers.qtpl:90
	} else {
//line views/readers.qtpl:90
		qw422016.N().S(`
		    `)
//line views/readers.qtpl:91
		streamnonExistentHyphaNotice(qw422016, h, u)
//line views/readers.qtpl:91
		qw422016.N().S(`
		`)
//line views/readers.qtpl:92
	}
//line views/readers.qtpl:92
	qw422016.N().S(`
	</article>
	<section class="prevnext">
		`)
//line views/readers.qtpl:95
	if prevHyphaName != "" {
//line views/readers.qtpl:95
		qw422016.N().S(`
		<a class="prevnext__el prevnext__prev" href="/hypha/`)
//line views/readers.qtpl:96
		qw422016.E().S(prevHyphaName)
//line views/readers.qtpl:96
		qw422016.N().S(`" rel="prev">← `)
//line views/readers.qtpl:96
		qw422016.E().S(util.BeautifulName(path.Base(prevHyphaName)))
//line views/readers.qtpl:96
		qw422016.N().S(`</a>
		`)
//line views/readers.qtpl:97
	}
//line views/readers.qtpl:97
	qw422016.N().S(`
		`)
//line views/readers.qtpl:98
	if nextHyphaName != "" {
//line views/readers.qtpl:98
		qw422016.N().S(`
		<a class="prevnext__el prevnext__next" href="/hypha/`)
//line views/readers.qtpl:99
		qw422016.E().S(nextHyphaName)
//line views/readers.qtpl:99
		qw422016.N().S(`" rel="next">`)
//line views/readers.qtpl:99
		qw422016.E().S(util.BeautifulName(path.Base(nextHyphaName)))
//line views/readers.qtpl:99
		qw422016.N().S(` →</a>
		`)
//line views/readers.qtpl:100
	}
//line views/readers.qtpl:100
	qw422016.N().S(`
	</section>
`)
//line views/readers.qtpl:102
	StreamSubhyphaeHTML(qw422016, subhyphae)
//line views/readers.qtpl:102
	qw422016.N().S(`
</main>
`)
//line views/readers.qtpl:104
	streamsisterHyphaeHTML(qw422016, sisters)
//line views/readers.qtpl:104
	qw422016.N().S(`
</div>
`)
//line views/readers.qtpl:106
	streamviewScripts(qw422016)
//line views/readers.qtpl:106
	qw422016.N().S(`
`)
//line views/readers.qtpl:107
}

//line views/readers.qtpl:107
func WriteHyphaHTML(qq422016 qtio422016.Writer, rq *http.Request, h *hyphae.Hypha, contents string) {
//line views/readers.qtpl:107
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/readers.qtpl:107
	StreamHyphaHTML(qw422016, rq, h, contents)
//line views/readers.qtpl:107
	qt422016.ReleaseWriter(qw422016)
//line views/readers.qtpl:107
}

//line views/readers.qtpl:107
func HyphaHTML(rq *http.Request, h *hyphae.Hypha, contents string) string {
//line views/readers.qtpl:107
	qb422016 := qt422016.AcquireByteBuffer()
//line views/readers.qtpl:107
	WriteHyphaHTML(qb422016, rq, h, contents)
//line views/readers.qtpl:107
	qs422016 := string(qb422016.B)
//line views/readers.qtpl:107
	qt422016.ReleaseByteBuffer(qb422016)
//line views/readers.qtpl:107
	return qs422016
//line views/readers.qtpl:107
}

//line views/readers.qtpl:109
func StreamRevisionHTML(qw422016 *qt422016.Writer, rq *http.Request, h *hyphae.Hypha, contents, revHash string) {
//line views/readers.qtpl:109
	qw422016.N().S(`
`)
//line views/readers.qtpl:111
	sisters, subhyphae, _, _ := tree.Tree(h.Name)

//line views/readers.qtpl:112
	qw422016.N().S(`
`)
//line views/readers.qtpl:113
	StreamNavHTML(qw422016, rq, h.Name, "revision", revHash)
//line views/readers.qtpl:113
	qw422016.N().S(`
<div class="layout">
<main class="main-width">
	<article>
		<p>Please note that viewing binary parts of hyphae is not supported in history for now.</p>
		`)
//line views/readers.qtpl:118
	qw422016.N().S(NaviTitleHTML(h))
//line views/readers.qtpl:118
	qw422016.N().S(`
		`)
//line views/readers.qtpl:119
	qw422016.N().S(contents)
//line views/readers.qtpl:119
	qw422016.N().S(`
	</article>
`)
//line views/readers.qtpl:121
	StreamSubhyphaeHTML(qw422016, subhyphae)
//line views/readers.qtpl:121
	qw422016.N().S(`
</main>
`)
//line views/readers.qtpl:123
	streamsisterHyphaeHTML(qw422016, sisters)
//line views/readers.qtpl:123
	qw422016.N().S(`
</div>
`)
//line views/readers.qtpl:125
	streamviewScripts(qw422016)
//line views/readers.qtpl:125
	qw422016.N().S(`
`)
//line views/readers.qtpl:126
}

//line views/readers.qtpl:126
func WriteRevisionHTML(qq422016 qtio422016.Writer, rq *http.Request, h *hyphae.Hypha, contents, revHash string) {
//line views/readers.qtpl:126
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/readers.qtpl:126
	StreamRevisionHTML(qw422016, rq, h, contents, revHash)
//line views/readers.qtpl:126
	qt422016.ReleaseWriter(qw422016)
//line views/readers.qtpl:126
}

//line views/readers.qtpl:126
func RevisionHTML(rq *http.Request, h *hyphae.Hypha, contents, revHash string) string {
//line views/readers.qtpl:126
	qb422016 := qt422016.AcquireByteBuffer()
//line views/readers.qtpl:126
	WriteRevisionHTML(qb422016, rq, h, contents, revHash)
//line views/readers.qtpl:126
	qs422016 := string(qb422016.B)
//line views/readers.qtpl:126
	qt422016.ReleaseByteBuffer(qb422016)
//line views/readers.qtpl:126
	return qs422016
//line views/readers.qtpl:126
}

//line views/readers.qtpl:128
func streamviewScripts(qw422016 *qt422016.Writer) {
//line views/readers.qtpl:128
	qw422016.N().S(`
`)
//line views/readers.qtpl:129
	for _, scriptPath := range cfg.ViewScripts {
//line views/readers.qtpl:129
		qw422016.N().S(`
<script src="`)
//line views/readers.qtpl:130
		qw422016.E().S(scriptPath)
//line views/readers.qtpl:130
		qw422016.N().S(`"></script>
`)
//line views/readers.qtpl:131
	}
//line views/readers.qtpl:131
	qw422016.N().S(`
`)
//line views/readers.qtpl:132
}

//line views/readers.qtpl:132
func writeviewScripts(qq422016 qtio422016.Writer) {
//line views/readers.qtpl:132
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/readers.qtpl:132
	streamviewScripts(qw422016)
//line views/readers.qtpl:132
	qt422016.ReleaseWriter(qw422016)
//line views/readers.qtpl:132
}

//line views/readers.qtpl:132
func viewScripts() string {
//line views/readers.qtpl:132
	qb422016 := qt422016.AcquireByteBuffer()
//line views/readers.qtpl:132
	writeviewScripts(qb422016)
//line views/readers.qtpl:132
	qs422016 := string(qb422016.B)
//line views/readers.qtpl:132
	qt422016.ReleaseByteBuffer(qb422016)
//line views/readers.qtpl:132
	return qs422016
//line views/readers.qtpl:132
}
