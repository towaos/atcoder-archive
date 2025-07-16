Q = int(input())
stack = [0] * 100
result = []

for i in range(Q):
  query = list(map(int,input().split()))
  if query[0] == 1:
    x = query[1]
    stack.append(x)
  
  elif query[0] == 2:
    removed_card = stack.pop()
    result.append(removed_card)

print(*result)