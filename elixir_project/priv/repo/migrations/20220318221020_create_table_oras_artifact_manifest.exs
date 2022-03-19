defmodule Stockkeeper.Repo.Migrations.CreateTableOrasArtifactManifest do
  use Ecto.Migration

  def change do
    create table(:oras_artifact_manifest, primary_key: false) do
      add :id, :uuid, primary_key: true
      timestamps(type: :utc_datetime_usec)

      add :repository_id, references(:repository, type: :uuid, null: false)

      add :properties, :map, default: %{}
    end
  end
end
