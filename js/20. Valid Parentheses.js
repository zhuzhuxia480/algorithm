/**
 * @param {string} s
 * @return {boolean}
 */
var isValid = function (s) {
    let stack = [];
    stack.push(s[0]);
    let size = s.length;
    for (let i = 1; i < size; i++) {
        if (s[i] === '(' || s[i] === '[' || s[i] === '{') {
            stack.push(s[i]);
        } else {
            if (stack.length === 0) {
                return false;
            }
            if (s[i] === ')' && stack[stack.length - 1] === '(' ||
                s[i] === ']' && stack[stack.length - 1] === '[' ||
                s[i] === '}' && stack[stack.length - 1] === '{') {
                stack.pop();
                continue;
            }
            return false;
        }
    }
    return stack.length === 0
};