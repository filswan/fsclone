package service

import (
	"encoding/json"
	"fsclone/common"
	"fsclone/conf"
	"fsclone/domain"
	"fsclone/logs"
	"github.com/mitchellh/go-homedir"
	"gopkg.in/ini.v1"
	"strconv"
)

func getConfPath() string {
	rcloneConfigPath := conf.GetConfig().RcloneConfigPath
	if rcloneConfigPath == "" {
		rcloneConfigPath = common.RCLONE_DEFAULT_CONFIG_PATH
	}
	expandedDir, err := homedir.Expand(rcloneConfigPath)
	if err != nil {
		logs.GetLogger().Error(err)
		return ""
	}
	return expandedDir
}

func CreateAWStoGS2Conf(gs2AWS domain.AWStoGS2) error {
	cfg := ini.Empty()
	err := createGS2Conf(cfg, gs2AWS.GS2)
	if err != nil {
		return err
	}
	err = createAWSConf(cfg, gs2AWS.AWStest)
	if err != nil {
		return err
	}

	err = cfg.SaveTo(getConfPath())
	return err
}

func CreateGoogleDrivetoGS2Conf(googleDrivetoGS2 domain.GoogleDrivetoGS2) error {
	cfg := ini.Empty()
	err := createGoogleDriveConf(cfg, googleDrivetoGS2.GoogleDrive)
	if err != nil {
		return err
	}
	err = createGS2Conf(cfg, googleDrivetoGS2.GS2)
	if err != nil {
		return err
	}
	err = cfg.SaveTo(getConfPath())
	return err
}

func CreateOneDrivetoGS2Conf(oneDrivetoGS2 domain.OneDrivetoGS2) error {
	cfg := ini.Empty()
	err := createOneDriveConf(cfg, oneDrivetoGS2.OneDrive)
	if err != nil {
		return err
	}
	err = createGS2Conf(cfg, oneDrivetoGS2.GS2)
	if err != nil {
		return err
	}

	err = cfg.SaveTo(getConfPath())
	return err
}

func createGS2Conf(cfg *ini.File, gs2 domain.GS2) error {

	// add GS2 section
	var gs2name = gs2.ConfName
	_, err := cfg.NewSection(gs2name)
	if err != nil {
		return err
	}
	_, err = cfg.Section(gs2name).NewKey("type", gs2.Type)
	if err != nil {
		return err
	}
	_, err = cfg.Section(gs2name).NewKey("provider", gs2.Provider)
	if err != nil {
		return err
	}
	_, err = cfg.Section(gs2name).NewKey("env_auth", strconv.FormatBool(gs2.EnvAuth))
	if err != nil {
		return err
	}
	_, err = cfg.Section(gs2name).NewKey("access_key_id", gs2.AccessKeyID)
	if err != nil {
		return err
	}
	_, err = cfg.Section(gs2name).NewKey("secret_access_key", gs2.SecretAccessKey)
	if err != nil {
		return err
	}
	_, err = cfg.Section(gs2name).NewKey("endpoint", gs2.Endpoint)
	if err != nil {
		return err
	}
	_, err = cfg.Section(gs2name).NewKey("acl", gs2.ACL)
	if err != nil {
		return err
	}
	return err
}

func createAWSConf(cfg *ini.File, s3 domain.AWS_s3) error {
	var name = s3.ConfName
	_, err := cfg.NewSection(name)
	if err != nil {
		return err
	}
	_, err = cfg.Section(name).NewKey("type", s3.Type)
	if err != nil {
		return err
	}
	_, err = cfg.Section(name).NewKey("provider", s3.Provider)
	if err != nil {
		return err
	}
	_, err = cfg.Section(name).NewKey("env_auth", strconv.FormatBool(s3.EnvAuth))
	if err != nil {
		return err
	}
	_, err = cfg.Section(name).NewKey("region", s3.Region)
	if err != nil {
		return err
	}
	_, err = cfg.Section(name).NewKey("location_constraint", s3.LocationConstraint)
	if err != nil {
		return err
	}
	_, err = cfg.Section(name).NewKey("acl", s3.ACL)
	if err != nil {
		return err
	}
	_, err = cfg.Section(name).NewKey("access_key_id", s3.AccessKeyID)
	if err != nil {
		return err
	}
	_, err = cfg.Section(name).NewKey("secret_access_key", s3.SecretAccessKey)
	if err != nil {
		return err
	}
	return err
}

func createGoogleDriveConf(cfg *ini.File, googledrive domain.Google_Drive) error {
	var name = googledrive.ConfName
	_, err := cfg.NewSection(name)
	if err != nil {
		return err
	}
	_, err = cfg.Section(name).NewKey("type", googledrive.Type)
	if err != nil {
		return err
	}
	_, err = cfg.Section(name).NewKey("scope", googledrive.Scope)
	if err != nil {
		return err
	}
	bytes, err := json.Marshal(googledrive.Token)
	if err != nil {
		return err
	}
	token := string(bytes)
	_, err = cfg.Section(name).NewKey("token", token)
	if err != nil {
		return err
	}
	return err
}

func createOneDriveConf(cfg *ini.File, onedrive domain.OneDrive) error {
	var name = onedrive.ConfName
	_, err := cfg.NewSection(name)
	if err != nil {
		return err
	}
	_, err = cfg.Section(name).NewKey("type", onedrive.Type)
	if err != nil {
		return err
	}
	bytes, err := json.Marshal(onedrive.Token)
	if err != nil {
		return err
	}
	token := string(bytes)
	_, err = cfg.Section(name).NewKey("token", token)
	if err != nil {
		return err
	}
	_, err = cfg.Section(name).NewKey("drive_id", onedrive.DriveID)
	if err != nil {
		return err
	}
	_, err = cfg.Section(name).NewKey("drive_type", onedrive.DriveType)
	if err != nil {
		return err
	}
	return err
}
