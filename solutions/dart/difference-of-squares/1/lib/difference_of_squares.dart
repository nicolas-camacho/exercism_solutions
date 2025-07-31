class DifferenceOfSquares {
  double squareOfSum(double n) {
    double result = n * (n + 1) / 2;
    return result * result;
  }

  double sumOfSquares(double n) {
    return n * (n + 1) * (2 * n + 1) / 6;
  }

  double differenceOfSquares(double n) {
    return squareOfSum(n) - sumOfSquares(n);
  }
}
