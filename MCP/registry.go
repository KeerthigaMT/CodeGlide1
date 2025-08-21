package main

import (
	"github.com/healthcare/mcp-server/config"
	"github.com/healthcare/mcp-server/models"
	tools_api "github.com/healthcare/mcp-server/tools/api"
	tools_es "github.com/healthcare/mcp-server/tools/es"
	tools_question "github.com/healthcare/mcp-server/tools/question"
	tools_pagenamemediatypeextension "github.com/healthcare/mcp-server/tools/pagenamemediatypeextension"
	tools_blog "github.com/healthcare/mcp-server/tools/blog"
	tools_glossary "github.com/healthcare/mcp-server/tools/glossary"
	tools_statenamemediatypeextension "github.com/healthcare/mcp-server/tools/statenamemediatypeextension"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_api.CreateGet_api_statesmediatypeextensionTool(cfg),
		tools_es.CreateGet_es_question_pagenamemediatypeextensionTool(cfg),
		tools_es.CreateGet_es_statenamemediatypeextensionTool(cfg),
		tools_question.CreateGet_question_pagenamemediatypeextensionTool(cfg),
		tools_pagenamemediatypeextension.CreateGet_pagenamemediatypeextensionTool(cfg),
		tools_api.CreateGet_api_blogmediatypeextensionTool(cfg),
		tools_es.CreateGet_es_pagenamemediatypeextensionTool(cfg),
		tools_api.CreateGet_api_glossarymediatypeextensionTool(cfg),
		tools_api.CreateGet_api_topicsmediatypeextensionTool(cfg),
		tools_es.CreateGet_es_glossary_pagenamemediatypeextensionTool(cfg),
		tools_blog.CreateGet_blog_pagenamemediatypeextensionTool(cfg),
		tools_es.CreateGet_es_blog_pagenamemediatypeextensionTool(cfg),
		tools_glossary.CreateGet_glossary_pagenamemediatypeextensionTool(cfg),
		tools_api.CreateGet_api_questionsmediatypeextensionTool(cfg),
		tools_api.CreateGet_api_articlesmediatypeextensionTool(cfg),
		tools_statenamemediatypeextension.CreateGet_statenamemediatypeextensionTool(cfg),
	}
}
