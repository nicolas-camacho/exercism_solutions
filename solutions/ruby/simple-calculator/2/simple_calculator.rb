# frozen_string_literal: true

# This is a simple calculator that can perform addition, multiplication and division.
class SimpleCalculator
  ALLOWED_OPERATIONS = ['+', '/', '*']

  class UnsupportedOperation < StandardError; end

  def self.calculate(first_operand, second_operand, operation)
    new(first_operand, second_operand, operation).calculate
  end

  def initialize(first_operand, second_operand, operation)
    @first_operand = first_operand
    @second_operand = second_operand
    @operation = operation
  end

  def calculate
    raise UnsupportedOperation unless ALLOWED_OPERATIONS.include?(@operation)
    raise ArgumentError unless @first_operand.is_a?(Numeric) && @second_operand.is_a?(Numeric)
    return 'Division by zero is not allowed.' if @operation == '/' && @second_operand.zero?

    result = @first_operand.public_send(@operation, @second_operand)

    format('%<first_operand>d %<operation>s %<second_operand>d = %<result>d',
           first_operand: @first_operand,
           operation: @operation,
           second_operand: @second_operand,
           result: result)
  end
end
