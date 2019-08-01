package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	toml "github.com/sioncojp/tomlssm"
	"go.uber.org/zap"
)

var (
	log   Logger
	botID string
	pid   = "/tmp/go-slack-auto-comment.pid"
)

// Logger ... zap logger
type Logger struct {
	logger *zap.Logger
	sugar  *zap.SugaredLogger
}

// Config ... tomlのデータ
type Config struct {
	BotID             string   `toml:"bot_id"`
	BotToken          string   `toml:"bot_token"`
	VerificationToken string   `toml:"verification_token"`
	Actions           []Action `toml:"action"`
}

// Action ... actionのデータ
type Action struct {
	ChannelID string `toml:"channel"`
	In        string `toml:"in"`
	Out       string `toml:"out"`
}

// LoadToml ... ディレクトリ配下にあるtomlファイルを読み込む
func LoadToml(dir, region string) (*Config, error) {
	// 末尾が / で終わってなければ追加
	if string(dir[len(dir)-1]) != "/" {
		dir = dir + "/"
	}

	// load config. ディレクトリ配下の設定ファイルを結合して読み込む
	files, _ := ioutil.ReadDir(dir)
	openFiles := make([]io.Reader, len(files)*2)

	// ファイル間の結合の際、改行を加える
	for i := 0; i < len(files); i++ {
		num := int(2 * float64(i))
		if i == 0 {
			num = 0
		}
		openFiles[num], _ = os.Open(fmt.Sprintf("%s%s", dir, files[i].Name()))
		openFiles[num+1] = strings.NewReader("\n")
	}

	reader := io.MultiReader(openFiles...)

	var config Config
	if _, err := toml.DecodeReader(reader, &config, region); err != nil {
		return nil, err
	}
	return &config, nil
}
