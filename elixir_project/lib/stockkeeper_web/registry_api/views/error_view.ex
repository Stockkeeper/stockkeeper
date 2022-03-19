defmodule StockkeeperWeb.RegistryAPI.ErrorView do
  use StockkeeperWeb, :view

  def render("error.json", %{code: code, message: message}) do
    %{
      errors: [
        %{
          code: code,
          message: message,
          details: "TODO: Add debugging details, such as the resource name, resource ID, etc."
        }
      ]
    }
  end

  def render("unsupported.json", _assigns) do
    %{
      errors: [
        %{
          code: "UNSUPPORTED",
          message: "This endpoint is not yet implemented."
        }
      ]
    }
  end
end
