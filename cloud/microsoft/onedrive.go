package microsoft

import (
	"fmt"
	"fsclone/logs"
	"fsclone/service"
)

func SyncDataWithOneDrive(gs2DirectoryName string, oneDirectoryName string, gs2ConfName string, oneConfName string) error {
	oneTime, err := service.GetOneDriveTime(oneDirectoryName, oneConfName)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	fmt.Println("oneTime" + oneTime)

	gs2Time, err := service.GetGs2Time(gs2DirectoryName, gs2ConfName)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	fmt.Println("gs2Time" + gs2Time)

	err = service.TimeCompare(gs2Time, oneTime, gs2DirectoryName, oneDirectoryName, gs2ConfName, oneConfName)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	return nil
}
