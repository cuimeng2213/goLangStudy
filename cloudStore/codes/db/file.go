package db

import (
	"fmt"
	mydb "goLangStudy/cloudStore/codes/db/mysql"
)

type TableFile struct {
	FileHash string
	FileName string
	FileSize int64
	FileAddr string
}

func OnFileUploadFinished(filehash string, filename string, filesize int64, fileaddr string) bool {
	stmt, err := mydb.DBConn().Prepare("insert ignore into tbl_file(`file_sha1`,`file_name`,`file_size`,`file_addr`, `status`) values(?,?,?,?,1)")
	defer stmt.Close()
	if err != nil {
		fmt.Println("failed to prepare statement, err: ", err.Error())
		return false
	}

	ret, err := stmt.Exec(filehash, filename, filesize, fileaddr)
	if err != nil {
		fmt.Println("failed insert into tbl_file, err: ", err)
		return false
	}
	//判断是否插入重复了
	if rf, err := ret.RowsAffected(); nil == err {
		if rf <= 0 {
			fmt.Printf("FIle with hash %s has been upload before\n", filehash)
		}
		return true
	}
	return false
}

func GetFileMeta(filehash string) (*TableFile, error) {
	stmt, err := mydb.DBConn().Prepare("select file_sha1, file_addr, file_name, file_size from tbl_file where file_sha1=? and status=1 limit 1")
	defer stmt.Close()
	if err != nil {
		fmt.Printf("GetFileMeta error: %v\n", err)
		return nil, err
	}
	tfile := TableFile{}
	//将读取的数据赋值给TableFIle对象
	err = stmt.QueryRow(filehash).Scan(&tfile.FileHash, &tfile.FileAddr, &tfile.FileName, &tfile.FileSize)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &tfile, nil
}
