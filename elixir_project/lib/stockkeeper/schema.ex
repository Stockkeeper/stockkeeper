defmodule Stockkeeper.Schema do
  @moduledoc """
  A customized Ecto schema:
  - Uses UUIDs for primary and foreign keys.
  - Uses UTC datetimes with microsecond precision for timestamps.
  """

  defmacro __using__(_) do
    quote do
      use Ecto.Schema
      @type t :: %__MODULE__{}
      @primary_key {:id, Ecto.UUID, autogenerate: true}
      @foreign_key_type Ecto.UUID
      @derive {Phoenix.Param, key: :id}
      @timestamps_opts [type: :utc_datetime_usec]
    end
  end
end
