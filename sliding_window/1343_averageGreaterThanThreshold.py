from math import inf
class Solution():
    def numOfSubarrays(self, arr: list[int], k: int, threshold: int) -> int:
        total = 0
        ans = 0
        
        for i, num in enumerate(arr):
            total += num
            
            if i < k - 1:
                continue
            
            if total >= k * threshold:
                ans += 1
                
            total -= arr[i - k + 1]
            
        return ans
    
if __name__ == "__main__":
    test = Solution()
    print(test.numOfSubarrays([2,2,2,2,5,5,5,8], 3, 4))