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
                int subtraction = referenceNumber - 1;
		int division = subtraction / 9;
		return division == 0 ? referenceNumber + 9 : (referenceNumber - 9 * division + 1) * Math.Pow(10, division) - 1;
            }
            else
            {
                Console.WriteLine("Reference number must be greater than zero");
                return referenceNumber;
            }
        }
}
