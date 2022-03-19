defmodule Stockkeeper.Repo.Migrations.CreateIndexNameOnRepository do
  use Ecto.Migration

  def change do
    create index(:repository, [:name], unique: true)
  end
end
