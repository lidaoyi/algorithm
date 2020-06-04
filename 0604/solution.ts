function productExceptSelf(nums: number[]): number[] {
    let len = nums.length;
    let left = new Array(len);
    left[0] = 1;
    let output = new Array(len);
    output[len - 1] = 1;

    for (let i = 1; i < len; i++) {
        left[i] = left[i - 1] * nums[i - 1];
        output[len - 1 - i] = output[len - i] * nums[len - i]
    }

    for (let i = 0; i < len; i++) {
        output[i] = left[i] * output[i]
    }
    return output
}