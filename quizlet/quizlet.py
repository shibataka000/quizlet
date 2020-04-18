"""
Convert text in english lesson e-mail into english words list to import them to quizlet.
"""

from re import match
from collections import namedtuple

Word = namedtuple("Word", ("word_en", "word_jp", "kind", "express", "example"))


def read():
    """
    Read text in english lesson e-mail.
    """
    text, blank_line_count = [], 0
    while blank_line_count < 2:
        row = input()
        if row == "":
            blank_line_count += 1
        else:
            blank_line_count = 0
            text.append(row)
    return text


def parse(text):
    """
    Convert text in english lesson e-mail into english words list.
    """
    lined_text = []
    for row in text:
        row = row.lstrip()
        if row.startswith("e.g."):
            lined_text[-1] += " " + row
        else:
            lined_text.append(row)
    words = []
    for row in lined_text:
        r = match('([\\w\\s]+)\\((\\w+)\\)\\s*(\\/.*\\/)?\\s*(.+)\\/(.+)e\\.g\\.(.+)', row)
        if r is None:
            print(f"[WARN] Parsing failed: {row}")
            continue
        p = [x.strip() if x else "" for x in r.groups()]
        word = Word(p[0], p[3], p[1], p[4], p[5])
        words.append(word)
    return words


def output(words):
    """
    Print english words list.
    """
    for word in words:
        print(f"{word.word_en}（{word.word_jp}）\t{word.express}")
