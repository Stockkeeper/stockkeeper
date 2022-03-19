defmodule Stockkeeper.Repo.Migrations.CreateTableUploadSession do
  use Ecto.Migration

  def change do
    create table(:upload_session, primary_key: false) do
      add :id, :uuid, primary_key: true
      timestamps(type: :utc_datetime_usec)

      add :repository_id, references(:repository, type: :uuid, null: false)

      add :location, :string, null: false
      add :is_multichunk, :boolean, null: false
    end
  end
end
