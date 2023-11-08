nums = [1,1,1,5,5,5,5,5,5]
target = 14
smaller, bigger = sum(nums[:3]), sum(nums[:3])
nums.sort()
for i in range(len(nums) - 2):
    l = i + 1
    r = len(nums) - 1
    while l < r:
        triplet = [nums[i], nums[l], nums[r]]
        summa = sum(triplet)
        if summa == target:
            print(target)
            break
        elif summa < target:
            if target - bigger > target - summa > 0:
                bigger = summa
            l += 1
        elif summa > target:
            if abs(target - summa) < abs(target - smaller):
                smaller = summa
            r -= 1

print(bigger if abs(bigger - target) <= abs(smaller - target) else smaller)
