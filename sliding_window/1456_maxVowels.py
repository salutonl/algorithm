class Solution:
    def maxVowels(self, s: str, k: int) -> int:
        vowels = set("aeiou")
        ans = 0
        for i in range(k - 1):
            if s[i] in vowels:
                ans += 1
        res = 0
        for i in range(k - 1, len(s)):
            if s[i] in vowels:
                ans += 1
            res = max(res, ans)
            if s[i - k + 1] in vowels:
                ans -= 1
        return res
    
    
if __name__ == "__main__":
    test = Solution()
    print(test.maxVowels("ramadan", 2))
    s = "asdfsdfs"
    for i, c in enumerate(s):
        print(i, c)