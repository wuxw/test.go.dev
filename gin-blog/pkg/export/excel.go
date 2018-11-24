package export
import "test.go.dev/gin-blog/pkg/setting"
func GetExcelFullUrl(name string) string {
	return setting.AppSetting.PrefixUrl + "/" + GetExcelPath() + name
}
func GetExcelPath() string {
	return setting.AppSetting.ExportSavePath
}
func GetExcelFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetExcelPath()
}