package flag

import (
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	"golang.org/x/term"
	"os"
	"server/global"
	"server/model/appTypes"
	"server/model/database"
	"server/utils"
	"syscall"
)

// Admin 用于创建一个管理员用户
func Admin() error {
	var user database.User
	// 提示用户输入邮箱
	fmt.Print("输入邮箱：")
	var email string
	_, err := fmt.Scanln(&email)
	if err != nil {
		return fmt.Errorf("")
	}
	user.Email = email
	// 获取标准输入的文件描述符
	fd := int(syscall.Stdin)
	// 关闭回显，确保密码不会在终端显示
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		return err
	}
	defer term.Restore(fd, oldState) // 恢复终端状态
	fmt.Print("输入密码：")
	// 读取第一次输入的密码
	password, err := readPassword()
	fmt.Println()
	if err != nil {
		return err
	}
	fmt.Println("再次输入确认密码：")
	rePassword, err := readPassword()
	fmt.Println()
	if err != nil {
		return err
	}
	if password != rePassword {
		return errors.New("前后密码不匹配。")
	}
	if len(password) < 3 || len(password) > 20 {
		return errors.New("密码应该在8到20个字符内。")
	}
	// 填充用户数据
	user.UID = uuid.Must(uuid.NewV4())
	user.Username = global.Config.Website.Name
	user.Password = utils.BcryptHash(password)
	user.RoleID = appTypes.Admin
	user.Avatar = "/image/avatar.jpg"
	user.Address = global.Config.Website.Address
	// 在数据库中创建管理员用户
	if err := global.DB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

// readPassword 用于读取密码并且避免回显
func readPassword() (string, error) {
	var password string
	var buf [1]byte
	// 持续读取字符直到遇到换行符为止
	for {
		_, err := os.Stdin.Read(buf[:])
		if err != nil {
			return "", err
		}
		char := buf[0]
		// 检查是否为回车键，若是则终止输入
		if char == '\n' || char == '\r' {
			break
		}
		// 将输入的字符附加到密码字符串中
		password += string(char)
	}
	return password, nil
}
