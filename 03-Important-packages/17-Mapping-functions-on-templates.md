# Mapping functions on templates

File `content.html`

```html
{{ template "header.html" }}

<h1>Courses</h1>

<table>
  <thead>
    <tr>
      <th>Name</th>
      <th>Hours</th>
    </tr>
  </thead>
  <tbody>
    {{ range .}}
    <!-- the . means the informations are on the same struct -->
    <tr>
      <td>{{ course.name | ToUpper}}</td>
      <td>{{ course.hours }}</td>
    </tr>
    {{ end }}
  </tbody>
</table>

{{ template "footer.html" }}
```

```go
package main

import (
  "html/template"
  "os"
)

type Course struct{
  Name string
  Hours: int
}

type Couses []Course

func ToUpper(s string) string{
  return strings.ToUpper(s)
}

func main(){
  templates := []string{
    "header.html",
    "content.html"
    "footer.html"
  }

  t := template.New("content.html")
  t.Funcs(template.FuncMap("ToUpper": To Upper))
  t = template.Must(t.ParseFiles(templates...))

  err := t.Execute(os.Stdout, Courses{
    {"Go", 40},
    {"Java", 30},
    {"Python", 35}
  })
  if err != nil {
    panic(err)
  }
}
```
