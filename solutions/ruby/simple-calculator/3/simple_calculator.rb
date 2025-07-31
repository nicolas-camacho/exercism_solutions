# frozen_string_literal: true

# This is a simple calculator that can perform addition, multiplication and division.
class SimpleCalculator
  ALLOWED_OPERATIONS = ['+', '/', '*']
  MESSAGE_FORMAT = '%<first_operand>i %<operation>s %<second_operand>i = %<result>i'

  class UnsupportedOperation < StandardError; end

  def self.calculate(first_operand, second_operand, operation)
    raise ArgumentError unless [first_operand, second_operand].all? { |operand| operand.is_a?(Integer) }

    new(first_operand, second_operand, operation).calculate
  end

  def initialize(first_operand, second_operand, operation)
    @first_operand = first_operand
    @second_operand = second_operand
    @operation = operation
  end

  def calculate
    raise UnsupportedOperation unless ALLOWED_OPERATIONS.include?(@operation)

    begin
      result = @first_operand.public_send(@operation, @second_operand)
    rescue ZeroDivisionError
      return 'Division by zero is not allowed.'
    end

    format(MESSAGE_FORMAT, first_operand: @first_operand, operation: @operation, second_operand: @second_operand,
                           result: result)
  end
end
