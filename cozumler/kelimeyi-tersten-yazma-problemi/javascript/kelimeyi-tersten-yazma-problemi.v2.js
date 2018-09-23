const reverseString = str => str ? reverseString(str.substr(1)) + str[0] : ''

console.log(reverseString('Hello World!')); // !dlroW olleH
