defmodule Stockkeeper.Application do
  # See https://hexdocs.pm/elixir/Application.html
  # for more information on OTP Applications
  @moduledoc false

  use Application

  @impl true
  def start(_type, _args) do
    children = [
      # Start the Ecto repository
      Stockkeeper.Repo,
      # Start the Telemetry supervisor
      StockkeeperWeb.Telemetry,
      # Start the PubSub system
      {Phoenix.PubSub, name: Stockkeeper.PubSub},
      # Start the Endpoint (http/https)
      StockkeeperWeb.Endpoint
      # Start a worker by calling: Stockkeeper.Worker.start_link(arg)
      # {Stockkeeper.Worker, arg}
    ]

    # See https://hexdocs.pm/elixir/Supervisor.html
    # for other strategies and supported options
    opts = [strategy: :one_for_one, name: Stockkeeper.Supervisor]
    Supervisor.start_link(children, opts)
  end

  # Tell Phoenix to update the endpoint configuration
  # whenever the application is updated.
  @impl true
  def config_change(changed, _new, removed) do
    StockkeeperWeb.Endpoint.config_change(changed, removed)
    :ok
  end
end
