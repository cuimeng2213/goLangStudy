# 分布式文件上传服务

### 要做什么？

- 基于golang语言实现分布式文件上传服务
- 重点结合开源存储ceph及公有云（阿里OSS）支持断电续传及妙传功能
- 微服务化及容器化部署（docker相关）

### 使用的工具

- Redis / RabbitMQ
- Docker （微服务部署）/ Kubernets（服务编排工具）
- 分布式对象存储服务Ceph
- 阿里云OSS对象存储服务

### 收获

- 文件分块断电上传与秒传功能
- 对象从Ceph迁移到阿里云OSS的经验

### 记录上传文件信息

```
type FileMeta struct {

}
```

### 文件元信息查询接口：

```
func FileQueryHandler(w http.ResponseWriter, r *http.Request) {
   r.ParseForm()
   filehash := r.Form.Get("filehash")
   //从缓存中获取对应文件的数据
}
```

### 文件下载接口

```
func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	filehash := r.Form.Get("filehash")
	//加载存储到云端本地存储的文件内容返回给客户端
}
```

文件Meta更新（重命名）接口

```
func FileUpdateMetaHandler(){
	r.ParseForm()
	file_hash := r.Form.Get("filehash")
	new_filename := r.Form.Get("filename")
	//文件重命名
}
```

文件删除接口

```
func DeleteFileHandler(){
	
}
```

## Mysql知识与文件元数据

### 文件表设计

```
CREATE TABLE `tbl_file`(
	`id` int(11) NOT NULL AUTO_INCREMENT,
	`file_sha1` char(40) NOT NULL DEFAULT '' COMMENT `文件hash`,
	`file_name` varchar(256) NOT NULL DEFAULT '' COMMENT `文件名`,
	`file_size` bigint(20) DEFAULT '0' COMMENT `文件大小`,
	`file_addr` varchar(1024) NOT NULL DEFAULT '' COMMENT `文件存储位置`,
	`create_at` datetime DEFAULT NOW() COMMENT `创建日期`,
	`update_at` datetime  DEFAULT NOW() on update current_timestamp() COMMENT `更新日期`,
	`status` int(11) NOT NULL DEFAULT '0' COMMENT `状态（可用/禁用/已删除等状态）`,
	`ext1` int(11) DEFAULT '0' COMMENT `备用字段1`,
	`ext2` text COMMENT `备用字段2`,
	PRIMARY KEY ('id'),
	UNIQUE KEY `idx_file_hash` (`file_sha1`),
	KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```

### Go访问Mysql

​	使用Go的标准接口：database/sql

```
package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var (
	db  *sql.DB
	err error
)

func init() {
	db, err = sql.Open("mysql", "root:cm2213@tcp(127.0.0.1:3306)/fileserver?charset=utf8")
	db.SetMaxOpenConns(1000)
	err = db.Ping()
	if err != nil {
		fmt.Printf("fail to connect mysql: %v \n", err)
		os.Exit(1)
		return
	}
}

//DBConn: 返回数据库连接对象
func DBConn() *sql.DB {
	return db
}

```

### 账号系统功能

支持用户注册登陆

支持session鉴权

用户数据资源隔离