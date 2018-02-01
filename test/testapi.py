#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/1/31 下午9:34
# @Author  : wudizhangzhi

import sys

sys.path.append('..')
import unittest
from unittest.mock import MagicMock, Mock

from hupu.hupuapp import HupuApp
from hupu.messages.messages import *
from hupu.hupulivewebsocket import HupuSocket

class TestAPIMethods(unittest.TestCase):
    def setUp(self):
        self.hupuapp = HupuApp()

    # def test_1_upper(self):
    #     self.assertEqual('foo'.upper(), 'FOO')
    #
    # def test_2_isupper(self):
    #     self.assertTrue('FOO'.isupper())
    #     self.assertFalse('Foo'.isupper())
    #
    # def test_3_split(self):
    #     s = 'hello world'
    #     self.assertEqual(s.split(), ['hello', 'world'])
    #     # check that s.split fails when the separator is not a string
    #     with self.assertRaises(TypeError):
    #         s.split(2)

    # @unittest.expectedFailure
    # def test_4_fail(self):
    #     self.assertEqual(1, 0, 'broken')
    def test_getGames(self):
        games = self.hupuapp.getGames()
        with self.subTest(i=0):
            self.assertGreater(len(games), 0)
        for i, game in enumerate(games):
            with self.subTest(i=1 + i):
                self.assertIsInstance(game, Game)

    def test_getIpAdress(self):
        host, port = self.hupuapp.getIpAdress()
        self.assertNotEqual((host, port), (None, None))

    def test_teamdata(self):
        teamranks = self.hupuapp.getDatas()
        self.assertGreater(len(teamranks), 0)

    def test_getPlayerDataInGenernal(self):
        datas_regular = self.hupuapp.getPlayerDataInGenernal()
        with self.subTest(i=0):
            self.assertGreater(len(datas_regular), 0)
        datas_injury = self.hupuapp.getPlayerDataInGenernal(datatype='injury')
        with self.subTest(i=1):
            self.assertGreater(len(datas_injury), 0)
        datas_daily = self.hupuapp.getPlayerDataInGenernal(datatype='daily')
        with self.subTest(i=2):
            self.assertGreater(len(datas_daily), 0)

    def test_getNews(self):
        news = self.hupuapp.getNews()
        self.assertGreater(len(news), 0)

    def test_gettoken(self):
        host, port = self.hupuapp.getIpAdress()
        hs = HupuSocket(game=None, client=self.hupuapp.client, host=host, port=port)
        token = hs.get_token()
        self.assertEqual(len(token), 20)


if __name__ == '__main__':
    unittest.main()
