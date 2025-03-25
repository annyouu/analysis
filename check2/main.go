package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
)

func main() {
	// 解析対象のコード
	src := `
		package mypkg

		func f() { println("f") }
		func G() { println("G") }
		func d() { println("d") }

		func main() {
			G()
			d()
		}
	`

	// ファイルセットを作成
	fset := token.NewFileSet()

	// ソースコードを解析してASTを取得する
	node, err := parser.ParseFile(fset, "example.go", src, parser.AllErrors)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 型情報を格納する構造体を定義
	info := &types.Info{
		Defs: make(map[*ast.Ident]types.Object),
		Uses: make(map[*ast.Ident]types.Object),
	}

	// 型チェックを行うための設定
	conf := types.Config{
		Importer: importer.Default(),
	}

	// 型チェックを実行
	_, err = conf.Check("example", fset, []*ast.File{node}, info)
	if err != nil {
		fmt.Println("Typeエラー:", err)
		return
	}

	// info.Usesにあるオブジェクトのセットを作る
	usedObjs := make(map[types.Object]bool)
	for _, obj := range info.Uses {
		usedObjs[obj] = true
	}

	// 使用されていない識別子の検出を行う
	for _, obj := range info.Defs {
		if obj == nil {
			continue
		}

		// パッケージ名を除外
		if _, ok := obj.(*types.PkgName); ok {
			continue
		}

		// main関数は除外
		if fn, ok := obj.(*types.Func); ok {
			if fn.Name() == "main" {
				continue
			}
		}

		// 定義されたやつが使われているかを確認する
		if !usedObjs[obj] {
			fmt.Printf("使用されていない識別子: %s\n", obj.Name())
		}
	}
}
