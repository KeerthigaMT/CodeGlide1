package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/healthcare/mcp-server/config"
	"github.com/healthcare/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Get_statenamemediatypeextensionHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		mediaTypeExtensionVal, ok := args["mediaTypeExtension"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: mediaTypeExtension"), nil
		}
		mediaTypeExtension, ok := mediaTypeExtensionVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: mediaTypeExtension"), nil
		}
		stateNameVal, ok := args["stateName"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: stateName"), nil
		}
		stateName, ok := stateNameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: stateName"), nil
		}
		url := fmt.Sprintf("%s/%s%s", cfg.BaseURL, mediaTypeExtension, stateName)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// No authentication required for this endpoint
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateGet_statenamemediatypeextensionTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_stateNamemediaTypeExtension",
		mcp.WithDescription("Returns pages content."),
		mcp.WithString("mediaTypeExtension", mcp.Required(), mcp.Description("Omiting the param causes html to be returned.")),
		mcp.WithString("stateName", mcp.Required(), mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_statenamemediatypeextensionHandler(cfg),
	}
}
