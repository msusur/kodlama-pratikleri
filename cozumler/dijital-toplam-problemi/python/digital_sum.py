def solve(num):
    p = (num-1) // 9
    return num+9 if p == 0 else (num - 9*p + 1) * (10 ** p) - 1

if __name__ == '__main__':
    print (solve(5))
    print (solve(9))
    print (solve(21))
    print (solve(27))
    print (solve(28))

