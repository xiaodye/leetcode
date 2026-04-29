from collections import defaultdict
from typing import List


class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        # key: 排序后的字符串，value: []，收集的字符串数组

        map_ = defaultdict(list)

        for i in range(len(strs)):
            # 字符排序
            key = "".join(sorted(strs[i]))
            map_[key].append(strs[i])

        return list(map_.values())
