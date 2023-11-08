num = 3999

romans = {1: 'I', 5: 'V', 10: 'X', 50: 'L', 100: 'C', 500: 'D', 1000: 'M',
          4: 'IV', 9: 'IX', 40: 'XL', 90: 'XC', 400: 'CD', 900: 'CM'}
res = ''

length = len(str(num))
integers = [int(e) for e in str(num)]

for i in integers:
    if i <= 5 and length == 1:
        res += 'I' * i
        break

    if i >= 5 and length == 1:
        res += 'V' + 'I' * (i - 5)
        break

    digit = i * 10 ** (length-1)
    if digit > 1000:
        res += 'M' * (digit // 1000)
        length -= 1
        continue
    a = romans.get(int(digit))
    res += a
    length -= 1

print(res)
