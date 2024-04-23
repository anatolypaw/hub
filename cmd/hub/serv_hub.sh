GOARCH=amd64 GOOS=linux go build hub.go
upx hub
lftp -p22 sftp://$USER:$PASS@10.0.4.5:~  -e "put hub; bye"
rm hub