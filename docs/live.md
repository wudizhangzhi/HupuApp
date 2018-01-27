# 虎扑文字直播


## websocket 流程
```python
token: yNnIwGBS7PxDpZAIteuD
receive: 1::

send: 2:::
receive: 2::

send: 1::/nba_v1
receive: 1::/nba_v1

send: 5::/nba_v1:{"args":[{"room":"NBA_HOME"}],"name":"join"}
receive: 5::/nba_v1:{"name":"wall","args":[{"room":"NBA_HOME","result":[{"begin_time":"1516665600","process":"已结束","home_score":"112","away_score":"107","gid":"153722","status":"1"},{"home_score":"104","status":"1","process":"已结束","gid":"153723","begin_time":"1516667400","away_score":"90"},{"home_score":"132","status":"1","process":"已结束","gid":"153724","begin_time":"1516669200","away_score":"128"},{"begin_time":"1516669200","process":"已结束","home_score":"99","away_score":"90","gid":"153725","status":"1"},{"begin_time":"1516669200","process":"已结束","home_score":"109","away_score":"105","gid":"153726","status":"1"},{"home_score":"105","status":"1","process":"已结束","gid":"153727","begin_time":"1516669200","away_score":"101"},{"begin_time":"1516671000","process":"已结束","home_score":"98","away_score":"75","gid":"153728","status":"1"},{"begin_time":"1516672800","process":"已结束","home_score":"104","away_score":"101","gid":"153729","status":"1"},{"begin_time":"1516678200","process":"已结束","home_score":"118","away_score":"126","gid":"153730","status":"1"},{"gid":"10004674","process":"已结束","status":"1","begin_time":"1516708805","home_score":"0","away_score":"0"}]}]}

send: 5::/nba_v1:{"args":[{"roomid":-1,"gid":100721,"pid":429,"room":"NBA_PLAYBYPLAY_CASINO"}],"name":"join"}
receive: 5::/nba_v1:{"name":"wall","args":[{"room":"CBA_PLAYBYPLAY_CASINO","pid":429,"result":{},"type":"cba","gid":"100721","room_live_type":"1"}]}

receive: 5::/nba_v1:{"name":"wall","args":[{"room":"CBA_PLAYBYPLAY_CASINO","gid":"1
            00721","result":{"scoreboard":{"gid":"100721","status":{"id":2,"desc":"\xe8\
            xbf\x9b\xe8\xa1\x8c\xe4\xb8\xad"},"process":"\xe7\xac\xac\xe4\xba\x8c\xe8\x8
            a\x82 2:28","home_score":"52","away_score":"36","race_v":[0,-3,-4,-4,0,-2,-4
            ,-3,-5,-5,-9,-9,-11,-13,-14,-14,-13,-13,-13,-16,-15,-15]}}}]}
### closed ###
```


## 直播message
```python
{
  "name": "wall",
  "args": [
    {
      "room": "NBA_PLAYBYPLAY_CASINO",
      "gid": "153733",
      "status": "2",
      "result": {
        # 竞猜
        "casino_update": [
          {
            "casino_id": 412051,
            "desc": "485人参与，402人选「能」",
            "max_bet": 0,
            "status": {
              "id": 1,
              "desc": "竞猜中"
            },
            "user_count": 485,
            "user_win_coins": 0
          },
          {
            "casino_id": 412050,
            "desc": "580人参与，473人选「能」",
            "max_bet": 0,
            "status": {
              "id": 1,
              "desc": "竞猜中"
            },
            "user_count": 580,
            "user_win_coins": 0
          },
        ],
        # 计分板
        "scoreboard": {
          "home_tid": "21",
          "home_score": "4",
          "away_tid": "10",
          "away_score": "6",
          "process": "第一节 9:35",
          "status": "2"
        },

        # 直播内容
        "data": [
          {
            "a": [
              {
                "rowId": 0,
                "content": {
                  "event": "谁先得到25分？",
                  "end_time": "情兽",
                  "team": 0,
                  "u": "27341070",
                  "type": 1,
                  "casino": {
                    "casino_id": 412053,
                    "content": "谁先得到25分？",
                    "max_bet": 0,
                    "user_count": 0,
                    "desc": "0人参与",
                    "status": {
                      "id": 1,
                      "desc": "竞猜中"
                    },
                    "answers": [
                      {
                        "answer_id": 1,
                        "title": "骑士"
                      },
                      {
                        "answer_id": 2,
                        "title": "马刺"
                      }
                    ]
                  },
                  "t": 1516756189
                }
              }
            ],
            "d": [

            ]
          }
        ]
      },
      "online": "43万"
    }
  ]
}
```