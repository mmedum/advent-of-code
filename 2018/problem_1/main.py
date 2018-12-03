# https://adventofcode.com/2018/day/1
# Starting with a frequency of zero, what is the resulting frequency
# after all of the changes in frequency have been applied?


def main():
    with open('input', 'r') as inputs:
        freq = 0
        for str_input in inputs.readlines():
            freq = perform_freq_change(str_input, freq)

        print(freq)


def perform_freq_change(str_input, current_freq):
    opr = str_input[0:1]
    freq_change = int(str_input[1:])

    if opr == '-':
        current_freq -= freq_change
    else:
        current_freq += freq_change
    return current_freq


if __name__ == '__main__':
    main()
