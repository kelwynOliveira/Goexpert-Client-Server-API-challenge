# Templates with webserver

File `template.html`

```html
...

<body>
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
</body>
...
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
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
    t := template.Must(template.New("CourseTemplate").ParseFiles("template.html"))
    err := t.Execute(w, Courses{
      {"Go", 40},
      {"Java", 30},
      {"Python", 35}
    })
  })
  http.ListenAndServe(":8282", nil)
}
```
