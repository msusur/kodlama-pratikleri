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

console.log(scoreToResultCode(30)) // FD
console.log(scoreToResultCode(90)) // AA
console.log(scoreToResultCode(66)) // CB
