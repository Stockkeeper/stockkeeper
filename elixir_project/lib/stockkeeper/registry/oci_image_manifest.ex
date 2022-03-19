defmodule Stockkeeper.Registry.OCIImageManifest do
  use Stockkeeper.Schema

  schema "oci_image_manifest" do
    timestamps()

    belongs_to :repository, Stockkeeper.Registry.Repository, type: Ecto.UUID

    field :properties, :map, default: %{}
  end
end
