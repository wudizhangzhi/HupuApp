#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/1/31 下午9:23
# @Author  : wudizhangzhi


def test():
    import sys
    print(sys.getdefaultencoding())
    print('我们')
    print(u'我们')
    print(u'我们'.encode('utf8'))


if __name__ == '__main__':
    test()
