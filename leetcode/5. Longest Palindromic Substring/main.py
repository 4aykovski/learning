def longest_palindrome(self, s: str) -> str:
    res = ''

    def get_pal(s, left, right):
        while left >= 0 and right < len(s) and s[left] == s[right]:
            left -= 1
            right += 1
        return s[left + 1:right]

    for i in range(len(s)):
        pal1, pal2 = get_pal(s, i, i), get_pal(s, i, i + 1)
        if len(pal1) > len(res):
            res = pal1
        if len(pal2) > len(res):
            res = pal2

    return res
