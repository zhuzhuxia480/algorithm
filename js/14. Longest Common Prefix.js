/**
 * @param {string[]} strs
 * @return {string}
 */
var longestCommonPrefix = function(strs) {
    let preFixLen = 0;
    let size = strs[0].length;
    for (let i = 0; i < size; i++) {
        if (!checkWithIndex(preFixLen, strs)) {
            break;
        }
        preFixLen++;
    }
    return strs[0].substring(0, preFixLen);
};


var checkWithIndex = function (index, strs) {
    if (index >= strs[0].length) {
        return false;
    }
    let size = strs.length;
    let flag = strs[0][index];
    for (let i = 1; i < size; i++) {
        if (index >= strs[i].length) {
            return false;
        }
        if (flag !== strs[i][index]) {
            return false;
        }
    }
    return true;
};