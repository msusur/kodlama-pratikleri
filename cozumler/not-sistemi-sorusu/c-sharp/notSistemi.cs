public string Score(int numericScore)
{
    if (numericScore >= 0 && numericScore <= 100)
    {
        Dictionary<int, string> scores = new Dictionary<int, string> { { 30, "FF" }, { 50, "FD" }, { 60, "DD" }, { 65, "CC" }, { 70, "CB" }, { 80, "BB" }, { 90, "BA" }, { 101, "AA" } };

        foreach (var key in scores.Keys)
            if (numericScore < key)
                return scores[key];
    }
    return "";
}

[TestMethod]
public void TestNote()
{    
    Assert.AreEqual("FD", Score(30));
    Assert.AreEqual("AA", Score(90));
    Assert.AreEqual("CB", Score(66));
}
