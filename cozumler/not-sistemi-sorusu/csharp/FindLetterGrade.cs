using System;
public class Program
{
    public static void Main(string[] args)
    {
        Console.WriteLine(LetterGrade(66));
    }
    static string LetterGrade(int numberGrade)
    {
        return numberGrade >= 90 ? "AA" : numberGrade >= 80 ? "BA" : numberGrade >= 70 ? "BB" : numberGrade >= 65 ? "CB" : numberGrade >= 60 ? "CC" : numberGrade >= 50 ? "DD" : numberGrade >= 30 ? "FD" : numberGrade >= 0 ? "FF" : "Geçersiz Not";
    }
}