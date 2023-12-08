# Templates with external file

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
  course := Course{"Go", 40}
  t := template.Must(template.New("CourseTemplate").ParseFiles("template.html"))
  err := t.Execute(os.Stdout, course{
    {"Go", 40},
    {"Java", 30},
    {"Python", 35}
  })
  if err != nil{
    panic(err)
  }
}
```

> `.ParseFiles([]string["template.html", "otherfile.something", ....])`
