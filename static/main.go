package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"go/types"
	"go/ast"
	"go/importer"
)

func main() {

	// 解析するソースコード
	src := `
		package main
		func main() {
			var v interface{} = "hello"
			n := v.(int) // 型アサーションのミス
			println(n)
		}
	`

	// 字句解析と構文解析
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "example.go", src, parser.AllErrors)
	if err != nil {
		fmt.Println("構文解析エラー:", err)
		return
	}

	// 型チェック
	conf := types.Config{Importer: importer.Default()}
	info := &types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
	}
	_, err = conf.Check("main", fset, []*ast.File{node}, info)
	if err != nil {
		fmt.Println("型チェックエラー:", err)
	} else {
		fmt.Println("型チェック成功")
	}

	// 型アサーションの検出と型チェック
	ast.Inspect(node, func(n ast.Node) bool {
		if assert, ok := n.(*ast.TypeAssertExpr); ok {
			// 型アサーションの左辺（インターフェース型）と右辺（アサーションする型）の確認
			xType := info.Types[assert.X].Type // 型アサーションの左辺の型
			if xType != nil {
				// 型アサーションの右辺の型情報を取得
				assertType := assert.Type
				if assertType != nil {
					// ast.TypeAssertExpr の右辺の型情報を取得
					var assertTypeInfo types.Type
					switch t := assertType.(type) {
					case *ast.Ident:
						// 型識別子の処理（例えば `int` や `string`）
						if t.Name == "int" {
							assertTypeInfo = types.Typ[types.Int]
						} else if t.Name == "string" {
							assertTypeInfo = types.Typ[types.String]
						} else {
							// 追加する型があればここで対応
							fmt.Printf("未知の型: %s\n", t.Name)
							return true
						}
					}
					// 型が一致しない場合、警告を出す
					if !types.AssignableTo(xType, assertTypeInfo) {
						fmt.Printf("警告: 型アサーションの型が一致しません。左: %s 右: %s\n", xType, assertTypeInfo)
					}
				}
			}
		}
		return true
	})
}
