package main

import (
	"fmt"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplates() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

type Count struct {
	Count int
}

type Contact struct {
	Name  string
	Email string
}

type Contacts = []Contact

type Data struct {
	Contacts Contacts
}

func newContact(name string, email string) Contact {
	return Contact{
		Name:  name,
		Email: email,
	}
}

func newData() Data {
	return Data{
		Contacts: []Contact{
			newContact("Tom", "tobarrientos@gmail.com"),
			newContact("Alberto", "alberto@gmail.com"),
		},
	}
}

func main() {
	fmt.Println("Hello World")

	e := echo.New()

	e.Use(middleware.Logger())

	e.Renderer = newTemplates()

	data := newData()

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", data)
	})

	e.POST("/contacts", func(c echo.Context) error {
		name := c.FormValue("name")
		email := c.FormValue("email")

		data.Contacts = append(data.Contacts, newContact(name, email))
		return c.Render(200, "display", data)
	})

	e.Logger.Fatal(e.Start(":3000"))

}
