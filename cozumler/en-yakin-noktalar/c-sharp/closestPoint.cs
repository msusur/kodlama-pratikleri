  public class MathProblems
  {
      public List<Point> closestPoints(Point[] points)
      {
          List<Point> tempPoints = new List<Point>();
          if (points != null && points.Length > 0)
          {
              tempPoints.Add(points[0]);
              for (int i = 1; i < points.Length; i++)
              {
                  if ((square(tempPoints[0].X) + square(tempPoints[0].Y)) > (square(points[i].X) + square(points[i].Y)))
                      tempPoints = new List<Point> { points[i] };
                  else if ((square(tempPoints[0].X) + square(tempPoints[0].Y)) == (square(points[i].X) + square(points[i].Y)))
                      tempPoints.Add(points[i]);
              }
          }
          return tempPoints;
      }

      private int square(int i)
      {
          return i * i;
      }
  }   
  public class Point
  {
      public int X { get; set; }
      public int Y { get; set; }
  }
