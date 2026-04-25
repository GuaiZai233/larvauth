package auth

import (
	"errors"
	"github.com/GuaiZai233/larvauth/internal/db"
	"regexp"
)

// username must start with a letter, be 3-20 chars long, contain only letters, digits or hyphen,
// and cannot end with a hyphen. Underscore is not allowed.
var usernameRegexp = regexp.MustCompile(`^[A-Za-z](?:[A-Za-z0-9]|-(?=[A-Za-z0-9])){2,19}$`)

func LoginUser(username string, password string) (string, error) {
	// get DB
	database := db.GetDB()

	// query if user exists
	user, err := FindUserByAccount(database, username)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("用户名/邮箱或密码错误") // 用户不存在
	}

	// verify password
	if !CheckPswHash(password, user.PasswordHash) {
		return "", errors.New("账号或密码错误")
	}

	// generate token
	token, err := GenerateToken(user.ID, user.Username)
	if err != nil {
		return "", errors.New("生成令牌失败")
	}

	return token, nil

}

func RegisterUser(username, password, email string) error {
	if !usernameRegexp.MatchString(username) {
		return errors.New("用户名只能包含字母、数字和连字符，且不能以连字符开头或结尾")
	}
	// get DB handle
	database := db.GetDB()

	exists, err := IsUserExist(database, username)
	if err != nil {
		return errors.New("数据查询失败")
	}
	if exists != nil {
		return errors.New("用户名已存在")
	}

	hashPassword, err := HashPassword(password)
	if err != nil {
		return errors.New("密码加密失败")
	}

	err = CreateUser(database, username, hashPassword, email)
	if err != nil {
		return err
	}

	return nil
}
