# https://adventofcode.com/2018/day/1
# Starting with a frequency of zero, what is the resulting frequency
# after all of the changes in frequency have been applied?
# What is the first frequency your device reaches twice?


def main():
    with open('input', 'r') as inputs:
        seen_freq = {}
        freq = 0
        seen_freq[freq] = 1
        searching_for_dup_freq = True
        freq_inputs = inputs.readlines()
        while searching_for_dup_freq:
            for str_input in freq_inputs:
                freq = perform_freq_change(str_input, freq)
                if searching_for_dup_freq and freq in seen_freq:
                    print(f'Already seen freq: {freq}')
                    searching_for_dup_freq = False
                seen_freq[freq] = 1
            print(f'End freq: {freq}')


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
