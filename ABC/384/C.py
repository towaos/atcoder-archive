N, Q = map(int, input().split())

query = []
nest = []

for i in range(N):
  nest.append([i+1])

for i in range(Q):
  x = list(map(int, input().split()))
  query.append(x)

for i in range(len(query)):
  if len(query[i]) == 1:
    count = 0
    for i in nest:
      if len(i) > 1:
        count += 1
    print(count)
  else:
    P, H = query[i][1], query[i][2] - 1
    for i in nest:
      if P in i:
        i.remove(P)
        break
    nest[H].append(P)