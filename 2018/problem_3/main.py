from collections import defaultdict


def main():
    with open('input', 'r') as f:
        fabric = defaultdict(lambda: defaultdict(lambda: 0))
        for line in f.readlines():
            x, y, width, height = parse_rect(line)

            for x_cord in range(x, x+width):
                for y_cord in range(y, y+height):
                    fabric[x_cord][y_cord] += 1

        result = 0
        for x in fabric:
            for y in fabric[x]:
                if fabric[x][y] > 1:
                    result += 1
        print(result)


def parse_rect(line):
    inputs = line.split('@')[1].split(':')
    temp_pos = inputs[0]
    temp_size = inputs[1]
    positions = temp_pos.split(',')
    size = temp_size.split('x')

    x_coord = int(positions[0])
    y_coord = int(positions[1])

    width = int(size[0])
    height = int(size[1])

    return x_coord, y_coord, width, height


if __name__ == '__main__':
    main()
