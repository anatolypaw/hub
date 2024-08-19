npm install --prefix ../internal/api/http_web/webpanel/
npm run build --prefix ../internal/api/http_web/webpanel/
go build -o hub ../cmd/hub/hub.go
./hub