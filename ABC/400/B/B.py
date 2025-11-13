N, M = map(int,input().split())

total = 0
result = "inf"

for i in range(M + 1):
  total += N ** i

if total <= 10**9:
  result = total

print(result)