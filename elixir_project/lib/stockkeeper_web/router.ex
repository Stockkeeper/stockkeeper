defmodule StockkeeperWeb.Router do
  use StockkeeperWeb, :router

  # Routes for Web Browser UI

  pipeline :browser do
    plug :accepts, ["html"]
    plug :fetch_session
    plug :fetch_live_flash
    plug :put_root_layout, {StockkeeperWeb.LayoutView, :root}
    plug :protect_from_forgery
    plug :put_secure_browser_headers
  end

  scope "/", StockkeeperWeb do
    pipe_through :browser

    get "/", PageController, :index
  end

  # Routes for OCI Registry API

  pipeline :registry_api do
    plug :accepts, ["json", "application/octet-stream"]
  end

  scope "/v2", StockkeeperWeb.RegistryAPI, as: :registry_api do
    pipe_through :registry_api

    get "/", SpecSupportController, :show

    get "/:repo_name/blobs/:digest", BlobController, :show
    delete "/:repo_name/blobs/:digest", BlobController, :delete

    post "/:repo_name/blobs/uploads", BlobUploadController, :create
    patch "/:repo_name/blobs/uploads/:reference", BlobUploadController, :update
    put "/:repo_name/blobs/uploads/:reference", BlobUploadController, :finish

    get "/:repo_name/manifests/:manifest_ref", ManifestController, :show
    put "/:repo_name/manifests/:manifest_ref", ManifestController, :update
    delete "/:repo_name/manifests/:manifest_ref", ManifestController, :delete

    get "/:repo_name/tags/list", UnsupportedController, :error
  end

  # Enables LiveDashboard only for development
  #
  # If you want to use the LiveDashboard in production, you should put
  # it behind authentication and allow only admins to access it.
  # If your application does not have an admins-only section yet,
  # you can use Plug.BasicAuth to set up some basic authentication
  # as long as you are also using SSL (which you should anyway).
  if Mix.env() in [:dev, :test] do
    import Phoenix.LiveDashboard.Router

    scope "/" do
      pipe_through :browser

      live_dashboard "/dashboard", metrics: StockkeeperWeb.Telemetry
    end
  end

  # Enables the Swoosh mailbox preview in development.
  #
  # Note that preview only shows emails that were sent by the same
  # node running the Phoenix server.
  if Mix.env() == :dev do
    scope "/dev" do
      pipe_through :browser

      forward "/mailbox", Plug.Swoosh.MailboxPreview
    end
  end
end
