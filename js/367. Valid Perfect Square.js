/**
 * @param {number} num
 * @return {boolean}
 */
var isPerfectSquare = function(num) {
    for (let i = 1; i <= num; i++) {
        let tmp = i * i;
        if (tmp === num) {
            return true;
        }
        else if (tmp > num) {
            return false;
        }
    }
    return false;
};