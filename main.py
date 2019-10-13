from flask import Flask, redirect
import socket

app = Flask(__name__)

@app.route('/')
def hello():
    return f'Hello from {socket.gethostname()}'

@app.route('/node')
def node():
    return socket.gethostname()

@app.route('/twitter')
def twitter():
    return redirect('https://twitter.com/bitcynth')

if __name__ == '__main__':
    app.run(host='0.0.0.0')