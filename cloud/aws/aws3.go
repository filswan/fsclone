package aws

import (
	"fmt"
	"migrates3/logs"
	"migrates3/service"
)

func SyncDataWithAWS3(gs2DirectoryName string, awsDirectoryName string, gs2ConfName string, awsConfName string) error {

	awsTime, err := service.GetAWSTime(awsDirectoryName, awsConfName)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	fmt.Println("awsTime" + awsTime)

	gs2Time, err := service.GetGs2Time(gs2DirectoryName, gs2ConfName)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	fmt.Println("gs2Time" + gs2Time)

	err = service.TimeCompare(gs2Time, awsTime, gs2DirectoryName, awsDirectoryName, gs2ConfName, awsConfName)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	return nil
}
