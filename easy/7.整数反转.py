class Solution:
    def reverse(self, x: int) -> int:
        res = 0
        tag = 1
        if x < 0:
            x = 0-x
            tag = -1
        while(x>0):
            tmp = x%10
            res = res*10 + tmp;
            x=x//10
        if res > 2147483647 or res == 2147483647:
            return 0
        elif res < -2147483648 or res == -2147483648:
            return 0
        else:
            return res*tag