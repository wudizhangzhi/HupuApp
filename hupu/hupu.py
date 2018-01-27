"""Hupu.
    Proudly presented by Hupu JRs.

Usage:
    hupu [-m MODE] [-u USERNAME] [-p PASSWORD] [-a APIVERSION]
    hupu -h | --help
    hupu -v | --version

Tips:
    Please hit Ctrl-C on the keyborad when you want to interrupt the game live.

Options:
    -u --username    input username.
    -p --password    input password.
    -a --apiversion  api version.[default: 7.1.15]
    -m --mode        run mode.[default: live, available: live news...]
    -h --help        Show this help message and exit.
    -v --version     Show version.
"""
import docopt
import requests

from api.live import LiveMinxin
from api.login import LoginMixin
from api.news import NewsMixin

try:
    # 不打印ssl警告
    requests.packages.urllib3.disable_warnings()
except ImportError:
    pass

from screen import Screen
from api import logger

log = logger.getLogger(__name__)

MODE_LIST = ['live', 'news']


class Hupu(LiveMinxin, NewsMixin, LoginMixin):
    def run(self):
        # 判断参数, 执行哪一种场景
        # 1.没参数
        # 2.有用户名密码 -- 登录

        # 默认进入比赛文字直播模式
        mode = self._kwargs.get('MODE', '') or 'live'
        mode = mode.lower()
        assert mode in MODE_LIST, AttributeError('Expect mode is {}, got {}.'.format(', '.join(MODE_LIST), mode))

        screen = Screen(self, client_id=self.client)  # 显示的屏幕

        if mode == 'live':  # 文字直播模式
            games = self.getGames()
            screen.set_screen(games)

        elif mode == 'news':  # 新闻模式
            news = self.getNews()
            screen.set_screen(news)

        # 设置模式， 开始监听
        screen.set_mode(mode)
        screen.listen()


def main():
    arguments = docopt.docopt(__doc__, version='Hupu 1.0')
    hupulive = Hupu(**arguments)
    hupulive.run()


if __name__ == '__main__':
    # hupu = HupuLive()
    # print(hupu.getNews().json())
    main()
