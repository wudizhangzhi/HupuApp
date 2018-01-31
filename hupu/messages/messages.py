# !/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/1/27 下午2:22
# @Author  : wudizhangzhi
# from __future__ import unicode_literals
from six import python_2_unicode_compatible

from hupu.messages.entries import StringEntry, IntEntry, ListEntry, to_text, PY3
from hupu.api import logger

log = logger.getLogger(__name__)


class BaseMessage(object):
    __type__ = 'base'

    def __init__(self, message):
        self.__dict__.update(message)

    def __repr__(self):
        return '({})'.format(self.__str__())


class LiveMessage(BaseMessage):
    """
      {
        "rowId": 0,
        "content": {
          "uid": "18492968",
          "event": "回传德章泰-默里！分球底角贝尔坦斯！！",
          "end_time": "夏天",
          "team": 0,
          "t": 1516756188
        }
      }
    """
    __type__ = 'live'
    event = StringEntry('content.event')
    uid = StringEntry('content.uid')
    end_time = StringEntry('content.end_time')
    team = IntEntry('content.team')
    t = IntEntry('content.t')

    def __str__(self):
        return '{end_time}:  {event}'.format(end_time=self.end_time, event=self.event)


# TODO
class CasinoMessage(BaseMessage):
    __type__ = 'casino'
    event = StringEntry('event')
    uid = StringEntry('uid')
    end_time = StringEntry('end_time')
    team = IntEntry('team')
    t = IntEntry('t')

    desc = StringEntry('desc')
    status = StringEntry('status.desc')
    answers = ListEntry('answer.title')


class ScoreBoard(BaseMessage):
    """
    "scoreboard": {
          "home_tid": "21",
          "home_score": "4",
          "away_tid": "10",
          "away_score": "6",
          "process": "第一节 9:35",
          "status": "2"
        },
    """
    __type__ = 'score'
    home_tid = IntEntry('home_tid')
    home_score = IntEntry('home_score')
    away_tid = IntEntry('away_tid')
    away_score = IntEntry('away_score')
    process = StringEntry('process')  # 时间
    status = IntEntry('status')

    def __str__(self):
        return '{away_score}:{home_score} {process}'.format(away_score=self.away_score,
                                                            home_score=self.home_score,
                                                            process=self.process)


class Game(BaseMessage):
    __type__ = 'game'
    home_name = StringEntry('home_name')
    away_name = StringEntry('away_name')
    home_score = IntEntry('home_score')
    away_score = IntEntry('away_score')
    process = StringEntry('process')
    gid = IntEntry('gid')

    def __str__(self):
        return "{away_name}  {away_score} vs {home_score}  {home_name}  {process}".format(
            away_name=self.away_name,
            home_name=self.home_name,
            away_score=self.away_score,
            home_score=self.home_score,
            process=self.process,
        )

    def __repr__(self):
        return "Game({} {})".format(self.__str__(), self.gid)


class SocketMessage(BaseMessage):
    __type__ = 'socket'
    room = StringEntry('room')
    gid = StringEntry('gid')  # 比赛的id
    status = StringEntry('status')
    pid = StringEntry('pid')  # TODO 好像是交流的id

    room_live_type = IntEntry('room_live_type')

    livemessges = []

    games = []

    # casinomessages = ListEntry('casinomessages')

    def __init__(self, message):
        self.room = message.pop('room', None)
        result = message.pop('result', {})
        if self.room == 'NBA_HOME':
            if isinstance(result, list):
                for r in result:
                    self.games.append(Game(r))
        else:
            # 计分板
            scoreboard = result.pop('scoreboard', {})
            self.scoreboard = ScoreBoard(scoreboard)
            # 直播消息
            data = result.pop('data', {})
            if data:
                _livemessge = data[0]['a']
                for _lm in _livemessge:
                    self.livemessges.append(LiveMessage(_lm))

        # TODO 竞猜消息

        super(SocketMessage, self).__init__(message)


# @python_2_unicode_compatible
class News(BaseMessage):
    """
    虎扑新闻
    """
    __type__ = 'news'
    news_type_list = [1, 5]

    nid = IntEntry('nid')
    title = StringEntry('title')
    summary = StringEntry('summary')
    uptime = IntEntry('uptime')
    type = IntEntry('type')  # 1:比赛新闻? 3:帖子?  5:主题帖?
    lights = IntEntry('lights')
    replies = IntEntry('replies')
    read = IntEntry('read')

    def __str__(self):
        return '{}'.format(self.title)

    @property
    def statistics_text(self):
        return '阅读: {read} 点亮: {lights} 回复: {replies}'.format(
            read=self.read,
            lights=self.lights,
            replies=self.replies,
        )


# @python_2_unicode_compatible
class NewsDetail(BaseMessage):
    """
    虎扑新闻正文
    """
    __type__ = 'newsdetail'

    url = StringEntry('url')
    title = StringEntry('title')

    content = StringEntry('offline_data.data.news.content')

    def __str__(self):
        return self.content


class TeamRank(BaseMessage):
    """
    排行
    """
    __type__ = 'teamrank'
    rank_type = StringEntry('rank_type')
    name = StringEntry('name')
    title = StringEntry('title')
    field = StringEntry('field')

    data = ListEntry('data')

    def __str__(self):
        return '{title} {name}'.format(
            title=self.title,
            name=self.name
        )

    @property
    def to_table(self):
        """
        排行  胜-负  胜率/胜场差  近况
        """
        table = []
        table.append(self.table_title)
        if self.rank_type in ['east', 'west']:
            for i, _data in enumerate(self.data):
                table.append('{rank}.{name:<10}{win:>3}-{lost:<8}{gb:>10}{strik:>6}'.format(
                    rank=i + 1,
                    name=to_text(_data.get('name')),
                    win=to_text(_data.get('win')),
                    lost=to_text(_data.get('lost')),
                    gb=to_text(_data.get('gb')),
                    strik=to_text(_data.get('strk')),
                ))
        else:
            for _data in self.data:
                table.append('{rank}.{team_name:<20}{value}'.format(
                    rank=to_text(_data.get('rank')),
                    team_name=to_text(_data.get('team_name')),
                    value=to_text(_data.get(self.field)),
                ))
        return table

    @property
    def table_title(self):
        if self.rank_type in ['east', 'west']:
            return '排行         胜-负      胜率/胜场差     近况'
        else:
            return '{name:<20}数据'.format(name=self.name)


class PlayData(BaseMessage):
    __type__ = 'playerdata'

    name = StringEntry('name')

    data = ListEntry('data')

    def __str__(self):
        return '{}'.format(self.name)

    @property
    def to_table(self):
        table = []
        for _data in self.data:
            if 'injury_detail_cn' in _data:
                table.append('{team_name}-{player_name} {injury_part_cn} {injury_detail_cn}'.format(
                    injury_part_cn=to_text(_data.get('injury_part_cn', '')),
                    player_name=to_text(_data.get('player_name')),
                    team_name=to_text(_data.get('team_name')),
                    injury_detail_cn=to_text(_data.get('injury_detail_cn')),
                ))
            else:
                table.append('{rank}.{player_name}({team_name}){val:>20}'.format(
                    rank=to_text(_data.get('rank', '')),
                    player_name=to_text(_data.get('player_name')),
                    team_name=to_text(_data.get('team_name')),
                    val=to_text(_data.get('val')),
                ))
        return table

    @property
    def title(self):
        return '球员数据--{}'.format(self.name)


def test_message():
    sb = ScoreBoard({'home_tid': 123})
    print(sb.home_tid)
    print(type(sb))


if __name__ == '__main__':
    test_message()
