package logging

import (
	"fmt"
	"go-gin-example/pkg/file"
	"go-gin-example/pkg/setting"
	"os"
	"time"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.AppSetting.RuntimeRootPath, setting.AppSetting.LogSavePath)
}

func getLogFilename() string {
	return fmt.Sprintf("%s%s.%s",
		setting.AppSetting.LogSaveName,
		time.Now().Format(setting.AppSetting.TimeFormat),
		setting.AppSetting.LogFileExt,
	)
}

//var (
//	LogSavePath = "runtime/logs/"
//	LogSaveName = "log"
//	LogFileExt  = "log"
//	TimeFormat  = "20060102"
//)

//func getLogFilePath() string {
//	return fmt.Sprintf("%s", LogSavePath)
//}

//func getLogFileFullPath() string {
//	prefixPath := getLogFilePath()
//	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)
//
//	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
//}

func openLogFile(fileName, filePath string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v", err)
	}

	src := dir + "/" + filePath
	perm := file.CheckPermission(src)
	if perm == true {
		return nil, fmt.Errorf("file.CheckPermisson Permisson denied src: %s", src)
	}

	err = file.IsNotExistMkDir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}

	f, err := file.Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("Fail to OpenFile: %v", err)
	}

	return f, nil
}

//func mkDir() {
//	dir, _ := os.Getwd()
//	err := os.MkdirAll(dir+"/"+getLogFilePath(), os.ModePerm)
//	if err != nil {
//		panic(err)
//	}
//}
