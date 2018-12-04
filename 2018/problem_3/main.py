from collections import defaultdict


def main():
    with open('input', 'r') as f:
        fabric = defaultdict(lambda: defaultdict(lambda: 0))
        lines = f.readlines()
        for line in lines:
            claim, x, y, width, height = parse_rect(line)

            for x_cord in range(x, x+width):
                for y_cord in range(y, y+height):
                    fabric[x_cord][y_cord] += 1

        result = 0
        for x in fabric:
            for y in fabric[x]:
                if fabric[x][y] > 1:
                    result += 1
        print(result)

        for line in lines:
            claim, x, y, width, height = parse_rect(line)
            free = check_free_rect(claim, x, y, width, height, fabric)
            if free:
                print(claim)
                break


def check_free_rect(claim, x, y, width, height, fabric):
    for x_cord in range(x, x+width):
        for y_cord in range(y, y+height):
            if fabric[x_cord][y_cord] > 1:
                return False
    return True


def parse_rect(line):
    inputs = line.split('@')
    claim = inputs[0]
    sub_input = inputs[1].split(':')
    temp_pos = sub_input[0]
    temp_size = sub_input[1]
    positions = temp_pos.split(',')
    size = temp_size.split('x')

    x_coord = int(positions[0])
    y_coord = int(positions[1])

    width = int(size[0])
    height = int(size[1])

    return claim, x_coord, y_coord, width, height


if __name__ == '__main__':
    main()
