#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/1/27 下午2:22
# @Author  : wudizhangzhi
from __future__ import print_function
from __future__ import absolute_import

import six


def get_value(instance, path, default=None):
    dic = instance.__dict__
    for entry in path.split('.'):
        dic = dic.get(entry)
        if dic is None:
            return default
    return dic or default


# TODO 太蠢了!!!
def get_value_list(instance, path, default=None):
    # list - > 循环调用 dict方法 返回下一层
    # dict - > 直接调用 返回下一层
    # 如果 path 没了 添加
    dic = instance.__dict__
    result = []

    def get_value(_instance, _path, buffer):
        if not _instance:
            return
        if isinstance(_instance, list):
            for _ins in _instance:
                get_value(_ins, _path, buffer)
        else:
            if len(_path) == 0:
                buffer.append(_instance)
            else:
                _instance = _instance.get(_path[0])
                get_value(_instance, _path[1:], buffer)

    get_value(dic, path.split('.'), result)
    return result or default


def to_text(value, encoding="utf-8"):
    if isinstance(value, six.text_type):
        # TODO 适配测试
        if not six.PY3:
            if isinstance(value, unicode):
                value = value.encode(encoding)
        return value
    if isinstance(value, six.binary_type):
        return value.decode(encoding)
    return six.text_type(value)


class BaseEntry(object):
    def __init__(self, entry, default=None):
        self._name = entry.split('.')[-1]
        self.entry = entry
        self.default = default

    def __set__(self, instance, value):
        instance.__dict__[self._name] = value


class IntEntry(BaseEntry):
    def __get__(self, instance, owner):
        try:
            return int(get_value(instance, self.entry, self.default))
        except TypeError:
            return


class StringEntry(BaseEntry):
    def __get__(self, instance, owner):
        v = get_value(instance, self.entry, self.default)
        if v:
            return to_text(v)
        else:
            return v


class ListEntry(BaseEntry):
    def __get__(self, instance, owner):
        try:
            return list(get_value_list(instance, self.entry, self.default))
        except TypeError:
            return


if __name__ == '__main__':
    class A(object):
        test = ListEntry('test')

        def __init__(self, message):
            self.__dict__.update(message)


    a = A({'test': [1, 2, 3, 4]})
    print(a.test)
    a.test.append(5)
    print(a.test)
