// Package cfg contains global variables that represent the current wiki
// configuration, including CLI options, configuration file values and header
// links.
package cfg

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/go-ini/ini"
)

// These variables represent the configuration. You are not meant to modify
// them after they were set.
// See https://mycorrhiza.wiki/hypha/configuration/fields for the
// documentation.
var (
	WikiName                string
	NaviTitleIcon           string
	UseSiblingHyphaeSidebar bool

	HomeHypha        string
	UserHypha        string
	HeaderLinksHypha string

	ListenAddr string
	URL        string

	UseAuth           bool
	AllowRegistration bool
	// AllowEmailAsLogin TODO: is not currently in use, "@" characted is banned in mycomarkup
	AllowEmailAsLogin bool
	RegistrationLimit uint64
	Locked            bool
	UseWhiteList      bool
	WhiteList         []string

	CommonScripts []string
	ViewScripts   []string
	EditScripts   []string

	// TelegramEnabled if both TelegramBotToken and TelegramBotName are not empty strings.
	TelegramEnabled  bool
	TelegramBotToken string
	TelegramBotName  string

	// OIDC config
	OidcEnabled             bool
	OidcClientId            string
	OidcClientSecret        string
	OidcProvider            string
	OidcProviderUri         string
	OidcPlusSeparatedScopes string
)

// WikiDir is a full path to the wiki storage directory, which also must be a
// git repo. This variable is set in parseCliArgs().
var WikiDir string

// Config represents a Mycorrhiza wiki configuration file. This type is used
// only when reading configs.
type Config struct {
	WikiName                string `comment:"This name appears in the header and on various pages."`
	NaviTitleIcon           string `comment:"This icon is used in the breadcrumbs bar."`
	UseSiblingHyphaeSidebar bool   `comment:"You are discouraged from using the sibling hyphae sidebar on new wikis. Enable it on old wikis that depend on it heavily."`
	Hyphae
	Network
	Authorization
	CustomScripts `comment:"You can specify additional scripts to load on different kinds of pages, delimited by a comma ',' sign."`
	Telegram      `comment:"You can enable Telegram authorization. Follow these instructions: https://core.telegram.org/widgets/login#setting-up-a-bot"`
	Oidc
}

// Hyphae is a section of Config which has fields related to special hyphae.
type Hyphae struct {
	HomeHypha        string `comment:"This hypha will be the main (index) page of your wiki, served on /."`
	UserHypha        string `comment:"This hypha is used as a prefix for user hyphae."`
	HeaderLinksHypha string `comment:"You can also specify a hypha to populate your own custom header links from."`
}

// Network is a section of Config that has fields related to network stuff.
type Network struct {
	ListenAddr string
	URL        string `comment:"Set your wiki's public URL here. It's used for OpenGraph generation and syndication feeds."`
}

// CustomScripts is a section with paths to JavaScript files that are loaded on
// specified pages.
type CustomScripts struct {
	// CommonScripts: everywhere...
	CommonScripts []string `delim:"," comment:"These scripts are loaded from anywhere."`
	// ViewScripts: /hypha, /rev
	ViewScripts []string `delim:"," comment:"These scripts are only loaded on view pages."`
	// Edit: /edit
	EditScripts []string `delim:"," comment:"These scripts are only loaded on the edit page."`
}

// Authorization is a section of Config that has fields related to
// authorization and authentication.
type Authorization struct {
	UseAuth           bool
	AllowRegistration bool
	AllowEmailAsLogin bool
	RegistrationLimit uint64   `comment:"This field controls the maximum amount of allowed registrations."`
	Locked            bool     `comment:"Set if users have to authorize to see anything on the wiki."`
	UseWhiteList      bool     `comment:"If true, WhiteList is used. Else it is not used."`
	WhiteList         []string `delim:"," comment:"Usernames of people who can log in to your wiki separated by comma."`

	// TODO: let admins enable auth-less editing
}

// Telegram is the section of Config that sets Telegram authorization.
type Telegram struct {
	TelegramBotToken string `comment:"Token of your bot."`
	TelegramBotName  string `comment:"Username of your bot, sans @."`
}

type Oidc struct {
	OidcEnabled             bool   `comment:"true."`
	OidcClientId            string `comment:"Oauth2 client_id "`
	OidcClientSecret        string `comment:"Oauth2 cient_secret "`
	OidcProvider            string `comment:"Oauth2 provider name "`
	OidcProviderUri         string `comment:"Oauth2 provider base URI "`
	OidcPlusSeparatedScopes string `comment:"Oauth2 scopes, plus-separated "`
}

// ReadConfigFile reads a config on the given path and stores the
// configuration. Call it sometime during the initialization.
func ReadConfigFile(path string) error {
	cfg := &Config{
		WikiName:                "Mycorrhiza Wiki",
		NaviTitleIcon:           "🍄",
		UseSiblingHyphaeSidebar: false,

		Hyphae: Hyphae{
			HomeHypha:        "home",
			UserHypha:        "u",
			HeaderLinksHypha: "",
		},
		Network: Network{
			ListenAddr: "127.0.0.1:1737",
			URL:        "",
		},
		Authorization: Authorization{
			UseAuth:           true,
			AllowRegistration: true,
			AllowEmailAsLogin: true,
			RegistrationLimit: 0,
			Locked:            true,
			UseWhiteList:      false,
			WhiteList:         []string{},
		},
		CustomScripts: CustomScripts{
			CommonScripts: []string{},
			ViewScripts:   []string{},
			EditScripts:   []string{},
		},
		Telegram: Telegram{
			TelegramBotToken: "",
			TelegramBotName:  "",
		},

		Oidc: Oidc{
			OidcEnabled:             false,
			OidcClientId:            "",
			OidcClientSecret:        "",
			OidcProvider:            "",
			OidcProviderUri:         "",
			OidcPlusSeparatedScopes: "",
		},
	}

	f, err := ini.Load(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			f = ini.Empty()

			// Save the default configuration
			err = f.ReflectFrom(cfg)
			if err != nil {
				return fmt.Errorf("Failed to serialize the config: %w", err)
			}

			// Disable key-value auto-aligning, but retain spaces around '=' sign
			ini.PrettyFormat = false
			ini.PrettyEqual = true
			if err = f.SaveTo(path); err != nil {
				return fmt.Errorf("Failed to save the config file: %w", err)
			}
		} else {
			return fmt.Errorf("Failed to open the config file: %w", err)
		}
	}

	// Map the config file to the config struct. It'll do nothing if the file
	// doesn't exist or is empty.
	if err := f.MapTo(cfg); err != nil {
		return err
	}

	// Map the struct to the global variables
	WikiName = cfg.WikiName
	NaviTitleIcon = cfg.NaviTitleIcon
	UseSiblingHyphaeSidebar = cfg.UseSiblingHyphaeSidebar
	HomeHypha = cfg.HomeHypha
	UserHypha = cfg.UserHypha
	HeaderLinksHypha = cfg.HeaderLinksHypha
	if ListenAddr == "" {
		ListenAddr = cfg.ListenAddr
	}
	URL = cfg.URL
	UseAuth = cfg.UseAuth
	AllowRegistration = cfg.AllowRegistration
	AllowEmailAsLogin = cfg.AllowEmailAsLogin
	RegistrationLimit = cfg.RegistrationLimit
	Locked = cfg.Locked && cfg.UseAuth // Makes no sense to have the lock but no auth
	UseWhiteList = cfg.UseWhiteList
	WhiteList = cfg.WhiteList
	CommonScripts = cfg.CommonScripts
	ViewScripts = cfg.ViewScripts
	EditScripts = cfg.EditScripts
	TelegramBotToken = cfg.TelegramBotToken
	TelegramBotName = cfg.TelegramBotName
	TelegramEnabled = (TelegramBotToken != "") && (TelegramBotName != "")

	OidcEnabled = cfg.OidcEnabled
	OidcClientId = cfg.OidcClientId
	OidcClientSecret = cfg.OidcClientSecret
	OidcProvider = cfg.OidcProvider
	OidcProviderUri = cfg.OidcProviderUri
	OidcPlusSeparatedScopes = cfg.OidcPlusSeparatedScopes

	// This URL makes much more sense. If no URL is set or the protocol is forgotten, assume HTTP.
	if (URL == "") || !strings.Contains(URL, ":") {
		URL = "http://" + ListenAddr
	}

	return nil
}
