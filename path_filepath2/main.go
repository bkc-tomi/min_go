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
		fmt.Println(r)
		if r.Method == "GET" {
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
		}

		// ファイルアップロードの際もpathとfilepathをしっかり使い分ける。
		// 特にcurlコマンドでサーバー内のファイルを操作される可能性もある。
		if r.Method == "POST" {
			stream, header, err := r.FormFile("myfile")
			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
				return
			}
			// 安全に処理をするためにfilepath.Baseを使ってファイル名を取り出して操作するようにする。
			filename := filepath.Base(header.Filename)
			p := filepath.Join(cwd, "files", filename)
			fmt.Println(p)
			f, err := os.Create(p)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
				return
			}
			defer f.Close()
			io.Copy(f, stream)
			http.Redirect(w, r, path.Join("/files", filename), 301)
		} else {
			http.Redirect(w, r, "/data/index.html", 301)
		}
	})

	http.ListenAndServe(":8080", nil)
}
