package utils

import (
	"html/template"
	"os"
	"path/filepath"
)

func GetCssAsset(name string) (res template.HTML) {
	err := filepath.Walk("public/assets", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Name() == name {
			res = template.HTML("<link rel=\"stylesheet\" href=\"/" + path + "\">")
		}
		return nil
	})
	if err != nil {
		return "Stylesheet not loading"
	}
	return
}
