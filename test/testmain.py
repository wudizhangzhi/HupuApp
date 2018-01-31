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


class TestStartMethods(unittest.TestCase):
    def setUp(self):
        self.hupuapp = HupuApp()

    # def test_upper(self):
    #     self.assertEqual('foo'.upper(), 'FOO')
    #
    # def test_isupper(self):
    #     self.assertTrue('FOO'.isupper())
    #     self.assertFalse('Foo'.isupper())
    #
    # def test_split(self):
    #     s = 'hello world'
    #     self.assertEqual(s.split(), ['hello', 'world'])
    #     # check that s.split fails when the separator is not a string
    #     with self.assertRaises(TypeError):
    #         s.split(2)
    #
    # @unittest.expectedFailure
    # def test_fail(self):
    #     self.assertEqual(1, 0, 'broken')

    def test_getGames(self):
        games = self.hupuapp.getGames()
        with self.subTest(i=0):
            self.assertGreater(len(games), 0)
        for i, game in enumerate(games):
            with self.subTest(i=1+i):
                self.assertIsInstance(game, Game)

    def test_websocket(self):
        pass


if __name__ == '__main__':
    unittest.main()
