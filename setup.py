"""
distutils/setuptools install script
"""

import os
from setuptools import setup, find_packages


def find_install_requires():
    """
    Find install_requires from requirements.txt
    """
    path = os.path.join(os.path.dirname(__file__), "requirements.txt")
    return [x.strip() for x in open(path) if not x.startswith("#")]


setup(
    name="quizlet",
    version="0.0.1",
    description="Convert e-mail text in my english lesson into english words list to import them to quizlet",
    url="https://github.com/shibataka000/quizlet",
    packages=find_packages(),
    install_requires=find_install_requires(),
    scripts=["bin/quizlet"],
)
