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
	p.Parameters.Width = 768
	p.Parameters.Height = 512
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
			_, _, im, err := n.Draw(nm + ",avatar")
			if err != nil {
				panic(err)
			}
			_ = os.WriteFile(target+"/"+nm+".png", im, 0644)
		}
	}
}
