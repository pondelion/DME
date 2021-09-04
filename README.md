# DME

## Build (android/arm64 target)

```bash
$ GOOS=android GOARCH=arm64 GOARM=7 go build -o droid_dme_service
```

## Run API server on android

```bash
$ adb push droid_dme_service /data/local/tmp
$ adb shell
$ su
# cd /data/local/tmp
# chmod 777 droid_dme_service
# ./droid_dme_service
```
