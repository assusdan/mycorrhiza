module github.com/bouncepaw/mycorrhiza

go 1.18

require (
	github.com/bouncepaw/mycomarkup/v4 v4.2.0
	github.com/coreos/go-oidc/v3 v3.2.0
	github.com/go-ini/ini v1.63.2
	github.com/gorilla/feeds v1.1.1
	github.com/gorilla/mux v1.8.0
	github.com/valyala/quicktemplate v1.7.0
	golang.org/x/crypto v0.0.0-20211108221036-ceb1ce70b4fa
	golang.org/x/exp v0.0.0-20220414153411-bcd21879b8fd
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd
	golang.org/x/oauth2 v0.0.0-20220524215830-622c5d57e401
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211
	golang.org/x/text v0.3.7
)

require (
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/kr/pretty v0.2.1 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	google.golang.org/appengine v1.6.6 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/square/go-jose.v2 v2.5.1 // indirect
)

// Use this trick to test local Mycomarkup changes, replace the path with yours,
// but do not commit the change to the path:
//  replace github.com/bouncepaw/mycomarkup/v4 v4.0.0 => "/Users/bouncepaw/GolandProjects/mycomarkup"

// Use this utility every time Mycomarkup gets a major update:
// https://github.com/marwan-at-work/mod
// Or maybe just ⌘⇧R every time, the utility is kinda weird.
