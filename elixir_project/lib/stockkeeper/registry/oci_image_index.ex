defmodule Stockkeeper.Registry.OCIImageIndex do
  use Stockkeeper.Schema

  schema "oci_image_index" do
    timestamps()

    belongs_to :repository, Stockkeeper.Registry.Repository, type: Ecto.UUID

    field :properties, :map, default: %{}
  end
end
