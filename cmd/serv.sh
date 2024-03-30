GOARCH=amd64 GOOS=linux go build hub.go

lftp -p22 sftp://$USER:$PASS@192.168.11.148:~  -e "put hub; bye"
rm hub