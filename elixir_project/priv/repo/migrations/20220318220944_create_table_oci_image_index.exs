defmodule Stockkeeper.Repo.Migrations.CreateTableOciImageIndex do
  use Ecto.Migration

  def change do
    create table(:oci_image_index, primary_key: false) do
      add :id, :uuid, primary_key: true
      timestamps(type: :utc_datetime_usec)

      add :repository_id, references(:repository, type: :uuid, null: false)

      add :properties, :map, default: %{}
    end
  end
end
