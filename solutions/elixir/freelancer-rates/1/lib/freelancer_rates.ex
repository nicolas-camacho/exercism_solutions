defmodule FreelancerRates do
  @doc """
  Calculate the daily rate given an hourly rate

  ## Examples

      iex> FreelancerRates.daily_rate(60)
      480.0
  """
  @spec daily_rate(number) :: float
  def daily_rate(hourly_rate) do
    hourly_rate * 8.0
  end

  @doc """
  Apply a discount to a given price

  ## Examples

      iex> FreelancerRates.apply_discount(150, 10)
      135.0
  """
  @spec apply_discount(number, number) :: float
  def apply_discount(before_discount, discount) do
    before_discount - (before_discount * discount / 100)
  end

  @doc """
  Calculate the monthly rate, and apply a discount

  ## Examples

      iex> FreelancerRates.monthly_rate(77, 10.5)
      12130
  """
  @spec monthly_rate(number, number) :: integer
  def monthly_rate(hourly_rate, discount) do
    ceil(apply_discount(daily_rate(hourly_rate) * 22, discount))
  end

  @doc """
  Calculate the number of days in a budget, given an hourly rate and a discount

  ## Examples

      iex> FreelancerRates.days_in_budget(20000, 80, 11.0)
      35.1
  """
  @spec days_in_budget(number, number, number) :: float
  def days_in_budget(budget, hourly_rate, discount) do
    Float.floor(budget / apply_discount(daily_rate(hourly_rate), discount), 1)
  end
end
