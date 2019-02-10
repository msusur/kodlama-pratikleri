static void Main(string[] args)
{
    var sentence = "Tekrar eden harfleri bulmaca";

    var result = JoinDictionaryElements(FindFrequencyOfCharactersInAString(sentence));

    Console.WriteLine(result);

	// Karmaşıklık: O(n)

    Console.ReadLine();
}

public static Dictionary<char, int> FindFrequencyOfCharactersInAString(string sentence)
{
    Dictionary<char, int> dictionary = new Dictionary<char, int>();

    for (int i = 0; i < sentence.Length; i++)
    {
        var character = char.ToLower(sentence[i]);
        if (!char.IsWhiteSpace(character))
        {
            if (dictionary.ContainsKey(character))
            {
                dictionary[character] += 1;
            }
            else
            {
                dictionary.Add(character, 1);
            }
        }
    }

    return dictionary;
}

public static string JoinDictionaryElements(Dictionary<char, int> dictionary)
{
    var lines = dictionary.Select(x => x.Key + ": " + x.Value);
    return string.Join(", ", lines);
}

[TestMethod]
public void UnitTest()
{
	var sentence = "Tekrar";

	var expectedOutput = "t: 1, e: 1, k: 1, r: 2, a: 1";

	var result = TekrarEdenHarfleriBulma.JoinDictionaryElements(TekrarEdenHarfleriBulma.FindFrequencyOfCharactersInAString(sentence));

	Assert.AreEqual(expectedOutput, result);
}