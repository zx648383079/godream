package main

import (
	"os"
)
type php struct {
    dir string
}

var p php;

func (p *php) getId(name string) string, error {
    fs, err := os.Open(p.dir + "/" + name)
    if (err != nil) {
        return err
    }
    
}