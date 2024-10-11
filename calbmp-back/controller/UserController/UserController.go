package UserController

import (
	"calbmp-back/Database"
	"calbmp-back/Params/UserParams"
	"calbmp-back/Res"
	"calbmp-back/dto"
	"calbmp-back/model"
	"calbmp-back/security"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

func Register(ctx *gin.Context) {
	DB := Database.GetDB()

	// 获取参数
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	// 数据验证
	if len(password) <= 6 {
		Res.FailMsg(ctx, "[!] The password must be at least six characters long")
		return
	}
	//log.Println(username, password)

	// 判断用户名是否存在
	if IsUsernameExist(DB, username) {
		Res.FailMsg(ctx, "[!] The user already exists")
		return
	}

	// 创建用户
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("[!] Encryption errors:", err)
		return
	}
	newUser := model.User{
		Username: username,
		Password: string(hashedPassword),
	}
	DB.Create(&newUser)

	// 返回结果
	Res.SuccessMsg(ctx, "[*]Register Successfully")
}

func Login(ctx *gin.Context) {
	// 获取参数
	DB := Database.GetDB()
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	log.Println("[!] Username:", username)
	log.Println("[!] Password:", password)

	// 判断用户名是否存在
	var user model.User
	DB.Debug().Where("username = ?", username).First(&user)
	log.Println("[!] Find out: ", user.Username, user.Password, user.ID)
	if user.ID == 0 {
		Res.FailMsg(ctx, "[!] User does not exist")
		return
	}

	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		Res.FailMsg(ctx, "[*] Password Error")
		return
	}

	// 发放token
	token, err := security.ReleaseToken(user)
	if err != nil {
		Res.FailMsg(ctx, "[!] System Error")
		log.Printf("token generate error : %v", err)
		return
	}

	// 返回结果
	Res.Success(ctx, gin.H{"token": token}, "[*] Login Successfully")
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	Res.Success(ctx, gin.H{"user": dto.ToUserDto(user.(model.User))}, "[*] Get User Info")
}

func IsUsernameExist(db *gorm.DB, username string) bool {
	var user model.User
	db.Where("username = ?", username).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

func ChangePassword(ctx *gin.Context) {
	var params UserParams.ChangePasswordParams
	err := ctx.BindJSON(&params)
	if err != nil {
		Res.FailMsg(ctx, "[ChangePassword] Params error")
		return
	}

	// check user existed
	var user model.User
	errFind := Database.DB.Where("username=?", params.Username).First(&user).Error
	if errFind == gorm.ErrRecordNotFound {
		log.Println("[ChangePassword] User not found")
		Res.FailMsg(ctx, "[ChangePassword] User not found")
		return
	}

	// make new password hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.ConfirmPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Println("[!] Encryption errors:", err)
		return
	}

	// change password
	user.Password = string(hashedPassword)

	// update
	Database.DB.Save(&user)

	Res.SuccessMsg(ctx, "[ChangePassword] successfully")
}
