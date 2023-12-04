# Compiling projects

- Runtime: Go code (where all the Go dependencies are)
- Your code

Runtime + Your code = Binary file

We can choose which platform we want to work.

## How to compile

Run `go build <file>`.

For windows: `GOOS=windows go build main.go`
It will generate a executable file that will work on windows.

- GOOS: Operational System
- GOARCH: Processor Architecture

> To find the list of possible platforms, run the following:
> `go tool dist list`

<a href="https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63" target="_blank">Go (Golang) GOOS and GOARCH</a>

<a href="https://www.digitalocean.com/community/tutorials/building-go-applications-for-different-operating-systems-and-architectures" target="_blank">Building Go Applications for Different Operating Systems and Architectures</a>
