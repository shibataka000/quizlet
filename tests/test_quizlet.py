"""
Tests for quizlet module.
"""

from quizlet.quizlet import parse


def test_parse():
    """
    Test parse function.
    """
    text = """broaden(v)            広げる、拡大する/ to make or become broad or broader; widen  e.g. broaden their knowledge, broaden one’s horizons視野を広げる
significantly(adv)   大いに、著しく、かなり/ in a statistically significant way; hugely, notably, considerably, remarkably
e.g. increase significantly, change significantly
municipal(n) /mjuːnísəpəl/    地方自治体の/ of or pertaining to a city, town, etc., or its local government  e.g. municipal elections, municipal policy"""
    expect = [
        {
            "word_en": "broaden",
            "word_jp": "広げる、拡大する",
            "kind": "v",
            "express": "to make or become broad or broader; widen",
            "example": "broaden their knowledge, broaden one’s horizons視野を広げる"
        },
        {
            "word_en": "significantly",
            "word_jp": "大いに、著しく、かなり",
            "kind": "adv",
            "express": "in a statistically significant way; hugely, notably, considerably, remarkably",
            "example": "increase significantly, change significantly"
        },
        {
            "word_en": "municipal",
            "word_jp": "地方自治体の",
            "kind": "n",
            "express": "of or pertaining to a city, town, etc., or its local government",
            "example": "municipal elections, municipal policy"
        }
    ]
    words = parse(text.split("\n"))
    for i, actual in enumerate(words):
        assert actual.word_en == expect[i]["word_en"]
        assert actual.word_jp == expect[i]["word_jp"]
        assert actual.kind == expect[i]["kind"]
        assert actual.express == expect[i]["express"]
        assert actual.example == expect[i]["example"]
