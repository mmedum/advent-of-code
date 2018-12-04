
def main():
    with open('input', 'r') as inputs:
        two_equal = 0
        three_equal = 0
        for line in inputs.readlines():
            two, three = scan_input(line)
            two_equal += two
            three_equal += three
        print(two_equal * three_equal)


def scan_input(line):
    two_equal = 0
    three_equal = 0
    chars = {}
    for c in line:
        if c in chars:
            nr_of_c = chars[c] + 1
            if nr_of_c == 2:
                two_equal += 1
            elif nr_of_c == 3:
                two_equal -= 1
                three_equal += 1
            chars[c] = nr_of_c
        else:
            chars[c] = 1
    if two_equal > 0 and three_equal > 0:
        return 1, 1
    elif two_equal > 0:
        return 1, 0
    elif three_equal > 0:
        return 0, 1
    else:
        return 0, 0


if __name__ == '__main__':
    main()
