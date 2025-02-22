// package main
//
// import (
//   "html/template"
//   "io"
//   "net/http"
//
//   "github.com/labstack/echo/v4"
//   "github.com/labstack/echo/v4/middleware"
// )
//
// type TemplateRenderer struct {
//   templates * template.Template
// }

// func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
//     return t.templates.ExecuteTemplate(w, name, data)
// }


//func main() {
//   e := echo.New()
//  e.get("/", func(c echo.Context) error {
//    return c.String(http.StatusOk, "Hello world!")
//  })
//  e.Logger.Fatal(e.Start(":123"))
//}

