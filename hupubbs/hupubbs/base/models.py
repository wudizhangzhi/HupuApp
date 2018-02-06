import datetime

from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy import (Column, Integer, String, CHAR, TEXT, DATE, DATETIME, ForeignKey, create_engine)
from sqlalchemy.dialects.mysql import LONGTEXT
from sqlalchemy.orm import (sessionmaker, relationship)

# 创建对象的基类:
Base = declarative_base()
# 初始化数据库连接:
try:
    engine = create_engine('mysql+pymysql://root:@localhost:3306/hupu?charset=utf8mb4', echo=False)
except ModuleNotFoundError:
    engine = create_engine('mysql+mysqldb://root:@localhost:3306/hupu?charset=utf8mb4', echo=False)


class BBSPost(Base):  # 帖子
    __tablename__ = 'BBSPost'

    id = Column(Integer, primary_key=True, nullable=False)
    title = Column(TEXT, nullable=False)
    author_id = Column(Integer, ForeignKey('JRS.id'), nullable=False)  # 作者

    bbsid = Column(Integer, nullable=False, unique=True)
    post_time = Column(DATETIME, nullable=True)

    content = Column(LONGTEXT)
    createtime = Column(DATETIME, nullable=False, default=datetime.datetime.now())

    post = relationship("BBSPostComment")

    def __repr__(self):
        return "<BBSPost(title='{}')>".format(self.title)


class BBSPostComment(Base):  # 帖子评论
    __tablename__ = 'BBSPostComment'

    id = Column(Integer, primary_key=True, nullable=False)
    comment_id = Column(Integer, unique=True)

    post_id = Column(Integer, ForeignKey('BBSPost.id'), nullable=False)  # 帖子id

    author_id = Column(Integer, ForeignKey('JRS.id'), nullable=False)  # 作者
    comment_time = Column(DATETIME, nullable=True)

    content = Column(LONGTEXT)
    createtime = Column(DATETIME, nullable=False, default=datetime.datetime.now())

    def __repr__(self):
        return "<BBSPost(title='{}')>".format(self.comment_id)


class JRS(Base):
    __tablename__ = 'JRS'

    id = Column(Integer, primary_key=True, nullable=False)
    uid = Column(String(20), nullable=False, unique=True)

    name = Column(String(length=255), nullable=False)
    avatar = Column(String(length=255), nullable=True)

    createtime = Column(DATETIME, nullable=False, default=datetime.datetime.now())

    post = relationship('BBSPost')
    author = relationship("BBSPostComment")

    def __repr__(self):
        return "<JRS(name='{}')>".format(self.name)


def init_db():
    drop_db()
    Base.metadata.create_all(engine)


def drop_db():
    Base.metadata.drop_all(engine)


def test():
    # 创建DBSession类型:
    DBSession = sessionmaker(bind=engine)
    # 创建session对象:
    session = DBSession()
    query = session.query(BBSPost)
    print(query)
    r = query.filter(BBSPost.author_id == 12).order_by(BBSPost.createtime.desc()).all()
    print(r)
    session.rollback()

    # insert jrs
    # jrs = JRS(uid=123456, name='jack')
    # session.add(jrs)
    # session.commit()
    jrs = session.query(JRS).one()
    print(jrs)
    print(jrs.post)
    post = session.query(BBSPost).first()
    print(post)

    session.close()


if __name__ == '__main__':
    init_db()
    # print(BBSPost.__name__)
    # test()
    # print(BBSPost.__table__)
