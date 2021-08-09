# ytl-bounce
A small application that redirects a YouTube link with a channel ID to the popout chat for a live stream on that channel.
This lets you avoid pasting a new YouTube Live link into OBS during each stream.

## Setup
Click the deploy button below to deploy this application on your own Netlify site.

[![Deploy to Netlify](https://www.netlify.com/img/deploy/button.svg)](https://app.netlify.com/start/deploy?repository=https://github.com/karashiiro/ytl-bounce)

You will be prompted to enter a repo name for Netlify to create a copy of this repo on your GitHub account.

After a minute, the website will be ready for use.

## Usage
Set your browser widget URL to `https://yoururl.netlify.app/.netlify/functions/bounce?c=ChannelID`, replacing `yoururl` with the start of
your own Netlify site's URL and `ChannelID` with your own YouTube channel ID.
