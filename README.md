#Migration With S3 storage

A component of Synchronized data between fs3 and googledrive,or fs3 between aws3.
It compares the latest time of the file in you specified bucket in s3 and google drive or aws3,
Then synchronize the bucket with the updated time to the bucket with the older time.

## Getting Started
1. &ensp; install rclone 
```console
#sudo apt update
#sudo apt install rclone
```
2. &ensp;  clone code to $GOPATH/src
```console
#git clone http://192.168.88.183/filswan/gs2migrations3
```
3. &ensp; find the config file of gs2migrations3.
```console
#cd $GOPATH/src/gs2migrations3/conf/
#vim conf.toml
```
4. &ensp; input raclone config path to the config file.<br>
&ensp; rclone default config path is : ~/.config/rclone/rclone.conf<br>
&ensp; if your rclone config path is default path,<br>
&ensp; You can ignore the steps to modify the configuration above.<br>
```console
rcloneConfigPath = "~/.config/rclone/rclone.conf"
```
5. &ensp; run gs2migrations3
```console
#go run ./main/main.go
```
6. call the api <br>
   Put the following address in the address bar of postman, modify the json data in the param example to your own data, <br>
   put it in the request body of postman, and then send the request using the post method. <br>


   ##### http://127.0.0.1:8083/api/v1/migration/sync/GS2SyncGoogleDrive
   - [param example](http://192.168.88.183/filswan/gs2migrations3/blob/master/postman/GS2syncGoogleDrive.json) - 💌 show the request body for sync between s3 and google drive.


   ##### http://127.0.0.1:8083/api/v1/migration/sync/GS2SyncAWS3 
   - [param example](http://192.168.88.183/filswan/gs2migrations3/blob/master/postman/GS2syncAWS.json) - 💌 show the request body for sync between s3 and aws3. 


## Docker usage
in the os terminal,run cmd below
```console
#make build
#docker build -t nbai/rclone .
#docker run -p 8083:8083 nbai/rclone
```

## Versioning and Releases

Feature branches and master are designated as **unstable** which are internal-only development builds.

Periodically a build will be designated as **stable** and will be assigned a version number by tagging the repository
using Semantic Versioning in the following format: `vMajor.Minor.Patch`.

## Other Topics
- [how to apply google drive api token](http://192.168.88.212:8090/display/PD/how+to+apply+google+drive+api+token) 
- [how to apply aws s3 api token](http://192.168.88.212:8090/display/PD/how+to+apply+aws+s3+token)

