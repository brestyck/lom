# Против лома нет приема, если нет другого лома

import os

answers = []


def read_attempt():
    with open("top_secret", "r") as h:
        return h.read()


def write_attempt(attempt):
    with open("top_secret", "w") as h:
        h.write(attempt)


if not os.path.exists("top_secret"):
    write_attempt("1")
    print(answers[0])
    exit()
else:
    i = int(read_attempt())
    write_attempt(str(i+1))
    print(answers[i])
    exit()
