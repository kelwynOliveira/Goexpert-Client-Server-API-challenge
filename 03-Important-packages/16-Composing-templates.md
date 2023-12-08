# Composing templates

File `header.html`

```html
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Document</title>
  </head>
  <body></body>
</html>
```

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
      <td>{{ course.name }}</td>
      <td>{{ course.hours }}</td>
    </tr>
    {{ end }}
  </tbody>
</table>

{{ template "footer.html" }}
```

File `footer.html`

```html
</body>
</html>
```

```go
package main

import (
  "os"
  "text/template"
)

type Course struct{
  Name string
  Hours: int
}

type Couses []Course


func main(){
  templates := []string{
    "header.html",
    "content.html"
    "footer.html"
  }

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
    t := template.Must(template.New(content.html).ParseFiles(templates...))
    err := t.Execute(w, Courses{
      {"Go", 40},
      {"Java", 30},
      {"Python", 35}
    })
  })
  http.ListenAndServe(":8282", nil)
}
```
