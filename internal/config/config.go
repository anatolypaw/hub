package config

import (
	"bytes"
	"encoding/json"
	"os"
)

type Config struct {
	filename string
	P        Params
}

type Params struct {
	MongoUri string
	DbName   string
}

// Значения по умолчанию для конфигурации
var DefaultConfig = Params{
	MongoUri: "mongodb://localhost:27017/",
	DbName:   "molocode",
}

func New(file string) Config {
	return Config{
		filename: file,
		P:        Params{},
	}
}

// Функция для загрузки конфигурации из файла
func (c *Config) Load() error {
	var config Params

	// Попытка чтения файла конфигурации
	file, err := os.ReadFile(c.filename)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(file)

	// Декодируем содержимое файла JSON в структуру конфигурации
	decoder := json.NewDecoder(reader)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&config)
	if err != nil {
		return err
	}

	c.P = config
	return nil
}

// Функция для сохранения конфигурации в файл
func (c *Config) Save() error {
	// Кодируем конфигурацию в формат JSON
	encodedConfig, err := json.MarshalIndent(c.P, "", "    ")
	if err != nil {
		return err
	}

	// Записываем конфигурацию в файл
	err = os.WriteFile(c.filename, encodedConfig, 0644)
	if err != nil {
		return err
	}

	return nil
}
