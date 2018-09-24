const findDuplicateChars = sentence => {
  return sentence.toLowerCase()
          .replace(/\s+/g,"")
          .split("")
          .sort()
          .join("")
          .match(/(.)\1+/g)
          .reduce(function(result, item) {
            result[item[0]] = item.length; 
            return result;
          }, {});
}

console.log(findDuplicateChars('Tekrar eden harfleri bulmaca!'))
// {a: 4, e: 4, l: 2, r: 4}