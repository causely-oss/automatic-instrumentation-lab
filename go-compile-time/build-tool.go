package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"os/exec"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Usage: %s <go-command> [args...]\n", os.Args[0])
		os.Exit(1)
	}

	// Check if this is a build command
	if args[0] == "build" {
		instrumentFibonacci()
	}

	// Call the real go tool
	cmd := exec.Command("go", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		os.Exit(1)
	}
}

func instrumentFibonacci() {
	// Parse the source file
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "app/fibonacci.go", nil, parser.ParseComments)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing file: %v\n", err)
		return
	}

	// Check if already instrumented
	for _, decl := range file.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok && fn.Name.Name == "trace_fibonacci" {
			fmt.Println("Already instrumented: app/fibonacci.go")
			return
		}
	}

	// Find the fibonacci function and add defer
	ast.Inspect(file, func(n ast.Node) bool {
		if fn, ok := n.(*ast.FuncDecl); ok && fn.Name.Name == "fibonacci" {
			// Add defer statement
			deferStmt := &ast.DeferStmt{
				Call: &ast.CallExpr{
					Fun: &ast.CallExpr{
						Fun: &ast.Ident{Name: "trace_fibonacci"},
					},
				},
			}
			fn.Body.List = append([]ast.Stmt{deferStmt}, fn.Body.List...)
		}
		return true
	})

	// Add timing function
	timingFunc := &ast.FuncDecl{
		Recv: nil,
		Name: &ast.Ident{Name: "trace_fibonacci"},
		Type: &ast.FuncType{
			Params: &ast.FieldList{},
			Results: &ast.FieldList{
				List: []*ast.Field{
					{
						Type: &ast.FuncType{
							Params:  &ast.FieldList{},
							Results: &ast.FieldList{},
						},
					},
				},
			},
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.AssignStmt{
					Lhs: []ast.Expr{&ast.Ident{Name: "start"}},
					Tok: token.DEFINE,
					Rhs: []ast.Expr{
						&ast.CallExpr{
							Fun: &ast.SelectorExpr{
								X:   &ast.Ident{Name: "time"},
								Sel: &ast.Ident{Name: "Now"},
							},
						},
					},
				},
				&ast.ReturnStmt{
					Results: []ast.Expr{
						&ast.FuncLit{
							Type: &ast.FuncType{
								Params:  &ast.FieldList{},
								Results: &ast.FieldList{},
							},
							Body: &ast.BlockStmt{
								List: []ast.Stmt{
									&ast.AssignStmt{
										Lhs: []ast.Expr{&ast.Ident{Name: "duration"}},
										Tok: token.DEFINE,
										Rhs: []ast.Expr{
											&ast.CallExpr{
												Fun: &ast.SelectorExpr{
													X:   &ast.Ident{Name: "time"},
													Sel: &ast.Ident{Name: "Since"},
												},
												Args: []ast.Expr{&ast.Ident{Name: "start"}},
											},
										},
									},
									&ast.ExprStmt{
										X: &ast.CallExpr{
											Fun: &ast.SelectorExpr{
												X:   &ast.Ident{Name: "fmt"},
												Sel: &ast.Ident{Name: "Printf"},
											},
											Args: []ast.Expr{
												&ast.BasicLit{
													Kind:  token.STRING,
													Value: "\"Function fibonacci took: %v\\n\"",
												},
												&ast.Ident{Name: "duration"},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	file.Decls = append(file.Decls, timingFunc)

	// Add time import to existing import declaration
	var importDecl *ast.GenDecl
	for _, decl := range file.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok && genDecl.Tok == token.IMPORT {
			importDecl = genDecl
			break
		}
	}

	// Add time import to existing declaration
	importDecl.Specs = append(importDecl.Specs, &ast.ImportSpec{Path: &ast.BasicLit{Kind: token.STRING, Value: "\"time\""}})

	// Write the modified file back
	f, err := os.Create("app/fibonacci.go")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating file: %v\n", err)
		return
	}
	defer f.Close()

	printer.Fprint(f, fset, file)
	fmt.Println("Instrumented: app/fibonacci.go")
}
