from .entries import StringEntry, IntEntry, ListEntry

from api import logger

log = logger.getLogger(__name__)


class BaseMessage(object):
    __type__ = 'base'

    def __init__(self, message):
        self.__dict__.update(message)


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
    away_name = StringEntry('home_name')
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


class NewsDetail(BaseMessage):
    """
    虎扑新闻正文
    """
    __type__ = 'newsdetail'

    url = StringEntry('url')
    title = StringEntry('title')

    content = StringEntry('offline_data.data.news.content')


def test_message():
    sb = ScoreBoard({'home_tid': 123})
    print(sb.home_tid)
    print(type(sb))


if __name__ == '__main__':
    test_message()
