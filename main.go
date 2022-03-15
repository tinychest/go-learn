package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println(getAppPath())
}

func getAppPath() string {
	f, _ := exec.LookPath(os.Args[0])
	p, _ := filepath.Abs(f)
	path := p[:strings.LastIndex(p, string(os.PathSeparator))]
	return path
}
