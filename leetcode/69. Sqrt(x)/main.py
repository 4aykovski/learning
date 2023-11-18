def mySqrt(x: int) -> int:
    prev_root = 0
    for i in range(1, x // 2 + 3):
        if prev_root*prev_root <= x < i*i:
            return prev_root
        else:
            prev_root = i


if __name__ == '__main__':
    mySqrt(1)
    print(1//2)
