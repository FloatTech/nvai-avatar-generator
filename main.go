package main

import (
	"io/fs"
	"os"

	nvai "github.com/FloatTech/AnimeAPI/novelai"
)

func main() {
	key := os.Getenv("NOVELAI_API_KEY")
	if key == "" {
		panic("nil api key")
	}
	p := nvai.NewDefaultPayload()
	p.Parameters.Width = 700
	p.Parameters.Height = 430
	n := nvai.NewNovalAI(key, p)
	err := n.Login()
	if err != nil {
		panic(err)
	}
	e, err := fs.ReadDir(os.DirFS(os.Args[1]), ".")
	if err != nil {
		panic(err)
	}
	target := os.Args[2]
	err = os.MkdirAll(target, 0755)
	if err != nil {
		panic(err)
	}
	for _, d := range e {
		if d.IsDir() {
			nm := d.Name()
			_, _, img, err := n.Draw(nm + ",avatar")
			if err != nil {
				panic(err)
			}
			_ = os.WriteFile(target+"/"+nm+".png", img, 0644)
		}
	}
}
