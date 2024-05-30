#!/bin/bash

echo "🔄 Копирование html шаблонов и статики"
cp -R ../internal/web/static/ ./static/
cp -R ../internal/web/templates/ ./templates/
go build ../cmd/hub/hub.go
chmod +x hub
./hub
echo "✅ Скрипт выполнен успешно"