package main

import (
	"fmt"
	"io/fs"
	"os"

	nvai "github.com/FloatTech/AnimeAPI/novelai"
	"github.com/FloatTech/zbputils/process"
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
			fmt.Println("生成:", nm)
			for i := 0; i < 8; i++ {
				_, _, im, err := n.Draw(nm + ",1girl,looking_at_viewer,half_body,cute,bule_archive,extremely_detailed_CG_unity_8k_wallpaper,illustration,high_quality,lan")
				if err != nil {
					fmt.Println("ERROR:", err)
					process.SleepAbout1sTo2s()
					continue
				}
				err = os.WriteFile(target+"/"+nm+".png", im, 0644)
				if err != nil {
					panic(err)
				}
				fmt.Println("成功")
				process.SleepAbout1sTo2s()
				break
			}
		}
	}
}
