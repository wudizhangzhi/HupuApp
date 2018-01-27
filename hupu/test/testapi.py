#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/1/27 上午4:16
# @Author  : wudizhangzhi


import sys

sys.path.append('..')


def test_news_detail():
    from news import NewsMixin
    n = NewsMixin()
    print(n.getNewsDetailSchema(2255215))



if __name__ == '__main__':
    test_news_detail()