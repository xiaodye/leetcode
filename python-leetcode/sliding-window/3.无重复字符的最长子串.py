class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        if len(s) == 0:
            return 0

        # 滑动窗口
        i = 0
        j = 0
        max_len = 0
        char_set = set()

        while j < len(s):
            while j < len(s) and s[j] not in char_set:
                char_set.add(s[j])
                max_len = max(max_len, len(char_set))
                j += 1

            # 收缩窗口
            char_set.remove(s[i])
            i += 1

        return max_len
