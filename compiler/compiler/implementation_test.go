package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// print directory recurisvely

func printDir(cwd string) {
	files, _ := os.ReadDir(cwd)
	fmt.Printf("dir: %s\n", cwd)
	for _, file := range files {
		fmt.Printf("  - %s \n", file.Name())
	}
}
func TestImplementation(t *testing.T) {
	// read cwd
	cwd, _ := os.Getwd()
	printDir(cwd)
	printDir(cwd + "/faux_mysql_dir")
	printDir(cwd + "/faux_mysql_dir/bin")
	printDir(cwd + "/faux_mysql_dir/lib")

	mysql_dir_envar := os.Getenv("MYSQL_DIR")

	fmt.Printf("MYSQL_DIR %s\n", mysql_dir_envar)
	testIfFileExists(mysql_dir_envar)
	// drop first directory segment
	mysql_from_envar_fixed := mysql_dir_envar[strings.Index(mysql_dir_envar, "/")+1:]
	fmt.Printf("MYSQL_DIR (fixed) %s\n", mysql_from_envar_fixed)
	// test if file exists
	testIfFileExists(mysql_from_envar_fixed)

	mysql, _ := filepath.Abs(cwd + "/faux_mysql_dir/bin/mysql")
	fmt.Printf("mysql: %s\n", mysql)

	testIfFileExists(mysql)

	t.Fail()
}

func testIfFileExists(filename string) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Printf("file does not exist: %s\n", filename)
	} else {
		fmt.Printf("file exists: %s\n", filename)
	}
}
