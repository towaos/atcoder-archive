S = input()
count = 0
num = 1

for i in S:
  if (num == 1):
    if (i == "i"):
      num = 2
    elif (i != "i"):
      count += 1
      num = 1
  if (num == 2):
    if (i == "o"):
      num = 1
    elif (i != "o"):
      count += 1
      num = 2

if (num == 2):
  count += 1

print(count)