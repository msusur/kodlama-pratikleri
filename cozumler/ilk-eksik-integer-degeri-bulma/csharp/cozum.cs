using System;
					
public class Program
{
	public static void Main()
	{		
		int[] _array = { 1,2,0 };  //Input
		int result= FindMinValue(_array);
		Console.WriteLine(Output  + result);
	}
	static int FindMinValue(int[] values)
	{
		int minPositiveVal=1;
		for(int i=0; i < values.Length; i++)
		{	
			var _res = Array.IndexOf(values,minPositiveVal);
			if(_res==-1) break; else minPositiveVal++;	
		}					
		return minPositiveVal;
	}
}