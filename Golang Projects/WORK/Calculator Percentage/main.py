import threading
import time

scoresIELTS = [4.0, 4.5, 5.0, 5.5, 6.0, 6.5, 7.0, 7.5, 8.0, 8.5, 9.0]
scoresTOEFL = [[10, 31], [32, 34], [35, 45], [46, 59], [60, 78], [79, 93],
               [94, 101], [102, 109], [110, 114], [115, 117], [118, 120]]


# def LOOP(liste):
#     x = 0
#     while x < len(liste):
#         print(liste[x])
#         x += 1

# LOOP(scoresIELTS)
# LOOP(scoresTOEFL)


matrix = [[0 for j in range(4)] for i in range(11)]


def LOOPI2(liste):
    x = 0
    while x < len(liste):
        matrix[x][0] = liste[x]
        matrix[x][1] = round((liste[x]/9*100), 1)
        x += 1


def LOOPT2(liste):
    x = 0
    while x < len(liste):
        matrix[x][2] = liste[x]
        matrix[x][3] = round((liste[x][0]/120*100), 1)
        x += 1


LOOPI2(scoresIELTS)
LOOPT2(scoresTOEFL)

print("\n\tIELTS\t\t  ---  \t\t\tTOEFL\t\n")
for i in matrix:
    print(f"{i[0]}\t >> \t%{i[1]}\t  ---  \t{i[2]}\t >> \t%{i[3]}\n")
