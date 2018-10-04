const findNearestPoints = (points=[], k=points.length) => {
  const distance = (x, y) => Math.sqrt(x ** 2 + y ** 2)

  return points
    .sort((p, _p) => distance(...p) > distance(..._p))
    .slice(0, k)
}

// Test
const result = JSON.stringify(findNearestPoints([[1, 2], [5, 7], [8, 9], [2, 3], [2, 2]], 3))
const expected = JSON.stringify([[1, 2], [2, 2], [2, 3]])
console.assert(result === expected)
