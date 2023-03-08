package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/philips-labs/go-unleash-api/api"
)

func resourceProject() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Provides a resource for managing unleash projects.",

		CreateContext: resourceProjectCreate,
		ReadContext:   resourceProjectRead,
		UpdateContext: resourceProjectUpdate,
		DeleteContext: resourceProjectDelete,

		// The descriptions are used by the documentation generator and the language server.
		Schema: map[string]*schema.Schema{
			"project_id": {
				Description: "The project id of the unleash project.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "The project name.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "The project description.",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func resourceProjectCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.ApiClient)

	var diags diag.Diagnostics

	project := &api.Project{
		Id:          d.Get("project_id").(string),
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
	}

	createdProject, resp, err := client.Projects.CreateProject(*project)
	if resp == nil {
		return diag.FromErr(fmt.Errorf("response is nil: %v", err))
	}
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(createdProject.Id)
	readDiags := resourceUserRead(ctx, d, meta)
	if readDiags != nil {
		diags = append(diags, readDiags...)
	}

	return diags
}

func resourceProjectRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.ApiClient)

	var diags diag.Diagnostics

	projectId := d.Get("project_id").(string)

	foundProject, _, err := client.Projects.GetProjectById(projectId)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(foundProject.Name)
	d.Set("name", foundProject.Name)
	d.Set("description", foundProject.Description)

	return diags
}

func resourceProjectUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.ApiClient)

	var diags diag.Diagnostics

	projectId := d.Get("project_id").(string)

	project := &api.Project{
		Id:          d.Get("project_id").(string),
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
	}

	_, _, err := client.Projects.UpdateProject(projectId, *project)
	if err != nil {
		return diag.FromErr(err)
	}

	readDiags := resourceProjectRead(ctx, d, meta)
	if readDiags != nil {
		diags = append(diags, readDiags...)
	}

	return diags
}

func resourceProjectDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*api.ApiClient)

	var diags diag.Diagnostics

	projectId := d.Get("project_id").(string)

	_, err := client.Projects.DeleteProject(projectId)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}
