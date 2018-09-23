const findMinMissingPositive = arr => {
  const max = Math.max(...arr)
  let rv

  for (let i = 1; i < max; i++) {
    if (arr.indexOf(i) > -1)
      continue
    
    rv = i
    break
  }

  return rv || max + 1
}

console.log(findMinMissingPositive([-4,4,6,2,7,3,1])) // 5
console.log(findMinMissingPositive([1,2,0])) // 3
