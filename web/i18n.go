package web

import (
	duxI18n "github.com/duxweb/go-fast/i18n"
	"github.com/labstack/echo/v4"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func I18nHandler() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			accept := c.Request().Header.Get("Accept-Language")
			if accept == "" {
				accept = "en-US"
			}
			t, _, _ := language.ParseAcceptLanguage(accept)

			var langTag language.Tag
			if len(t) > 0 {
				langTag = t[0]
			} else {
				langTag = language.English
			}

			c.Set("i18n", i18n.NewLocalizer(duxI18n.Bundle, accept))
			c.Set("lang", langTag.String())
			return next(c)
		}
	}
}
