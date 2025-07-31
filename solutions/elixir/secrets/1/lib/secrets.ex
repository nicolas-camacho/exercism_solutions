defmodule Secrets do
  import Bitwise

  @doc """
  Adds a secret number to a given number

  ## Examples

      iex> adder = Secrets.secret_add(2)
      iex> adder.(2)
      4
  """
  @spec secret_add(integer) :: (integer -> integer)
  def secret_add(secret), do: fn number -> secret + number end

  @doc """
  Subtracts a secret number from a given number

  ## Examples

      iex> subtractor = Secrets.secret_subtract(2)
      iex> subtractor.(3)
      1
  """

  @spec secret_subtract(integer) :: (integer -> integer)
  def secret_subtract(secret), do: fn number -> number - secret end

  @doc """
  Multiplies a secret number with a given number

  ## Examples

      iex> multiplier = Secrets.secret_multiply(7)
      iex> multiplier.(3)
      21
  """
  @spec secret_multiply(integer) :: (integer -> integer)
  def secret_multiply(secret), do: fn number -> number * secret end

  @doc """
  Divides a given number by a secret number

  ## Examples

      iex> divider = Secrets.secret_divide(3)
      iex> divider.(32)
      10
  """
  @spec secret_divide(integer) :: (integer -> integer)
  def secret_divide(secret), do: fn number -> div(number, secret) end

  @doc """
  Performs a bitwise AND operation with a secret number and a given number

  ## Examples

      iex> ander = Secrets.secret_and(1)
      iex> ander.(2)
      0
  """
  def secret_and(secret), do: fn number -> band(number, secret) end

  @doc """
  Performs a bitwise XOR operation with a secret number and a given number

  ## Examples

      iex> xorer = Secrets.secret_xor(1)
      iex> xorer.(3)
      2
  """
  @spec secret_xor(integer) :: (integer -> integer)
  def secret_xor(secret), do: fn number -> bxor(number, secret) end

  @doc """
  Combines two secret functions

  ## Examples

      iex> multiply = Secrets.secret_multiply(7)
      iex> divide = Secrets.secret_divide(3)
      iex> combined = Secrets.secret_combine(multiply, divide)
      iex> combined.(6)
      14
  """
  @spec secret_combine((integer -> integer), (integer -> integer)) :: (integer -> integer)
  def secret_combine(secret_function1, secret_function2),
    do: fn number -> secret_function2.(secret_function1.(number)) end
end
