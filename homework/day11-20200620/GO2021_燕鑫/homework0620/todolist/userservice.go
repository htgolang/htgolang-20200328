package todolist

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"homework0620/tools"
	"math/rand"
	"os"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type UserService struct {
	usercount int64
	lastid    int64
	db        *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	usersrv := &UserService{db: db}
	usersrv.SetUsercountAndLastId()
	return usersrv
}

func (u *UserService) GetAllUser() []*user {
	users := make([]*user, 0)
	u.db.Find(&users)
	return users
}

func (u *UserService) CreateUser(username string, password string) error {
	if username == "root" {
		return errors.New("YOU CAN'T CREATE USER NAMED ROOT!!!")
	}
	var ucount int
	u.db.Model(&user{}).Where("username=?", username).Count(&ucount)
	if ucount != 0 {
		return errors.New("The user has already existed!")
	}
	newuser := &user{
		Username:   username,
		Password:   "",
		Salt:       "",
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
		UpdateTime: time.Now().Format("2006-01-02 15:04:05"),
	}

	randbytes := make([]byte, 10)
	chars := "01234569abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i := 0; i < 10; i++ {
		randbytes[i] = chars[rand.Intn(len(chars))]
	}
	newuser.Salt = string(randbytes)

	randbytes = append(randbytes, []byte(password)...)
	passmd5 := md5.Sum(randbytes)
	newuser.Password = hex.EncodeToString(passmd5[:])
	u.db.Create(newuser)
	u.usercount++
	u.lastid++
	return nil
}

func (u *UserService) DeleteUser(user *user) error {
	err := u.db.Delete(user).Error
	if err != nil {
		return err
	}
	u.usercount--
	return nil
}

func (u *UserService) VerifyUser(username string, password string) (string, error) {
	//verify superuser
	if username == "root" {
		f, err := os.Open(tools.ROOTFILE)
		if err != nil {
			return "", err
		}
		br := bufio.NewReader(f)
		line, _, _ := br.ReadLine()
		linestr := strings.TrimSpace(string(line))
		superuserinfo := strings.Fields(linestr)
		if password == superuserinfo[1] {
			return "YOU ARE SUPERUSER!PLEASE OPERATE CAREFULLY!", nil
		} else {
			return "", errors.New("USERNAME or PASSWORD is not right,please enter again!")
		}
	}

	user := &user{}
	u.db.Where("username=?", username).Find(user)
	if user.Id == 0 {
		return "", errors.New("USERNAME or PASSWORD is not right,please enter again!")
	}

	passbytes := []byte(user.Salt)
	passbytes = append(passbytes, []byte(password)...)
	passmd5 := md5.Sum(passbytes)
	if user.Password == hex.EncodeToString(passmd5[:]) {
		return fmt.Sprintf("Welcome %s", username), nil
	} else {
		return "", errors.New("USERNAME or PASSWORD is not right,please enter again!")
	}
}

func (u *UserService) SetUsercountAndLastId() {
	var tableLines int64 = 0
	u.db.Model(&user{}).Count(&tableLines)
	user := &user{}
	u.db.Last(user)
	u.lastid = user.Id
	u.usercount = tableLines
}
