package todolist

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type UserService struct {
	users      []*user
	accounturl string
	usercount  int64
	lastid     int64
}

func NewUserService(accounturl string) *UserService {
	usersrv := &UserService{accounturl: accounturl}
	usersrv.users = usersrv.getAllUser()
	if usersrv.users == nil {
		usersrv.lastid = 0
	} else {
		usersrv.lastid, _ = strconv.ParseInt(usersrv.users[len(usersrv.users)-1].Id, 0, 0)
	}
	return usersrv
}

func (u *UserService) getAllUser() []*user {
	_, err := os.Stat(u.accounturl)
	if os.IsNotExist(err) {
		f, _ := os.Create(u.accounturl)
		f.Close()
		return nil
	}
	f, err := os.Open(u.accounturl)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer f.Close()
	users := make([]*user, 0)
	br := bufio.NewReader(f)
	for {
		var oneuser user
		line, _, eof := br.ReadLine()
		if eof == io.EOF {
			break
		}
		_ = json.Unmarshal(line, &oneuser)
		users = append(users, &oneuser)
		u.usercount++
	}
	return users
}

func (u *UserService) CreateUser(username string, password string) error {
	if username == "root" {
		return errors.New("YOU CAN'T CREATE USER NAMED ROOT!!!")
	}
	for _, user := range u.users {
		if user.Username == username {
			return errors.New("The users exists, please change another username!")
		}
	}
	newuser := user{
		strconv.FormatInt(u.lastid+1, 10),
		username,
		"",
		"",
		time.Now().Format("2006-01-02 15:04:05"),
		time.Now().Format("2006-01-02 15:04:05"),
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

	f, err := os.OpenFile(u.accounturl, os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	br := bufio.NewWriter(f)
	userline, _ := json.Marshal(newuser)
	_, _ = br.Write(userline)
	_, _ = br.WriteString("\n")
	_ = br.Flush()

	u.users = append(u.users, &newuser)
	u.usercount++
	u.lastid++
	return nil
}

func (u *UserService) DeleteUser(username string) error {
	if u.usercount == 0 {
		return errors.New("The specified users doesn't exist!")
	}
	users := u.users
	for index, user := range users {
		if user.Username == username {
			for i := index; i < len(users)-1; i++ {
				users[i] = users[i+1]
			}
			break
		}
		if index == len(users)-1 {
			return errors.New("The specified users doesn't exist!")
		}
	}
	users = users[:len(users)-1]
	f, err := os.OpenFile(u.accounturl, os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	br := bufio.NewWriter(f)
	for _, user := range users {
		userline, _ := json.Marshal(*user)
		_, _ = br.Write(userline)
		_, _ = br.WriteString("\n")
		_ = br.Flush()
	}
	u.users = users
	u.usercount--
	return nil
}

func (u *UserService) VerifyUser(username string, password string) (string, error) {
	//verify superuser
	if username == "root" {
		f, err := os.Open(ROOTFILE)
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

	for _, user := range u.users {
		if user.Username == username {
			passbytes := []byte(user.Salt)
			passbytes = append(passbytes, []byte(password)...)
			passmd5 := md5.Sum(passbytes)
			if user.Password == hex.EncodeToString(passmd5[:]) {
				return fmt.Sprintf("Welcome %s", username), nil
			} else {
				return "", errors.New("USERNAME or PASSWORD is not right,please enter again!")
			}
		}
	}
	return "", errors.New("USERNAME or PASSWORD is not right,please enter again!")
}
