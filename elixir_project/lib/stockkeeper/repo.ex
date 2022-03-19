defmodule Stockkeeper.Repo do
  use Ecto.Repo,
    otp_app: :stockkeeper,
    adapter: Ecto.Adapters.Postgres
end
