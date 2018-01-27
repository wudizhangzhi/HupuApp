# coding=utf8
import re
import json
from scrapy.selector import Selector

try:
    from scrapy.spiders import Spider
except:
    from scrapy.spiders import BaseSpider as Spider
from scrapy.utils.response import get_base_url
from scrapy.spiders import CrawlSpider, Rule
from scrapy.linkextractors import LinkExtractor as sle
from collections import defaultdict

# from .log import *

'''
1. 默认取sel.css()[0]，如否则需要'__unique':False or __list:True
2. 默认字典均为css解析，如否则需要'__use':'dump'表明是用于dump数据
'''


class CommonSpider(CrawlSpider):
    auto_join_text = False
    ''' # css rule example:
    all_css_rules = {
        '.zm-profile-header': {
            '.zm-profile-header-main': {
                '__use':'dump',
                'name':'.title-section .name::text',
                'sign':'.title-section .bio::text',
                'location':'.location.item::text',
                'business':'.business.item::text',
                'employment':'.employment.item::text',
                'position':'.position.item::text',
                'education':'.education.item::text',
                'education_extra':'.education-extra.item::text',
            }, '.zm-profile-header-operation': {
                '__use':'dump',
                'agree':'.zm-profile-header-user-agree strong::text',
                'thanks':'.zm-profile-header-user-thanks strong::text',
            }, '.profile-navbar': {
                '__use':'dump',
                'asks':'a[href*=asks] .num::text',
                'answers':'a[href*=answers] .num::text',
                'posts':'a[href*=posts] .num::text',
                'collections':'a[href*=collections] .num::text',
                'logs':'a[href*=logs] .num::text',
            },
        }, '.zm-profile-side-following': {
            '__use':'dump',
            'followees':'a.item[href*=followees] strong::text',
            'followers':'a.item[href*=followers] strong::text',
        }
    }
    '''

    # Extract content without any extra spaces.
    # NOTE: If content only has spaces, then it would be ignored.
    def extract_item(self, sels):
        # TODO bug  [u'[\u52a8\u6f2b]']
        contents = []
        for i in sels:
            content = re.sub(r'\s+', ' ', i.extract())
            if content != ' ':
                contents.append(content)
        return contents

    def extract_items(self, sel, rules, item):
        for nk, nv in rules.items():
            if nk in ('__use', '__list'):
                continue
            if nk not in item:
                item[nk] = []
            if sel.css(nv):
                # item[nk] += [i.extract() for i in sel.css(nv)]
                # Without any extra spaces:
                item[nk] += self.extract_item(sel.css(nv))
            else:
                item[nk] = []

    # 1. item是一个单独的item，所有数据都聚合到其中 *merge
    # 2. 存在item列表，所有item归入items
    def traversal(self, sel, rules, item_class, item, items):
        # print 'traversal:', sel, rules.keys()
        if item is None:
            item = item_class()
        if '__use' in rules:
            if '__list' in rules:
                unique_item = item_class()
                self.extract_items(sel, rules, unique_item)
                items.append(unique_item)
            else:
                self.extract_items(sel, rules, item)
        else:
            for nk, nv in rules.items():
                for i in sel.css(nk):
                    self.traversal(i, nv, item_class, item, items)

    DEBUG = True

    def debug(self, sth):
        if self.DEBUG == True:
            print(sth)

    def deal_text(self, sel, item, force_1_item, k, v):
        # if v.endswith('::text') and self.auto_join_text:
        if v.endswith('string(.)') and self.auto_join_text:
            text = ' '.join(self.extract_item(sel.css(v.replace('string(.)', '')). \
                                              xpath('string(.)')))
            item[k] = text.replace('\n', '').replace('\r', '')
        elif '::' in v and self.auto_join_text:
            item[k] = ' '.join(self.extract_item(sel.css(v)))
        else:
            _items = self.extract_item(sel.css(v))
            if force_1_item:
                if len(_items) >= 1:
                    item[k] = _items[0]
                else:
                    item[k] = ''
            else:
                item[k] = _items

    keywords = set(['__use', '__list'])

    def traversal_dict(self, sel, rules, item_class, item, items, force_1_item, selector=None):
        item = {}
        for k, v in rules.items():
            if not isinstance(v, dict):
                if k in self.keywords:
                    continue
                if type(v) == list:
                    continue

                self.deal_text(sel, item, force_1_item, k, v)

            else:
                # TODO 名称和规则分开
                tmp = k.split('|')
                assert len(tmp) == 2
                k_name = tmp[0].strip()
                k_sel = tmp[1].strip()
                item[k_name] = []
                _css_selector = sel.css(k_sel) if k_sel else sel.css('')
                for sel_sub in _css_selector:
                    # for i in sel.xpath(k):
                    # print(k, v)
                    self.traversal_dict(sel_sub, v, item_class, item, item[k_name], force_1_item)
        items.append(item)

    def dfs(self, sel, rules, item_class, force_1_item):
        if sel is None:
            return []

        items = []
        if item_class != dict:
            self.traversal(sel, rules, item_class, None, items)
        else:
            self.traversal_dict(sel, rules, item_class, None, items, force_1_item)

        return items

    def parse_with_rules(self, response, rules, item_class, force_1_item=False):
        return self.dfs(Selector(response), rules, item_class, force_1_item)

    ''' # use parse_with_rules example:
    def parse_people_with_rules(self, response):
        item = self.parse_with_rules(response, self.all_css_rules, ZhihuPeopleItem)
        item['id'] = urlparse(response.url).path.split('/')[-1]
        info('Parsed '+response.url) # +' to '+str(item))
        return item
    '''


class BaseRuleSpider(CrawlSpider):
    CSS = 0
    XPATH = 1

    auto_join_text = False

    def parse_with_rules(self, response, rules, rule_type=CSS):
        items = []
        self.scan_rules(Selector(response), rules, items, rule_type)
        return items

    def scan_rules(self, selector, rules, items, rule_type=CSS):
        item = defaultdict(list)  # TODO 是否用 __setattr__
        for key, value in rules.items():
            if isinstance(value, dict):  # 后面还是字典,还需要扫描
                tmp = key.split('|')
                assert len(tmp) == 2, ValueError('Expected format keyname|rule, get: {}'.format(key))
                key_name = tmp[0]
                key_rule = tmp[1]
                # item[key_name] = []
                if rule_type == self.CSS:
                    sel_sub = selector.css(key_rule)
                else:
                    sel_sub = selector.xpath(key_rule)
                for sel in sel_sub:
                    self.scan_rules(sel, value, item[key_name], rule_type)

            else:  # 不需要继续扫描了
                self.deal_text(selector, key, value, item)

        items.append(item)

    def deal_text(self, selector, key, rule, item):
        if rule.endswith('string(.)'):
            content = []
            for sel in selector.xpath(rule.replace('string(.)', '')):
                _c = sel.xpath('string(.)').extract()
                content.extend(_c)
        else:
            content = selector.xpath(rule).extract()
        if self.auto_join_text:
            content = ''.join(content)

        item[key] = self.handler_text(content)

    def handler_text(self, text):
        # re.sub(r'\s+', '')
        if isinstance(text, list):
            return [i.strip() for i in text if i.strip()]
        else:
            return text.strip()
