# Binance Bot

A terminal (bot) that parses commands and executes buy/sell order(s) on Binance Future Trading platform. The bot is intended to be fast and reliable for placing trade order quickly. Ideal to host the bot on a VPS located in Tokyo, Japan for faster trading. Or use from a location where ping-rate is lower to api.binance.com

### Valid Command Examples:

1) buy 300 eth        : Buy 300 USDT Worth Of ETH
2) sell 500 eth       : Sell 500 USDT Worth Of ETH At Market Price
3) exit eth           : Exit  Currently Open ETH position
4) cancel eth         : Cancel All Pending Orders For ETH
5) buy 500 btc 19000  : Buy 500 USDT Worth Of BTC At Limit Price of 19000
6) sell 200 bnb 350   : Sell 500 USDT Worth Of BNB At Limit Price of 350

---

Put api key and secret in `settings.json`

### Important
_*_ Command Can Be sent In Uppercase or Lowercase

_*_ Command Must Match Its Format To Process It Properly
