#!/usr/bin/env python
# -*- coding: utf-8 -*-
import math
from math import sqrt

coordinate = [(3, 1), (3, 4), (4, 2)]

def distances(coordinate):
    values = []
    parsing = []
    distance = []
    for i in range(len(coordinate)):
        parsing.append(" ".join(str(x) for x in coordinate[i]))
        calculateValue = parsing[i].split(" ")
        firstValue = math.pow(int(calculateValue[0]), 2)
        secondValue = math.pow(int(calculateValue[1]), 2)
        distance.append(sqrt(int(firstValue) + int(secondValue)))
    return distance

def matchResult(coordinate):
        new_dict = dict(zip(coordinate,distances(coordinate)))
        return sorted(new_dict, key=new_dict.__getitem__)

print("En yakindan noktadan en uzak noktaya siralanmis hali:",matchResult(coordinate))

