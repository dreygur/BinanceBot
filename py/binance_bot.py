import ccxt
import config
import time
import traceback

#Binance API
binance = ccxt.binance({
  'apiKey': config.api_key,
  'secret': config.api_secret,
  'enableRateLimit': True,  # enable built-in rate limiter
  'options': {
    'defaultType': 'future',
  },
})
binance.set_sandbox_mode(True)


#Market Enter Position
def market_enter_position(currency_pair,side,size):
  print("Size:", size)
  return binance.create_order(currency_pair, type='market', side=side, amount= size)


#Market Enter Position
def market_exit_position(currency_pair,side,size):
  params = {"reduceOnly": True }
  return binance.create_order(currency_pair, type='market', side=side, amount=size,params=params)

#Limit Enter Position
def limit_enter_position(currency_pair,side,size,entry_price):
  print("Size:", size)
  print(currency_pair,side,size,entry_price)
  return binance.create_order(currency_pair, type='limit', side=side, amount=size,price=entry_price)

#Fetch position
def fetch_open_position_data(currency_pair):
  position_data = binance.fetch_positions(symbols=[currency_pair])[0]['info']
  if position_data['positionSide'] == "BOTH":
    reverse_trade_side = "SELL"
    if float(position_data['positionAmt'])<0:
      reverse_trade_side = "BUY"
    #Required Data
    required_data = [position_data['symbol'],reverse_trade_side,abs(float(position_data['positionAmt']))]
  return required_data

#Cancel Active Position
def exit_position(currency_pair):
  try:
    #Fetch Open Position Data
    position_data = fetch_open_position_data(currency_pair)
    #Exit the position
    exit_position = market_exit_position(currency_pair,position_data[1],position_data[2])
  except Exception as e:
    print("Error: "+str(e))


#Calculate Market Order Lot Size
def get_market_order_lot_size(currency_pair,usdt_size):
  #Last Price
  last_price = binance.fetch_ticker(currency_pair)['last']
  lot_size = (usdt_size / last_price)
  return lot_size

#Calculate Limit Order Lot Size
def get_limit_order_lot_size(usdt_size,limit_price):
  lot_size = (usdt_size / limit_price)
  print(lot_size)
  return lot_size


help_string = '''
Valid Command Examples:
...............................
1> buy  300 eth   :-   Buy 300 USDT Worth Of ETH.

2> sell  500 xrp  :-   Sell 500 USDT Worth Of XRP At Market Price.

3> exit doge :-   Exit  Currently Open DOGE position.

4> cancel ada :- Cancel All Pending Orders For ADA

5> buy 500 btc 19000 :-   Buy 500 USDT Worth Of BTC At Limit Price of 19000.

6> sell 200 bnb 350 :-  Sell 500 USDT Worth Of BNB At Limit Price of 350.

...............................

-Command Can Be sent In Uppercase or Lowercase-
-Command Must Match Its Format To Process It Properly-
'''

#Process Command
def process_command(raw_string):
  try:
    #Command Start Time
    start_time = time.time()
    #Process String
    raw_string = raw_string.lower().split(" ")
    data_list = [items.strip() for items in raw_string]
    print(len(data_list))
    #If Market Order
    if len(data_list)==3:
      currency_pair = data_list[2].upper()+"/USDT"
      trade_side = data_list[0].upper()
      usdt_size = float(data_list[1])
      print(currency_pair, trade_side, usdt_size)
      try:
        enter_trade = market_enter_position(currency_pair,trade_side,get_market_order_lot_size(currency_pair,usdt_size))
        #Response Time
        response_time = time.time()
        time_taken = int((response_time-start_time)*1000)
        #Get Data
        msg_string = "\n*** Market Order Filled ***\nSymbol: #sym\nSide: #side\nFill Price: #price\nSize: #size\nUSD Value: #usd\nTime Taken: #time\n**********\n"
        msg_string = msg_string.replace("#sym",enter_trade['info']['symbol']).replace("#side",enter_trade['info']['side']).replace("#price",enter_trade['info']['avgPrice'])
        msg_string = msg_string.replace("#size",enter_trade['info']['origQty']).replace("#usd",enter_trade['info']['cumQuote']).replace("#time",str(time_taken))
        print(msg_string)
      except Exception as e:
        print("Error: "+str(e))

    #If Limit Order
    if len(data_list)== 4:
      print(len(data_list))

      currency_pair = data_list[2].upper()+"USDT"
      trade_side = data_list[0].upper()
      usdt_size = float(data_list[1])
      entry_price = float(data_list[3])
      print(currency_pair, trade_side, usdt_size, entry_price)
      try:
        limit_enter_position(currency_pair,trade_side,get_limit_order_lot_size(usdt_size,entry_price),entry_price)
        #Response Time
        response_time = time.time()
        time_taken = int((response_time-start_time)*1000)
        print("\nLimit Order Executed Successfully\nTime Taken: "+str(time_taken))
      except Exception as e:
        print("Error: "+str(e))

    #If Exit Order
    if len(data_list)== 2:
      # Exit Positions
      if data_list[0]== "exit":
        currency_pair = data_list[1].upper()+"USDT"
        exit_position(currency_pair)
        #Response Time
        response_time = time.time()
        time_taken = int((response_time-start_time)*1000)
        msg_string = "\nPosition Closed For: "+currency_pair+"\nTime Taken: "+str(time_taken)+"\n"
        print(msg_string)

      #Cancel Orders
      if data_list[0]== "cancel":
        currency_pair = data_list[1].upper()+"USDT"
        binance.cancel_all_orders(symbol=currency_pair)
        #Response Time
        response_time = time.time()
        time_taken = int((response_time-start_time)*1000)
        msg_string= "\nOrders Canceled For "+currency_pair+"\nTime Taken: "+str(time_taken)+"\n"
        print(msg_string)


  except:
      print(traceback.format_exc())
      print("\n\n__________Invalid Command__________")
      print(help_string)



#Main
def main():
  print("__________WELCOME__________\n\n")
  while True:
    raw_string  = input('Enter Command: ')
    #Process Command
    process_command(raw_string)


main()

