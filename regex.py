#!/usr/bin/python3

import re

M = int(input())
for v in range(M):
    print(v)
    N, P = (int(n) for n in input().split())
    table = {}
    words = []
    for i in range(10):
        codes = list(input().split()[1])
        for c in codes:
            table[c] = str(i)
    for _ in range(N):
        decode = ""
        word = list(input().split()[1])
        for c in word:
            decode += table[c]
        words.append('('+decode+')')
    reg = re.compile('|'.join(words))
    for _ in range(P):
        hum = input().split()[1]
        print(reg.search(hum))
