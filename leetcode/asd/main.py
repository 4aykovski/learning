nums = [1,2]
largest_sum = 0
for index, i in enumerate(nums):
    last_sum = 0
    for j in nums[index:]:
        curr_sum = last_sum + j
        if curr_sum > largest_sum:
            largest_sum = curr_sum
        last_sum = curr_sum
print(largest_sum)