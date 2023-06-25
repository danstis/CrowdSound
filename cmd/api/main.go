package main

const (
	addr = ":9200"
)

func main() {
	a := App{}
	a.Run(addr)
}
