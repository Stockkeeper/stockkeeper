defmodule Stockkeeper.Registry.Repository do
  use Stockkeeper.Schema
  import Ecto.Changeset
  alias __MODULE__
  alias Stockkeeper.Registry.Blob

  @blob_upload_base_url "https://registry.io"

  schema "repository" do
    timestamps()

    field :name, :string, default: ""

    has_many :oci_image_indexes, Stockkeeper.Registry.OCIImageIndex
    has_many :oci_image_manifests, Stockkeeper.Registry.OCIImageManifest
    has_many :oras_artifact_manifests, Stockkeeper.Registry.ORASArtifactManifest
  end

  @spec new_changeset(t(), map()) :: Changeset.t(t())
  def new_changeset(repository, attrs) do
    repository
    |> cast(attrs, [:name])
    |> validate_name()
  end

  @spec validate_name(Changeset.t(t())) :: Changeset.t(t())
  defp validate_name(changeset) do
    changeset
    |> validate_required(:name, message: "Name cannot be blank.")
    |> validate_length(:name, max: 128, message: "Name cannot be more than 128 characters.")
  end

  @spec get_by_name(String.t()) :: Repository.t() | nil
  def get_by_name(name) do
    Repo.get_by(Repository, name: name)
  end

  @spec open_blob(t()) :: {:ok, Blob.t()} | {:error, :repository_not_found}
  def open_blob(repository) do
    id = Ecto.UUID.generate()

    %Blob{
      id: id,
      repository_id: repository.id
    }
    |> Repo.insert()
    |> case do
      {:ok, blob} ->
        {:ok, blob}

      {:error, _changeset} ->
        {:error, :repository_not_found}
    end
  end
end
