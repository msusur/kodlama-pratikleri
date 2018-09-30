const scoreToResultCode = inputScore => {
  const scores = [90, 80, 70, 65, 60, 50, 30, 0]
  const codes = ['AA', 'BA', 'BB', 'CB', 'CC', 'DD', 'FD', 'FF']
  let code

  scores.every((score, idx) => {
    if (inputScore >= score) code = codes[idx]
    return !code
  })

  return code
}

// Test
const testMapping = { 30: 'FD', 90: 'AA', 66: 'CB' }
Object.keys(testMapping).forEach(score => 
  console.assert(scoreToResultCode(score) === testMapping[score])
)
