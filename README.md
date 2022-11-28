# LineBot with meme

## About The Project
There is a meme line bot which you can get random mem image！

## Built With
+ HTTP framework: https://github.com/gin-gonic/gin
+ Config: https://github.com/spf13/viper
+ Mongo driver: https://github.com/mongodb/mongo-go-driver
+ Command line tools: https://github.com/spf13/cobra
+ Line Bot SDK: https://github.com/line/line-bot-sdk-go

## Getting Started
This is an example of how you setting up your project locally.

### Installation
+ [Install Go](https://go.dev/dl/)
+ [Install Docker Desktop](https://www.docker.com/products/docker-desktop/)
+ [Sign up Line Bot Offical Account](https://tw.linebiz.com/login/) 
+ [Sing up and install Ngrok](https://ngrok.com/)

### Run 
Below is an example of how you can run the meme line bot api.
1. Clone the repo
```
https://github.com/Terence1105/LineBot.git
```
2. Init Ngrok connect
```
ngrok http 8080
```
3. Init mongoDB, build go project and run exe file
```
make initLineBot
```
4. Init command line tool which you can get user message logs from mongoDB.
```
make cli
GetAllMessageLogCmd
GetMessageLogByUserIdCmd -u user_it
```

## Usage
Setting config with LINE_BOT_CHANNEL_TOKEN, and LINE_BOT_CHANNEL_SECRET
LineBot/config/config.yaml

Starting Ngrok with http 8080
![](https://clouding.city/tool/ngrok/http-8080.png)

Setting line bot webhook url 
![image](https://user-images.githubusercontent.com/61178586/204192417-7bb0150b-b487-4d98-83af-e52457bedaf3.png)

Send message to meme line bot, then you will get a meme image
![](https://i.imgur.com/DkWIwDu.png)

## Reference
+ [Line Message API Document](https://developers.line.biz/en/docs/messaging-api/overview/)
+ [Memes API](https://imgflip.com/api)

## Contact
Terence Shih - https://www.linkedin.com/in/terence-shih-077a721b1/