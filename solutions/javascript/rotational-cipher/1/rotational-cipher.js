//
// This is only a SKELETON file for the 'Rotational Cipher' exercise. It's been provided as a
// convenience to get you started writing code faster.
//
const ALPHABET = "abcdefghijklmnopqrstuvwxyz"

export const rotate = (text, number) => {
  return [...text].map((char) => {
    if (!char.match(/[a-z]/i)) {
      return char
    }

    let isUpperCase = char.match(/[A-Z]/)

    let index = isUpperCase ? 
      ALPHABET.indexOf(char.toLowerCase()) : 
      ALPHABET.indexOf(char);
    
    let newIndex = (index + number) % ALPHABET.length

    return isUpperCase ? 
      ALPHABET[newIndex].toUpperCase() : 
      ALPHABET[newIndex]
  }).join("");
};