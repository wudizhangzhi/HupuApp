# -*- coding: utf-8 -*-

# Define your item pipelines here
#
# Don't forget to add your pipeline to the ITEM_PIPELINES setting
# See: http://doc.scrapy.org/en/latest/topics/item-pipeline.html
import json
from collections import OrderedDict
from sqlalchemy import create_engine
from sqlalchemy.orm import scoped_session, sessionmaker
from hupubbs.base.models import JRS, BBSPostComment, BBSPost
from scrapy import log
from scrapy.exceptions import DropItem


def scrapydebug(msg, level=None):
    if not level:
        level = log.DEBUG
    log.msg(msg, level=level)


class HupubbsPipeline(object):
    def process_item(self, item, spider):
        items = OrderedDict(item)
        # TODO handler items
        self.save(items)
        return item

    def open_spider(self, spider):
        # 初始化数据库连接:
        charset = 'utf8mb4'
        connect_string = 'mysql+{connector}://root:@localhost:3306/hupu?charset={charset}'
        try:
            engine = create_engine(connect_string.format(connector='pymysql', charset=charset),
                                   echo=False)
        except ModuleNotFoundError:
            engine = create_engine(connect_string.format(connector='mysqldb', charset=charset), echo=False)
        # 创建DBSession类型:
        DBSession = scoped_session(sessionmaker(bind=engine))
        # 创建session对象:
        self.session = DBSession()

    def close_spider(self, spider):
        self.session.close()

    def save(self, item):
        post_title = item['post_title']
        post_author = item['post_author']
        post_content = item['post_content']
        post_id = item['post_id']
        post_uid = item['post_uid']
        post_avatar = item['post_avatar']
        post_datetime = item['post_datetime']

        jrs = self.session.query(JRS).filter(JRS.uid == post_uid).first()
        if not jrs:  # 新建
            jrs = self.create_jrs(name=post_author, uid=post_uid, avatar=post_avatar)

        # 判断帖子
        post = self.session.query(BBSPost).filter(BBSPost.bbsid == post_id).first()

        if not post:  # 新建
            post = self.create_post(title=post_title, author_id=jrs.id, bbsid=post_id, post_time=post_datetime,
                                    content=post_content)

        for comment in item['comment']:
            comment_datetime = comment['comment_datetime']
            comment_content = comment['comment_content']
            comment_uid = comment['comment_uid']
            comment_author = comment['comment_author']
            comment_avatar = comment['comment_avatar']
            comment_id = comment['comment_id']
            # TODO delete
            if len(comment_uid.split()) > 1:
                # scrapydebug(
                #     'comment_uid: {!r} comment_author: {!r} comment_content: {!r}'.format(comment_uid, comment_author,
                #                                                                           comment_content),
                #     level=log.ERROR)
                scrapydebug('='*30, level=log.ERROR)
                scrapydebug(json.dumps(comment, indent=4), level=log.ERROR)
                scrapydebug('post_id: {}'.format(post_id), level=log.ERROR)
                scrapydebug('=' * 30)
            _jr = self.session.query(JRS).filter(JRS.uid == comment_uid).first()

            if not _jr:
                _jr = self.create_jrs(comment_author, comment_uid, comment_avatar)
            if not comment_id:
                scrapydebug('跳过: comment_uid: {}  comment_author: {}'.format(comment_uid, comment_author))
                continue
            # 判断回复是否存在
            _comment = self.session.query(BBSPostComment).filter(BBSPostComment.comment_id == comment_id).first()
            if not _comment:
                c = self.create_comment(comment_id=comment_id,
                                        post_id=post.id, author_id=_jr.id,
                                        comment_time=comment_datetime,
                                        content=comment_content)

    def create_jrs(self, name, uid, avatar):
        jrs = JRS(name=name, uid=uid, avatar=avatar)
        self.session.add(jrs)
        self.session.commit()
        scrapydebug('保存: {!r}'.format(jrs))
        return jrs

    def create_post(self, title, author_id, bbsid, post_time, content):
        # if not (title, author_id, bbsid, post_time):
        #     raise Exception('数据不全: title: {} author_id: {}')
        post_time = post_time or None
        post = BBSPost(title=title, author_id=author_id, bbsid=bbsid, post_time=post_time,
                       content=content)
        self.session.add(post)
        self.session.commit()
        scrapydebug('保存: {!r}'.format(post))
        return post

    def create_comment(self, comment_id, post_id, author_id, comment_time, content):
        comment_time = comment_time or None
        comment = BBSPostComment(comment_id=comment_id, post_id=post_id,
                                 author_id=author_id, comment_time=comment_time,
                                 content=content)
        self.session.add(comment)
        self.session.commit()
        scrapydebug('保存: {!r}'.format(comment))
        return comment
