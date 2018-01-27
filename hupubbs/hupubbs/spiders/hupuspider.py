import sys

sys.path.append("..")
import re
import json

from scrapy.selector import Selector
import time

try:
    from scrapy.spiders import Spider
except:
    from scrapy.spiders import BaseSpider as Spider
from scrapy.utils.response import get_base_url
from scrapy.spiders import CrawlSpider, Rule
from scrapy.linkextractors import LinkExtractor as sle
from hupubbs.base.spider import CommonSpider, BaseRuleSpider


class Hupu_Spider(CommonSpider):
# class Hupu_Spider(BaseRuleSpider):
    name = "hupu_spider"
    allowed_domains = ["bbs.hupu.com"]
    start_urls = [
        "https://bbs.hupu.com/",
    ]

    rules = [
        Rule(sle(allow=("bbs.hupu.com/[\d\-]+.html")), callback='parse_bbs', follow=True),
        # Rule(sle(allow=("bbs.hupu.com/*")), follow=True),
    ]

    # 帖子
    bbs_css_rules = {
        'post_title': '#j_data::text',
        'post_author': '#tpc > div > div.floor_box > div.author > div.left > a::text',
        'post_content': '#tpc > div > div.floor_box > table.case > tbody > tr > td > div.quote-content string(.)',
        'post_uid': '#tpc > div > div.user > div.j_u::attr(uid)',
        'post_avatar': '#tpc > div > div.user > div.j_u > a.headpic > img::attr(src)',
        'post_datetime': '#tpc > div > div.floor_box > div.author > div.left > span.stime::text',
        'comment|div.floor': {
            'comment_id': 'div.author > div.left > span[class*="f444"]::attr(pid)',
            'comment_datetime': 'div.author div.left > span.stime::text',
            'comment_content': 'table > tbody > tr > td string(.)',
            'comment_uid': 'div.user > div.j_u::attr(uid)',
            # 'comment_author': 'div.author div.left a.u::text',
            'comment_author': 'div.user > div.j_u::attr(uname)',
            'comment_avatar': 'div.user > div.j_u > a.headpic > img::attr(src)',
        }
    }

    # test_xpath_rules = {
    #     'post_title': '//*[@id="j_data"]/@data-title',
    #     'comment|//div[@class="floor"]': {
    #         'comment_content': './div/div[2]/table/tbody/tr/td string(.)',
    #     }
    # }

    auto_join_text = True

    def parse_bbs(self, response):
        # print('Parse ' + response.url)
        match = re.findall(r'bbs.hupu.com/(\d+)+?[\-\d]*.html', response.url)
        assert match
        post_id = match[0]
        x = self.parse_with_rules(response, self.bbs_css_rules, dict)
        # x = self.parse_with_rules(response, self.test_xpath_rules, rule_type=self.XPATH)

        for i in x:
            i['post_id'] = post_id
        # print(json.dumps(x, ensure_ascii=False, indent=2))
        return x
