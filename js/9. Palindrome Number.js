/**
 * @param {number} x
 * @return {boolean}
 */
const isPalindrome = function (x) {
    let str = String(x);
    const strLen = str.length;
    for (let i = 0; i < strLen; i++) {
        if (str[i] !== str[strLen - 1 - i]) {
            return false;
        }
    }
    return true;
};

console.log("test")