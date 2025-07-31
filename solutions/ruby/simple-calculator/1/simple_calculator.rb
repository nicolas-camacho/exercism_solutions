# frozen_string_literal: true

# This is a simple calculator that can perform addition, multiplication and division.
class SimpleCalculator
  ALLOWED_OPERATIONS = ['+', '/', '*'].freeze

  class UnsupportedOperation < StandardError; end

  def self.calculate(first_operand, second_operand, operation)
    raise UnsupportedOperation unless ALLOWED_OPERATIONS.include?(operation)
    raise ArgumentError if !first_operand.is_a?(Numeric) || !second_operand.is_a?(Numeric)
    return 'Division by zero is not allowed.' if operation == '/' && second_operand.zero?

    result = first_operand.public_send(operation, second_operand)

    "#{first_operand} #{operation} #{second_operand} = #{result}"
  end
end
