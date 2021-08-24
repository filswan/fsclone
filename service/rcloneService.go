package service

import (
	"fmt"
	"migrates3/common"
	"migrates3/logs"
	"strings"
)

func RcloneGs2SyncS3Drive(gs2DirectoryName string, s3DirectoryName string, GS2Struct string, S3Struct string) error {
	cmd := "rclone sync " + GS2Struct + ":" + gs2DirectoryName + " " + S3Struct + ":" + s3DirectoryName
	fmt.Println(cmd)
	_, errMsg, err := common.ExecCommand(cmd)
	if err != nil {
		logs.GetLogger().Error(err, errMsg) //nolint:goimports,gofmt
		return err
	}
	return nil
}
func RcloneS3DriveSyncGs2(gs2DirectoryName string, s3DirectoryName string, GS2Struct string, S3Struct string) error {
	cmd := "rclone sync " + S3Struct + ":" + s3DirectoryName + " " + GS2Struct + ":" + gs2DirectoryName
	fmt.Println(cmd)
	_, errMsg, err := common.ExecCommand(cmd)
	if err != nil {
		logs.GetLogger().Error(err, errMsg) //nolint:gofmt
		return err
	}
	return nil
}

func TimeCompare(gs2Time string, s3Time string, gs2DirectoryName string, s3DirectoryName string, GS2Struct string, S3Struct string) error {
	var err error
	if s3Time == "" {
		err = RcloneGs2SyncS3Drive(gs2DirectoryName, s3DirectoryName, GS2Struct, S3Struct)
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}
	} else if s3Time > gs2Time {
		err = RcloneS3DriveSyncGs2(gs2DirectoryName, s3DirectoryName, GS2Struct, S3Struct)
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}
	}
	if gs2Time == "" {
		err = RcloneS3DriveSyncGs2(gs2DirectoryName, s3DirectoryName, GS2Struct, S3Struct)
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}
	} else if gs2Time > s3Time {
		err = RcloneGs2SyncS3Drive(gs2DirectoryName, s3DirectoryName, GS2Struct, S3Struct)
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}
	} else {
		fmt.Println("GS2 Equal S3")
	}
	return nil
}

func GetAWSTime(awsDirectoryName string, awsConfName string) (string, error) {
	var err error
	cmd := "rclone lsl " + awsConfName + ":" + awsDirectoryName
	fmt.Println("cmd " + cmd)

	strs, errMsg, err := common.ExecCommand(cmd)
	if err != nil {
		logs.GetLogger().Error(err, errMsg)
		return "", err
	}
	lastFileTime, err := getLastFileTime(strings.Split(strings.TrimSpace(strs), "\n"))
	if err != nil {
		logs.GetLogger().Error(err)
		return "", err
	}

	return lastFileTime, err
}

func GetGoogleTime(googDirectoryName string, googConfName string) (string, error) {
	cmd := "rclone lsl " + googConfName + ":" + googDirectoryName
	fmt.Println("cmd " + cmd)
	strs, errMsg, err := common.ExecCommand(cmd)
	if err != nil {
		logs.GetLogger().Error(err, errMsg)
		return "", err
	}
	lastFileTime, err := getLastFileTime(strings.Split(strings.TrimSpace(strs), "\n"))
	if err != nil {
		logs.GetLogger().Error(err)
		return "", err
	}
	return lastFileTime, err
}

func GetOneDriveTime(oneDirectoryName string, oneConfName string) (string, error) {
	var err error
	cmd := "rclone lsl " + oneConfName + ":" + oneDirectoryName
	fmt.Println("cmd " + cmd)
	strs, err := common.ExecCommand2(cmd)
	if err != nil {
		logs.GetLogger().Error(err)
		return "", err
	}
	var oneTimeList []string
	fmt.Println(strs)
	for _, v := range strs {
		if v != "" {
			v := strings.TrimSpace(v)
			strsNew := strings.Split(v, " ")
			strTemp := strsNew[1] + " " + strsNew[2]
			fmt.Println("OneDrive file time =" + strTemp)
			oneTimeList = append(oneTimeList, strTemp)
		}
	}
	fmt.Println("OneDriveTimeList ", oneTimeList)
	common.BubbleSort(oneTimeList)
	if len(oneTimeList) <= 0 {
		return "", nil
	}
	oneTime := oneTimeList[0]
	fmt.Println("before" + oneTime)
	oneTime = strings.Replace(oneTime, ".000000000", "", -1)
	fmt.Println("after" + oneTime)
	oneTimeStamp, _ := common.TimeStamp(oneTime)
	fmt.Println("One Drive TimeStamp ", oneTimeStamp)
	return oneTimeStamp, err
}

func GetGs2Time(gs2DirectoryName string, gs2ConfName string) (string, error) {
	var err error
	cmd := "rclone lsl " + gs2ConfName + ":" + gs2DirectoryName
	fmt.Println("cmd " + cmd)
	strs, errMsg, err := common.ExecCommand(cmd)
	if err != nil {
		logs.GetLogger().Error(err, errMsg)
		return "", err
	}

	lastFileTime, err := getLastFileTime(strings.Split(strings.TrimSpace(strs), "\n"))
	if err != nil {
		logs.GetLogger().Error(err)
		return "", err
	}
	return lastFileTime, err
}

func getLastFileTime(allFileTimes []string) (lastTime string, err error) {
	var fileTimeList []string
	for _, v := range allFileTimes {
		if v != "" {
			v := strings.TrimSpace(v)
			strsNew := strings.Split(v, " ")
			strTemp := strsNew[1] + " " + strsNew[2]
			fileTimeList = append(fileTimeList, strTemp)
		}
	}
	fmt.Println("fileTimeList ", fileTimeList)
	common.BubbleSort(fileTimeList)
	if len(fileTimeList) <= 0 {
		return "", nil
	}
	//lastFileTime := fileTimeList[0]
	fmt.Println("before" + fileTimeList[0])
	tmpLastFileTime := strings.Replace(fileTimeList[0], ".000000000", "", -1)
	fmt.Println("after" + tmpLastFileTime)
	lastTimeStamp, err := common.TimeStamp(tmpLastFileTime)
	if err != nil {
		return "", err
	}
	return lastTimeStamp, nil
}
