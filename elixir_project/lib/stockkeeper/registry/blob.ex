defmodule Stockkeeper.Registry.Blob do
  use Stockkeeper.Schema
  import Ecto.Query
  alias __MODULE__
  alias Stockkeeper.Registry.{Chunk, Repository}
  alias Stockkeeper.Repo

  schema "blob" do
    timestamps()

    has_many :chunks, Stockkeeper.Registry.Chunk

    field :digest, :string
    field :size, :integer
    field :location, :string
    field :is_ready, :boolean, default: false
  end

  @spec write_chunk(t(), Chunk.t()) :: :ok | :error
  def write_chunk(blob, chunk) do
    chunk
    |> Map.put(:blob_id, blob.id)
    |> Repo.insert()
    |> case do
      {:ok, _chunk} -> :ok
      {:error, _changeset} -> :error
    end
  end

  @spec close(t()) :: :ok | :error
  def close(blob) do
    blob
    |> Map.put(:is_ready, true)
    |> Repo.update()
    |> case do
      {:ok, _blob} -> :ok
      {:error, _changeset} -> :error
    end
  end

  @spec delete(t()) :: :ok | :error
  def delete(blob) do
    blob
    |> Repo.delete()
    |> case do
      {:ok, _blob} -> :ok
      {:error, _changeset} -> :error
    end
  end
end
