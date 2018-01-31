#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/1/27 下午2:22
# @Author  : wudizhangzhi


import hupu
from setuptools import setup, find_packages

install_requires = open("requirements.txt").readlines()

setup(
    name='HupuApp',
    packages=find_packages(exclude=('test',)),
    package_data={'hupu': ['tac.csv']},
    version=hupu.__version__,
    author=hupu.__author__,
    author_email='554330595@qq.com',
    url="https://github.com/wudizhangzhi/HupuApp",
    # download_url='https://github.com/wudizhangzhi/HupuApp/archive/1.0.0.tar.gz',
    description='Proudly presented by Hupu JRs',
    license="MIT",
    # py_modules=['hupu'],
    keywords='hupu',
    entry_points={
        "console_scripts": ["HupuApp = hupu.hupuapp:start", ]
    },
    python_requires='>=2.6,',
    install_requires=install_requires,
    include_package_data=True,
    data_files=[
        ('HupuApp', ['requirements.txt', ]),
    ],
    classifiers=[
        # How mature is this project? Common values are
        #   3 - Alpha
        #   4 - Beta
        #   5 - Production/Stable
        'Development Status :: 3 - Alpha',
        'Programming Language :: Python :: 2',
        'Programming Language :: Python :: 3',
    ],
)
