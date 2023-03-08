package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/philips-labs/go-unleash-api/api"
)

func resourceProjectUser() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Provides a resource for managing unleash project users.",

		CreateContext: resourceProjectUserAdd,
		UpdateContext: resourceProjectUserUpdate,
		DeleteContext: resourceProjectUserDelete,

		// The descriptions are used by the documentation generator and the language server.
		Schema: map[string]*schema.Schema{
			"user_id": {
				Description: "The user's id.",
				Type:        schema.TypeInt,
				Required:    true,
			},
			"project_id": {
				Description: "The project id.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"role_id": {
				Description: "The project user's role.",
				Type:        schema.TypeInt,
				Required:    true,
			},
		},
	}
}

func resourceProjectUserAdd(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.ApiClient)

	var diags diag.Diagnostics

	userID := d.Get("user_id").(int)
	projectID := d.Get("project_id").(string)
	roleID := d.Get("role_id").(int)

	_, _, err := client.Projects.AddUserProject(userID, projectID, roleID)
	if err != nil {
		return diag.FromErr(err)
	}

	readDiags := resourceProjectRead(ctx, d, meta)
	if readDiags != nil {
		diags = append(diags, readDiags...)
	}

	return diags
}

func resourceProjectUserDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.ApiClient)

	var diags diag.Diagnostics

	userID := d.Get("user_id").(int)
	projectID := d.Get("project_id").(string)
	roleID := d.Get("role_id").(int)

	_, _, err := client.Projects.DeleteUserProject(projectID, userID, roleID)
	if err != nil {
		return diag.FromErr(err)
	}

	readDiags := resourceProjectRead(ctx, d, meta)
	if readDiags != nil {
		diags = append(diags, readDiags...)
	}

	return diags
}

func resourceProjectUserUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.ApiClient)

	var diags diag.Diagnostics

	projectID := d.Get("project_id").(string)
	userID := d.Get("user_id").(int)
	roleID := d.Get("role_id").(int)

	_, _, err := client.Projects.UpdateUserProject(projectID, userID, roleID)
	if err != nil {
		return diag.FromErr(err)
	}

	readDiags := resourceProjectRead(ctx, d, meta)
	if readDiags != nil {
		diags = append(diags, readDiags...)
	}

	return diags
}
