# coding=utf8


import sys

sys.path.append('..')

from hupu import start
from hupu.terminalsize import get_terminal_size

if __name__ == '__main__':
    # sizex, sizey = get_terminal_size()
    # print('width =', sizex, 'height =', sizey)
    start()
