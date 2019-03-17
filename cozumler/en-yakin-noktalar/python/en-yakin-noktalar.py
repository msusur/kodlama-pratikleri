import math


class Point:
    def __init__(self, point, length):
        self.point = point
        self.length = length

    def __str__(self):
        return str(self.point)


class Facebook:
    def __init__(self):
        self.points = [[1, 2], [5, 7], [8, 9], [2, 3], [2, 2]]  # (x2,y2)
        self.origin = (0, 0)  # (x1,y1)

    def find_nearest_neighbor(self, k_number):
        return sorted([self.metric(point) for point in self.points], key=lambda k: k.length)[:k_number]

    def metric(self, point):
        return Point(point,
                     math.sqrt((math.pow(point[0] - self.origin[0], 2) + math.pow(point[1] - self.origin[1], 2))))


if __name__ == "__main__":
    k = 3
    result = Facebook().find_nearest_neighbor(k)
    print("En yakın {} nokta ->".format(k), [p.point for p in result])
