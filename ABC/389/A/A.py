S = input()
total = 1

for i in S:
  if i.isdigit():
    total *= int(i)

print(total)