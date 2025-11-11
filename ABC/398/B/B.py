A = list(map(int,input().split()))

result = "No"
# seven_card = list(set(A))
# five_card = []

# for i in seven_card:
#   card = A.count(i)
#   five_card.append(card)
#   if 3 in five_card and 2 in five_card:
#     result = "Yes"
for i in set(A):
  if A.count(i) >= 3:
    for j in set(A):
      if i != j and A.count(i) >= 2:
        result = "Yes"

print(result)