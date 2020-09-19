package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

/*
	作成したファイルを削除する際には、作成した手順と逆に行わないとエラーになることがある。
	これはファイル操作に限らずある。Goではdeferが用意されている。
	deferを使用すると呼び出したスコープを抜ける際に、呼び出された順番と逆の順番で処理を実行する。

	今回の場合
	1.ディレクトリの作成
	2.ファイルストリームの作成
	3.ファイルストリームを閉じる
	4.ディレクトリごと削除する
*/
func makeFileandDelete() error {
	cwd, _ := os.Getwd()
	p := filepath.Join(cwd, "newdir")
	fmt.Println(p)
	err := os.MkdirAll(p, 0755)
	if err != nil {
		return err
	}
	time.Sleep(time.Second)
	// (2) 次にディレクトリが削除される
	defer os.RemoveAll(p)
	filenameWithPath := filepath.Join(p, "newfile.txt")
	fmt.Println(filenameWithPath)
	f, err := os.Create(filenameWithPath)
	if err != nil {
		return err
	}

	// (1) 最初にファイルハンドルが閉じられる。
	defer f.Close()
	time.Sleep(time.Second)
	return nil
}

func main() {
	makeFileandDelete()
}
