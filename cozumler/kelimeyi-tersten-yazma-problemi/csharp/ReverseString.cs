using System;
class MainClass {
  public static void Main (string[] args) {
    Console.WriteLine(ReverseString("Bunu tersten yaz bakalım"));
  }
  static string nText="";  
  static string ReverseString(string text)
  {    
    if(text.Length>0)
    {
      nText+=text.Substring(text.Length-1,  1);      
      ReverseString(text.Substring(0,text.Length-1));
    }   
    return nText;
  }
}