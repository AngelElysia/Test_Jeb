package util

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// 通过接口访问全局的嵌入文件系统
var StaticFS embed.FS

// SetStaticFS 由主程序调用，设置嵌入的文件系统
func SetStaticFS(embedFS embed.FS) {
	StaticFS = embedFS
}

// ExtractAssets 使用 embed 提取所有嵌入的静态资源到指定目录
func ExtractAssets(outputDir string) error {
	return fs.WalkDir(StaticFS, "static", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		
		// 跳过目录
		if d.IsDir() {
			return nil
		}

		// 读取嵌入的文件内容
		data, err := StaticFS.ReadFile(path)
		if err != nil {
			return err
		}

		// 构建目标路径
		destPath := filepath.Join(outputDir, path)
		
		// 创建必要的目录结构
		err = os.MkdirAll(filepath.Dir(destPath), os.ModePerm)
		if err != nil {
			return err
		}

		// 写入文件
		err = os.WriteFile(destPath, data, os.ModePerm)
		if err != nil {
			fmt.Printf("Write to the file failed, please close the ide and try again: %s:warning: %v\n", destPath, err)
			return err
		}
		
		return nil
	})
}

// Asset 提供与 go-bindata 兼容的接口，读取指定文件的内容
func Asset(name string) ([]byte, error) {
	// 确保路径以 static/ 开头
	if name != "" && name[0] != 's' {
		name = "static/" + name
	}
	return StaticFS.ReadFile(name)
}

// AssetNames 返回所有嵌入文件的路径列表，兼容 go-bindata 接口
func AssetNames() []string {
	var names []string
	err := fs.WalkDir(StaticFS, "static", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			names = append(names, path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error walking embedded files: %v\n", err)
	}
	return names
}

// AssetInfo 返回嵌入文件的信息，兼容 go-bindata 接口
func AssetInfo(name string) (fs.FileInfo, error) {
	// 确保路径以 static/ 开头
	if name != "" && name[0] != 's' {
		name = "static/" + name
	}
	
	file, err := StaticFS.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	
	return file.Stat()
}

// AssetDir 返回指定目录下的文件列表，兼容 go-bindata 接口
func AssetDir(name string) ([]string, error) {
	// 确保路径以 static/ 开头
	if name != "" && name[0] != 's' {
		name = "static/" + name
	}
	
	entries, err := StaticFS.ReadDir(name)
	if err != nil {
		return nil, err
	}
	
	var names []string
	for _, entry := range entries {
		names = append(names, entry.Name())
	}
	return names, nil
}