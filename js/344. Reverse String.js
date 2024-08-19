/**
 * @param {character[]} s
 * @return {void} Do not return anything, modify s in-place instead.
 */
var reverseString = function (s) {
    let sLength = s.length;
    for (let i = 0; i < sLength / 2; i++) {
        let tmp = s[i];
        s[i] = s[sLength - i - 1];
        s[sLength - 1 - i] = tmp;
    }
};