package gclone

import (
	"fmt"
	"github.com/acehinnnqru/gclone/src/g"
	"log"
	"path/filepath"
)

type Option struct {
	Root         string
	Repositories []string
}

func Clone(opt Option) {
	for _, repository := range opt.Repositories {
		addr, e := g.Parse(repository)
		if e != nil {
			log.Println(e)
			continue
		}
		e = g.Clone(repository, filepath.Join(opt.Root, fmt.Sprintf("%v/%v/%v", addr.Server, addr.Namespace, addr.Repository)))
		if e != nil {
			log.Println(e)
			continue
		}
	}
}
