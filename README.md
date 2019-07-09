# go-slack-auto-comment
<img src="https://github.com/sioncojp/go-slack-auto-comment/blob/master/docs/go-slack-auto-comment01.png" width="300">

---

<img src="https://github.com/sioncojp/go-slack-auto-comment/blob/master/docs/go-slack-auto-comment02.png" width="300">


tomlに記載されたルールにマッチしたら、スレッドで指定されたコメントを返してくれます


## Prepare

Customize Slack -> API -> Your Apps -> Create New App

でAppを作り、

- Bot UsersをONにし、Web版SlackからBOT ID
- OAuth & PermissionsでBot User OAuth Access Token
- BasicInformationにあるVerification Token

を取得してください。

https://github.com/sioncojp/go-slack-auto-comment/tree/master/examples

を参考にconfigを設定してください。

Parameter Storeを使ってる場合は

`ssm://SSMのAlias名` と書けばIAM, KMS権限があればDecodeしてくれます。

## Usage

```shell
### build
$ make build

### help
$ ./bin/go-slack-auto-comment help
NAME:
   go-slack-auto-comment - A new cli application

USAGE:
   go-slack-auto-comment [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --config-dir value, -c value  Load configuration *.toml in target dir
   --region value, -r value      Setting AWS region for tomlssm (default: "ap-northeast-1")
   --help, -h                    show help
   --version, -v                 print the version

### run。configが入ってるdirectoryを指定
$ ./bin/go-slack-auto-comment -c examples/
{"level":"info","ts":1562750380.530795,"caller":"go-slack-auto-comment/main.go:64","msg":"start..."}
```

# License
The MIT License

Copyright Shohei Koyama / sioncojp 

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.