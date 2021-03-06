package main

import (
	"errors"
	"regexp"

	"github.com/JoshuaDoes/duckduckgolang"
	"github.com/JoshuaDoes/go-soundcloud"
	"github.com/JoshuaDoes/go-wolfram"
	"github.com/bwmarrin/discordgo"
	"github.com/google/go-github/github"
	"github.com/koffeinsource/go-imgur"
	"github.com/nishanths/go-xkcd"
	"google.golang.org/api/youtube/v3"
)

//Bot data structs
type BotClients struct {
	DuckDuckGo *duckduckgo.Client
	GitHub     *github.Client
	Imgur      imgur.Client
	SoundCloud *soundcloud.Client
	Wolfram    *wolfram.Client
	XKCD       *xkcd.Client
	YouTube    *youtube.Service
}
type BotData struct {
	BotClients      BotClients
	BotKeys         BotKeys               `json:"botKeys"`
	BotName         string                `json:"botName"`
	BotOwnerID      string                `json:"botOwnerID"`
	BotOptions      BotOptions            `json:"botOptions"`
	BotToken        string                `json:"botToken"`
	CommandPrefix   string                `json:"cmdPrefix"`
	CustomResponses []CustomResponseQuery `json:"customResponses"`
	CustomStatuses  []CustomStatus        `json:"customStatuses"`
	DebugMode       bool                  `json:"debugMode"`
	//RecoverCrashes       bool                  `json:"recoverCrashes"`
	//SendOwnerStackTraces bool                  `json:"sendOwnerStackTraces"`

	DiscordSession *discordgo.Session
	Commands       map[string]*Command
}
type BotKeys struct {
	DuckDuckGoAppName    string `json:"ddgAppName"`
	ImgurClientID        string `json:"imgurClientID"`
	SoundCloudAppVersion string `json:"soundcloudAppVersion"`
	SoundCloudClientID   string `json:"soundcloudClientID"`
	WolframAppID         string `json:"wolframAppID"`
	YouTubeAPIKey        string `json:"youtubeAPIKey"`
}
type BotOptions struct {
	SendTypingEvent   bool     `json:"sendTypingEvent"`
	UseDuckDuckGo     bool     `json:"useDuckDuckGo"`
	UseGitHub         bool     `json:"useGitHub"`
	UseImgur          bool     `json:"useImgur"`
	UseSoundCloud     bool     `json:"useSoundCloud"`
	UseWolframAlpha   bool     `json:"useWolframAlpha"`
	UseXKCD           bool     `json:"useXKCD"`
	UseYouTube        bool     `json:"useYouTube"`
	WolframDeniedPods []string `json:"wolframDeniedPods"`
	YouTubeMaxResults int      `json:"youtubeMaxResults"`
}
type CustomResponseQuery struct {
	Expression string `json:"expression"`
	Regexp     *regexp.Regexp
	Responses  []CustomResponseReply `json:"response"`
}
type CustomResponseReply struct {
	Response string `json:"text"`
	ImageURL string `json:"imageURL"`
}
type CustomStatus struct {
	Type   int    `json:"type"`
	Status string `json:"status"`
	URL    string `json:"url,omitempty"`
}

func (configData *BotData) PrepConfig() error {
	//Bot config checks
	if configData.BotName == "" {
		return errors.New("config:{botName: \"\"}")
	}
	if configData.BotToken == "" {
		return errors.New("config:{botName: \"\"}")
	}
	if configData.CommandPrefix == "" {
		return errors.New("config:{cmdPrefix: \"\"}")
	}

	//Bot key checks
	if configData.BotOptions.UseDuckDuckGo && configData.BotKeys.DuckDuckGoAppName == "" {
		return errors.New("config:{botOptions:{useDuckDuckGo: true}} not permitted, config:{botKeys:{ddgAppName: \"\"}}")
	}
	if configData.BotOptions.UseImgur && configData.BotKeys.ImgurClientID == "" {
		return errors.New("config:{botOptions:{useImgur: true}} not permitted, config:{botKeys:{imgurClientID: \"\"}}")
	}
	if configData.BotOptions.UseSoundCloud && configData.BotKeys.SoundCloudAppVersion == "" {
		return errors.New("config:{botOptions:{useSoundCloud: true}} not permitted, config:{botKeys:{soundcloudAppVersion: \"\"}}")
	}
	if configData.BotOptions.UseSoundCloud && configData.BotKeys.SoundCloudClientID == "" {
		return errors.New("config:{botOptions:{useSoundCloud: true}} not permitted, config:{botKeys:{soundcloudClientID: \"\"}}")
	}
	if configData.BotOptions.UseWolframAlpha && configData.BotKeys.WolframAppID == "" {
		return errors.New("config:{botOptions:{useWolframAlpha: true}} not permitted, config:{botKeys:{wolframAppID: \"\"}}")
	}
	if configData.BotOptions.UseYouTube && configData.BotKeys.YouTubeAPIKey == "" {
		return errors.New("config:{botOptions:{useYouTube: true}} not permitted, config:{botKeys:{youtubeAPIKey: \"\"}}")
	}

	//Custom response checks
	for i, customResponse := range configData.CustomResponses {
		regexp, err := regexp.Compile(customResponse.Expression)
		if err != nil {
			return err
		} else {
			configData.CustomResponses[i].Regexp = regexp
		}
	}
	return nil
}
