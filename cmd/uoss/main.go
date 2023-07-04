package main

import (
	"context"
	"fmt"
	"gitlab.mihoyo.com/infosys/uoss/internal/aliyunoss"
	"gitlab.mihoyo.com/ysTools/db"
	"log"
	"os"
	"time"
)

const defaultTimeout = 10 * time.Minute

func dieF(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}

func main() {
	const argNum = 4

	if len(os.Args) != argNum {
		dieF("Usage: %s <bucket> <name> <path>", os.Args[0])
	}

	name := os.Args[2]
	//path := os.Args[3]

	// 读取阿里云ram配置
	// cred, err := db.QueryCredential("aliyunram")
	// 读取阿里云ak配置
	cred, err := db.QueryCredential("aliyunoss")
	if err != nil {
		log.Print("Unable to get credfrom db failed: %s", err)
		log.Print("Use the default ak")
	} else {
		aliyunoss.AccessKeyID = cred.AccessKeyID
		aliyunoss.AccessKeySecret = cred.AccessKeySecret
		aliyunoss.Endpoint = cred.Endpoint
		aliyunoss.Bucket = cred.Bucket
	}

	c, err := aliyunoss.NewClient()
	if err != nil {
		dieF("fail to create client: %s", err)
	}

	// 列举文件，并将文件存入数据库。
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	info, err := c.ListObjectsOfCurrentDir(ctx, name)
	if err != nil {
		dieF("fail to save file: %w", err)
	}

	err = db.UpdateOSSObjects(aliyunoss.Bucket, info)
	if err != nil {
		dieF("fail to save file: %w", err)
	}

	fmt.Print("Successfully written data")
}
