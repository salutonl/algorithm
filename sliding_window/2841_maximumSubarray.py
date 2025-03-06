from math import inf
class Solution():
    def maxSum(self, nums: list[int], m: int, k: int) -> int:
        counter = {}
        ans = -inf
        sum = 0
        
        for i, n in enumerate(nums):
            sum += n
            if str(n) in counter:
                counter[str(n)] += 1
            else:
                counter[str(n)] = 1
                
            if i < k - 1:
                continue
            
            if len(counter) >= m:
                ans = max(ans, sum)
                
            sum -= nums[i - k + 1]
            counter[str(nums[i - k + 1])] -= 1
            if counter[str(nums[i - k + 1])] <= 0:
                counter.pop(str(nums[i - k + 1]))
                
        return 0 if ans == -inf else ans
    
if __name__ == "__main__":
    test = Solution()
    print(test.maxSum([1,2,1,2,1,2,1], 3, 3))