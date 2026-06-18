## Two Sum

- Function takes a slice of integers nums and a target integer
- Return indices of the two numbers such that they add up to target

##### Workflow

nums = [3, 2, 4], target = 6

seen = {}  (empty map)

i=0, num=3:
  needed = 6-3 = 3
  seen[3]? No (ok=false)
  seen[3] = 0  → map now has {3:0}

i=1, num=2:
  needed = 6-2 = 4
  seen[4]? No (ok=false)
  seen[2] = 1  → map now has {3:0, 2:1}

i=2, num=4:
  needed = 6-4 = 2
  seen[2]? Yes! j=1, ok=true
  return [1, 2]  ✅   (nums[1]=2, nums[2]=4, 2+4=6)