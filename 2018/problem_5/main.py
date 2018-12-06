from string import ascii_lowercase


def main():
    with open('input', 'r') as f:
        lines = []
        for line in f:
            lines.append(line.rstrip('\n'))

        reduced_polymer = react(lines[0])
        print(len(reduced_polymer))

        best_length = len(reduced_polymer)
        for c in ascii_lowercase:
            test_polymer = reduced_polymer.replace(c.upper(), '').replace(c.lower(), '')

            length_test_polymer = len(react(test_polymer))

            if length_test_polymer < best_length:
                best_length = length_test_polymer

        print(best_length)


def react(polymer):
    stack = []
    for c in polymer:
        if stack and should_react(c, stack[-1]):
            stack.pop()
        else:
            stack.append(c)
    return ''.join(stack)


def should_react(a, b):
    return ((a.lower() == b.lower()) and ((a.isupper() and b.islower()) or (a.islower() and b.isupper())))


if __name__ == '__main__':
    main()
