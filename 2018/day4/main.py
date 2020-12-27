from collections import defaultdict, Counter


class Guard:

    def __init__(self, name):
        self._name = name
        self._time_asleep = 0
        self._minute_occurence = defaultdict(int)

    def add_sleep(self, start_time, end_time):
        for x in range(start_time, end_time):
            self._minute_occurence[x] += 1
            self._time_asleep += 1

    def get_most_sleeping_minute(self):
        return Counter(self._minute_occurence).most_common(1)

    @property
    def time_asleep(self):
        return self._time_asleep

    @property
    def name(self):
        return self._name

    def __repr__(self):
        out = '{} slept {} mins, where {} is the minute'.format(self._name, self._time_asleep,
                                                                self.get_most_sleeping_minute())
        return out


def main():
    with open('input', 'r') as f:
        lines = []
        for line in f:
            lines.append(line.rstrip('\n'))
        lines.sort()
        most_sleeping_guard = None

        guards = {}
        current_guard = None
        start_time = -1
        for line in lines:
            splits = line.split(' ')

            if len(splits) > 4:
                name = int(splits[3][1:])
                if name not in guards:
                    guards[name] = Guard(name)
                current_guard = guards[name]

            guard_function = splits[2]
            minute = int(line[15:17])
            if guard_function == 'falls':
                start_time = minute
            elif guard_function == 'wakes':
                current_guard.add_sleep(start_time, minute)
                if most_sleeping_guard is None:
                    most_sleeping_guard = current_guard
                elif current_guard.time_asleep > most_sleeping_guard.time_asleep:
                    most_sleeping_guard = current_guard
        print(most_sleeping_guard)
        print(most_sleeping_guard.name * most_sleeping_guard.get_most_sleeping_minute()[0][0])

        guard = -1
        minute = (-1, -1)
        for k, v in guards.items():
            same_minute = v.get_most_sleeping_minute()
            if not same_minute:
                continue
            same_minute = same_minute[0]
            if same_minute[1] > minute[1]:
                guard = k
                minute = same_minute

        print(f'Guard {guard} is most asleep on minute {minute}')
        print(guard * minute[0])


if __name__ == '__main__':
    main()
