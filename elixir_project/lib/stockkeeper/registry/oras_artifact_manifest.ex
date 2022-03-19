defmodule Stockkeeper.Registry.ORASArtifactManifest do
  use Stockkeeper.Schema

  schema "oras_artifact_manifest" do
    timestamps()

    belongs_to :repository, Stockkeeper.Registry.Repository, type: Ecto.UUID

    field :properties, :map, default: %{}
  end
end
