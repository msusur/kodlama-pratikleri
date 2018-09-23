const multipleOccurencesOfLetters = sentence => {
  const letters = {}
  sentence = sentence.toLowerCase()

  for (let i = 0; i < sentence.length; i++) {
    const char = sentence[i]

    if (letters.char || !char.match(/[a-z]/)) continue

    letters[char] = sentence.match(new RegExp(char, 'g')).length
  }
  
  return Object.keys(letters)
    .filter(k => letters[k] > 1)
    .map(k => ({ [k]: letters[k] }))
}

console.log(multipleOccurencesOfLetters('Tekrar eden harfleri bulmaca!'))
// [ { e: 4 }, { r: 4 }, { a: 4 }, { l: 2 } ]
