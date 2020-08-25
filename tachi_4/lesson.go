package main

import (
	"crypto/rand"
	"io/ioutil"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

type config struct {
	Token []byte
}

func main() {
	data, err := ioutil.ReadFile("config")
	if err != nil {
		panic(err)
	}

	/*
		config := &config{}
		err = xml.Unmarshal(data, config)
		if err != nil {
			panic(err)
		}
	*/

	bot, err := discordgo.New("Bot " + strings.ReplaceAll(string(data), "\r\n", ""))
	rand.Read(data)
	if err != nil {
		panic(err)
	}

	err = bot.Open()
	if err != nil {
		panic(err)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	/*
		_, err = bot.ChannelMessageSend("397094353257889806", "Hey this is a test mesage")
		if err != nil {
			return
		} */
	/*
		channel, err := bot.UserChannelCreate("187301662518411264")
		if err != nil {
			return
		}

		_, err = bot.ChannelMessageSend(channel.ID, "Taci made me do this")
		if err != nil {
			return
		}
	*/

	bot.Close()
	return
}
