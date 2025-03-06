D = input()

directions = {
  "N": "S",
  "S": "N",
  "E": "W",
  "W": "E",
  "NE": "SW",
  "SE": "NW",
  "NW": "SE",
  "SW": "NE"
}

print(directions[D])