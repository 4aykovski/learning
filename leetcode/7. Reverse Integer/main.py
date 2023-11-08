x = 1534236469
max_int = 2 ** 31 - 1
min_int = -2 ** 31
res = 0
count = 0
while x != 0:
    digit = x % 10 if x > 0 else x % -10
    res = res * 10 + digit
    x = x // 10 if x > 0 else -(x // -10)

if res >= max_int or res <= min_int:
    print(0)

else:
    print(res)
