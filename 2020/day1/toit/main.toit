import .file as file
import reader show BufferedReader

main args:
  lines := get_input("../input")
  print(part_a(lines))
  print(part_b(lines))

part_a lines -> int:
  for i := 0; i < lines.size - 2; i++:
    for j:= i + 1; j < lines.size - 1; j++:
      a := int.parse lines[i]
      b := int.parse lines[j]
      if a + b == 2020:
        return a * b
  return 0

part_b lines -> int:
  for i := 0; i < lines.size - 2; i++:
    for j:= i + 1; j < lines.size - 1; j++:
      for x:= j + 1; x < lines.size; x++:
        a := int.parse lines[i]
        b := int.parse lines[j]
        c := int.parse lines[x]
        if a + b + c == 2020:
          return a * b * c
  return 0

get_input path -> List:
  reader := BufferedReader(file.Stream.for_read path)

  lines := []
  while line := reader.read_line:
    lines.add line
  return lines
