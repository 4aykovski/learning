nums = [-1, 0, 1, 2, -1, -4]

if len(nums) == 3:
    if sum(nums) == 0:
        print(nums)
    else:
        print([])

res = []

element = nums[0]
print(element)
element = nums[len(nums)-1-0]
print(element)
# for i in range(len(nums) // 2):
#     element1 = nums[i]
#     element2 = nums[:-i]
#     for j in range(len(nums)):
#         element3 = nums[j]
#         if element1 + element2 + element3 == 0:
#             res.append([element1, element2, element3])
