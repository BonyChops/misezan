pattern = {
    '6': '9',
    '9': '6',
    '2': '5',
    '5': '2',
    '.': '.',
}


def calc(a: int | float, b: int | float) -> float:
    if type(a) not in (int, float) or type(b) not in (int, float):
        raise TypeError("a and b must be int or float")

    if a == b:
        return 0

    min, max = (a, b) if a < b else (b, a)

    if max >= min * 100:
        return max - min*17

    min_str, max_str = (str(min), str(max))
    if len(min_str) != len(max_str):
        return max

    is_special = False
    for i in range(len(min_str)):
        if min_str[i] not in pattern:
            return max
        if pattern[min_str[i]] != max_str[len(max_str) - i - 1]:
            return max
        if min_str[i] == '2' or min_str[i] == '5':
            is_special = True

    return 1.1 if is_special else 11
