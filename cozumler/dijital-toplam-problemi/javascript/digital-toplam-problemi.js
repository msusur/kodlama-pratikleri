const minNumberOfDigitalSum = x => {
	const a = Math.ceil(x / 9) - 1
	
  return !a ? x + 9 : parseInt((x % 9 || 9) + Array(a).fill(9).join(''))
}

console.log(minNumberOfDigitalSum(19)) // 199
