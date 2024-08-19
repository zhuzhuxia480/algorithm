/**
 Do not return anything, modify s in-place instead.
 */
function reverseString(s: string[]): void {
    let slen: number = s.length;
    for (let i: number = 0; i < slen / 2; i++) {
        let tmp: string = s[i];
        s[i] = s[slen - 1 - i];
        s[slen - 1 - i] = tmp;
    }
};