#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/2/6 上午11:18
# @Author  : wudizhangzhi
# @File    : Word2Vec.py

import sys
import re
import redis
import jieba
import jieba.posseg as pseg
from zhon.hanzi import punctuation as zhon_punctuation
from PIL import Image
import numpy as np
import matplotlib.pyplot as plt
from wordcloud import WordCloud, STOPWORDS

sys.path.append('../..')
from string import punctuation, digits, ascii_letters
from hupubbs.base.models import *

pool = redis.ConnectionPool(host='localhost', port=6379, db=0)
redis_client = redis.Redis(connection_pool=pool)

char_not_append = punctuation + digits + zhon_punctuation + ascii_letters

exclude_word = ['mhupucom', 'iPhone', '客户端', 'Android', '手机', '发自', '此帖', 'MR', '的']
flag_list = ['n', 'nr', 'v', 'a', 'vn', 'ns']


def pure_text(text):
    # return re.sub(r"[%s]" % char_not_append, '', text)
    return re.sub(r"[^\u4e00-\u9fa5]+", '', text)


def count_words():
    jieba.load_userdict('dict.txt')

    # 创建DBSession类型:
    DBSession = sessionmaker(bind=engine)
    # 创建session对象:
    session = DBSession()
    query = session.query(BBSPostComment)
    # count = 0
    # result = query.filter(BBSPostComment.id < 20).order_by(BBSPostComment.createtime.desc()).all()
    result = query.all()

    for c in result:
        content = pure_text(c.content)
        # content_list = [j for j in jieba.cut(content, cut_all=False) if j.strip()]
        content_list = [(j, flag) for j, flag in pseg.cut(content) if j.strip()]
        for _char, flag in content_list:
            if len(_char) == 1 or _char in ascii_letters or _char in exclude_word or flag not in flag_list:
                continue
            frequent = redis_client.zincrby('frequent', _char)
            print("{}: {}".format(_char, int(frequent)))

    session.close()


def word2vec():
    pass


def showwordcloud_from_frequent(text_frequent):
    alice_mask = np.array(Image.open("/Users/admin/alice_mask.png"))

    stopwords = set(STOPWORDS)
    stopwords.add("said")

    wc = WordCloud(background_color="white",
                   max_words=2000,
                   mask=alice_mask,
                   stopwords=stopwords,
                   font_path="/Users/admin/Arial.ttf", )
    # generate word cloud
    wc.generate_from_frequencies(text_frequent)

    # store to file
    wc.to_file("alice.png")

    # show
    plt.imshow(wc, interpolation='bilinear')
    plt.axis("off")
    plt.figure()


def display_hupu_word_cloud():
    generator = redis_client.zscan_iter('frequent', score_cast_func=int)
    vocabulary = {}
    while True:
        try:
            _char, score = next(generator)
            _char = _char.decode('utf8')
            # vocabulary.append({_char: score})
            vocabulary[_char] = score
            print('{} :{}'.format(_char, score))
        except StopIteration as e:
            break

    showwordcloud_from_frequent(vocabulary)


def ludingji(filepath):
    f = open(filepath, 'r', encoding='gbk')
    for line in f.readlines():
        # 分词,每一个都要
        word_list = jieba.cut(pure_text(line.strip()), cut_all=False)
        # TODO 分隔符?
        # 存储
        pass


if __name__ == '__main__':
    # count_words()
    display_hupu_word_cloud()
