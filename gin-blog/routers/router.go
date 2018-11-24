package routers

/* import (
	"test.go.dev/gin-blog/pkg/setting"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
	return r
}
 */
import (
    "net/http"
    "github.com/gin-gonic/gin"
    "test.go.dev/gin-blog/middleware"
    "test.go.dev/gin-blog/pkg/export"
    "test.go.dev/gin-blog/routers/api"
    "test.go.dev/gin-blog/routers/api/v1"
    "test.go.dev/gin-blog/pkg/setting"
    "test.go.dev/gin-blog/pkg/upload"
    "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)
func InitRouter() *gin.Engine {
    r := gin.New()
    r.Use(gin.Logger())
    r.Use(gin.Recovery())
    gin.SetMode(setting.ServerSetting.RunMode)

    r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
    r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))

    //获取jwt的token

    r.GET("/auth", api.GetAuth)
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    r.POST("/upload", api.UploadImage)
    //导出标签
    r.POST("/tags/export", v1.ExportTag)
    r.POST("/tags/import", v1.ImportTag)

   // apiv1 := r.Group("/api/v1")
   // {
   apiv1 := r.Group("/api/v1")


    apiv1.Use(jwt.JWT())
    {
        //获取标签列表
        apiv1.GET("/tags", v1.GetTags)



        //新建标签
        apiv1.POST("/tags", v1.AddTag)
        //更新指定标签
        apiv1.PUT("/tags/:id", v1.EditTag)
        //删除指定标签
        apiv1.DELETE("/tags/:id", v1.DeleteTag)


        apiv1.GET("/articles", v1.GetArticles)
        //获取指定文章
        apiv1.GET("/articles/:id", v1.GetArticle)
        //新建文章
        apiv1.POST("/articles", v1.AddArticle)
        //更新指定文章
        apiv1.POST("/articles/:id/edit", v1.EditArticle)
        //删除指定文
        apiv1.POST("/articles/:id/delete", v1.DeleteArticle)
    }
    return r
}