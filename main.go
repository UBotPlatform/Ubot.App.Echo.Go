package main

import (
	"regexp"

	ubot "github.com/UBotPlatform/UBot.Common.Go"
)

var api *ubot.AppApi
var MatchPattern = regexp.MustCompile(`^\s*(?:复述|复读|重复|echo)\s*(.*?)\s*$`)

func onReceiveChatMessage(bot string, msgType ubot.MsgType, source string, sender string, message string, info ubot.MsgInfo) (ubot.EventResultType, error) {
	match := MatchPattern.FindStringSubmatch(message)
	if match != nil {
		_ = api.SendChatMessage(bot, msgType, source, sender, match[1])
		return ubot.CompleteEvent, nil
	}
	return ubot.IgnoreEvent, nil
}

func main() {
	err := ubot.HostApp("Echo.Go", func(e *ubot.AppApi) *ubot.App {
		api = e
		return &ubot.App{
			OnReceiveChatMessage: onReceiveChatMessage,
		}
	})
	ubot.AssertNoError(err)
}
