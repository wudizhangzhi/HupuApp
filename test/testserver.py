#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/1/28 上午12:43
# @Author  : wudizhangzhi

# encoding: utf-8

'''
测试用websocket服务器
'''
# !/usr/bin/env python

# !/usr/bin/env python

import asyncio
import datetime
import random
import websockets
import json


def read_log(filepath):
    with open(filepath, 'r') as f:
        lines = f.readlines()
    for line in lines:
        if 'receive:' not in line:
            continue
        yield line.split('receive:')[-1].strip()


def run_server():
    reader = read_log('/Users/admin/hupu.log.1')

    async def time(websocket, path):
        while True:
            response = next(reader)
            print('response : {}'.format(response))
            await websocket.send(response)
            await asyncio.sleep(random.random() * 3)

    start_server = websockets.serve(time, '127.0.0.1', 5000)

    asyncio.get_event_loop().run_until_complete(start_server)
    asyncio.get_event_loop().run_forever()


def socket_io():
    from flask import Flask, render_template
    from flask_socketio import SocketIO, emit

    app = Flask(__name__)
    app.config['SECRET_KEY'] = 'secret!'
    socketio = SocketIO(app)

    @app.route('/')
    def index():
        return render_template('index.html')

    @socketio.on('connect', namespace='/test')
    def test_connect():
        emit('my response', {'data': 'Connected', 'count': 0})

    @socketio.on('disconnect', namespace='/test')
    def test_disconnect():
        print('Client disconnected')

    @socketio.on('my event')
    def test_message(message):
        emit('my response', {'data': 'got it!'})

    socketio.run(app, debug=True)


if __name__ == '__main__':
    run_server()
    # socket_io()
