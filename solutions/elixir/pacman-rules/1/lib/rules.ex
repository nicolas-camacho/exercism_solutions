defmodule Rules do
  @doc """
  Determines if pacman can eat a ghost

  ## Examples

    iex> Rules.eat_ghost?(false, true)
    false
  """
  @spec eat_ghost?(power_pellet_active? :: boolean(), touching_ghost? :: boolean()) :: boolean()
  def eat_ghost?(power_pellet_active?, touching_ghost?) do
    power_pellet_active? and touching_ghost?
  end

  @doc """
  Determines if pacman scored

  ## Examples

    iex> Rules.score?(true, true)
    true
  """
  @spec score?(touching_power_pellet? :: boolean(), touching_dot? :: boolean()) :: boolean()
  def score?(touching_power_pellet?, touching_dot?) do
    touching_power_pellet? or touching_dot?
  end

  @doc """
  Tells if pacman losed

  ## Examples

    iex> Rules.lose?(false, true)
    true
  """
  @spec lose?(power_pellet_active? :: boolean(), touching_ghost? :: boolean()) :: boolean()
  def lose?(power_pellet_active?, touching_ghost?) do
    not power_pellet_active? and touching_ghost?
  end

  @doc """
  Tells if pacman win

  ## Examples

    iex> Rules.win?(false, true, false)
    false
  """
  @spec win?(
          has_eaten_all_dots? :: boolean(),
          power_pellet_active? :: boolean(),
          touching_ghost :: boolean()
        ) :: boolean()
  def win?(has_eaten_all_dots?, power_pellet_active?, touching_ghost?) do
    has_eaten_all_dots? and not lose?(power_pellet_active?, touching_ghost?)
  end
end
