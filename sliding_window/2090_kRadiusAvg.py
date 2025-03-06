class Solution():
    def getAverages(self, num: list[int], k: int) -> list[int]:
        if len(num) < 2 * k + 1:
            return [-1]*len(num)
        
        ans = [-1] * len(num)
        total = 0
        for i in range(len(num)):
            total += num[i]
            
            if i < 2 * k:
                continue
            
            ans[i - k] = total // (2 * k + 1)
            total -= num[i - 2 * k]
        return ans
    
if __name__ == "__main__":
    test = Solution()
    print(test.getAverages([7,4,3,9,1,8,5,2,6], 3))