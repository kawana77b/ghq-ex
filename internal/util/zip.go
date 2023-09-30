package util

import (
	"archive/zip"
	"errors"
	"fmt"
	"os"
	"time"
)

// zipファイルを作成するビルダー
type ZipFileBuilder struct {
	files []struct {
		name     string
		realPath string
	}
}

func NewZipFileBuilder() *ZipFileBuilder {
	builder := &ZipFileBuilder{}
	builder.files = make([]struct {
		name     string
		realPath string
	}, 0)

	return builder
}

// filesの内容を文字列で返す
func (b *ZipFileBuilder) String() string {
	str := "ZipFileBuilder{\n"
	for _, file := range b.files {
		str += fmt.Sprintf("File{file: %s, realPath: %s},\n", file.name, file.realPath)
	}
	str += "\n}"

	return str
}

// ファイルを追加する
// name: zipに含めるファイル名. / で区切るとディレクトリを作成できる
// realPath: 実際のファイルパス
func (b *ZipFileBuilder) AddFile(name string, realPath string) {
	b.files = append(b.files, struct {
		name     string
		realPath string
	}{
		name:     name,
		realPath: realPath,
	})
}

// 指定の場所にzipファイルを作成する
func (b *ZipFileBuilder) Create(dst string) error {
	if len(b.files) == 0 {
		return errors.New("no files to zip")
	}

	if len(dst) == 0 {
		return errors.New("destination is empty")
	}

	// ファイルの存在チェック
	for _, file := range b.files {
		fi, err := os.Stat(file.realPath)
		if err != nil {
			return err
		}

		if fi.IsDir() {
			return errors.New("directory is not supported")
		}
	}

	files := make([]struct {
		Name     string
		RealPath string
	}, 0)

	for _, file := range b.files {
		files = append(files, struct {
			Name     string
			RealPath string
		}{
			Name:     file.name,
			RealPath: file.realPath,
		})
	}

	return createZipFile(dst, files)
}

// zipファイルを作成する
func createZipFile(dst string, fileInfos []struct {
	Name     string // ZIPファイルに含めるファイル名. ディレクトリ名を含めることもできる
	RealPath string // 実際のファイルパス
}) error {
	zipFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for _, file := range fileInfos {
		// file.RealPathからファイル内容を読み込む
		fileReader, err := os.Open(file.RealPath)
		if err != nil {
			return err
		}
		defer fileReader.Close()

		content := make([]byte, 0)
		fileReader.Write(content)

		// ファイル内容をZIPアーカイブに追加
		fileWriter, err := zipWriter.CreateHeader(&zip.FileHeader{
			Name:     file.Name,
			Method:   zip.Deflate,
			Modified: time.Now(),
		})
		if err != nil {
			return err
		}

		_, err = fileWriter.Write(content)
		if err != nil {
			return err
		}
	}

	return nil
}
