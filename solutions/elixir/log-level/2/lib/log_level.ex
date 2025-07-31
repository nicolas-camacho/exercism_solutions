defmodule LogLevel do
  @doc """
  Converts a log level to a label.

  ## Examples

      iex> LogLevel.to_label(0, false)
      :trace
  """
  @spec to_label(level :: integer, legacy? :: boolean) :: atom
  def to_label(level, legacy?) do
    cond do
      level == 0 && !legacy? -> :trace
      level == 1 -> :debug
      level == 2 -> :info
      level == 3 -> :warning
      level == 4 -> :error
      level == 5 && !legacy? -> :fatal
      true -> :unknown
    end
  end

  @doc """
  Determines the recipient of an alert based on the log level.

  ## Examples

      iex> LogLevel.alert_recipient(-1, true)
      :dev1
  """
  @spec alert_recipient(level :: integer, legacy? :: boolean) :: atom | boolean
  def alert_recipient(level, legacy?) do
    label = to_label(level, legacy?)
    cond do
      label == :fatal || label == :error -> :ops
      label == :unknown && legacy? -> :dev1
      label == :unknown && !legacy? -> :dev2
      true -> false
    end
  end
end
