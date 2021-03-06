# clinet-discord

### A Discord bot intended for conversation and fact-based queries

----

## Using the official development version of Clinet (based on this repo)

1. Click on [this link](https://discordapp.com/api/oauth2/authorize?client_id=374546169755598849&permissions=8&scope=bot) to invite `Clinet` into your Discord server.
    1. The `Administrator` permission is not required. The invite link above
    only requests the permission for when new features arrive that require
    individual permissions, saving time from going through and adding missing
    permissions. Should you still wish to manage these permissions yourself,
    do not grant the bot the `Administrator` permission. At a later date, the
    bot will give errors for permissions it does not have when asked to do
    something, and should sending the error message fail it will warn the
    server owner in a DM.
2. (Optional) Join the official [Clinet Discord server](https://discord.gg/qkbKEWT) to get updates on the latest features and changes Clinet has to offer. It even sports bot testing channels to test-drive the bot before deciding if it's right for your community!
3. Enjoy!

## What does it do?

After the `Clinet` bot is invited to your Discord server, it will immediately
begin listening for certain keywords within conversations to trigger certain
events.

`Clinet` will listen for a message that tags it and contains a question mark (?) suffix to
detect when it is being queried with a question. It begins by checking a list of RegEx expressions
stored in the bot configuration to look for configurable responses that bot hosters can set
for their specific instance of the bot. If nothing is found, it then continues on to query
DuckDuckGo's Instant Answers API for a possible response. If DuckDuckGo comes up short, it
finally queries Wolfram|Alpha as a last resort (as Wolfram|Alpha's API services are limited for
non-paying developers). Should no responses be found from any of the three sources, Clinet tells
the user that there was an error finding the response.

Additionally, `Clinet` supports various functionalities not available in question-response queries
using commands. Commands are prefixed by default using `cli$` and sometimes take parameters to
control the output and action of the command.

Finally, `Clinet` has some very useful message management features proven to be successful in the
servers it resides in. If you send a query and make a mistake, for example you misspell a word,
forget to tag the bot, or forget a question mark, you don't have to send a whole new message - just
edit the previous message with the fix and Clinet happily responds to the updated message. Adding
onto this, you can even edit a message that had a successful response - Clinet will happily edit its
response message with the new response to the updated query. Lastly, if you delete your original query
message, Clinet will help with the chat cleanup and delete its response message. These message
management features work for both question queries and commands, and soon will be interwoven into
music playback commands.

## Commands

The default configuration of Clinet uses `cli$` as the command prefix. This can be configured in
the bot configuration for bot hosts, however remains a unique command prefix that should never
interfere with another bot's default command prefix.

All of Clinet's commands respond using a rich embed with all fields inlined to save chat screen
estate on desktop and web versions of Discord while maintaining a clean output everywhere.

For a list of available commands, use the `cli$help` command in a server with Clinet.

----

## Rolling your own locally

In order to run `Clinet` locally, you must have already installed a working Golang
environment on your development system and installed the package dependencies that
Clinet relies on to fully function. Clinet is currently built using Golang `1.10.3`.

### Dependencies

| Package Name |
| ------------ |
| [duckduckgolang](https://github.com/JoshuaDoes/duckduckgolang) |
| [go-soundcloud](https://github.com/JoshuaDoes/go-soundcloud) |
| [go-wolfram](https://github.com/JoshuaDoes/go-wolfram) |
| [discordgo](https://github.com/bwmarrin/discordgo) |
| [github](https://github.com/google/go-github/github) |
| [dca](https://github.com/jonas747/dca) |
| [go-imgur](https://github.com/koffeinsource/go-imgur) |
| [go-klogger](https://github.com/koffeinsource/go-klogger) |
| [go-xkcd](https://github.com/nishanths/go-xkcd) |
| [cron](https://github.com/robfig/cron) |
| [ytdl](https://github.com/rylio/ytdl) |
| [transport](https://google.golang.org/api/googleapi/transport) |
| [youtube](https://google.golang.org/api/youtube/v3) |
| [urbandictionary](https://github.com/JoshuaDoes/urbandictionary) |

### Building

`Clinet` is built using a compiler wrapper known as `govvv`, and opts to use an
altered version to support additional things. govvv acts as a git version injector
for the output compiled binary, taking current statuses of the git repo Clinet is
in and injecting them into uninitialized strings in the main source file to be used
in the command `cli$version`. Simply follow the instructions on the [govvv](https://github.com/JoshuaDoes/govvv) repo page
to learn how to install and use it, then run `govvv build` in the Clinet workspace
directory.

### Acquiring necessary API keys

Clinet's functionality relies on a set of different API keys and access tokens, and without them sports less features to interact with and use. The official bot has all of these already, but if you're looking to roll your own instance of the bot you'll need to acquire these on your own (an exercise left up to you).

| Services | Requirements |
| -------- | ------------ |
| Wolfram\|Alpha | App ID |
| DuckDuckGo | App name (can be anything) |
| YouTube | Search API key |
| Imgur | Client ID |
| SoundCloud | Client ID and app version |

### Writing the configuration

`Clinet` stores its configuration in a file named `config.json` using the JSON data
structure. It has a number of configurable variables and will always globally
override a server's settings if it disables a feature.

The following is an example configuration file:
```JSON
{
	"botToken": "[insert bot token here]",
	"botName": "Clinet",
	"botOwnerID": "[insert bot owner's user ID here]",
	"cmdPrefix": "cli$",
	"botKeys": {
		"wolframAppID": "[insert Wolfram|Alpha app ID here]",
		"ddgAppName": "Clinet",
		"youtubeAPIKey": "[insert YouTube API key here]",
		"imgurClientID": "[insert Imgur client ID here]",
		"soundcloudClientID": "[insert SoundCloud client ID here]",
		"soundcloudAppVersion": "[insert SoundCloud app version here]"
	},
	"botOptions": {
		"sendTypingEvent": true,
		"useDuckDuckGo": true,
		"useGitHub": true,
		"useImgur": true,
		"useSoundCloud": true,
		"useWolframAlpha": true,
		"useXKCD": true,
		"useYouTube": true,
		"wolframDeniedPods": [
			"Locations",
			"Nearby locations",
			"Local map",
			"Inferred local map",
			"Inferred nearest city center",
			"IP address",
			"IP address registrant",
			"Clocks"
		],
		"youtubeMaxResults": 5
	},
	"debugMode": false,
	"customResponses": [
		{
			"expression": "(.*)(?i)raining(.*)tacos(.*)",
			"response": [
				{
					"text": "https://www.youtube.com/watch?v=npjF032TDDQ"
				}
			]
		}
	],
	"customStatuses": [
		{
			"type": 0,
			"status": "`cli$help`!"
		},
		{
			"type": 0,
			"status": "experimentally!"
		},
		{
			"type": 0,
			"status": "in the Go Playground!"
		},
		{
			"type": 0,
			"status": "a nearly finished me!"
		},
		{
			"type": 1,
			"status": "some tunes with users!"
		},
		{
			"type": 2,
			"status": "rewritten code!",
			"url": "https://github.com/JoshuaDoes/clinet-discord"
		}
	]
}
```

Most of the above configuration options should be self-explanatory, but here's some explanations for a few of the less guessable ones:

| Variable | Description |
| -------- | ----------- |
| `botToken` | The token of the bot account Clinet should log into. Can be acquired by [creating an application and then declaring it as a bot user](https://discordapp.com/developers/applications/me/create) and/or [selecting a pre-existing bot user application and acquiring the bot token under the `APP BOT USER` section](https://discordapp.com/developers/applications/me). |
| `botOwnerID` | The user ID of the bot owner. Can be acquired by enabling developer mode on Discord, right clicking your user in a server's user list, and clicking `Copy ID`. If Clinet crashes and recovers from the crash, the error and a full stack trace will be directly messaged to whatever user this option is set to. |
| `botOptions` -> `sendTypingEvent` | Whether or not to send a typing notification in a channel containing a query or command for Clinet to respond to. Helpful for queries or commands that take a little longer than usual to respond to so users know the bot isn't broken. |
| `botOptions` -> `wolframDeniedPods` | An array of pod titles to skip over when creating a list of responses to use in a rich embed response from a Wolfram\|Alpha query. The default list is highly recommended for bot hosters concerned with the privacy of the bot's host location. |
| `botOptions` -> `youtubeMaxResults` | The total amount of results to display per page for YouTube searches via the `cli$youtube search` command. Maximum of 253. |
| `debugMode` | Debug mode enables various console debugging features, such as chat output and other detailed information about what Clinet is up to. |
| `customResponses` | Stored as objects in an array, custom responses are exactly what the name depicts. Each object contains an `expression` variable, which stores a valid regular expression, and a `response` array, which itself contains objects randomly selected by the main program for different `text` responses each time the custom response is queried. |
| `customStatuses` | Stored as objects in an array, custom statuses are used to set the bot's presence. Each object contains a `type` variable, which stores integers 0, 1, and 2, which are "Playing", "Listening to", and "Streaming" respectively, and a `status` variable, which stores the status text to use. If the type is set to 2, you can also set a `url` variable to use as the stream URL. |

The configuration file by default will never be included in git commits, as declared by `.gitignore`. This is to prevent accidental leakage of API keys and bot tokens.

### Running `Clinet`

Finally, to run Clinet, simply type `./clinet-discord` in your terminal/shell or `.\clinet-discord.exe` in your command prompt. If everything goes well, you can find your bot user application and generate an OAuth2 URL to invite the bot into various servers in which you have the `Administrator` permission of.

### Panic recovery

If Clinet ever crashes from a panic, custom-made panic recovery will save the crash message to `crash.txt` and the stack trace to `stacktrace.txt` in the bot's working directory. When Clinet is next started up, it will send the crash message and the file of the stack trace to the user specified in the configuration option `botOwnerID` and proceed to delete the two files.

Running Clinet by itself will spawn a "master" process with a few small jobs: Spawning a "bot" process, restarting the "bot" process if it exits for any reason, and closing the "bot" process if the "master" process ever exits for any reason. This is to ensure that, even if the "bot" process crashes, Clinet can continue running and instantly report the crash to the user specified in the configuration option `botOwnerID`.

### States

If you close Clinet after running it long enough for it to merely exist on Discord, you'll notice a new folder called `state`. This folder contains "states" of various structs within Clinet's memory, stored in JSON format. Upon reopening Clinet, these state files are then loaded into memory so Clinet can (for the most part) return to its original "state" before it was closed. States were added as helpers to panic recovery so users can continue with what they were doing.

----

## Support
For help and support with Clinet, visit the [Clinet Discord server](https://discord.gg/qkbKEWT) and ask for an online developer.

## License
The source code for Clinet is released under the MIT License. See LICENSE for more details.

## Donations
All donations are highly appreciated. They help me pay for the server costs to keep Clinet running and even help point my attention span to Clinet to fix issues and implement newer and better features!

[![Donate](https://img.shields.io/badge/Donate-PayPal-green.svg)](https://paypal.me/JoshuaDoes)
