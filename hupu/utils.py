# -*- coding: utf-8 -*-

import json
import os
import random
from hashlib import md5

from colored import stylize
from six import text_type

from api import logger
from messages.messages import LiveMessage, ScoreBoard

log = logger.getLogger(__name__)

HUPU_SALT = 'HUPU_SALT_AKJfoiwer394Jeiow4u309'

TAC_LIST = ['35651900', '35666503', '91054200',
            '35537803', '44831527', '86489400',
            '35084240', '13004008', '35103090',
            '35332802']


def luhn_residue(digits):
    return sum(sum(divmod(int(d) * (1 + i % 2), 10))
               for i, d in enumerate(digits[::-1])) % 10


def get_random_Imei(N=None, filename=None):
    if not N:
        N = 15
    return getImei(N, get_random_tac(filename))


def getImei(N, tac=None):
    '''
    IMEI就是移动设备国际身份码，我们知道正常的手机串码IMEI码是15位数字，
    由TAC（6位，型号核准号码）、FAC（2位，最后装配号）、SNR（6位，厂商自行分配的串号）和SP（1位，校验位）。
    tac数据库: https://www.kaggle.com/sedthh/typeallocationtable/data
    :param N:
    :return:
    '''
    part = ''.join(str(random.randrange(0, 9)) for _ in range(N - 1))
    if tac:
        part = tac + part[len(tac):]
    res = luhn_residue('{}{}'.format(part, 0))
    return '{}{}'.format(part, -res % 10)


def get_random_tac(filename=None):
    if not filename:
        filename = 'tac.csv'
    if not os.path.exists(filename):
        return random.choice(TAC_LIST)
    with open(filename, 'r') as f:
        lines = f.readlines()
        line = random.choice(lines)
        return line.split(',')[0]


def get_android_id():
    # 固定 adb shell settings get secure android_id 随机64位数字的16进制
    result = ''
    for _ in range(64):
        result += random.choice(['0', '1'])
    return hex(int(result, base=2))[2:]


def getSortParam(**kwargs):
    result = ''
    kwargs_sorted = sorted(kwargs)
    for key in kwargs_sorted:
        if len(result) > 0:
            result += '&'
        result += '='.join((key, str(kwargs.get(key))))
    result += HUPU_SALT
    return md5(result.encode('utf8')).hexdigest()


def parser_message(message):
    response = []
    scoreboard = None
    try:
        if isinstance(message, text_type):
            message = json.loads(message)
        args = message['args']
        result = args[0]['result']
        if 'scoreboard' in result:
            _scoreboard = result['scoreboard']
            scoreboard = ScoreBoard(_scoreboard)
        if 'data' in result:
            for i in result['data'][0]['a']:
                lm = LiveMessage(i['content'])
                response.append(lm)
    except Exception as e:
        log.error(e)
    return scoreboard, response


def colored_text(text, *style):
    return stylize(text, *style)
