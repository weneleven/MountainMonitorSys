import json
import time
import websocket

try:
    import thread
except ImportError:
    import _thread as thread

def on_message(ws, message):
    print(message)

def on_error(ws, error):
    print(error)

def on_close(ws, close_status_code, close_msg):
    print("### closed ###")

def on_open(ws):
    def run(*args):
        with open('new_data.json', 'r', encoding='utf-8') as file:
            jsonData = json.load(file)

#改数据传送的频率等等,改这个函数就行
        for data in jsonData:
            time.sleep(1)
            ws.send(json.dumps(data))
        ws.close()
    thread.start_new_thread(run, ())

if __name__ == "__main__":
    websocket.enableTrace(True)
    ws = websocket.WebSocketApp("ws://localhost:3000/ws/sensor",
                                on_open=on_open,
                                on_message=on_message,
                                on_error=on_error,
                                on_close=on_close)
    ws.run_forever()
