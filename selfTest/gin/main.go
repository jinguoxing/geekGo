package main

import (
    "fmt"
    "geekGo/selfTest/gin/common/form_validator"
    "geekGo/selfTest/gin/middleware"
    "geekGo/selfTest/gin/vaildatorx"
    "github.com/gin-gonic/gin"
    "log"
    "net/http"
)

type loginRequest struct {
    Username string `form:"username" binding:"required"`
    Password string `form:"password" binding:"required,max=16,min=6"`
    // 需要使用自定义校验方法checkDate做参数校验的字段Date
    Date string `form:"date" binding:"required,datetime=2006-01-02,checkDate"`
}

func main() {
    // if err := transInit("zh"); err != nil {
    // 	fmt.Printf("init trans failed, err:%v\n", err)
    // 	return
    // }
    //err := form_validator.SetupValidator()
    //
    //if err != nil {
    //    fmt.Println(err)
    //}

    router := gin.Default()

    router.Use(middleware.Translations())
    router.POST("/user/login", login)

    errx := router.Run(":8888")
    if errx != nil {
        log.Println("failed")
    }
}

func login(c *gin.Context) {
    var req loginRequest
    form_validator.RegisterValidator(c)
    vaild, errs := vaildatorx.BindAndValid(c, &req)

    if !vaild {
        fmt.Printf("app.BindAndValid errs: %v", errs)
        c.JSON(http.StatusOK, gin.H{
            "msg": errs,
        })
        return
    }

    //if err := c.ShouldBind(&req); err != nil {
    //
    //    fmt.Printf("app.BindAndValid errs: %v", errs)
    //    c.JSON(http.StatusOK, gin.H{
    //        "msg": err.Error(),
    //    })
    //    return
    //v := c.Value("trans")
    //trans, _ := v.(ut.Translator)
    //// 获取validator.ValidationErrors类型的errors
    //errs, ok := err.(validator.ValidationErrors)
    //if !ok {
    //    // 非validator.ValidationErrors类型错误直接返回
    //    c.JSON(http.StatusOK, gin.H{
    //        "msg": err.Error(),
    //    })
    //    return
    //}
    //// validator.ValidationErrors类型错误则进行翻译
    //c.JSON(http.StatusOK, gin.H{
    //    "msg": errs.Translate(trans),
    //})
    //return
    //}
    //login 操作省略
    c.JSON(http.StatusOK, gin.H{
        "code": 0,
        "msg":  "success",
    })
}
