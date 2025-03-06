from math import inf


class Solution():
    def findMaxAverage(self, num: list[int], k: int) -> float:
        ans = -inf
        total = 0
        for i in range(len(num)):
            total += num[i]
            
            if i < k - 1:
                continue
            
            ans = max(ans, total)
            
            total -= num[i - k + 1]
            
        return ans/k
    
if __name__ == "__main__":
    test = Solution()
    print(test.findMaxAverage([1,2,4,5,3,6,7,8,6,3], 2))