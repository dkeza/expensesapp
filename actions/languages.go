package actions

import (
	"log"
	"net/http"
	"time"

	"github.com/gobuffalo/buffalo"
)

// LanguagesChange default implementation.
func LanguagesChange(c buffalo.Context) error {
	lng := c.Param("lang")
	log.Println(lng)

	f := struct {
		Language string `form:"lang"`
		URL      string `form:"url"`
	}{}
	f.Language = lng
	f.URL = "/"
	//if err := c.Bind(&f); err != nil {
	//return errors.WithStack(err)
	//}

	// Set new current language using a cookie, for instance
	cookie := http.Cookie{
		Name:   "lang",
		Value:  f.Language,
		MaxAge: int((time.Hour * 24 * 265).Seconds()),
		Path:   "/",
	}
	http.SetCookie(c.Response(), &cookie)

	// Update language for the flash message
	T.Refresh(c, f.Language)

	//c.Flash().Add("success", T.Translate(c, "users.language-changed", f))

	return c.Redirect(302, f.URL)

	//return c.Redirect(302, "/")
}
