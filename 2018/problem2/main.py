
def main():
    with open('input', 'r') as inputs:
        two_equal = 0
        three_equal = 0
        common_ids(inputs)
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


def common_ids(inputs):
    lines = inputs.readlines()
    done = False
    for i in range(0, len(lines)-1):
        for j in range(i+1, len(lines)):
            i_line = lines[i]
            j_line = lines[j]
            number_of_chars_diff = 0
            char_diff_pos = -1
            for index, char in enumerate(i_line):
                if char != j_line[index]:
                    number_of_chars_diff += 1
                    char_diff_pos = index
                    if number_of_chars_diff > 1:
                        break
            if number_of_chars_diff == 1:
                result = i_line[0:char_diff_pos] + i_line[char_diff_pos+1:]
                print(result)
                done = True
                break
        if done:
            break


if __name__ == '__main__':
    main()
