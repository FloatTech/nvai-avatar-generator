package main

import (
	"fmt"
	"io/fs"
	"math/rand"
	"os"
	"regexp"
	"time"

	nvai "github.com/FloatTech/AnimeAPI/novelai"
	"github.com/FloatTech/floatbox/binary"
	"github.com/FloatTech/zbputils/process"
)

var (
	regre = regexp.MustCompile(`.Register\(\n?\t*"(.+)",`)
	regsv = regexp.MustCompile(`[sS]ervice[nN]ame = "(.+)"`)
)

func main() {
	rand.Seed(time.Now().UnixNano())
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
	plugins := []string{}
	err = fs.WalkDir(os.DirFS(os.Args[1]), ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		data, err := os.ReadFile(os.Args[1] + "/" + path)
		if err != nil {
			return err
		}
		for _, m := range regre.FindAllStringSubmatch(binary.BytesToString(data), -1) {
			plugins = append(plugins, m...)
		}
		for _, m := range regsv.FindAllStringSubmatch(binary.BytesToString(data), -1) {
			plugins = append(plugins, m...)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	target := os.Args[2]
	err = os.MkdirAll(target, 0755)
	if err != nil {
		panic(err)
	}
	for _, nm := range plugins {
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
