res = ''
dy = 1
y = 0
numRows = 3
strings = ['']*numRows
s = "PAYPALISHIRING"

for element in s:
    strings[y] += element

    if y == 0:
        dy = 1
    elif y == numRows-1:
        dy = -1

    y += dy

for i in range(numRows):
    res += strings[i]

print(*strings, sep='')
# res = ''
# dy = 1
# y = 0
# matrix = []
# numRows = 3
# s = "PAYPALISHIRING"
# for i in range(numRows):
#     matrix.append([])
#
# for element in s:
#     if y == 0:
#         dy = 1
#         matrix[y].append(element)
#         y += dy
#     elif y == numRows-1:
#         dy = -1
#         matrix[y].append(element)
#         y += dy
#     else:
#         matrix[y].append(element)
#         y += dy
#
#
# for i in range(len(matrix)):
#     res += ''.join(matrix[i])
#
# print(res)