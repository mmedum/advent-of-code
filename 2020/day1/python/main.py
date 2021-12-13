from itertools import combinations

def part_a(data):
    for combo in combinations(data, 2):
        if combo[0] + combo[1] == 2020:
            part1 = combo[0] * combo[1]
            break

    return max(i*j if i + j == 2020 else 0 for i in data for j in data)


def part_b(data):
    for combo in combinations(data, 3):
        if combo[0] + combo[1] + combo[2] == 2020:
            part2 = combo[0] * combo[1] * combo[2]
            break

        return max(i*j*k if i + j + k == 2020 else 0 for i in data for j in data for k in data)


if __name__ == "__main__":
    with open("../input") as file:
        data = [int(line) for line in file]
        a = part_a(data)
        b = part_b(data)
        print(f"A: {a}, B: {b}")
