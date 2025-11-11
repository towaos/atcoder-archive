N = int(input())

result = ""

if N % 2 == 0:
  num = N // 2 - 1
  for i in range(N):
    if num == i or num + 1 == i:
      result += "="
    else:
      result += "-"
else:
  num = N // 2
  for i in range(N):
    if num == i:
      result += "="
    else:
      result += "-"

print(result)

# 別解

# N = int(input())
# if N % 2 == 1:
#   s = '-' * (N // 2) + '=' + '-' * (N // 2)
# else:
#   s = '-' * (N // 2 - 1) + '==' + '-' * (N // 2 - 1)
# print(s)