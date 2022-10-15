## ✋ No longer maintained 🪦 
Alternative can be found at https://github.com/cvzi/play

---
---

# playshields [![Go Report Card](https://goreportcard.com/badge/github.com/cvzi/playshields?style=for-the-badge)](https://goreportcard.com/report/github.com/cvzi/playshields)

A Go app that offers JSON data from Play Store for a [shields.io endpoint](https://shields.io/endpoint), which can easily be deployed to Heroku.

## Running Locally

Make sure you have [Go](http://golang.org/doc/install) version 1.13 or newer and the [Heroku Toolbelt](https://toolbelt.heroku.com/) installed.

```sh
$ git clone https://github.com/cvzi/playshields.git
$ cd playshields
$ go build -o bin/playshields -v . # or `go build -o bin/playshields.exe -v .` in git bash
...
$ heroku local
```

Your app should now be running on [localhost:5000](http://localhost:5000/).

On windows run with:

```sh
$ run.bat
```

## Deploying to Heroku

```sh
$ heroku create
$ heroku config:set GOVERSION=go1.18.2
$ heroku config:set GIN_MODE=release
$ git push heroku master
$ heroku open
```

or

[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)
