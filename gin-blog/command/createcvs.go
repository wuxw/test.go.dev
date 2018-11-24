package command

import (
	"encoding/csv"
	"fmt"
	"os"
	"test.go.dev/gin-blog/pkg/export"
	"test.go.dev/gin-blog/pkg/setting"
)

func main() {
	setting.Setup();

	fmt.Println(export.GetExcelFullPath()+"test.csv")
	f, err := os.Create(export.GetExcelFullPath()+"test.csv")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer  f.Close()
	//\xEF\xBB\xBF 是 UTF-8 BOM 的 16 进制格式 在这里的用处是标识文件的编码格式，通常会出现在文件的开头，因此第一步就要将其写入。如果不标识 UTF-8 的编码格式的话，写入的汉字会显示为乱码
	f.WriteString("\xEF\xBB\xBF")

	w := csv.NewWriter(f)
	data := [][]string{
		{"1" , "test1" , "test1-1"},
		{"2" , "test2" , "test2-1"},
		{"3" , "test3" , "test3-1"},
	}

	w.WriteAll(data)
}