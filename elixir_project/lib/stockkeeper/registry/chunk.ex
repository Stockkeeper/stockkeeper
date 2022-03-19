defmodule Stockkeeper.Registry.Chunk do
  use Stockkeeper.Schema

  schema "chunk" do
    timestamps()

    belongs_to :blob, Stockkeeper.Registry.Blob, type: Ecto.UUID

    field :size, :integer
    field :range_start, :integer
    field :range_end, :integer
    field :location, :string
  end
end
