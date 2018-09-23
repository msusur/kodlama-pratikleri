import itertools
from typing import Iterable


def first_missing_positive(iterable: Iterable[int]) -> int:
    numbers = set(iterable)
    return next(number for number in itertools.count(1) if number not in numbers)

assert first_missing_positive([1,2,0]) == 3
assert first_missing_positive([3,4,-1,1]) == 2
assert first_missing_positive([7,8,9,11,12]) == 1
