defmodule Stockkeeper.Repo.Migrations.CreateTableRepository do
  use Ecto.Migration

  def change do
    create table(:repository, primary_key: false) do
      add :id, :uuid, primary_key: true
      timestamps(type: :utc_datetime_usec)

      add :name, :string, null: false
    end
  end
end
