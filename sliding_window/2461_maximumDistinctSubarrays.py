from math import inf
from collections import Counter
class Solution():
    def maximumSubarraySum(self, nums: list[int], k: int) -> int:
        cnt = Counter(nums[:k - 1])
        total = sum(nums[:k - 1])
        ans = -inf
        
        for in_, out in zip(nums[k - 1:], nums):
            total += in_
            cnt[in_] += 1
            
            if len(cnt) == k:
                ans = max(ans, total)
                
            total -= out
            cnt[out] -= 1
            if cnt[out] == 0:
                cnt.pop(out)
        
        return 0 if ans == -inf else ans
    
if __name__ == "__main__":
    test = Solution()
    print(test.maximumSubarraySum([1,5,4,2,9,9,9], 3))
        