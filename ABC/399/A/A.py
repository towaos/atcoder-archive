N = int(input())
S = input()
T = input()

count = 0

for i in range(N):
  if S[i] != T[i]:
    count += 1

print(count)