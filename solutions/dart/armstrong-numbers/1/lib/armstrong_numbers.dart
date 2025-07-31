class ArmstrongNumbers {
  bool isArmstrongNumber(String number) {
    BigInt sum = BigInt.zero;
    int length = number.length;
    for (var i = 0; i < length; i++) {
      sum += BigInt.from(int.parse(number[i])).pow(length);
    }

    return sum == BigInt.parse(number);
  }
}
