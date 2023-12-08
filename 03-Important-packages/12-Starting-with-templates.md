# Starting with templates

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
  tmp := template.New("CourseTemplate")
  tmo, err := tmp.Parse("Course: {{.Name}} - Hours: {{.Hours}}")
  if err != nil{
    panic(err)
  }

  err := tmp.Execute(os.Stdout, course)
  if err != nil{
    panic(err)
  }
}
```
