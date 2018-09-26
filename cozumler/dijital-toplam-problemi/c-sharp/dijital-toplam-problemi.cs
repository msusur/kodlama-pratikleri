public class MathProblems
{
	static void Main(string[] args)
        {
            int result = DigitalSum(5);
            Console.WriteLine(result);
        }
        
        public static int DigitalSum(int referenceNumber)
        {
            if (referenceNumber > 0)
            {
                if (referenceNumber < 10)
                    return 10 + referenceNumber - 1;
                else if (referenceNumber < 19)
                    return 10 * (referenceNumber - 9) + 9;
                else
                {
                    var mod = referenceNumber % 9;
                    int division = referenceNumber / 9;
                    var result = 0;
                    for (int i = division - 1; i >= 0; i--)
                    {
                        result = result + 9 * (int)Math.Pow(10, i);
                    }
                    return result + mod * (int)Math.Pow(10, division);
                }
            }
            else
            {
                Console.WriteLine("Reference number must be greater than zero");
                return referenceNumber;
            }
        }
}
