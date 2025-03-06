from math import inf
class Solution():
    def minimumRecolors(self, blocks: str, k: int) -> int:
        
        ans = inf
        whiteNum = 0
        for i, c in enumerate(blocks):
            if c == "W":
                whiteNum += 1
                
            if i < k - 1:
                continue
            
            ans = min(ans, whiteNum)
            
            if blocks[i - k + 1] == "W":
                whiteNum -= 1
                
        return ans
    
if __name__ == "__main__":
    test = Solution()
    print(test.minimumRecolors("WBWBBBW", 2))