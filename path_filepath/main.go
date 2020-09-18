package main

import (
	"log"
	"os"
	"os/user"
	"path/filepath" //　PC内のディレクトリ構成のパスを指定(物理パス)
	// http, ftpのパス指定には"path"を使う(論理パス)
)

// homeに/gonfig/myappのディレクトリが作成される。
func main() {
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	// home/gonfig/myappを指定
	dir := filepath.Join(u.HomeDir, "gonfig", "myapp")
	//サブディレクトリも含めて作成したい時に使用する。
	// MkdirAll(path, permission as unix)
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		log.Fatal(err)
	}
}
