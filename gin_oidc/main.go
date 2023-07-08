package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
	//oidc support
	"context"
	"github.com/google/uuid"
	"github.com/zitadel/oidc/v2/pkg/client/rp"
	httphelper "github.com/zitadel/oidc/v2/pkg/http"
	"github.com/zitadel/oidc/v2/pkg/oidc"
	"net/http"
)

//You have to wrap the oidc helper function,
//because gin uses *gin.Context for routes
//and not (writer,request) as net/http would
func WrapF(f http.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		newCtx := context.WithValue(c.Request.Context(), "session", c)
		f(c.Writer, c.Request.WithContext(newCtx))
	}
}

func main() {
	Host := "http://localhost:7070/"
	ClientID := "myclient"
	ClientSecret := "jDDQD43FFSXXXXXXXXX"
	Issuer := "https://example.com/realms/myrealm"
	OidcCookieKey := "temp1234temp1234"
	callbackPath := "auth/callback"

	//oidc
	state := func() string {
		a := uuid.New().String()
		return a
	}

	redirectURI := Host + callbackPath
	scopes := strings.Split("email openid", " ")
	key := []byte(OidcCookieKey)
	cookieHandler := httphelper.NewCookieHandler(key, key, httphelper.WithUnsecure())
	options := []rp.Option{
		rp.WithCookieHandler(cookieHandler),
		rp.WithVerifierOpts(rp.WithIssuedAtOffset(5 * time.Second)),
	}
	if ClientSecret == "" {
		options = append(options, rp.WithPKCE(cookieHandler))
	}
	provider, err := rp.NewRelyingPartyOIDC(Issuer, ClientID, ClientSecret, redirectURI, scopes, options...)
	if err != nil {
		panic("error creating keycloak provider " + err.Error())
	}

	//Webserver
	gin.SetMode(gin.ReleaseMode)
	app := gin.Default()

	//Configure login session in gin
	store := cookie.NewStore([]byte(OidcCookieKey))
	app.Use(sessions.Sessions("provencesession", store))

	//This blocks access and redirects to /login if:
	// - you are not logged in and
	// - if you are not trying to login
	//  (either via /login or callback from oidc)
	app.Use(func(c *gin.Context) {
		s := sessions.Default(c)
		name := s.Get("name")

		if c.FullPath() == "/"+callbackPath {
			return
		}
		if c.FullPath() == "/login" {
			return
		}

		if name == nil {
			c.Redirect(307, Host+"login")
			c.Abort()
		}
	})

	//Redirect to OIDC login page with this
	app.GET("/login", gin.WrapF(rp.AuthURLHandler(state, provider)))

	app.GET("/logout", func(c *gin.Context) {
		s := sessions.Default(c)
		idt := s.Get("idtoken").(string)
		_, err := rp.EndSession(provider, idt, Host+"loggedout", "")
		if err != nil {
			c.String(500, "ERROR during oidc logout!")
			fmt.Println(err)
			return
		}
		s.Clear()
		err = s.Save()
		if err != nil {
			c.String(500, "ERROR during session save!")
			fmt.Println(err)
			return
		}
		c.Redirect(301, Host+"loggedout")
	})
	setSessionFunc := func(w http.ResponseWriter, r *http.Request, tokens *oidc.Tokens[*oidc.IDTokenClaims], state string, rp rp.RelyingParty, info *oidc.UserInfo) {
		c := r.Context().Value("session").(*gin.Context)
		s := sessions.Default(c)
		s.Set("name", info.UserInfoProfile.PreferredUsername)
		s.Set("idtoken", tokens.IDToken)
		s.Save()
		http.Redirect(w, r, Host+"home", 301)
	}
	app.GET(callbackPath, WrapF(rp.CodeExchangeHandler(rp.UserinfoCallback(setSessionFunc), provider)))

	// Routes
	app.GET("/", func(c *gin.Context) {
		c.String(200, "OK")
	})
	app.GET("/home", func(c *gin.Context) {
		c.String(200, "You are viewing /home")
	})

	fmt.Println("Now Listening on 7070")
	fmt.Println(app.Run(":7070"))
}
