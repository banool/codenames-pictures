# Codenames Pictures 🕵🏼‍♂️🕵🏾‍♀️

[![GoDoc](https://godoc.org/github.com/jbowens/codenames?status.svg)](https://godoc.org/github.com/jbowens/codenames)

**99% of the credit belongs to [jbowens](https://github.com/jbowens) for this wonderful creation!**

Codenames implements a web app for generating and displaying boards for the <a href="https://en.wikipedia.org/wiki/Codenames_(board_game)">Codenames</a> board game. Generated boards are shareable and will update as words are revealed. The board can be viewed either as a spymaster or an ordinary player.

This is a modified version of the original Codenames game where you use pictures instead of words. Look below for instructions on how to use your own images!

A hosted version of the app is available at [codenames.dport.me](https://codenames.dport.me). This is just running on a crusty old laptop that also hosts like 10 other sites so go easy on it.

![Spymaster view of board](https://raw.githubusercontent.com/banool/codenames-pictures/master/screenshot.png)

## How to run this yourself
Firstly, make sure you have go installed. There are good resources for this [on](https://ahmadawais.com/install-go-lang-on-macos-with-homebrew/) [the](https://www.digitalocean.com/community/tutorials/how-to-install-go-on-ubuntu-18-04) [net](https://www.reddit.com/r/golang/comments/79nnq2/go_development_using_wsl_in_win_10/). 

### Docker Instructions
```
# Build
docker build -t codenames-pictures .

# Make sure docker-compose is installed
docker-compose up -d --build
```

### Just installing
These instructions will just grab the binary for you:
```
cd $GOPATH
go get github.com/davidMGS/codenames-pictures/...
go install github.com/davidMGS/codenames-pictures/...
```
You'll still need to set up dependencies following this:
```
cd bin
ln -s ../src/github.com/davidMGS/codenames-pictures/assets_codenames
./codenames
```

Now go follow the instructions for adding images below.


### Developing
If you plan to make changes, you'll want to grab the source and build it yourself:
```
cd $GOPATH
go get github.com/davidMGS/codenames-pictures/...
cd src/github.com/davidMGS/codenames-pictures
# To build the server code.
go build github.com/davidMGS/codenames-pictures/...
# To build the binary.
go build github.com/davidMGS/codenames-pictures/cmd/...
# Run the binary.
./codenames
```

I just use this little one liner for the last three steps:
```
go build github.com/davidMGS/codenames-pictures/... && go build github.com/davidMGS/codenames-pictures/cmd/... && ./codenames 9000; rm codenames
```

You can optionally specify a port (the default is 9001):
```
./codenames 8000
```

Now go follow the instructions for adding images below.

## Loading up images
If you followed the steps above, you should now have a `codenames` binary with an `assets` folder. You can add your own images to `assets/images`. You can also add further sub-directories, it's scanned recursively. They should be square, but beyond that you can really do what you want. It's okay for the image to have transparent backgrounds, both work :) There need to be at least 20 images, but of course the more the better! 🏙🛣🛤🏭🖼🗾🌁🌃🌄🌅🌆🌇🌈🌉🌌🌠🎆🎇🎑!!!

There is support for using remote images! You specify the link for this when creating the game in the lobby. There are 3 types of supported links:

### Text file with absolute links
```
https://mysite.com/links.txt
```
This is a file with absolute website links in it, one per line. For example:
```
https://site.com/image.jpg
http://images.org/cat.png
```
The way I check for it being absolute is whether the link on the first line contains `http`. Janky I know but Go is hard okay.

### Text file with relative links
```
https://mysite.com/links.txt
```
This is a file with links relative to the location of the text file, one per line. For example:
```
image.jpg
cat.png
```
These will resolve to:
```
https://mysite.com/image.jpg
https://mysite.com/cat.png
```

### Link to directory listing
```
https://mysite.com/images/
```
This has the worst support. I try to parse a directory listing (like what is produced by nginx for a directory of files) and extract any link by looking for anchor (`<a>`) tags.
