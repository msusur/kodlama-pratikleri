class ReverseText
    {
        static void Main(string[] args)
        {
            string text = "Muvaffakiyetsizleştiricileştiriveremeyebileceklerimizdenmişsinizcesine";
            Console.WriteLine(Reverse(text));
        }

        public static string Reverse(string text)
        {
            var s= text[text.Length - 1];
            if (text.Length == 1)
                return text;
            return text[text.Length - 1] + Reverse(text.Substring(0,text.Length-1));
        }
    }
