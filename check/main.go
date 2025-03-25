package main

import (
	"fmt"
	"go/parser"
	"go/ast"
	"go/token"
	"go/types"
)

func main() {
	src := `
		package main
		func main() {
			var x int = 10
		}
	`

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "", src, parser.AllErrors)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// 型情報を格納する
	info := &types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
	}

	// 型チェック
	conf := types.Config{}
	_, err = conf.Check("example", fset, []*ast.File{node}, info)
	if err != nil {
		fmt.Println("型チェックエラー:", err)
		return
	}

	// xの型を取得
	for expr, typ := range info.Types {
		fmt.Printf("式: %v, 型: %v\n", fset.Position(expr.Pos()), typ.Type)
	}
}