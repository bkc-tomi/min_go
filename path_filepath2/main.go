package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// httpは論理パスなのでpathを使う
		ok, err := path.Match("/data/*.html", r.URL.Path)
		fmt.Println(r.URL.Path)
		if err != nil || !ok {
			http.NotFound(w, r)
			fmt.Println("path match error")
			return
		}

		// 以降は物理パスを扱うのでpath/filepathを使う
		name := filepath.Join(cwd, "data", filepath.Base(r.URL.Path))
		/*
			上記のfilepath.Baseをpath.Baseに書き換え、windowsで
			http://localhost:8080/data/..\main.go
			にアクセスするとmain.goファイルにアクセス出来てしまう。
			filepathとpathをしっかり使い分ける。
		*/
		fmt.Println(name)
		f, err := os.Open(name)
		if err != nil {
			http.NotFound(w, r)
			fmt.Println("file not found")
			return
		}
		defer f.Close()
		io.Copy(w, f)
	})

	http.ListenAndServe(":8080", nil)
}
