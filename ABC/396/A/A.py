N = int(input())
A = list(map(int, input().split()))

count = 1
result = False

for i in range(1, N):
  if A[i] == A[i - 1]:
    count += 1
    if count >= 3:
      result = True
      break
  else:
    count = 1

if result:
  print("Yes")
else:
  print("No")