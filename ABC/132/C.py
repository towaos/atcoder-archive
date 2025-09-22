N = int(input())
d = list(map(int,input().split()))

d.sort()

ABC = d[N//2 - 1]
ARC = d[N//2]

if (ABC >= ARC):
  result = 0
else:
  result = ARC - ABC

print(result)