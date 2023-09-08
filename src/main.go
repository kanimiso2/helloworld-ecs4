package main

import (
	"math/rand"
	"text/template"

	"github.com/labstack/echo/v4"
)

func todoHandler(c echo.Context) error {
	tmpl := template.Must(template.ParseFiles("index.html"))
	data := struct {
		Question string
	}{
		Question: "あなたの質問",
	}
	return tmpl.Execute(c.Response().Writer, data)
}
func main() {
	e := echo.New()
	e.GET("/", todoHandler)
	e.POST("/process", func(c echo.Context) error {
		question := c.FormValue("question")
		fortune := generateFortune()
		tmpl := template.Must(template.ParseFiles("uranai.html"))
		data := struct {
			Question string
			Result   string
		}{
			Question: question,
			Result:   fortune,
		}
		return tmpl.Execute(c.Response().Writer, data)
	})
	e.Logger.Fatal(e.Start(":80"))
}
func generateFortune() string {
	fortunes := []string{
		"大吉",
		"中吉",
		"小吉",
		"吉",
		"凶",
	}
	index := rand.Intn(len(fortunes))
	return fortunes[index]
}
