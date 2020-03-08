package db

import (
	"fmt"
	mydb "goLangStudy/cloudStore/codes/db/mysql"
)

type User struct {
	UserName   string
	LastActive string
	LoginAt    string
}

// UserSignup : 新建用户
func UserSignup(username, pwd string) bool {
	stmt, err := mydb.DBConn().Prepare("insert ignore into tbl_user(`user_name`,`user_pwd`) values(?,?)")
	defer stmt.Close()
	if err != nil {
		fmt.Printf("UserSignup error: %v\n", err)
		return false
	}
	ret, err := stmt.Exec(username, pwd)
	if err != nil {
		fmt.Printf("failed to insert: %v \n", err)
		return false
	}
	//检测是否插入成功
	if rf, err := ret.RowsAffected(); nil == err && rf > 0 {
		return true
	}
	return false
}

func UserSignin(username, password string) bool {
	stmt, err := mydb.DBConn().Prepare("select * from tbl_user where username=? limit 1")
	if err != nil {
		return false
	}
	rows, err := stmt.Query(username)
	if err != nil {
		return false
	} else if rows == nil {
		return false
	}
	pRows := mydb.ParseRows(rows)
	if len(pRows) > 0 && string(pRows[0]["user_pwd"].([]byte)) == password {
		return true
	}
	return false
}

func UpdateToken(username string, token string) bool {
	stmt, err := mydb.DBConn().Prepare("replace into tbl_token(`user_name`,`user_token`) values(?,?)")
	if err != nil {
		fmt.Printf("Prepare tbl_token sql failed: %v \n", err)
		return false
	}
	_, err = stmt.Exec(username, token)
	if err != nil {
		fmt.Printf("replace into tbl_token failed: %v \n", err)
		return false
	}

	return true
}
func GetUserinfo(username string) (User, error) {
	stmt, err := mydb.DBConn().Prepare("select `user_name`, `login_at` form tbl_user where user_name=?")
	if err != nil {
		return User{}, err
	}
	user := User{}
	_, err := stmt.Exec(username).Scan(&user.UserName, &user.LoginAt)
	if err != nil {
		return User{}, err
	}
	return user, nil
}
