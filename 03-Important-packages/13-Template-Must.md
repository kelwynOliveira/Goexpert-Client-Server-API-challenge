# Template Must

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

func main(){
  course := Course{"Go", 40}
  t := template.Must(template.New("CourseTemplate").Parse("Course: {{.Name}} - Hours: {{.Hours}}"))
  err := t.Execute(os.Stdout, course)
  if err != nil{
    panic(err)
  }
}
```
