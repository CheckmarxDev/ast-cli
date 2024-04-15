package commands

import (
	"encoding/json"

	featureFlagsConstants "github.com/checkmarx/ast-cli/internal/constants/feature-flags"
	commonParams "github.com/checkmarx/ast-cli/internal/params"
	"github.com/checkmarx/ast-cli/internal/shared"
	"github.com/checkmarx/ast-cli/internal/wrappers"
	"github.com/spf13/cobra"
)

func updateGroupValues(input *[]byte, cmd *cobra.Command, groupsWrapper wrappers.GroupsWrapper) ([]*wrappers.Group, error) {
	groupListStr, _ := cmd.Flags().GetString(commonParams.GroupList)
	groups, err := shared.CreateGroupsMap(groupListStr, groupsWrapper)
	if err != nil {
		return groups, err
	}
	if !wrappers.FeatureFlags[featureFlagsConstants.AccessManagementEnabled] {
		var info map[string]interface{}
		_ = json.Unmarshal(*input, &info)
		info["groups"] = shared.GetGroupIds(groups)
		*input, _ = json.Marshal(info)
	}
	return groups, nil
}
