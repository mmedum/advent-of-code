import .file as file
import reader show BufferedReader

main args:
  lines := get_input(args[0])
  print lines

get_input path -> List:
  reader := BufferedReader(file.Stream.for_read path)

  lines := []
  while line := reader.read_line:
    lines.add line
  return lines
