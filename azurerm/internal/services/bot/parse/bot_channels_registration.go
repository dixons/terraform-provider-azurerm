package parse

import (
	"fmt"

	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
)

type BotChannelsRegistrationId struct {
	ResourceGroup string
	// BotName       string
	Name string
}

func BotChannelsRegistrationID(input string) (*BotChannelsRegistrationId, error) {
	id, err := azure.ParseAzureResourceID(input)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Unable to parse Bot Channels Registration ID %q: %+v", input, err)
	}

	service := BotChannelsRegistrationId{
		ResourceGroup: id.ResourceGroup,
	}

	if service.Name, err = id.PopSegment("botServices"); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &service, nil
}