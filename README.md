# ytl-bounce
A small application that redirects a YouTube link with a channel ID to the popout chat for a live stream on that channel.
This lets you avoid pasting a new YouTube Live link into OBS during each stream.

## Setup
[![Deploy to Netlify](https://www.netlify.com/img/deploy/button.svg)](https://app.netlify.com/start/deploy?repository=https://github.com/karashiiro/ytl-bounce)

## Usage
Set your browser widget URL to `https://yoururl.netlify.app/.netlify/functions/bounce?c=ChannelID`, replacing `ChannelID` with your
own channel's ID.
