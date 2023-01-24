scoresIELTS = [4.0, 4.5, 5.0, 5.5, 6.0, 6.5, 7.0, 7.5, 8.0, 8.5, 9.0]
scoresTOEFL = [[0, 31], [32, 34], [35, 45], [46, 59], [60, 78], [79, 93],
               [94, 101], [102, 109], [110, 114], [115, 117], [118, 120]]


def LOOP(liste):
    x = 0
    while x < len(liste):
        print(liste[x])
        x += 1


LOOP(scoresIELTS)
LOOP(scoresTOEFL)
