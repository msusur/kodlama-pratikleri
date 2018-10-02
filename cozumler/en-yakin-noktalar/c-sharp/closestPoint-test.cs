[TestMethod]
public void TestClosestPoint1()
{
    MathProblems problems = new MathProblems();
    Point p1 = new Point { X = 0, Y = 1 };
    Point p2 = new Point { X = -1, Y = 0 };

    Point[] array = new Point[] { new Point { X = 1, Y = 2 }, p1, new Point { X = -1, Y = 1 } , p2 };
    var expectedList = new List<Point> { p1, p2 };
    var result = problems.closestPoints(array);
    CollectionAssert.AreEqual(expectedList, result);
}
