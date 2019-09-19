# Crypto-Visualizer-API
The backend API Gateway (reverse proxy) for Cypto-Visualizer

## usage:

All AJAX requests toward https://crypto-visualizer-api.herokuapp.com/ will forward to https://min-api.cryptocompare.com/data/ with `api_key` set.

example:
request to `https://crypto-visualizer-api.herokuapp.com/price?fsym=BTC&tsyms=USD,JPY,EUR`
will equal to `https://min-api.cryptocompare.com/data/price?fsym=BTC&tsyms=USD,JPY,EUR?api_key=XXXXXXXXXXXXXXXXX`

## backend

The data is fetched from [cyptocompare api](https://min-api.cryptocompare.com/) and reverse-proxy as API gateway.
You should create your own `.env` in the root folder. example:
```
PROXY_URL=https://min-api.cryptocompare.com/data/
PORT=8000
SUFFIX=api_key=XXXXXXXXXXXXXXXXXXXX
```
after that, run `source .env` to load the variable into environment

## deployment

This api service uses [Heroku](https://www.heroku.com) for deployment.
some tools and command used:
* [heroku config](https://github.com/xavdid/heroku-config): quickly push/pull environment variables to/from heroku
* [Godep](https://github.com/tools/godep): automatic generate dependencies for go project
* heroku command:
    * Login: `heroku login`
    * Create a new heroku project: `heroku create [yourAppName]`
    * Deploy master to heroku: `git push heroku master` 
* a `Procfile` (with no extension) is required for heroku to know how to start the app. Here, `web: [yourRepositoryName]` will tell heroku that is app will be an http app.

