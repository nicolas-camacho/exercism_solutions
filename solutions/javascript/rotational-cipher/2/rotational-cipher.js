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

    let index = ALPHABET.indexOf(char.toLowerCase());

    let newIndex = (index + number) % ALPHABET.length

    return char.match(/[A-Z]/) ? 
      ALPHABET[newIndex].toUpperCase() : 
      ALPHABET[newIndex]
  }).join("");
};