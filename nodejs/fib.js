module.exports = {
  fibonacci: function fibonacci(n) {
    if (n <= 1) return n;
    return module.exports.fibonacci(n - 1) + module.exports.fibonacci(n - 2);
  },
};
