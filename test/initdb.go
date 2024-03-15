package test

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

var curDir string
var dbDir string

func init() {
	/* 设置数据库目录 */
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		log.Panic()
	}
	curDir = filepath.Dir(file)
	dbDir = filepath.Join(filepath.Dir(file), "testdb")
}

func DbDir() string {
	return dbDir
}

func InitDB() error {
	/* 如果目录存在，删除目录 */
	err := os.RemoveAll(dbDir)
	if err != nil {
		return err
	}

	/* 执行初始化命令 */
	initdb := fmt.Sprintf("go run . -d %s", dbDir)
	cmd := exec.Command("cmd", "/C", initdb)
	cmd.Dir = filepath.Join(curDir, "..", "cmd", "initdb")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
		return err
	}
	return nil
}
