# Crypto-Visualizer-API
The backend API Gateway (reverse proxy) for Cypto-Visualizer

## backend

The data is fetched from [cyptocompare api](https://min-api.cryptocompare.com/) and reverse-proxy as API gateway.
You should create your own `.env` in the root folder. example:
```
PROXY_URL=https://min-api.cryptocompare.com/data/
PORT=8000
SUFFIX=api_key=XXXXXXXXXXXXXXXXXXXX
```

