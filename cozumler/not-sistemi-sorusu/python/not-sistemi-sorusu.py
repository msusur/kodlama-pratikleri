def get_letter_grade(score):
    grades = [(90, 'AA'), (80, 'BA'), (70, 'BB'), (65, 'CB'), (60, 'CC'), (50, 'DD'), (30, 'FD'), (0, 'FF')]
    for i in range(len(grades)):
        if score >= grades[i][0]:
            return grades[i][1]


assert get_letter_grade(30) == "FD"
assert get_letter_grade(90) == "AA"
assert get_letter_grade(66) == "CB"
