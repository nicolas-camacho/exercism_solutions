defmodule GuessingGame do
  def compare(secret_number, guess) when secret_number === guess do
    "Correct"
  end

  def compare(_secret_number, guess) when guess === :no_guess do
    "Make a guess"
  end

  def compare(secret_number, guess) do
    cond do
      guess + 1 === secret_number || guess - 1 === secret_number ->
        "So close"

      guess > secret_number ->
        "Too high"

      guess < secret_number ->
        "Too low"
    end
  end

  def compare(_secret_number) do
    "Make a guess"
  end
end
