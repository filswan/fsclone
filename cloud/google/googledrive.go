package google

import (
	"fmt"
	"migrates3/logs"
	"migrates3/service"
)

func SyncDataWithGoogleDrive(gs2DirectoryName string, googDirectoryName string, gs2ConfName string, googConfName string) error {
	googTime, err := service.GetGoogleTime(googDirectoryName, googConfName)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	fmt.Println("googTime" + googTime)

	gs2Time, err := service.GetGs2Time(gs2DirectoryName, gs2ConfName)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	fmt.Println("gs2Time" + gs2Time)

	err = service.TimeCompare(gs2Time, googTime, gs2DirectoryName, googDirectoryName, gs2ConfName, googConfName)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	return nil
}
