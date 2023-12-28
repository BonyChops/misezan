from src.misezan_py import misezan
import json
import os
import pytest


# def test_cases():
#     with open(os.path.join(os.path.dirname(os.path.realpath(__file__)), "..", "..", "test", "test_cases.json")) as file:
#         case_sets = json.load(file)
#         for cases in case_sets.values():
#             for case in cases:
#                 assert misezan.calc(case['a'], case['b']) == case['ans']
#                 assert misezan.calc(case['b'], case['a']) == case['ans']


with open(os.path.join(os.path.dirname(os.path.realpath(__file__)), "..", "..", "test", "test_cases.json")) as file:
    case_sets = json.load(file)

    @pytest.mark.parametrize("case", case_sets['basic'] + case_sets['extra'] + case_sets['extraFloat'])
    def test_eval(case):
        assert misezan.calc(case['a'], case['b']) == case['ans']
