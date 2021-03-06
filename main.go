// amigobot is the default implementation for the amigo IRC bot.
package main

import (
	"flag"

	"github.com/leonelquinteros/amigo"
	"github.com/leonelquinteros/gorand"
)

func main() {
	defaultChannel := "#amigobot"
	defaultNick, _ := gorand.GetAlphaString(8)
	defaultPassword, _ := gorand.GetAlphaNumString(32)

	// Load Config
	host := flag.String("h", "chat.freenode.net:6667", "Host: The IRC host to connect to")
	channel := flag.String("c", defaultChannel, "Channel: The IRC #channel to join after connect to the IRC server")
	nick := flag.String("n", defaultNick, "Nick: The nick to use on the IRC server")
	password := flag.String("p", defaultPassword, "Password: The master password to perform authenticated commands without beign a registered master of the bot")
	flag.Parse()

	if *host == "" || *channel == "" || *nick == "" || *password == "" {
		println("Usage:")
		println("amigo -h=irc.host.com:6667 -c=#channel_name -n=nick -p=masterpassword")
		return
	}

	a := new(amigo.Amigo)
	a.EhAmigo(*host, *channel, *nick, *password)
}
