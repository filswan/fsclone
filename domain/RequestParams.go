package domain

type AWStoGS2 struct {
	GS2     GS2    `json:"GS2"`
	AWStest AWS_s3 `json:"AWStest"`
}

type OneDrivetoGS2 struct {
	GS2      GS2      `json:"GS2"`
	OneDrive OneDrive `json:"OneDrive"`
}

type GoogleDrivetoGS2 struct {
	GS2         GS2          `json:"GS2"`
	GoogleDrive Google_Drive `json:"Google_Drive"`
}

type AWS_s3 struct {
	ConfName           string `json:"confName"`
	Type               string `json:"type"`
	Provider           string `json:"provider"`
	EnvAuth            bool   `json:"env_auth"`
	AccessKeyID        string `json:"access_key_id"`
	SecretAccessKey    string `json:"secret_access_key"`
	Region             string `json:"region"`
	LocationConstraint string `json:"location_constraint"`
	Bucket             string `json:"bucket"`
	ACL                string `json:"acl"`
}

type GS2 struct {
	ConfName        string `json:"confName"`
	Type            string `json:"type"`
	Provider        string `json:"provider"`
	EnvAuth         bool   `json:"env_auth"`
	AccessKeyID     string `json:"access_key_id"`
	SecretAccessKey string `json:"secret_access_key"`
	Bucket          string `json:"bucket"`
	Endpoint        string `json:"endpoint"`
	ACL             string `json:"acl"`
}

type Google_Drive struct {
	ConfName string `json:"confName"`
	Type     string `json:"type"`
	Scope    string `json:"scope"`
	Bucket   string `json:"bucket"`
	Token    Token  `json:"token"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	Expiry       string `json:"expiry"`
}

type OneDrive struct {
	ConfName  string `json:"confName"`
	Type      string `json:"type"`
	Token     Token  `json:"token"`
	Bucket    string `json:"bucket"`
	DriveID   string `json:"drive_id"`
	DriveType string `json:"drive_type"`
}
